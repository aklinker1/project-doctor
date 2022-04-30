package exec

import (
	"github.com/riywo/loginshell"
)

var defaultShell string

func Shell() (string, error) {
	if defaultShell == "" {
		shell, err := loginshell.Shell()
		if err != nil {
			return "", err
		}
		defaultShell = shell
	}
	return defaultShell, nil
}
