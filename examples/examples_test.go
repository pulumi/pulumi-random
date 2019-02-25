// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestExamples(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err, "expected a valid working directory: %v", err) {
		return
	}

	base := integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/random",
		},
	}

	shortTests := []integration.ProgramTestOptions{
		base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "simple")}),
	}

	longTests := []integration.ProgramTestOptions{}

	tests := shortTests
	if !t.Short() {
		tests = append(tests, longTests...)
	}

	for _, ex := range tests {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, &example)
		})
	}
}
