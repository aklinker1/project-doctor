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
	fmt.Fprintf(os.Stderr, Error("%v\n"), err)
	println()
	os.Exit(1)
}
