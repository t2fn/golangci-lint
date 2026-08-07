package recvcheck

import (
	"github.com/raeperd/recvcheck"

	"github.com/t2fn/golangci-lint/v2/pkg/config"
	"github.com/t2fn/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.RecvcheckSettings) *goanalysis.Linter {
	var cfg recvcheck.Settings

	if settings != nil {
		cfg.DisableBuiltin = settings.DisableBuiltin
		cfg.Exclusions = settings.Exclusions
	}

	return goanalysis.
		NewLinterFromAnalyzer(recvcheck.NewAnalyzer(cfg)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
