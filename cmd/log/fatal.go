package log

import (
	"fmt"
	"os"
)

func CheckFatal(err interface{}) {
	if err == nil {
		return
	}

	println()
	fmt.Fprintln(
		os.Stderr,
		"Error:",
		err,
	)
	println()
	os.Exit(1)
}
