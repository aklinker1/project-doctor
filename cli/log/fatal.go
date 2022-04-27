package log

import (
	"fmt"
	"os"

	"github.com/aklinker1/project-doctor/cli/errors"
)

func CheckFatal(err error) {
	if err == nil {
		return
	}

	println()
	fmt.Fprintf(os.Stderr, Error("%v\n"), err)
	println()
	os.Exit(errors.ExitCode(err))
}
