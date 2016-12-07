package goclitemplate

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, Haskell!"
	actual := Hello("Haskell")
	if expected != actual {
		t.Errorf("\nexpected: %s\nactual: %s\n", expected, actual)
	}
	return
}
