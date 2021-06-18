package main

import (
	"os"
	"testing"

	ex "github.com/srjnm/demoinfocs-golang/examples"
)

// Just make sure the example runs
func TestBouncyNades(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}

	os.Args = []string{"cmd", "-demo", "../../test/cs-demos/default.dem"}

	ex.RedirectStdout(func() {
		main()
	})
}
