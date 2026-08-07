package containedctx

import (
	"github.com/sivchari/containedctx"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(containedctx.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
