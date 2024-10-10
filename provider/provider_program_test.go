//go:build !go && !nodejs && !python && !dotnet
// +build !go,!nodejs,!python,!dotnet

package random

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/optproviderupgrade"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/opttest"
)

const providerName = "random"
const defaultBaselineVersion = "4.16.0"

var programs = []string{
	"test-programs/index_randomstring",
	"test-programs/index_randomshuffle",
	"test-programs/index_randombytes",
	"test-programs/index_randomuuid",
	"test-programs/index_randompassword",
	"test-programs/index_randomid",
	"test-programs/index_randominteger",
	"test-programs/index_randompet",
}

func TestUpgradeCoverage(t *testing.T) {
	providertest.ReportUpgradeCoverage(t)
}

func testProviderUpgrade(t *testing.T, dir string) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without credentials")
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	test := pulumitest.NewPulumiTest(t, dir,
		opttest.DownloadProviderVersion(providerName, defaultBaselineVersion),
		opttest.LocalProviderPath(providerName, filepath.Join(cwd, "..", "bin")),
	)
	result := providertest.PreviewProviderUpgrade(t, test, providerName, defaultBaselineVersion,
		optproviderupgrade.DisableAttach())
	assertpreview.HasNoReplacements(t, result)
}

func testProgram(t *testing.T, dir string) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without credentials")
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	test := pulumitest.NewPulumiTest(t, dir,
		opttest.LocalProviderPath(providerName, filepath.Join(cwd, "..", "bin")),
		opttest.SkipInstall(),
	)
	test.Up(t)
}

func TestPrograms(t *testing.T) {
	for _, p := range programs {
		t.Run(p, func(t *testing.T) {
			testProgram(t, p)
		})
	}
}

func TestProgramsUpgrade(t *testing.T) {
	for _, p := range programs {
		t.Run(p, func(t *testing.T) {
			testProviderUpgrade(t, p)
		})
	}
}
