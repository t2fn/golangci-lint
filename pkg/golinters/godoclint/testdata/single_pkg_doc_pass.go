//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/single_pkg_doc.yml
//golangcitest:expected_exitcode 0

package single_pkg_doc_test

import "fmt"

// Package single_pkg_doc_test has a godoc.
func SinglePkgDocPass() { fmt.Println("pass") }
