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

package events_examples

// [START import]
import (
	"fmt"

	"github.com/nitrictech/go-sdk/api/events"
)

// [END import]

func publishEvent() {
	// [START snippet]
	// Create a new storage client
	ec, err := events.New()

	if err != nil {
		// handle client creation error...
	}

	pEvt, err := ec.Topic("my-topic").Publish(&events.Event{
		Payload: map[string]interface{}{
			"test": "event",
		},
		PayloadType: "my-payload",
	})

	if err != nil {
		// handle topic publish error
	}

	fmt.Println("Published event: ", pEvt)
	// [END snippet]
}
