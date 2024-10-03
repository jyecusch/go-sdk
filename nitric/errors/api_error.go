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

package errors

import (
	"errors"
	"fmt"

	multierror "github.com/missionMeteora/toolkit/errors"
	"google.golang.org/grpc/status"

	"github.com/nitrictech/go-sdk/nitric/errors/codes"
)

type ApiError struct {
	code  codes.Code
	msg   string
	cause error
}

func (a *ApiError) Unwrap() error {
	return a.cause
}

func (a *ApiError) Error() string {
	if a.cause != nil {
		// If the wrapped error is an ApiError than these should unwrap
		return fmt.Sprintf("%s: %s: \n %s", a.code.String(), a.msg, a.cause.Error())
	}

	return fmt.Sprintf("%s: %s", a.code.String(), a.msg)
}

// FromGrpcError - translates a standard grpc error to a nitric api error
func FromGrpcError(err error) error {
	if s, ok := status.FromError(err); ok {
		errList := &multierror.ErrorList{}
		errList.Push(err)
		for _, item := range s.Details() {
			errList.Push(fmt.Errorf("%v", item))
		}

		return &ApiError{
			code:  codes.Code(s.Code()),
			msg:   s.Message(),
			cause: errList.Err(),
		}
	}

	return &ApiError{
		code:  codes.Unknown,
		msg:   "error from grpc library",
		cause: err,
	}
}

// Code - returns a nitric api error code from an error or Unknown if the error was not a nitric api error
func Code(err error) codes.Code {
	var apiErr *ApiError
	if ok := errors.As(err, &apiErr); ok {
		return apiErr.code
	}

	return codes.Unknown
}

// New - Creates a new nitric API error
func New(c codes.Code, msg string) error {
	return &ApiError{
		code: c,
		msg:  msg,
	}
}

// NewWithCause - Creates a new nitric API error with the given error as it's cause
func NewWithCause(c codes.Code, msg string, cause error) error {
	return &ApiError{
		code:  c,
		msg:   msg,
		cause: cause,
	}
}
