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
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-random/random"
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
		P:           random.Provider().(*schema.Provider),
		Name:        "random",
		Description: "A Pulumi package to safely use randomness in Pulumi programs.",
		Keywords:    []string{"pulumi", "random"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-random",
		Resources: map[string]*tfbridge.ResourceInfo{
			"random_id":      {Tok: randomResource(randomMod, "RandomId")},
			"random_pet":     {Tok: randomResource(randomMod, "RandomPet")},
			"random_shuffle": {Tok: randomResource(randomMod, "RandomShuffle")},
			"random_string":  {Tok: randomResource(randomMod, "RandomString")},
			"random_integer": {Tok: randomResource(randomMod, "RandomInteger")},
			"random_uuid":    {Tok: randomResource(randomMod, "RandomUuid")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^0.17.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.17.0,<0.18.0",
			},
		},
	}
}
