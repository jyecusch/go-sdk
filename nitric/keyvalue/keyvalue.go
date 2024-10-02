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

package keyvalue

import (
	"fmt"

	"github.com/nitrictech/go-sdk/nitric/workers"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
)

type KvStorePermission string

const (
	KvStoreSet    KvStorePermission = "set"
	KvStoreGet    KvStorePermission = "get"
	KvStoreDelete KvStorePermission = "delete"
)

var KvStoreEverything []KvStorePermission = []KvStorePermission{KvStoreSet, KvStoreGet, KvStoreDelete}

type KvStore interface {
	// Allow requests the given permissions to the key/value store.
	Allow(permission KvStorePermission, permissions ...KvStorePermission) KvStoreClientIface
}

type kvstore struct {
	name         string
	manager      *workers.Manager
	registerChan <-chan workers.RegisterResult
}

// NewKv creates a new key/value store resource with the given name.
func NewKv(name string) *kvstore {
	kvstore := &kvstore{
		name:         name,
		manager:      workers.GetDefaultManager(),
		registerChan: make(chan workers.RegisterResult),
	}

	kvstore.registerChan = kvstore.manager.RegisterResource(&v1.ResourceDeclareRequest{
		Id: &v1.ResourceIdentifier{
			Type: v1.ResourceType_KeyValueStore,
			Name: name,
		},
		Config: &v1.ResourceDeclareRequest_KeyValueStore{
			KeyValueStore: &v1.KeyValueStoreResource{},
		},
	})

	return kvstore
}

func (k *kvstore) Allow(permission KvStorePermission, permissions ...KvStorePermission) KvStoreClientIface {
	allPerms := append([]KvStorePermission{permission}, permissions...)

	actions := []v1.Action{}
	for _, perm := range allPerms {
		switch perm {
		case KvStoreGet:
			actions = append(actions, v1.Action_KeyValueStoreRead)
		case KvStoreSet:
			actions = append(actions, v1.Action_KeyValueStoreWrite)
		case KvStoreDelete:
			actions = append(actions, v1.Action_KeyValueStoreDelete)
		default:
			panic(fmt.Sprintf("KvStorePermission %s unknown", perm))
		}
	}

	registerResult := <-k.registerChan

	if registerResult.Err != nil {
		panic(registerResult.Err)
	}

	err := k.manager.RegisterPolicy(registerResult.Identifier, actions...)
	if err != nil {
		panic(err)
	}

	client, err := NewKvStoreClient(k.name)
	if err != nil {
		panic(err)
	}

	return client
}
