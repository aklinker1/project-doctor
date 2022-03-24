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
	println()

	fmt.Println(log.SectionHeader("Tools"))
	project := config.ProjectConfig
	toolErrors := []error{}
	for _, toolJson := range project.Tools {
		tool := config.ParseTool(toolJson)
		status := tool.DisplayName

		// Do the work
		stop, spin := log.BrailSpinner(status)
		go spin()
		currentVersion, err := tool.Verify()
		stop(err)

		if errors.Is(err, config.NotInPathError) {
			fmt.Println("    Not installed")
			err = tool.AttemptInstall()
		}
		if errors.Is(err, config.WrongVersionError) {
			fmt.Printf("    Installed version: %s\n", currentVersion)
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
		log.CheckFatal(fmt.Errorf("Found problems%s with %d tool%s", plural, len(toolErrors), plural))
	}
}
