package cmd

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cmd/config"
	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func doctor(cmd *cobra.Command, args []string) {
	fmt.Printf(log.Title("Project Doctor: %s\n"), config.Dirname())
	fmt.Println(log.Dim("Setting up tools, validating local env"))
	println(chalk.Blue.Color("Hello world"))
}
