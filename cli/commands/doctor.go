package commands

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/config"
	"github.com/spf13/cobra"
)

type DoctorCommand struct {
	Executor cli.ChecksValidator
	UI       cli.UI
	Project  cli.Project
}

func (c *DoctorCommand) Run(cmd *cobra.Command, args []string) {
	c.UI.PrintTitle("Project Doctor: %s", config.Dirname())
	c.UI.PrintSubtitle("Checking up on your local development environment")

	c.checksSection()
	c.commandsSection()
}

func (c *DoctorCommand) checksSection() error {
	checks := c.Project.Checks
	if len(checks) == 0 {
		return nil
	}

	c.UI.EmptyLine()
	c.UI.PrintSection("Checks")
	c.UI.EmptyLine()
	errs := c.Executor.Validate(checks)
	if len(errs) > 0 {
		return &cli.Error{
			ExitCode: cli.EINTERNAL,
			Message:  fmt.Sprintf("Some checks failed (%d total)", len(errs)),
			Op:       "commands.DoctorCommand.checksSection",
		}
	}
	return nil
}

func (c *DoctorCommand) commandsSection() {
	commands := c.Project.Commands
	if len(commands) == 0 {
		return
	}

	c.UI.EmptyLine()
	c.UI.PrintSection("Commands")

	for _, command := range commands {
		c.UI.EmptyLine()
		c.UI.PrintLabel("  %s", command.Name)
		for _, run := range command.Run {
			c.UI.PrintQuote("  ", run)
		}
	}
}
