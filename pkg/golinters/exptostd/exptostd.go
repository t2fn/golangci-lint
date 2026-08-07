package exptostd

import (
	"github.com/ldez/exptostd"

	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(exptostd.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
