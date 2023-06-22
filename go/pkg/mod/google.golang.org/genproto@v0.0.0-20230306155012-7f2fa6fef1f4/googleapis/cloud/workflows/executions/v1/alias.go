// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package executions aliases all exported identifiers in package
// "cloud.google.com/go/workflows/executions/apiv1/executionspb".
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package executions

import (
	src "cloud.google.com/go/workflows/executions/apiv1/executionspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/workflows/executions/apiv1/executionspb
const (
	ExecutionView_BASIC                      = src.ExecutionView_BASIC
	ExecutionView_EXECUTION_VIEW_UNSPECIFIED = src.ExecutionView_EXECUTION_VIEW_UNSPECIFIED
	ExecutionView_FULL                       = src.ExecutionView_FULL
	Execution_ACTIVE                         = src.Execution_ACTIVE
	Execution_CALL_LOG_LEVEL_UNSPECIFIED     = src.Execution_CALL_LOG_LEVEL_UNSPECIFIED
	Execution_CANCELLED                      = src.Execution_CANCELLED
	Execution_FAILED                         = src.Execution_FAILED
	Execution_LOG_ALL_CALLS                  = src.Execution_LOG_ALL_CALLS
	Execution_LOG_ERRORS_ONLY                = src.Execution_LOG_ERRORS_ONLY
	Execution_STATE_UNSPECIFIED              = src.Execution_STATE_UNSPECIFIED
	Execution_SUCCEEDED                      = src.Execution_SUCCEEDED
)

// Deprecated: Please use vars in: cloud.google.com/go/workflows/executions/apiv1/executionspb
var (
	ExecutionView_name                                         = src.ExecutionView_name
	ExecutionView_value                                        = src.ExecutionView_value
	Execution_CallLogLevel_name                                = src.Execution_CallLogLevel_name
	Execution_CallLogLevel_value                               = src.Execution_CallLogLevel_value
	Execution_State_name                                       = src.Execution_State_name
	Execution_State_value                                      = src.Execution_State_value
	File_google_cloud_workflows_executions_v1_executions_proto = src.File_google_cloud_workflows_executions_v1_executions_proto
)

// Request for the
// [CancelExecution][google.cloud.workflows.executions.v1.Executions.CancelExecution]
// method.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type CancelExecutionRequest = src.CancelExecutionRequest

// Request for the
// [CreateExecution][google.cloud.workflows.executions.v1.Executions.CreateExecution]
// method.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type CreateExecutionRequest = src.CreateExecutionRequest

// A running instance of a
// [Workflow](/workflows/docs/reference/rest/v1/projects.locations.workflows).
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution = src.Execution

// Defines possible views for execution resource.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type ExecutionView = src.ExecutionView

// Describes the level of platform logging to apply to calls and call
// responses during workflow executions.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution_CallLogLevel = src.Execution_CallLogLevel

// Error describes why the execution was abnormally terminated.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution_Error = src.Execution_Error

// A collection of stack elements (frames) where an error occurred.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution_StackTrace = src.Execution_StackTrace

// A single stack element (frame) where an error occurred.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution_StackTraceElement = src.Execution_StackTraceElement

// Position contains source position information about the stack trace element
// such as line number, column number and length of the code block in bytes.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution_StackTraceElement_Position = src.Execution_StackTraceElement_Position

// Describes the current state of the execution. More states might be added in
// the future.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type Execution_State = src.Execution_State

// ExecutionsClient is the client API for Executions service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type ExecutionsClient = src.ExecutionsClient

// ExecutionsServer is the server API for Executions service.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type ExecutionsServer = src.ExecutionsServer

// Request for the
// [GetExecution][google.cloud.workflows.executions.v1.Executions.GetExecution]
// method.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type GetExecutionRequest = src.GetExecutionRequest

// Request for the [ListExecutions][] method.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type ListExecutionsRequest = src.ListExecutionsRequest

// Response for the
// [ListExecutions][google.cloud.workflows.executions.v1.Executions.ListExecutions]
// method.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type ListExecutionsResponse = src.ListExecutionsResponse

// UnimplementedExecutionsServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/workflows/executions/apiv1/executionspb
type UnimplementedExecutionsServer = src.UnimplementedExecutionsServer

// Deprecated: Please use funcs in: cloud.google.com/go/workflows/executions/apiv1/executionspb
func NewExecutionsClient(cc grpc.ClientConnInterface) ExecutionsClient {
	return src.NewExecutionsClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/workflows/executions/apiv1/executionspb
func RegisterExecutionsServer(s *grpc.Server, srv ExecutionsServer) {
	src.RegisterExecutionsServer(s, srv)
}
