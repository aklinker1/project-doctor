package log

import (
	"fmt"
	"os"
)

func Debug(debugging bool, format string, args ...interface{}) {
	if !debugging {
		return
	}

	fmt.Fprintln(
		os.Stderr,
		Dim(fmt.Sprintf(format, args...)),
	)
}
