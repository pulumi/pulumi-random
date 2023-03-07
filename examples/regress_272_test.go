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

func TestRegress272(t *testing.T) {
	// States for RandomString may be contaminated by https://github.com/pulumi/pulumi-random/issues/258 bug where
	// the state is missing the __meta version marker. The test here ensures the provider handles such state without
	// failing.
	diff := `
	{
	  "method": "/pulumirpc.ResourceProvider/Diff",
	  "request": {
	    "id": "RbLI*Y#u>m@*Xz[s01s[oV89Pi8]xTRQ",
	    "urn": "urn:pulumi:repro2::simple-random::random:index/randomString:RandomString::string",
	    "olds": {
	      "id": "RbLI*Y#u>m@*Xz[s01s[oV89Pi8]xTRQ",
	      "length": 32,
	      "lower": true,
	      "minLower": 0,
	      "minNumeric": 0,
	      "minSpecial": 0,
	      "minUpper": 0,
	      "number": true,
	      "numeric": true,
	      "result": "RbLI*Y#u>m@*Xz[s01s[oV89Pi8]xTRQ",
	      "special": true,
	      "upper": true
	    },
	    "news": {
	      "length": 32
	    }
	  },
	  "response": {
             "changes": "DIFF_NONE"
          }
	}
	`
	testutils.Replay(t, providerServer(t), diff)

	// But editing length to 6 in the source program makes a replace plan.
	replaceDiff := `
	{
	  "method": "/pulumirpc.ResourceProvider/Diff",
	  "request": {
	    "id": "RbLI*Y#u>m@*Xz[s01s[oV89Pi8]xTRQ",
	    "urn": "urn:pulumi:repro2::simple-random::random:index/randomString:RandomString::string",
	    "olds": {
	      "id": "RbLI*Y#u>m@*Xz[s01s[oV89Pi8]xTRQ",
	      "length": 32,
	      "lower": true,
	      "minLower": 0,
	      "minNumeric": 0,
	      "minSpecial": 0,
	      "minUpper": 0,
	      "number": true,
	      "numeric": true,
	      "result": "RbLI*Y#u>m@*Xz[s01s[oV89Pi8]xTRQ",
	      "special": true,
	      "upper": true
	    },
	    "news": {
	      "length": 6
	    }
	  },
	  "response": {
	    "replaces": [
	      "length"
	    ],
	    "changes": "DIFF_SOME",
	    "diffs": [
	      "length"
	    ]
	  }
	}`
	testutils.Replay(t, providerServer(t), replaceDiff)

	// And subsequently the resource is created, this time populating __meta.
	create := `
	{
	  "method": "/pulumirpc.ResourceProvider/Create",
	  "request": {
	    "urn": "urn:pulumi:repro2::simple-random::random:index/randomString:RandomString::string",
	    "properties": {
	      "length": 6
	    }
	  },
	  "response": {
	    "id": "*",
	    "properties": {
	      "__meta": "*",
	      "id": "*",
	      "length": 6,
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
	  }
	}
        `
	testutils.Replay(t, providerServer(t), create)
}
