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

package faas

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/nitrictech/apis/go/nitric/v1"
	"github.com/nitrictech/go-sdk/api/errors"
	"github.com/nitrictech/go-sdk/api/errors/codes"
	"github.com/nitrictech/go-sdk/constants"
)

type ApiWorkerOptions struct {
	ApiName     string
	Path        string
	HttpMethods []string
}

type Frequency string //= "days" | "hours" | "minutes";

var Frequencies = []Frequency{"days", "hours", "minutes"}

type RateWorkerOptions struct {
	Description string
	Rate        int
	Frequency   Frequency
}

type SubscriptionWorkerOptions struct {
	Topic string
}

type HandlerBuilder interface {
	Http(...HttpMiddleware) HandlerBuilder
	Event(...EventMiddleware) HandlerBuilder
	Default(...TriggerMiddleware) HandlerBuilder
	WithApiWorkerOpts(ApiWorkerOptions) HandlerBuilder
	WithRateWorkerOpts(RateWorkerOptions) HandlerBuilder
	WithSubscriptionWorkerOpts(SubscriptionWorkerOptions) HandlerBuilder
	Start() error
}

type HandlerProvider interface {
	GetHttp() HttpMiddleware
	GetEvent() EventMiddleware
	GetDefault() TriggerMiddleware
}

type faasClientImpl struct {
	http                   HttpMiddleware
	apiWorkerOpts          ApiWorkerOptions
	event                  EventMiddleware
	rateWorkerOpts         RateWorkerOptions
	subscriptionWorkerOpts SubscriptionWorkerOptions
	trig                   TriggerMiddleware
}

func (f *faasClientImpl) Http(mwares ...HttpMiddleware) HandlerBuilder {
	f.http = ComposeHttpMiddlware(mwares...)
	return f
}

func (f *faasClientImpl) GetHttp() HttpMiddleware {
	return f.http
}

func (f *faasClientImpl) Event(mwares ...EventMiddleware) HandlerBuilder {
	f.event = ComposeEventMiddleware(mwares...)
	return f
}

func (f *faasClientImpl) GetEvent() EventMiddleware {
	return f.event
}

func (f *faasClientImpl) Default(mwares ...TriggerMiddleware) HandlerBuilder {
	f.trig = ComposeTriggerMiddleware(mwares...)
	return f
}

func (f *faasClientImpl) GetDefault() TriggerMiddleware {
	return f.trig
}

func (f *faasClientImpl) Start() error {
	// Fail if no handlers were provided
	conn, err := grpc.Dial(
		constants.NitricAddress(),
		constants.DefaultOptions()...,
	)

	if err != nil {
		return errors.NewWithCause(
			codes.Unavailable,
			"faas.Start: Unable to reach FaasServiceServer",
			err,
		)
	}

	fsc := pb.NewFaasServiceClient(conn)

	return f.startWithClient(fsc)
}

func (f *faasClientImpl) startWithClient(fsc pb.FaasServiceClient) error {
	// Fail if no handlers were provided
	if f.http == nil && f.event == nil && f.trig == nil {
		return fmt.Errorf("no valid handlers provided")
	}

	if stream, err := fsc.TriggerStream(context.TODO()); err == nil {
		initRequest := &pb.InitRequest{}

		if len(f.apiWorkerOpts.HttpMethods) > 0 {
			initRequest.Worker = &pb.InitRequest_Api{
				Api: &pb.ApiWorker{
					Api:     f.apiWorkerOpts.ApiName,
					Path:    f.apiWorkerOpts.Path,
					Methods: f.apiWorkerOpts.HttpMethods,
				},
			}
		}
		if f.rateWorkerOpts.Rate > 0 {
			initRequest.Worker = &pb.InitRequest_Schedule{
				Schedule: &pb.ScheduleWorker{
					Key: f.rateWorkerOpts.Description,
					Cadence: &pb.ScheduleWorker_Rate{
						Rate: &pb.ScheduleRate{
							Rate: fmt.Sprintf("%d %s", f.rateWorkerOpts.Rate, string(f.rateWorkerOpts.Frequency)),
						},
					},
				},
			}
		}
		if len(f.subscriptionWorkerOpts.Topic) > 0 {
			initRequest.Worker = &pb.InitRequest_Subscription{
				Subscription: &pb.SubscriptionWorker{
					Topic: f.subscriptionWorkerOpts.Topic,
				},
			}
		}

		// Let the membrane know the function is ready for initialization
		err := stream.Send(&pb.ClientMessage{
			Content: &pb.ClientMessage_InitRequest{
				InitRequest: initRequest,
			},
		})
		if err != nil {
			return err
		}

		errChan := make(chan error)

		// Start faasLoop in a go routine
		go faasLoop(stream, f, errChan)

		return <-errChan
	} else {
		return err
	}
}

// Creates a new HandlerBuilder
func New() HandlerBuilder {
	return &faasClientImpl{}
}

func (f *faasClientImpl) WithApiWorkerOpts(opts ApiWorkerOptions) HandlerBuilder {
	f.apiWorkerOpts = opts
	return f
}

func (f *faasClientImpl) WithRateWorkerOpts(opts RateWorkerOptions) HandlerBuilder {
	f.rateWorkerOpts = opts
	return f
}

func (f *faasClientImpl) WithSubscriptionWorkerOpts(opts SubscriptionWorkerOptions) HandlerBuilder {
	f.subscriptionWorkerOpts = opts
	return f
}
