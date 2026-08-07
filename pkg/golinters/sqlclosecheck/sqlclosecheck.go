package sqlclosecheck

import (
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewDeferOnlyAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
