//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/require_pkg_doc.yml

package require_pkg_doc_test // want `package should have a godoc`

import "fmt"

func RequirePkgDocFail() { fmt.Println("fail") }
