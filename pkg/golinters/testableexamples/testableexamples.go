package testableexamples

import (
	"github.com/maratori/testableexamples/pkg/testableexamples"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(testableexamples.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
