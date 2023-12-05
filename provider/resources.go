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

package random

import (
	"fmt"
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"path/filepath"
	"unicode"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-random/provider/v4/pkg/version"
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

//go:embed cmd/pulumi-resource-random/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the random package.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:            pf.ShimProvider(shim.NewProvider()),
		Name:         "random",
		Description:  "A Pulumi package to safely use randomness in Pulumi programs.",
		Keywords:     []string{"pulumi", "random"},
		License:      "Apache-2.0",
		Homepage:     "https://pulumi.io",
		Repository:   "https://github.com/pulumi/pulumi-random",
		Version:      version.Version,
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"random_id": {Tok: randomResource(randomMod, "RandomId")},

			"random_password": {
				Tok:  randomResource(randomMod, "RandomPassword"),
				Docs: &tfbridge.DocInfo{ImportDetails: string(docPasswordImport)},
			},

			"random_pet":     {Tok: randomResource(randomMod, "RandomPet")},
			"random_shuffle": {Tok: randomResource(randomMod, "RandomShuffle")},
			"random_integer": {Tok: randomResource(randomMod, "RandomInteger")},
			"random_uuid":    {Tok: randomResource(randomMod, "RandomUuid")},

			"random_string": {
				Tok: randomResource(randomMod, "RandomString"),
				PreStateUpgradeHook: func(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
					// States for RandomString may be contaminated by
					// https://github.com/pulumi/pulumi-random/issues/258 bug where the state is
					// missing the version marker. Pretend that these states are at V1, which is the
					// best guess. V1->V2/V3 migrations seem idempotent, this is probably safe.
					if args.PriorStateSchemaVersion == 0 {
						return 1, args.PriorState, nil
					}
					return args.PriorStateSchemaVersion, args.PriorState, nil
				},
				Docs: &tfbridge.DocInfo{ImportDetails: string(docStringImport)},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				Requires: map[string]string{
					"pulumi": ">=3.0.0,<4.0.0",
				}}
			i.PyProject.Enabled = true
			return i
		})(),

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
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"random": "Random",
			},
		},
	}

	makeToken := func(module, name string) (string, error) {
		return tks.MakeStandard(randomPkg)(module, "Random"+name)
	}

	prov.MustComputeTokens(tks.SingleModule("random_", randomMod, makeToken))
	prov.MustApplyAutoAliases()

	return prov
}
