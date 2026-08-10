//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/require_func_docs.yml

package require_func_docs_test

import "fmt"

// BadFunc1 has no documentation at all. // want `require-func-docs: BadFunc1: undocumented parameter\(s\): \"a\", \"b\"`
func BadFunc1(a, b int) int { //nolint:unused
	return a + b
}

// BadFunc2 has named parameters but no documentation at all. // want `require-func-docs: BadFunc2: undocumented parameter\(s\): \"x\"`
func BadFunc2(x int) (msg string) { //nolint:unused
	return fmt.Sprintf("%d", x)
}

// BadFunc3 has named parameters but no documentation at all. // want `require-func-docs: BadFunc3: undocumented parameter\(s\): \"name\"`
func BadFunc3(name string) (greeting string) { //nolint:unused
	return fmt.Sprintf("Hello, %s!", name)
}
