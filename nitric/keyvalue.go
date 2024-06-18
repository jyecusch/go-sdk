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
	"context"
	"fmt"

	"github.com/nitrictech/go-sdk/api/keyvalue"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
)

type KvStorePermission string

const (
	KvStoreWrite  KvStorePermission = "write"
	KvStoreRead   KvStorePermission = "read"
	KvStoreDelete KvStorePermission = "delete"
)

var KvStoreEverything []KvStorePermission = []KvStorePermission{KvStoreWrite, KvStoreRead, KvStoreDelete}

type KvStore interface {
	Allow(KvStorePermission, ...KvStorePermission) (keyvalue.Store, error)
}

type kvstore struct {
	name    string
	manager Manager
}

func NewKv(name string) *kvstore {
	return &kvstore{
		name:    name,
		manager: defaultManager,
	}
}

// NewQueue registers this queue as a required resource for the calling function/container.
func (k *kvstore) Allow(permission KvStorePermission, permissions ...KvStorePermission) (keyvalue.Store, error) {
	allPerms := append([]KvStorePermission{permission}, permissions...)

	return defaultManager.newKv(k.name, allPerms...)
}

func (m *manager) newKv(name string, permissions ...KvStorePermission) (keyvalue.Store, error) {
	rsc, err := m.resourceServiceClient()
	if err != nil {
		return nil, err
	}

	colRes := &v1.ResourceIdentifier{
		Type: v1.ResourceType_KeyValueStore,
		Name: name,
	}

	dr := &v1.ResourceDeclareRequest{
		Id: colRes,
		Config: &v1.ResourceDeclareRequest_KeyValueStore{
			KeyValueStore: &v1.KeyValueStoreResource{},
		},
	}
	_, err = rsc.Declare(context.Background(), dr)
	if err != nil {
		return nil, err
	}

	actions := []v1.Action{}
	for _, perm := range permissions {
		switch perm {
		case KvStoreRead:
			actions = append(actions, v1.Action_KeyValueStoreRead)
		case KvStoreWrite:
			actions = append(actions, v1.Action_KeyValueStoreWrite)
		case KvStoreDelete:
			actions = append(actions, v1.Action_KeyValueStoreDelete)
		default:
			return nil, fmt.Errorf("KvStorePermission %s unknown", perm)
		}
	}

	_, err = rsc.Declare(context.Background(), functionResourceDeclareRequest(colRes, actions))
	if err != nil {
		return nil, err
	}

	if m.kvstores == nil {
		m.kvstores, err = keyvalue.New()
		if err != nil {
			return nil, err
		}
	}

	return m.kvstores.Store(name), nil
}
