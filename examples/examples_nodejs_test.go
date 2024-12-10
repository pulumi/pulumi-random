//go:build nodejs || all
// +build nodejs all

// Copyright 2016-2023, Pulumi Corporation.  All rights reserved.
package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"

	"github.com/pulumi/pulumi-random/examples/v4/internal/testutil"
)

func TestSimpleTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestRandomBytes(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "bytes"),
		})

	integration.ProgramTest(t, &test)
}

func TestProviderUpdate(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple", "ts"),
		})

	testutil.ProviderUpdateTest(t, "pulumi-resource-random", test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	return getBaseOptions(t).With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/random",
		},
	})
}
