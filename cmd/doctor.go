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
	for _, toolJson := range project.Tools {
		tool := config.ParseTool(toolJson)
		status := tool.DisplayName

		// Do the work
		stop, spin := log.BrailSpinner(status)
		go spin()
		err := tool.Verify()
		stop(err)

		if errors.Is(err, config.NotInPathError) {
			err = tool.AttemptInstall()
		}
		log.CheckFatal(err)
	}
}
