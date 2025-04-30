// Copyright 2016-2023, Pulumi Corporation.
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

package examples

import (
	"testing"

	testutils "github.com/pulumi/providertest/replay"
)

func TestRegress393(t *testing.T) {
	// Ensure that no error is thrown for RandomPassword with length >72
	// Test values extracted from the grpc.json file produced by regress_393/run.sh
	createPreview := `
	{
		"method": "/pulumirpc.ResourceProvider/Create",
		"request": {
			"urn": "urn:pulumi:test::test::random:index/randomPassword:RandomPassword::username",
			"properties": {
				"length": 256
			},
			"preview": true
		},
		"response": {
			"properties": {
				"*": "*",
				"bcryptHash": "04da6b54-80e4-46f7-96ec-b56ff0331ba9",
				"id": "04da6b54-80e4-46f7-96ec-b56ff0331ba9",
				"length": 256,
				"lower": true,
				"minLower": 0,
				"minNumeric": 0,
				"minSpecial": 0,
				"minUpper": 0,
				"number": true,
				"numeric": true,
				"result": "04da6b54-80e4-46f7-96ec-b56ff0331ba9",
				"special": true,
				"upper": true
			}
		},
		"metadata": {
			"kind": "resource",
			"mode": "client",
			"name": "random"
		}
	}
	`
	testutils.Replay(t, providerServer(t), createPreview)

	create := `
	{
		"method": "/pulumirpc.ResourceProvider/Create",
		"request": {
			"urn": "urn:pulumi:test::test::random:index/randomPassword:RandomPassword::username",
			"properties": {
				"length": 256
			}
		},
		"response": {
			"id": "none",
			"properties": {
				"*": "*",
				"__meta": "{\"schema_version\":\"3\"}",
				"bcryptHash": "*",
				"id": "none",
				"length": 256,
				"lower": true,
				"minLower": 0,
				"minNumeric": 0,
				"minSpecial": 0,
				"minUpper": 0,
				"number": true,
				"numeric": true,
				"result": "*",
				"special": true,
				"upper": true
			}
		},
		"metadata": {
			"kind": "resource",
			"mode": "client",
			"name": "random"
		}
	}
	`
	testutils.Replay(t, providerServer(t), create)
}
