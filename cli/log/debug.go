package log

import (
	"fmt"
	"os"
)

var (
	IsDebugging bool
)

func Debug(format string, args ...interface{}) {
	if !IsDebugging {
		return
	}

	fmt.Fprintln(
		os.Stderr,
		Dim(fmt.Sprintf(format, args...)),
	)
}
