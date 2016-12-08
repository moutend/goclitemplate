// Author: Yoshiyuki Koyanagi <moutend@gmail.com>
// License: mIT

// Package goclitemplate provides useless functions.
package goclitemplate

import (
	"fmt"
)

var revision = "latest"
var version = "latest"

// Version returns version and revision as string.
func Version() string {
	return fmt.Sprintf("%s-%s", version, revision)
}

// Hello returns a greeting message like "Hello, James!".
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
