package goprintffuncname

import (
	"github.com/golangci/go-printf-func-name/pkg/analyzer"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.Analyzer).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
