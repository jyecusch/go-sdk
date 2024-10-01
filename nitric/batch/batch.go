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

package batch

import (
	"fmt"

	"github.com/nitrictech/go-sdk/nitric/handlers"
	"github.com/nitrictech/go-sdk/nitric/workers"
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
	Allow(JobPermission, ...JobPermission) (*BatchClient, error)

	// Handler will register and start the job task handler that will be called for all task submitted to this job.
	// Valid function signatures for middleware are:
	//
	//	func()
	//	func() error
	//	func(*batch.Ctx)
	//	func(*batch.Ctx) error
	//	Handler[batch.Ctx]
	Handler(JobResourceRequirements, interface{})
}

type jobReference struct {
	name         string
	manager      *workers.Manager
	registerChan <-chan workers.RegisterResult
}

// JobResourceRequirements defines the resource requirements for a job
type JobResourceRequirements struct {
	// Cpus is the number of CPUs/vCPUs to allocate to the job
	Cpus float32
	// Memory is the amount of memory in MiB to allocate to the job
	Memory int64
	// Gpus is the number of GPUs to allocate to the job
	Gpus int64
}

// NewJob creates a new job resource with the give name.
func NewJob(name string) JobReference {
	job := &jobReference{
		name:    name,
		manager: workers.GetDefaultManager(),
	}

	job.registerChan = job.manager.RegisterResource(&v1.ResourceDeclareRequest{
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

func (j *jobReference) Allow(permission JobPermission, permissions ...JobPermission) (*BatchClient, error) {
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

	err := j.manager.RegisterPolicy(registerResult.Identifier, actions...)
	if err != nil {
		return nil, err
	}

	return NewBatchClient(j.name)
}

func (j *jobReference) Handler(reqs JobResourceRequirements, handler interface{}) {
	registrationRequest := &batchpb.RegistrationRequest{
		JobName: j.name,
		Requirements: &batchpb.JobResourceRequirements{
			Cpus:   reqs.Cpus,
			Memory: reqs.Memory,
			Gpus:   reqs.Gpus,
		},
	}

	typedHandler, err := handlers.HandlerFromInterface[Ctx](handler)
	if err != nil {
		panic(err)
	}

	opts := &jobWorkerOpts{
		RegistrationRequest: registrationRequest,
		Handler:             typedHandler,
	}

	worker := newJobWorker(opts)
	j.manager.AddWorker("JobWorker:"+j.name, worker)
}
