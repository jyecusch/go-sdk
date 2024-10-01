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

import (
	"context"
	errorsstd "errors"
	"io"

	"google.golang.org/grpc"

	"github.com/nitrictech/go-sdk/constants"
	"github.com/nitrictech/go-sdk/nitric/errors"
	"github.com/nitrictech/go-sdk/nitric/errors/codes"
	"github.com/nitrictech/go-sdk/nitric/websockets"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/websockets/v1"
)

type websocketWorker struct {
	client              v1.WebsocketHandlerClient
	registrationRequest *v1.RegistrationRequest
	handler             Handler[websockets.Ctx]
}
type websocketWorkerOpts struct {
	RegistrationRequest *v1.RegistrationRequest
	Handler             Handler[websockets.Ctx]
}

// Start implements Worker.
func (w *websocketWorker) Start(ctx context.Context) error {
	initReq := &v1.ClientMessage{
		Content: &v1.ClientMessage_RegistrationRequest{
			RegistrationRequest: w.registrationRequest,
		},
	}

	// Create the request stream and send the initial request
	stream, err := w.client.HandleEvents(ctx)
	if err != nil {
		return err
	}

	err = stream.Send(initReq)
	if err != nil {
		return err
	}
	for {
		var ctx *websockets.Ctx

		resp, err := stream.Recv()

		if errorsstd.Is(err, io.EOF) {
			err = stream.CloseSend()
			if err != nil {
				return err
			}

			return nil
		} else if err == nil && resp.GetRegistrationResponse() != nil {
			// There is no need to respond to the registration response
		} else if err == nil && resp.GetWebsocketEventRequest() != nil {
			ctx = websockets.NewCtx(resp)
			err = w.handler(ctx)
			if err != nil {
				ctx.WithError(err)
			}

			err = stream.Send(ctx.ToClientMessage())
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
}

func newWebsocketWorker(opts *websocketWorkerOpts) *websocketWorker {
	conn, err := grpc.NewClient(constants.NitricAddress(), constants.DefaultOptions()...)
	if err != nil {
		panic(errors.NewWithCause(
			codes.Unavailable,
			"NewWebsocketWorker: Unable to reach WebsocketHandlerClient",
			err,
		))
	}

	client := v1.NewWebsocketHandlerClient(conn)

	return &websocketWorker{
		client:              client,
		registrationRequest: opts.RegistrationRequest,
		handler:             opts.Handler,
	}
}
