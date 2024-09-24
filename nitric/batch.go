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

	"github.com/nitrictech/go-sdk/nitric/batch"

	batchpb "github.com/nitrictech/nitric/core/pkg/proto/batch/v1"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
)

// JobPermission defines the available permissions on a job
type JobPermission string

const (
	// JobSubmit is required to call Submit on a job.
	JobSubmit JobPermission = "submit"
)

type JobReference interface {
	// Allow requests the given permissions to the job.
	Allow(JobPermission, ...JobPermission) (*batch.BatchClient, error)

	// Handler will register and start the job task handler that will be called for all task submitted to this job.
	// Valid function signatures for middleware are:
	//
	//	func()
	//	func() error
	//	func(*batch.Ctx)
	//	func(*batch.Ctx) error
	//	func(*batch.Ctx) *batch.Ctx
	//	func(*batch.Ctx) (*batch.Ctx, error)
	//	func(*batch.Ctx, Handler[batch.Ctx]) *batch.Ctx
	//	func(*batch.Ctx, Handler[batch.Ctx]) error
	//	func(*batch.Ctx, Handler[batch.Ctx]) (*batch.Ctx, error)
	//	Middleware[batch.Ctx]
	//	Handler[batch.Ctx]
	Handler(JobResourceRequirements, ...interface{})
}

type jobReference struct {
	name         string
	manager      *manager
	registerChan <-chan RegisterResult
}

type JobResourceRequirements struct {
	Cpus   float32
	Memory int64
	Gpus   int64
}

// NewJob creates a new Job with the give name.
func NewJob(name string) JobReference {
	job := &jobReference{
		name:    name,
		manager: defaultManager,
	}

	job.registerChan = defaultManager.registerResource(&v1.ResourceDeclareRequest{
		Id: &v1.ResourceIdentifier{
			Type: v1.ResourceType_Job,
			Name: name,
		},
		Config: &v1.ResourceDeclareRequest_Job{
			Job: &v1.JobResource{},
		},
	})

	return job
}

func (j *jobReference) Allow(permission JobPermission, permissions ...JobPermission) (*batch.BatchClient, error) {
	allPerms := append([]JobPermission{permission}, permissions...)

	actions := []v1.Action{}
	for _, perm := range allPerms {
		switch perm {
		case JobSubmit:
			actions = append(actions, v1.Action_JobSubmit)
		default:
			return nil, fmt.Errorf("JobPermission %s unknown", perm)
		}
	}

	registerResult := <-j.registerChan
	if registerResult.Err != nil {
		return nil, registerResult.Err
	}

	err := j.manager.registerPolicy(registerResult.Identifier, actions...)
	if err != nil {
		return nil, err
	}

	return batch.NewBatchClient(j.name)
}

func (j *jobReference) Handler(reqs JobResourceRequirements, middleware ...interface{}) {
	registrationRequest := &batchpb.RegistrationRequest{
		JobName: j.name,
		Requirements: &batchpb.JobResourceRequirements{
			Cpus:   reqs.Cpus,
			Memory: reqs.Memory,
			Gpus:   reqs.Gpus,
		},
	}

	middlewares, err := interfacesToMiddleware[batch.Ctx](middleware)
	if err != nil {
		panic(err)
	}

	composedHandler := ComposeMiddleware(middlewares...)

	opts := &jobWorkerOpts{
		RegistrationRequest: registrationRequest,
		Middleware:          composedHandler,
	}

	worker := newJobWorker(opts)
	j.manager.addWorker("JobWorker:"+j.name, worker)
}
