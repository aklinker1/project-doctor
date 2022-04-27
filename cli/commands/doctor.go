package commands

import (
	"errors"
	"fmt"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/config"
	"github.com/aklinker1/project-doctor/cli/log"
	"github.com/spf13/cobra"
)

func doctor(cmd *cobra.Command, args []string) {
	fmt.Printf(log.Title("Project Doctor: %s\n"), config.Dirname())
	fmt.Println(log.Dim("Checking up on your local development environment"))

	checksSection()
	commandsSection()
}

func checksSection() error {
	project := config.ProjectConfig
	if len(project.Checks) == 0 {
		return nil
	}

	checkErrors := []error{}
	println()
	fmt.Println(log.SectionHeader("Tools"))
	println()
	for _, check := range project.Checks {
		status := check.DisplayName

		// Do the work
		stop, spin := log.BrailSpinner(status)
		go spin()
		err := check.Verify()
		stop(err)

		if errors.Is(err, config.NotInPathError) {
			fmt.Println("    Not installed")
			err = check.Fix()
		}
		if errors.Is(err, config.WrongVersionError) {
			fmt.Printf("    Installed version: %s\n", config.AsWrongVersionError(err).InstalledVersion)
		}
		if err != nil {
			fmt.Println("    Error:", cli.ErrorMessage(err))
			checkErrors = append(checkErrors, err)
		}
	}
	if len(checkErrors) > 0 {
		return &cli.Error{
			ExitCode: cli.EINTERNAL,
			Message:  fmt.Sprintf("Some checks failed (%d total)", len(checkErrors)),
			Op:       "commands.Doctor",
		}
	}
	return nil
}

func commandsSection() {
	project := config.ProjectConfig
	if len(project.Commands) == 0 {
		return
	}

	println()
	fmt.Println(log.SectionHeader("Commands"))
	for _, command := range project.Commands {
		println()
		fmt.Println(log.Dim("  " + log.Italic(command.Name)))
		for _, c := range command.Run {
			fmt.Println(log.Dim("  │ ") + c)
		}
	}
}
