//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/broken_doclink.yml
//golangcitest:expected_exitcode 0

package broken_doclink_test

import (
	"fmt"
)

// GoodFunc1 demonstrates valid documentation links.
// See [fmt.Println] for output formatting.
func GoodFunc1() { //nolint:unused
	fmt.Println("Hello")
}

// GoodFunc2 shows a link to an exported symbol in the same package.
// Use [GoodStruct.GetIdentifier] to access the identifier field.
type GoodStruct struct {
	identifier string
}

// GetIdentifier returns the unique identifier.
func (g GoodStruct) GetIdentifier() string { //nolint:unused
	return g.identifier
}
