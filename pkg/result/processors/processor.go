package processors

import (
	"github.com/t2fn/golangci-lint/v2/pkg/result"
)

const typeCheckName = "typecheck"

type Processor interface {
	Process(issues []*result.Issue) ([]*result.Issue, error)
	Name() string
	Finish()
}
