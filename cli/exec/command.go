package exec

import (
	"os/exec"
	"strings"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/utils"
)

func Command(ui cli.UI, shell string, command ...string) (string, error) {
	ui.Debug("Execute: %s -c '%s'", shell, strings.Join(command, " "))
	args := []string{"-c"}
	args = append(args, command...)
	out, err := exec.Command(shell, args...).CombinedOutput()
	if err != nil {
		ui.Debug("Command failed: %v", err)
	}
	output := utils.RemoveFinalNewline(string(out))
	ui.Debug("Output: %s", output)
	return output, err
}
