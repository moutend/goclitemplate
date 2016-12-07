package goclitemplate

import (
	"fmt"
)

var revision string = "latest"
var version string = "latest"

func Version() string {
	return fmt.Sprintf("%s-%s", version, revision)
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
