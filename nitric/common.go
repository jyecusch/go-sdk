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
	nitricv1 "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
)

func functionResourceDeclareRequest(subject *nitricv1.ResourceIdentifier, actions []nitricv1.Action) *nitricv1.ResourceDeclareRequest {
	return &nitricv1.ResourceDeclareRequest{
		Id: &nitricv1.ResourceIdentifier{
			Type: nitricv1.ResourceType_Policy,
		},
		Config: &nitricv1.ResourceDeclareRequest_Policy{
			Policy: &nitricv1.PolicyResource{
				Principals: []*nitricv1.ResourceIdentifier{
					{
						Type: nitricv1.ResourceType_Service,
					},
				},
				Actions:   actions,
				Resources: []*nitricv1.ResourceIdentifier{subject},
			},
		},
	}
}

type Details struct {
	// The identifier of the resource
	ID string
	// The provider this resource is deployed with (e.g. aws)
	Provider string
	// The service this resource is deployed on (e.g. ApiGateway)
	Service string
}
