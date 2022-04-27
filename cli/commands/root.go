package commands

import (
	"os"

	"github.com/aklinker1/project-doctor/cli/config"
	"github.com/aklinker1/project-doctor/cli/errors"
	"github.com/aklinker1/project-doctor/cli/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Checkup on your project's local development environment",
	Long: `doctor is a powerful tool for setting up, validating, and documenting all the different tools a project requires. Just run:

    doctor

in a directory with a "doctor.config.yml" file to get started`,
	Run: doctor,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	println()
	err := rootCmd.Execute()
	println()
	if exitCode := errors.ExitCode(err); exitCode != 0 {
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(config.Init)

	rootCmd.Flags().StringVarP(&config.ConfigFile, "config", "c", "", "The path to the project's config file")
	rootCmd.Flags().BoolVar(&log.IsDebugging, "debug", false, "Print all debug statements")
}
