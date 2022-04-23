package cmd

import (
	"errors"
	"fmt"

	"github.com/aklinker1/project-doctor/cmd/config"
	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/spf13/cobra"
)

func doctor(cmd *cobra.Command, args []string) {
	fmt.Printf(log.Title("Project Doctor: %s\n"), config.Dirname())
	fmt.Println(log.Dim("Checking up on your local development environment"))

	toolsSection()
	commandsSection()
}

func toolsSection() {
	project := config.ProjectConfig
	if len(project.Tools) == 0 {
		return
	}

	toolErrors := []error{}
	println()
	fmt.Println(log.SectionHeader("Tools"))
	println()
	for _, toolJson := range project.Tools {
		tool := config.ParseTool(toolJson)
		status := tool.DisplayName

		// Do the work
		stop, spin := log.BrailSpinner(status)
		go spin()
		err := tool.Verify()
		stop(err)

		if errors.Is(err, config.NotInPathError) {
			fmt.Println("    Not installed")
			err = tool.AttemptInstall()
		}
		if errors.Is(err, config.WrongVersionError) {
			fmt.Printf("    Installed version: %s\n", config.AsWrongVersionError(err).InstalledVersion)
		}
		if err != nil {
			fmt.Println("    Error:", err)
			toolErrors = append(toolErrors, err)
		}
	}
	if len(toolErrors) > 0 {
		plural := "s"
		if len(toolErrors) == 1 {
			plural = ""
		}
		log.CheckFatal(
			fmt.Errorf("Found problems%s with %d tool%s", plural, len(toolErrors), plural),
		)
	}
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
		cmd := config.ParseCommand(command)
		fmt.Println(log.Dim("  " + log.Italic(command.Name)))
		for _, c := range cmd.Command {
			fmt.Println(log.Dim("  │ ") + c)
		}
	}
}
