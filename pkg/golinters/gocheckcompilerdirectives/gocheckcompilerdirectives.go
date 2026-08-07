package gocheckcompilerdirectives

import (
	"4d63.com/gocheckcompilerdirectives/checkcompilerdirectives"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(checkcompilerdirectives.Analyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
