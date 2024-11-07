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
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	provider "github.com/pulumi/pulumi-random/provider/v4"
	"github.com/pulumi/pulumi-random/provider/v4/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	testutils "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/x/testing"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestRegress279(t *testing.T) {
	// Ensure that the correct unknown value sentinel is used for id and result in preview Create.
	create := `
	{
	  "method": "/pulumirpc.ResourceProvider/Create",
	  "request": {
	    "urn": "urn:pulumi:dev::repro::random:index/randomInteger:RandomInteger::k",
	    "properties": {
	      "max": 10,
	      "min": 0
	    },
	    "preview": true
	  },
	  "response": {
	    "properties": {
	      "id": "04da6b54-80e4-46f7-96ec-b56ff0331ba9",
	      "max": 10,
	      "min": 0,
	      "result": "04da6b54-80e4-46f7-96ec-b56ff0331ba9"
	    }
	  }
	}
	`
	testutils.Replay(t, providerServer(t), create)
}

func providerServer(t *testing.T) pulumirpc.ResourceProviderServer {
	ctx := context.Background()
	readFile := func(f string) []byte {
		c, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
			return nil
		}
		return c
	}

	meta := tfbridge.ProviderMetadata{
		PackageSchema:  readFile("../provider/cmd/pulumi-resource-random/schema.json"),
		BridgeMetadata: readFile("../provider/cmd/pulumi-resource-random/bridge-metadata.json"),
	}

	version.Version = "0.0.1"
	info := provider.Provider()

	server, err := tfbridge.NewProvider(ctx, info, meta)
	require.NoError(t, err)

	providerServer := plugin.NewProviderServer(server)
	return providerServer
}
