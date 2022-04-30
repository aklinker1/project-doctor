package exec

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cli"
)

func Which(ui cli.UI, shell string, executable string) (string, error) {
	output, err := Command(ui, shell, fmt.Sprintf("which %s", executable))
	if err != nil {
		// Assume errors mean it's not installed
		ui.Debug("which command errorred out: %v", err)
		return "", err
	}
	return output, nil
}
