package exec

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cmd/log"
)

func Which(shell string, executable string) string {
	output, err := Command(shell, fmt.Sprintf("which %s", executable))
	if err != nil {
		// Assume errors mean it's not installed
		log.Debug("which command errorred out: %v", err)
		return ""
	}
	return output
}
