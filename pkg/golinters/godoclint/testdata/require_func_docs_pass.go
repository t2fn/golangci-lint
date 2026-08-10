//golangcitest:args -Egodoclint
//golangcitest:config_path testdata/require_func_docs.yml
//golangcitest:expected_exitcode 0

package require_func_docs_test

import "fmt"

// GoodFunc1 adds two numbers together.
// Parameters:
//   - a: the first number
//   - b: the second number
// Returns:
//   - The sum of a and b.
func GoodFunc1(a, b int) int { //nolint:unused
	return a + b
}

// GoodFunc2 formats an integer as a string.
// Parameters:
//   - x: the input value to format
// Returns:
//   - A formatted string representation of x.
func GoodFunc2(x int) string { //nolint:unused
	return fmt.Sprintf("%d", x)
}

// GoodFunc3 creates a greeting message.
// Parameters:
//   - name: the person to greet
// Returns:
//   - A greeting message for the specified person.
func GoodFunc3(name string) string { //nolint:unused
	return fmt.Sprintf("Hello, %s!", name)
}
