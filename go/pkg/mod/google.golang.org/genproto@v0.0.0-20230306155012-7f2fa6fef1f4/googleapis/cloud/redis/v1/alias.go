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

// Package redis aliases all exported identifiers in package
// "cloud.google.com/go/redis/apiv1/redispb".
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package redis

import (
	src "cloud.google.com/go/redis/apiv1/redispb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/redis/apiv1/redispb
const (
	FailoverInstanceRequest_DATA_PROTECTION_MODE_UNSPECIFIED = src.FailoverInstanceRequest_DATA_PROTECTION_MODE_UNSPECIFIED
	FailoverInstanceRequest_FORCE_DATA_LOSS                  = src.FailoverInstanceRequest_FORCE_DATA_LOSS
	FailoverInstanceRequest_LIMITED_DATA_LOSS                = src.FailoverInstanceRequest_LIMITED_DATA_LOSS
	Instance_BASIC                                           = src.Instance_BASIC
	Instance_CONNECT_MODE_UNSPECIFIED                        = src.Instance_CONNECT_MODE_UNSPECIFIED
	Instance_CREATING                                        = src.Instance_CREATING
	Instance_DELETING                                        = src.Instance_DELETING
	Instance_DIRECT_PEERING                                  = src.Instance_DIRECT_PEERING
	Instance_DISABLED                                        = src.Instance_DISABLED
	Instance_FAILING_OVER                                    = src.Instance_FAILING_OVER
	Instance_IMPORTING                                       = src.Instance_IMPORTING
	Instance_MAINTENANCE                                     = src.Instance_MAINTENANCE
	Instance_PRIVATE_SERVICE_ACCESS                          = src.Instance_PRIVATE_SERVICE_ACCESS
	Instance_READY                                           = src.Instance_READY
	Instance_READ_REPLICAS_DISABLED                          = src.Instance_READ_REPLICAS_DISABLED
	Instance_READ_REPLICAS_ENABLED                           = src.Instance_READ_REPLICAS_ENABLED
	Instance_READ_REPLICAS_MODE_UNSPECIFIED                  = src.Instance_READ_REPLICAS_MODE_UNSPECIFIED
	Instance_REPAIRING                                       = src.Instance_REPAIRING
	Instance_SERVER_AUTHENTICATION                           = src.Instance_SERVER_AUTHENTICATION
	Instance_STANDARD_HA                                     = src.Instance_STANDARD_HA
	Instance_STATE_UNSPECIFIED                               = src.Instance_STATE_UNSPECIFIED
	Instance_TIER_UNSPECIFIED                                = src.Instance_TIER_UNSPECIFIED
	Instance_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED             = src.Instance_TRANSIT_ENCRYPTION_MODE_UNSPECIFIED
	Instance_UPDATING                                        = src.Instance_UPDATING
	RescheduleMaintenanceRequest_IMMEDIATE                   = src.RescheduleMaintenanceRequest_IMMEDIATE
	RescheduleMaintenanceRequest_NEXT_AVAILABLE_WINDOW       = src.RescheduleMaintenanceRequest_NEXT_AVAILABLE_WINDOW
	RescheduleMaintenanceRequest_RESCHEDULE_TYPE_UNSPECIFIED = src.RescheduleMaintenanceRequest_RESCHEDULE_TYPE_UNSPECIFIED
	RescheduleMaintenanceRequest_SPECIFIC_TIME               = src.RescheduleMaintenanceRequest_SPECIFIC_TIME
)

// Deprecated: Please use vars in: cloud.google.com/go/redis/apiv1/redispb
var (
	FailoverInstanceRequest_DataProtectionMode_name   = src.FailoverInstanceRequest_DataProtectionMode_name
	FailoverInstanceRequest_DataProtectionMode_value  = src.FailoverInstanceRequest_DataProtectionMode_value
	File_google_cloud_redis_v1_cloud_redis_proto      = src.File_google_cloud_redis_v1_cloud_redis_proto
	Instance_ConnectMode_name                         = src.Instance_ConnectMode_name
	Instance_ConnectMode_value                        = src.Instance_ConnectMode_value
	Instance_ReadReplicasMode_name                    = src.Instance_ReadReplicasMode_name
	Instance_ReadReplicasMode_value                   = src.Instance_ReadReplicasMode_value
	Instance_State_name                               = src.Instance_State_name
	Instance_State_value                              = src.Instance_State_value
	Instance_Tier_name                                = src.Instance_Tier_name
	Instance_Tier_value                               = src.Instance_Tier_value
	Instance_TransitEncryptionMode_name               = src.Instance_TransitEncryptionMode_name
	Instance_TransitEncryptionMode_value              = src.Instance_TransitEncryptionMode_value
	RescheduleMaintenanceRequest_RescheduleType_name  = src.RescheduleMaintenanceRequest_RescheduleType_name
	RescheduleMaintenanceRequest_RescheduleType_value = src.RescheduleMaintenanceRequest_RescheduleType_value
)

// CloudRedisClient is the client API for CloudRedis service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type CloudRedisClient = src.CloudRedisClient

// CloudRedisServer is the server API for CloudRedis service.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type CloudRedisServer = src.CloudRedisServer

// Request for
// [CreateInstance][google.cloud.redis.v1.CloudRedis.CreateInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type CreateInstanceRequest = src.CreateInstanceRequest

// Request for
// [DeleteInstance][google.cloud.redis.v1.CloudRedis.DeleteInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type DeleteInstanceRequest = src.DeleteInstanceRequest

// Request for [Export][google.cloud.redis.v1.CloudRedis.ExportInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type ExportInstanceRequest = src.ExportInstanceRequest

// Request for [Failover][google.cloud.redis.v1.CloudRedis.FailoverInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type FailoverInstanceRequest = src.FailoverInstanceRequest

// Specifies different modes of operation in relation to the data retention.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type FailoverInstanceRequest_DataProtectionMode = src.FailoverInstanceRequest_DataProtectionMode

// The Cloud Storage location for the output content
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type GcsDestination = src.GcsDestination

// The Cloud Storage location for the input content
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type GcsSource = src.GcsSource

// Request for
// [GetInstanceAuthString][google.cloud.redis.v1.CloudRedis.GetInstanceAuthString].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type GetInstanceAuthStringRequest = src.GetInstanceAuthStringRequest

// Request for [GetInstance][google.cloud.redis.v1.CloudRedis.GetInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type GetInstanceRequest = src.GetInstanceRequest

// Request for [Import][google.cloud.redis.v1.CloudRedis.ImportInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type ImportInstanceRequest = src.ImportInstanceRequest

// The input content
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type InputConfig = src.InputConfig
type InputConfig_GcsSource = src.InputConfig_GcsSource

// A Memorystore for Redis instance.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type Instance = src.Instance

// Instance AUTH string details.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type InstanceAuthString = src.InstanceAuthString

// Available connection modes.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type Instance_ConnectMode = src.Instance_ConnectMode

// Read replicas mode.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type Instance_ReadReplicasMode = src.Instance_ReadReplicasMode

// Represents the different states of a Redis instance.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type Instance_State = src.Instance_State

// Available service tiers to choose from
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type Instance_Tier = src.Instance_Tier

// Available TLS modes.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type Instance_TransitEncryptionMode = src.Instance_TransitEncryptionMode

// Request for
// [ListInstances][google.cloud.redis.v1.CloudRedis.ListInstances].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type ListInstancesRequest = src.ListInstancesRequest

// Response for
// [ListInstances][google.cloud.redis.v1.CloudRedis.ListInstances].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type ListInstancesResponse = src.ListInstancesResponse

// This location metadata represents additional configuration options for a
// given location where a Redis instance may be created. All fields are output
// only. It is returned as content of the
// `google.cloud.location.Location.metadata` field.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type LocationMetadata = src.LocationMetadata

// Maintenance policy for an instance.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type MaintenancePolicy = src.MaintenancePolicy

// Upcoming maintenance schedule. If no maintenance is scheduled, fields are
// not populated.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type MaintenanceSchedule = src.MaintenanceSchedule

// Node specific properties.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type NodeInfo = src.NodeInfo

// Represents the v1 metadata of the long-running operation.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type OperationMetadata = src.OperationMetadata

// The output content
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type OutputConfig = src.OutputConfig
type OutputConfig_GcsDestination = src.OutputConfig_GcsDestination

// Request for
// [RescheduleMaintenance][google.cloud.redis.v1.CloudRedis.RescheduleMaintenance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type RescheduleMaintenanceRequest = src.RescheduleMaintenanceRequest

// Reschedule options.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type RescheduleMaintenanceRequest_RescheduleType = src.RescheduleMaintenanceRequest_RescheduleType

// TlsCertificate Resource
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type TlsCertificate = src.TlsCertificate

// UnimplementedCloudRedisServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type UnimplementedCloudRedisServer = src.UnimplementedCloudRedisServer

// Request for
// [UpdateInstance][google.cloud.redis.v1.CloudRedis.UpdateInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type UpdateInstanceRequest = src.UpdateInstanceRequest

// Request for
// [UpgradeInstance][google.cloud.redis.v1.CloudRedis.UpgradeInstance].
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type UpgradeInstanceRequest = src.UpgradeInstanceRequest

// Time window in which disruptive maintenance updates occur. Non-disruptive
// updates can occur inside or outside this window.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type WeeklyMaintenanceWindow = src.WeeklyMaintenanceWindow

// Defines specific information for a particular zone. Currently empty and
// reserved for future use only.
//
// Deprecated: Please use types in: cloud.google.com/go/redis/apiv1/redispb
type ZoneMetadata = src.ZoneMetadata

// Deprecated: Please use funcs in: cloud.google.com/go/redis/apiv1/redispb
func NewCloudRedisClient(cc grpc.ClientConnInterface) CloudRedisClient {
	return src.NewCloudRedisClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/redis/apiv1/redispb
func RegisterCloudRedisServer(s *grpc.Server, srv CloudRedisServer) {
	src.RegisterCloudRedisServer(s, srv)
}
