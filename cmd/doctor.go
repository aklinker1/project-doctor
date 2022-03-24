package cmd

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cmd/config"
	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/spf13/cobra"
)

func doctor(cmd *cobra.Command, args []string) {
	fmt.Printf(log.Title("Project Doctor: %s\n"), config.Dirname())
	fmt.Println(log.Dim("Setting up tools, validating local env"))
	println()

	fmt.Println(log.Title("Tools"))
	project := config.ProjectConfig
	for _, toolJson := range project.Tools {
		tool := config.ParseTool(toolJson)
		status := tool.DisplayName
		stop, spin := log.BrailSpinner(status)
		go spin()

		err := tool.Verify()

		stop(err)
		log.CheckFatal(err)
	}
}
