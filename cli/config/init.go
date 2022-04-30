package config

import (
	"github.com/aklinker1/project-doctor/cli"
	"github.com/spf13/viper"
)

func Init() {
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Lookup config file
		viper.AddConfigPath(GetCWD())        // Look in the working directory
		viper.SetConfigType("yaml")          // Prefer yaml files
		viper.SetConfigName("doctor.config") // Default config filename without an extension
	}

	viper.AutomaticEnv() // read in environment variables that match

	// Validate and read in the config
	if err := viper.ReadInConfig(); err != nil {
		panic(cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Unable to find config file",
			Op:       "config.BeforeExecute",
			Err:      err,
		})
	}
}
