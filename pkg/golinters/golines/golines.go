package golines

import (
	"github.com/t2fn/golangci-lint/v2/pkg/config"
	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
	"github.com/t2fn/golangci-lint/v2/pkg/goformatters"
	golinesbase "github.com/t2fn/golangci-lint/v2/pkg/goformatters/golines"
	"github.com/t2fn/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.GoLinesSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(golinesbase.Name),
				"Checks if code is formatted, and fixes long lines",
				golinesbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
