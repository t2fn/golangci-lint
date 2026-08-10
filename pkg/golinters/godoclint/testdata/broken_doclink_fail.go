//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/broken_doclink.yml

package broken_doclink_test

import (
	"fmt"
)

// BadFunc1 references a non-existent symbol [UnknownPkg.UnknownType]. // want `broken-doclink`
func BadFunc1() { _ = fmt.Println }
