package commands

import (
	"os"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/config"
	"github.com/aklinker1/project-doctor/cli/executors"
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
	Run: func(cmd *cobra.Command, args []string) {
		ui := log.NewTerminalUI(config.IsDebug, config.IsColor)
		executor := &executors.SeriesExecutor{
			UI: ui,
		}
		defer recovr(ui)
		(&DoctorCommand{
			Executor: executor,
			UI:       ui,
			Project:  config.GetProject(),
		}).Run(cmd, args)
	},
}

func recovr(ui cli.UI) {
	paniced := recover()
	if paniced == nil {
		return
	}

	if err, ok := paniced.(error); ok {
		ui.Println("doctor failed: %s", cli.ErrorMessage(err))
		ui.EmptyLine()
		os.Exit(cli.ExitCode(err))
	} else {
		println("Unknown error:", err)
		ui.EmptyLine()
		os.Exit(cli.EINTERNAL)
	}
}

func init() {
	cobra.OnInitialize(config.Init)

	rootCmd.Flags().StringVarP(&config.ConfigFile, "config", "c", "", "The path to the project's config file")
	rootCmd.Flags().BoolVar(&config.IsDebug, "debug", false, "Print all debug statements")

	config.IsColor = true // TODO: https://rosettacode.org/wiki/Check_input_device_is_a_terminal
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
