// Copyright 2023 Nitric Technologies Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nitric

import "fmt"

type (
	Handler[T any]    func(context *T) (*T, error)
	Middleware[T any] func(context *T, next Handler[T]) (*T, error)
)

type chainedMiddleware[T any] struct {
	fun      Middleware[T]
	nextFunc Handler[T]
}

func dummyHandler[T any](ctx *T) (*T, error) {
	return ctx, nil
}

func interfacesToMiddleware[T any](mws []interface{}) ([]Middleware[T], error) {
	mwares := make([]Middleware[T], len(mws))
	for i, mw := range mws {
		mware, err := interfaceToMiddleware[T](mw)
		if err != nil {
			return nil, err
		}
		mwares[i] = mware
	}
	return mwares, nil
}

// interfaceToMiddleware - Converts a function to a Middleware
// Valid function types are:
// func()
// func() error
// func(*T)
// func(*T) error
// func(*T) *T
// func(*T) (*T, error)
// func(*T, Handler[T]) *T
// func(*T, Handler[T]) error
// func(*T, Handler[T]) (*T, error)
// Middleware[T]
// Handler[T]
// If the function is not a valid type, an error is returned
func interfaceToMiddleware[T any](mw interface{}) (Middleware[T], error) {
	var handlerType Middleware[T]
	switch typ := mw.(type) {
	case func():
		handlerType = func(ctx *T, next Handler[T]) (*T, error) {
			typ()
			return next(ctx)
		}
	case func() error:
		handlerType = func(ctx *T, next Handler[T]) (*T, error) {
			err := typ()
			if err != nil {
				return nil, err
			}
			return next(ctx)
		}
	case func(*T):
		handlerType = handlerToMware(func(ctx *T) (*T, error) {
			typ(ctx)
			return ctx, nil
		})
	case func(*T) error:
		handlerType = handlerToMware(func(ctx *T) (*T, error) {
			err := typ(ctx)
			if err != nil {
				return nil, err
			}
			return ctx, nil
		})
	case func(*T) *T:
		handlerType = handlerToMware(func(ctx *T) (*T, error) {
			return typ(ctx), nil
		})
	case func(*T) (*T, error):
		handlerType = handlerToMware(typ)
	case func(*T, Handler[T]) *T:
		handlerType = Middleware[T](func(ctx *T, next Handler[T]) (*T, error) {
			return typ(ctx, next), nil
		})
	case func(*T, Handler[T]) error:
		handlerType = Middleware[T](func(ctx *T, next Handler[T]) (*T, error) {
			err := typ(ctx, next)
			if err != nil {
				return nil, err
			}
			return next(ctx)
		})
	case func(*T, Handler[T]) (*T, error):
		handlerType = Middleware[T](typ)
	case Middleware[T]:
		handlerType = typ
	case Handler[T]:
		handlerType = handlerToMware(typ)

	default:
		return nil, fmt.Errorf("invalid middleware type: %T", mw)
	}

	return handlerType, nil
}

func handlerToMware[T any](h Handler[T]) Middleware[T] {
	return func(ctx *T, next Handler[T]) (*T, error) {
		ctx, err := h(ctx)
		if err == nil {
			return next(ctx)
		}
		return nil, err
	}
}

func (c *chainedMiddleware[T]) invoke(ctx *T) (*T, error) {
	// Chains are left open-ended so middleware can continue to be linked
	// If the chain is incomplete, set a chained dummy handler for safety
	if c.nextFunc == nil {
		c.nextFunc = dummyHandler[T]
	}

	return c.fun(ctx, c.nextFunc)
}

type middlewareChain[T any] struct {
	chain []*chainedMiddleware[T]
}

func (h *middlewareChain[T]) invoke(ctx *T, next Handler[T]) (*T, error) {
	if len(h.chain) == 0 {
		return nil, fmt.Errorf("there are no middleware in this chain")
	}
	// Complete the chain
	h.chain[len(h.chain)-1].nextFunc = next

	return h.chain[0].invoke(ctx)
}

// Compose - Takes a collection of middleware and composes it into a single middleware function
func Compose[T any](funcs ...Middleware[T]) Middleware[T] {
	mwareChain := &middlewareChain[T]{
		chain: []*chainedMiddleware[T]{},
	}

	var nextFunc Handler[T] = nil
	for i := len(funcs) - 1; i >= 0; i = i - 1 {
		if funcs[i] == nil {
			continue
		}

		cm := &chainedMiddleware[T]{
			fun:      funcs[i],
			nextFunc: nextFunc,
		}
		nextFunc = cm.invoke
		mwareChain.chain = append(mwareChain.chain, cm)
	}

	return mwareChain.invoke
}
