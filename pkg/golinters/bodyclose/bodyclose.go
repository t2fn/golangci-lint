package bodyclose

import (
	"github.com/timakin/bodyclose/passes/bodyclose"

	"github.com/t2fn/golangci-lint/v2/pkg/config"
	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.BodyCloseSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"check-consumption": settings.CheckConsumption,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(bodyclose.Analyzer).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
