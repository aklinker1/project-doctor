package exec

import (
	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/riywo/loginshell"
)

var defaultShell string

func Shell() string {
	if defaultShell == "" {
		shell, err := loginshell.Shell()
		if err != nil {
			panic(err)
		}
		log.CheckFatal(err)
		defaultShell = shell
	}
	return defaultShell
}
