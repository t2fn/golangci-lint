package gofmt

import (
	"github.com/t2fn/golangci-lint/v2/pkg/config"
	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
	"github.com/t2fn/golangci-lint/v2/pkg/goformatters"
	gofmtbase "github.com/t2fn/golangci-lint/v2/pkg/goformatters/gofmt"
	"github.com/t2fn/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.GoFmtSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(gofmtbase.Name),
				"Check if the code is formatted according to 'gofmt' command.",
				gofmtbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
