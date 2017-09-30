package something

import (
	"fmt"
	"runtime"
)

// BEGIN OMIT
// Veryusefulfunc does some very useful things.
func Veryusefulfunc() string {
	return fmt.Sprintf("Looks like I'm running on: %s\n", internalfunc())
}

func internalfunc() string {
	var r string
	switch runtime.GOOS {
	case "darwin", "freebsd", "linux":
		r = "Unix"
	case "windows":
		r = "Windows"
	default:
		r = "woah, didn't expect this"
	}
	return r
}

// END OMIT
