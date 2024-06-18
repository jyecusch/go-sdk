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

package topics

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/nitrictech/go-sdk/api/errors"
	"github.com/nitrictech/go-sdk/api/errors/codes"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/topics/v1"
	"github.com/nitrictech/protoutils"
)

type PublishOption = func(*v1.TopicPublishRequest)

// Topic for pub/sub async messaging.
type Topic interface {
	// Name returns the Topic name.
	Name() string

	// Publish will publish the provided events on the topic.
	Publish(context.Context, map[string]interface{}, ...PublishOption) error
}

type topicImpl struct {
	name        string
	topicClient v1.TopicsClient
}

func (s *topicImpl) Name() string {
	return s.name
}

// WithDelay - Delay event publishing by the given duration
func WithDelay(duration time.Duration) func(*v1.TopicPublishRequest) {
	return func(epr *v1.TopicPublishRequest) {
		epr.Delay = durationpb.New(duration)
	}
}

func (s *topicImpl) Publish(ctx context.Context, message map[string]interface{}, opts ...PublishOption) error {
	// Convert payload to Protobuf Struct
	payloadStruct, err := protoutils.NewStruct(message)
	if err != nil {
		return errors.NewWithCause(codes.InvalidArgument, "Topic.Publish", err)
	}

	event := &v1.TopicPublishRequest{
		TopicName: s.name,
		Message: &v1.TopicMessage{
			Content: &v1.TopicMessage_StructPayload{
				StructPayload: payloadStruct,
			},
		},
	}

	// Apply options to the event payload
	for _, opt := range opts {
		opt(event)
	}

	_, err = s.topicClient.Publish(ctx, event)
	if err != nil {
		return errors.FromGrpcError(err)
	}

	return nil
}
