package goimports

import (
	"github.com/t2fn/golangci-lint/v2/pkg/config"
	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
	"github.com/t2fn/golangci-lint/v2/pkg/goformatters"
	goimportsbase "github.com/t2fn/golangci-lint/v2/pkg/goformatters/goimports"
	"github.com/t2fn/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.GoImportsSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(goimportsbase.Name),
				"Checks if the code and import statements are formatted according to the 'goimports' command.",
				goimportsbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
