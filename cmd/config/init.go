package config

import (
	"os"

	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Init() {
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Lookup config file
		var err error
		Cwd, err = os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(Cwd)             // Look in the working directory
		viper.SetConfigType("yaml")          // Prefer yaml files
		viper.SetConfigName("doctor.config") // Default config filename without an extension
	}

	viper.AutomaticEnv() // read in environment variables that match

	// Validate and read in the config
	if err := viper.ReadInConfig(); err == nil {
		log.Debug(Debug, "Using project file: %s", viper.ConfigFileUsed())
		projectMap := viper.AllSettings()
		log.Debug(Debug, "Project as seen by Viper: %+v", projectMap)

		err = validateProject(projectMap)
		cobra.CheckErr(err)

		ProjectConfig, err = ParseProjectConfig(projectMap)
		cobra.CheckErr(err)
		log.Debug(Debug, "Received config: %+v", ProjectConfig)
	}
}
