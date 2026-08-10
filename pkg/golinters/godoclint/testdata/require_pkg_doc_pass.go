//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/require_pkg_doc.yml
//golangcitest:expected_exitcode 0

// Package require_pkg_doc_test has a godoc.
package require_pkg_doc_test

import "fmt"

func RequirePkgDocPass() { fmt.Println("pass") }
