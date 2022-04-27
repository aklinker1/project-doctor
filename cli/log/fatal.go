package log

import (
	"fmt"
	"os"

	"github.com/aklinker1/project-doctor/cli"
)

func CheckFatal(err error) {
	if err == nil {
		return
	}

	println()
	fmt.Fprintf(os.Stderr, Error("%v\n"), err)
	println()
	os.Exit(cli.ExitCode(err))
}
