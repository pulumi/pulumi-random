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

	testutils "github.com/pulumi/pulumi-terraform-bridge/testing/x"
)

// This makes sure that importing a random password works. The key method is Read.
//
//	pulumi import random:index/randomPassword:RandomPassword newPassword supersecret
func TestRegress160(t *testing.T) {
	testCase := `
	[
	  {
	    "method": "/pulumirpc.ResourceProvider/Configure",
	    "request": {
	      "args": {
		"version": "4.11.3"
	      },
	      "acceptSecrets": true,
	      "acceptResources": true
	    },
	    "response": {
	      "acceptSecrets": true,
	      "supportsPreview": true,
	      "acceptResources": true,
	      "acceptOutputs": true
	    },
	    "metadata": {
	      "kind": "resource",
	      "mode": "client",
	      "name": "random"
	    }
	  },
	  {
	    "method": "/pulumirpc.ResourceProvider/Read",
	    "request": {
	      "id": "supersecret",
	      "urn": "urn:pulumi:v8::re::random:index/randomPassword:RandomPassword::newPassword",
	      "properties": {}
	    },
	    "response": {
	      "id": "none",
	      "properties": {
		"__meta": "{\"schema_version\":\"3\"}",
		"bcryptHash": {
		  "4dabf18193072939515e22adb298388d": "1b47061264138c4ac30d75fd1eb44270",
		  "value": "*"
		},
		"id": "none",
		"length": 11,
		"lower": true,
		"minLower": 0,
		"minNumeric": 0,
		"minSpecial": 0,
		"minUpper": 0,
		"number": true,
		"numeric": true,
		"result": {
		  "4dabf18193072939515e22adb298388d": "1b47061264138c4ac30d75fd1eb44270",
		  "value": "supersecret"
		},
		"special": true,
		"upper": true
	      },
	      "inputs": {
		"length": 11,
		"lower": true,
		"number": true,
		"numeric": true,
		"special": true,
		"upper": true
	      }
	    }
	  }
	]`
	testutils.ReplaySequence(t, providerServer(t), testCase)
}
