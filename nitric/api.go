// Copyright 2021 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nitric

import (
	"net/http"
	"path"
	"strings"

	httpx "github.com/nitrictech/go-sdk/nitric/apis"

	apispb "github.com/nitrictech/nitric/core/pkg/proto/apis/v1"
	resourcev1 "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
)

// Route providers convenience functions to register a handler in a single method.
type Route interface {
	// All adds a handler for all HTTP methods to the route.
	All(handler interface{}, opts ...MethodOption)
	// Get adds a Get method handler to the route.
	Get(handler interface{}, opts ...MethodOption)
	// Put adds a Put method handler to the route.
	Patch(handler interface{}, opts ...MethodOption)
	// Patch adds a Patch method handler to the route.
	Put(handler interface{}, opts ...MethodOption)
	// Post adds a Post method handler to the route.
	Post(handler interface{}, opts ...MethodOption)
	// Delete adds a Delete method handler to the route.
	Delete(handler interface{}, opts ...MethodOption)
	// Options adds an Options method handler to the route.
	Options(handler interface{}, opts ...MethodOption)
	// ApiName returns the name of the API this route belongs to.
	ApiName() string
}

type route struct {
	path       string
	api        *api
	middleware Middleware[httpx.Ctx]
	manager    *manager
}

func composeRouteMiddleware(apiMiddleware Middleware[httpx.Ctx], routeMiddleware []Middleware[httpx.Ctx]) Middleware[httpx.Ctx] {
	allMiddleware := append([]Middleware[httpx.Ctx]{apiMiddleware}, routeMiddleware...)

	return ComposeMiddleware(allMiddleware...)
}

func (a *api) NewRoute(match string, middleware ...Middleware[httpx.Ctx]) Route {
	r, ok := a.routes[match]
	if !ok {
		r = &route{
			manager:    a.manager,
			path:       path.Join(a.path, match),
			api:        a,
			middleware: composeRouteMiddleware(a.middleware, middleware),
		}
	}

	return r
}

func (r *route) ApiName() string {
	return r.api.name
}

func (r *route) AddMethodHandler(methods []string, middleware interface{}, opts ...MethodOption) error {
	bName := path.Join(r.api.name, r.path, strings.Join(methods, "-"))

	// default methodOptions will contain OidcOptions passed to API instance and securityDisabled to false
	mo := &methodOptions{
		securityDisabled: false,
		security:         r.api.security,
	}

	for _, o := range opts {
		o(mo)
	}

	mw, err := interfaceToMiddleware[httpx.Ctx](middleware)
	if err != nil {
		panic(err)
	}

	composedHandler := ComposeMiddleware(r.middleware, mw)

	apiOpts := &apispb.ApiWorkerOptions{
		SecurityDisabled: mo.securityDisabled,
		Security:         map[string]*apispb.ApiWorkerScopes{},
	}

	if mo.security != nil && !mo.securityDisabled {
		for _, oidcOption := range mo.security {
			err := attachOidc(r.api.name, oidcOption)
			if err != nil {
				return err
			}

			apiOpts.Security[oidcOption.Name] = &apispb.ApiWorkerScopes{
				Scopes: oidcOption.Scopes,
			}
		}
	}

	registrationRequest := &apispb.RegistrationRequest{
		Path:    r.path,
		Api:     r.api.name,
		Methods: methods,
		Options: apiOpts,
	}

	wkr := newApiWorker(&apiWorkerOpts{
		RegistrationRequest: registrationRequest,
		Middleware:          composedHandler,
	})

	r.manager.addWorker("route:"+bName, wkr)

	return nil
}

func (r *route) All(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions}, handler, opts...)
}

func (r *route) Get(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodGet}, handler, opts...)
}

func (r *route) Post(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodPost}, handler, opts...)
}

func (r *route) Put(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodPut}, handler, opts...)
}

func (r *route) Patch(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodPatch}, handler, opts...)
}

func (r *route) Delete(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodDelete}, handler, opts...)
}

func (r *route) Options(handler interface{}, opts ...MethodOption) {
	_ = r.AddMethodHandler([]string{http.MethodOptions}, handler, opts...)
}

// Api Resource represents an HTTP API, capable of routing and securing incoming HTTP requests to handlers.
// path is the route path matcher e.g. '/home'. Supports path params via colon prefix e.g. '/customers/:customerId'
// handler the handler to register for callbacks.
//
// Note: to chain middleware use handler.ComposeHttpMiddlware()
type Api interface {
	// Get adds a Get method handler to the path with any specified opts.
	// Valid function signatures:
	//
	//	func()
	//	func() error
	//	func(*apis.Ctx)
	//	func(*apis.Ctx) error
	//	func(*apis.Ctx) *apis.Ctx
	//	func(*apis.Ctx) (*apis.Ctx, error)
	//	func(*apis.Ctx, Handler[apis.Ctx]) *apis.Ctx
	//	func(*apis.Ctx, Handler[apis.Ctx]) error
	//	func(*apis.Ctx, Handler[apis.Ctx]) (*apis.Ctx, error)
	//	Middleware[apis.Ctx]
	//	Handler[apis.Ctx]
	Get(path string, handler interface{}, opts ...MethodOption)
	// Put adds a Put method handler to the path with any specified opts.
	// Valid function signatures:
	//
	//	func()
	//	func() error
	//	func(*apis.Ctx)
	//	func(*apis.Ctx) error
	//	func(*apis.Ctx) *apis.Ctx
	//	func(*apis.Ctx) (*apis.Ctx, error)
	//	func(*apis.Ctx, Handler[apis.Ctx]) *apis.Ctx
	//	func(*apis.Ctx, Handler[apis.Ctx]) error
	//	func(*apis.Ctx, Handler[apis.Ctx]) (*apis.Ctx, error)
	//	Middleware[apis.Ctx]
	//	Handler[apis.Ctx]
	Put(path string, handler interface{}, opts ...MethodOption)
	// Patch adds a Patch method handler to the path with any specified opts.
	// Valid function signatures:
	//
	//	func()
	//	func() error
	//	func(*apis.Ctx)
	//	func(*apis.Ctx) error
	//	func(*apis.Ctx) *apis.Ctx
	//	func(*apis.Ctx) (*apis.Ctx, error)
	//	func(*apis.Ctx, Handler[apis.Ctx]) *apis.Ctx
	//	func(*apis.Ctx, Handler[apis.Ctx]) error
	//	func(*apis.Ctx, Handler[apis.Ctx]) (*apis.Ctx, error)
	//	Middleware[apis.Ctx]
	//	Handler[apis.Ctx]
	Patch(path string, handler interface{}, opts ...MethodOption)
	// Post adds a Post method handler to the path with any specified opts.
	// Valid function signatures:
	//
	//	func()
	//	func() error
	//	func(*apis.Ctx)
	//	func(*apis.Ctx) error
	//	func(*apis.Ctx) *apis.Ctx
	//	func(*apis.Ctx) (*apis.Ctx, error)
	//	func(*apis.Ctx, Handler[apis.Ctx]) *apis.Ctx
	//	func(*apis.Ctx, Handler[apis.Ctx]) error
	//	func(*apis.Ctx, Handler[apis.Ctx]) (*apis.Ctx, error)
	//	Middleware[apis.Ctx]
	//	Handler[apis.Ctx]
	Post(path string, handler interface{}, opts ...MethodOption)
	// Delete adds a Delete method handler to the path with any specified opts.
	// Valid function signatures:
	//
	//	func()
	//	func() error
	//	func(*apis.Ctx)
	//	func(*apis.Ctx) error
	//	func(*apis.Ctx) *apis.Ctx
	//	func(*apis.Ctx) (*apis.Ctx, error)
	//	func(*apis.Ctx, Handler[apis.Ctx]) *apis.Ctx
	//	func(*apis.Ctx, Handler[apis.Ctx]) error
	//	func(*apis.Ctx, Handler[apis.Ctx]) (*apis.Ctx, error)
	//	Middleware[apis.Ctx]
	//	Handler[apis.Ctx]
	Delete(path string, handler interface{}, opts ...MethodOption)
	// Options adds a Options method handler to the path with any specified opts.
	// Valid function signatures:
	//
	//	func()
	//	func() error
	//	func(*apis.Ctx)
	//	func(*apis.Ctx) error
	//	func(*apis.Ctx) *apis.Ctx
	//	func(*apis.Ctx) (*apis.Ctx, error)
	//	func(*apis.Ctx, Handler[apis.Ctx]) *apis.Ctx
	//	func(*apis.Ctx, Handler[apis.Ctx]) error
	//	func(*apis.Ctx, Handler[apis.Ctx]) (*apis.Ctx, error)
	//	Middleware[apis.Ctx]
	//	Handler[apis.Ctx]
	Options(path string, handler interface{}, opts ...MethodOption)
	// NewRoute creates a new Route object for the given path.
	NewRoute(path string, middleware ...Middleware[httpx.Ctx]) Route
}

type ApiDetails struct {
	Details
	URL string
}

type api struct {
	name          string
	routes        map[string]Route
	manager       *manager
	securityRules map[string]interface{}
	security      []OidcOptions
	path          string
	middleware    Middleware[httpx.Ctx]
}

// NewApi Registers a new API Resource.
//
// The returned API object can be used to register Routes and Methods, with Handlers.
func NewApi(name string, opts ...ApiOption) (Api, error) {
	a := &api{
		name:    name,
		routes:  map[string]Route{},
		manager: defaultManager,
	}

	// Apply options
	for _, o := range opts {
		o(a)
	}

	apiResource := &resourcev1.ApiResource{}

	// Attaching OIDC Options to API
	if a.security != nil {
		for _, oidcOption := range a.security {
			err := attachOidc(a.name, oidcOption)
			if err != nil {
				return nil, err
			}

			if apiResource.GetSecurity() == nil {
				apiResource.Security = make(map[string]*resourcev1.ApiScopes)
			}
			apiResource.Security[oidcOption.Name] = &resourcev1.ApiScopes{
				Scopes: oidcOption.Scopes,
			}
		}
	}
	// declare resource
	result := <-defaultManager.registerResource(&resourcev1.ResourceDeclareRequest{
		Id: &resourcev1.ResourceIdentifier{
			Name: name,
			Type: resourcev1.ResourceType_Api,
		},
		Config: &resourcev1.ResourceDeclareRequest_Api{
			Api: apiResource,
		},
	})
	if result.Err != nil {
		return nil, result.Err
	}

	return a, nil
}

// Get adds a Get method handler to the path with any specified opts.
// Note: to chain middleware use handler.ComposeHttpMiddlware()
func (a *api) Get(match string, handler interface{}, opts ...MethodOption) {
	r := a.NewRoute(match)

	r.Get(handler, opts...)
	a.routes[match] = r
}

// Post adds a Post method handler to the path with any specified opts.
// Note: to chain middleware use handler.ComposeHttpMiddlware()
func (a *api) Post(match string, handler interface{}, opts ...MethodOption) {
	r := a.NewRoute(match)

	r.Post(handler, opts...)
	a.routes[match] = r
}

// Patch adds a Patch method handler to the path with any specified opts.
// Note: to chain middleware use handler.ComposeHttpMiddlware()
func (a *api) Patch(match string, handler interface{}, opts ...MethodOption) {
	r := a.NewRoute(match)

	r.Patch(handler, opts...)
	a.routes[match] = r
}

// Put adds a Put method handler to the path with any specified opts.
// Note: to chain middleware use handler.ComposeHttpMiddlware()
func (a *api) Put(match string, handler interface{}, opts ...MethodOption) {
	r := a.NewRoute(match)

	r.Put(handler, opts...)
	a.routes[match] = r
}

// Delete adds a Delete method handler to the path with any specified opts.
// Note: to chain middleware use handler.ComposeHttpMiddlware()
func (a *api) Delete(match string, handler interface{}, opts ...MethodOption) {
	r := a.NewRoute(match)

	r.Delete(handler, opts...)
	a.routes[match] = r
}

// Options adds an Options method handler to the path with any specified opts.
// Note: to chain middleware use handler.ComposeHttpMiddlware()
func (a *api) Options(match string, handler interface{}, opts ...MethodOption) {
	r := a.NewRoute(match)

	r.Options(handler, opts...)
	a.routes[match] = r
}
