// Copyright 2021 Confluent Inc. All Rights Reserved.
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

/*
 * REST Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Contact: kafka-clients-proxy-team@confluent.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v3

// AclOperation the model 'AclOperation'
type AclOperation string

// List of AclOperation
const (
	ACLOPERATION_UNKNOWN          AclOperation = "UNKNOWN"
	ACLOPERATION_ANY              AclOperation = "ANY"
	ACLOPERATION_ALL              AclOperation = "ALL"
	ACLOPERATION_READ             AclOperation = "READ"
	ACLOPERATION_WRITE            AclOperation = "WRITE"
	ACLOPERATION_CREATE           AclOperation = "CREATE"
	ACLOPERATION_DELETE           AclOperation = "DELETE"
	ACLOPERATION_ALTER            AclOperation = "ALTER"
	ACLOPERATION_DESCRIBE         AclOperation = "DESCRIBE"
	ACLOPERATION_CLUSTER_ACTION   AclOperation = "CLUSTER_ACTION"
	ACLOPERATION_DESCRIBE_CONFIGS AclOperation = "DESCRIBE_CONFIGS"
	ACLOPERATION_ALTER_CONFIGS    AclOperation = "ALTER_CONFIGS"
	ACLOPERATION_IDEMPOTENT_WRITE AclOperation = "IDEMPOTENT_WRITE"
)