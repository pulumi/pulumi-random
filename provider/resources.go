// Copyright 2016-2018, Pulumi Corporation.
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

package random

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi-random/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-random/shim"
)

// all of the random token components used below.
const (
	randomPkg = "random"
	randomMod = "index"
)

// randomMember manufactures a type token for the random package and the given module and type.
func randomMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(randomPkg + ":" + mod + ":" + mem)
}

// randomType manufactures a type token for the random package and the given module and type.
func randomType(mod string, typ string) tokens.Type {
	return tokens.Type(randomMember(mod, typ))
}

// randomResource manufactures a standard resource token given a module and resource name.  It automatically uses the
// random package and names the file by simply lower casing the resource's first character.
func randomResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return randomType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the random package.
func Provider() tfbridge.ProviderInfo {
	return tfbridge.ProviderInfo{
		P:           shimv2.NewProvider(shim.NewProvider()),
		Name:        "random",
		Description: "A Pulumi package to safely use randomness in Pulumi programs.",
		Keywords:    []string{"pulumi", "random"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-random",
		Resources: map[string]*tfbridge.ResourceInfo{
			"random_id":       {Tok: randomResource(randomMod, "RandomId")},
			"random_password": {Tok: randomResource(randomMod, "RandomPassword")},
			"random_pet":      {Tok: randomResource(randomMod, "RandomPet")},
			"random_shuffle":  {Tok: randomResource(randomMod, "RandomShuffle")},
			"random_string":   {Tok: randomResource(randomMod, "RandomString")},
			"random_integer":  {Tok: randomResource(randomMod, "RandomInteger")},
			"random_uuid":     {Tok: randomResource(randomMod, "RandomUuid")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.15.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.15.0,<3.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", randomPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				randomPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				"random": "Random",
			},
		},
	}
}
