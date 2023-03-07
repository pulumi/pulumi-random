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

package testutil

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/engine"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	pulumi_testing "github.com/pulumi/pulumi/sdk/v3/go/common/testing"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/retry"
)

// Test that upgrading the binary and language SDK is safe and results in no-change preview and update.
//
// providerBinary: name of the binary that may be in PATH, such as "pulumi-resource-random".
//
// Ideally this should be expressed in ProgramTest facilities, but it is currently not possible.
//
// Also, currently this only works for Node Pulumi programs.
func ProviderUpdateTest(t *testing.T, providerBinary string, opts integration.ProgramTestOptions) {
	opts.SkipEmptyPreviewUpdate = true
	opts.SkipExportImport = true

	// The environment may have providerBinary in PATH, but the first update should use a different version.
	// Restrict the PATH so pulumi CLI does not see it.
	opts = opts.With(restrictPath(t, "node", "pulumi", "yarn"))

	// Capture workdir from PrepareProject, as integration.ProgramTester does not expose it.
	var workdir string

	// This will be called only once. Call yarn install. Assume package.json refs a prod dependency like this:
	//
	//    "@pulumi/random": "^4.0.0"
	//
	// Then this should pull the latest prod version of the provider Node SDK, and then Pulumi will auto-install the
	// matching provider binary version.
	opts.PrepareProject = func(info *engine.Projinfo) (err error) {
		workdir = info.Root
		yarn(t, opts, workdir, "install")
		return nil
	}

	pt := integration.ProgramTestManualLifeCycle(t, &opts)
	err := pt.TestLifeCyclePrepare()
	require.NoError(t, err)

	pt.TestFinished = false
	defer pt.TestCleanUp()

	err = pt.TestLifeCycleInitialize()
	require.NoError(t, err)

	defer func() {
		destroyErr := pt.TestLifeCycleDestroy()
		require.NoError(t, destroyErr)
	}()

	// This should preview and update the stack using the stable version of the provider and its SDK.
	err = pt.TestPreviewUpdateAndEdits()
	require.NoError(t, err)

	// Now bring the local binary being tested back in PATH.
	opts = opts.With(restrictPath(t, "node", "pulumi", "yarn", providerBinary))

	// Also bring local Node SDK versions in for testing via yarn link.
	for _, d := range opts.Dependencies {
		yarn(t, opts, workdir, "link", d)
	}

	// Test that local provider and SDK versions generate nop preview/udpate on the stack.
	err = pt.PreviewAndUpdate(workdir, "upgraded",
		false, /* shouldFail */
		true,  /* expectNopPreview */
		true /* expectNopUpdate */)
	require.NoError(t, err)

	pt.TestFinished = true
}

func yarn(t *testing.T, opts integration.ProgramTestOptions, wd string, args ...string) {
	name := "yarn " + strings.Join(args, " ")

	// Yarn will time out if multiple processes are trying to install packages at the same time.
	pulumi_testing.YarnInstallMutex.Lock()
	defer pulumi_testing.YarnInstallMutex.Unlock()
	t.Log("acquired yarn install lock")
	t.Log("released yarn install lock")

	var yarnBin string = opts.YarnBin
	if yarnBin == "" {
		var err error
		yarnBin, err = exec.LookPath("yarn")
		require.NoError(t, err)
		require.NotEmpty(t, yarnBin)
	}

	_, _, err := retry.Until(context.Background(), retry.Acceptor{
		Accept: func(try int, nextRetryTime time.Duration) (bool, interface{}, error) {
			runerr := integration.RunCommand(t, name, append([]string{yarnBin}, args...), wd, &opts)
			if runerr == nil {
				return true, nil, nil
			} else if _, ok := runerr.(*exec.ExitError); ok {
				// yarn failed, let's try again, assuming we haven't failed a few times.
				if try+1 >= 3 {
					return false, nil, fmt.Errorf("%v did not complete after %v tries", name, try+1)
				}

				return false, nil, nil
			}

			// someother error, fail
			return false, nil, runerr
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func restrictPath(t *testing.T, commands ...string) integration.ProgramTestOptions {
	paths := lookPaths(t, commands...)
	pathVar := strings.Join(paths, string(filepath.ListSeparator))
	return integration.ProgramTestOptions{
		Env: []string{"PATH=" + pathVar},
	}
}

func lookPaths(t *testing.T, commands ...string) (ret []string) {
	for _, c := range commands {
		cPath, err := exec.LookPath(c)
		require.NoError(t, err)
		require.NotEmpty(t, cPath)
		ret = append(ret, filepath.Dir(cPath))
	}
	return
}
