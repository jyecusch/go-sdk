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
	"fmt"
	"strings"

	"github.com/nitrictech/go-sdk/nitric/storage"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
)

type BucketPermission string

type bucket struct {
	name         string
	manager      *manager
	registerChan <-chan RegisterResult
}

type Bucket interface {
	// Allow requests the given permissions to the bucket.
	Allow(BucketPermission, ...BucketPermission) (*storage.BucketClient, error)

	// On registers a handler for a specific event type on the bucket.
	// Valid function signatures for handler are:
	//
	//	func()
	//	func() error
	//	func(*storage.Ctx)
	//	func(*storage.Ctx) error
	//	Handler[storage.Ctx]
	On(eventType storage.EventType, notificationPrefixFilter string, handler interface{})
}

const (
	BucketRead   BucketPermission = "read"
	BucketWrite  BucketPermission = "write"
	BucketDelete BucketPermission = "delete"
)

var BucketEverything []BucketPermission = []BucketPermission{BucketRead, BucketWrite, BucketDelete}

// NewBucket - Create a new Bucket resource
func NewBucket(name string) Bucket {
	bucket := &bucket{
		name:    name,
		manager: defaultManager,
	}

	bucket.registerChan = defaultManager.registerResource(&v1.ResourceDeclareRequest{
		Id: &v1.ResourceIdentifier{
			Type: v1.ResourceType_Bucket,
			Name: name,
		},
		Config: &v1.ResourceDeclareRequest_Bucket{
			Bucket: &v1.BucketResource{},
		},
	})

	return bucket
}

func (b *bucket) Allow(permission BucketPermission, permissions ...BucketPermission) (*storage.BucketClient, error) {
	allPerms := append([]BucketPermission{permission}, permissions...)

	actions := []v1.Action{}
	for _, perm := range allPerms {
		switch perm {
		case BucketRead:
			actions = append(actions, v1.Action_BucketFileGet, v1.Action_BucketFileList)
		case BucketWrite:
			actions = append(actions, v1.Action_BucketFilePut)
		case BucketDelete:
			actions = append(actions, v1.Action_BucketFileDelete)
		default:
			return nil, fmt.Errorf("bucketPermission %s unknown", perm)
		}
	}

	registerResult := <-b.registerChan
	if registerResult.Err != nil {
		return nil, registerResult.Err
	}

	err := b.manager.registerPolicy(registerResult.Identifier, actions...)
	if err != nil {
		return nil, err
	}

	return storage.NewBucketClient(b.name)
}

func (b *bucket) On(eventType storage.EventType, notificationPrefixFilter string, handler interface{}) {
	var blobEventType storagepb.BlobEventType
	switch eventType {
	case storage.WriteNotification:
		blobEventType = storagepb.BlobEventType_Created
	case storage.DeleteNotification:
		blobEventType = storagepb.BlobEventType_Deleted
	}

	registrationRequest := &storagepb.RegistrationRequest{
		BucketName:      b.name,
		BlobEventType:   blobEventType,
		KeyPrefixFilter: notificationPrefixFilter,
	}

	typedHandler, err := interfaceToHandler[storage.Ctx](handler)
	if err != nil {
		panic(err)
	}

	opts := &bucketEventWorkerOpts{
		RegistrationRequest: registrationRequest,
		Handler:             typedHandler,
	}

	worker := newBucketEventWorker(opts)

	b.manager.addWorker("bucketNotification:"+strings.Join([]string{
		b.name, notificationPrefixFilter, string(eventType),
	}, "-"), worker)
}
