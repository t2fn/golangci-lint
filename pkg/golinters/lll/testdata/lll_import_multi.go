//golangcitest:args -Elll
//golangcitest:config_path testdata/lll_import.yml
//golangcitest:expected_exitcode 0
package testdata

import (
	anotherVeryLongImportAliasNameForTest "github.com/t2fn/golangci-lint/v2/internal/golinters"
	veryLongImportAliasNameForTest "github.com/t2fn/golangci-lint/v2/internal/golinters"
)

func LllMultiImport() {
	_ = veryLongImportAliasNameForTest.NewLLL(nil)
	_ = anotherVeryLongImportAliasNameForTest.NewLLL(nil)
}
