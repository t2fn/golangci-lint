package arangolint

import (
	"go.augendre.info/arangolint/pkg/analyzer"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
