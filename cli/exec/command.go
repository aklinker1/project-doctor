package exec

import (
	"os/exec"
	"strings"

	"github.com/aklinker1/project-doctor/cli/log"
	"github.com/aklinker1/project-doctor/cli/utils"
)

func Command(shell string, command ...string) (string, error) {
	log.Debug("Execute: %s -c '%s'", shell, strings.Join(command, " "))
	args := []string{"-c"}
	args = append(args, command...)
	out, err := exec.Command(shell, args...).CombinedOutput()
	if err != nil {
		log.Debug("Command failed: %v", err)
	}
	output := utils.RemoveFinalNewline(string(out))
	log.Debug("Output: %s", output)
	return output, err
}
