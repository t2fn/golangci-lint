package test

import (
	"path/filepath"
	"testing"

	"github.com/t2fn/golangci-lint/v2/test/testshared"
	"github.com/t2fn/golangci-lint/v2/test/testshared/integration"
)

const testdataDir = "testdata"

func TestSourcesFromTestdata(t *testing.T) {
	integration.RunTestdata(t)
}

func TestTypecheck(t *testing.T) {
	testshared.SkipOnWindows(t)

	integration.RunTestSourcesFromDir(t, filepath.Join(testdataDir, "notcompiles"))
}
