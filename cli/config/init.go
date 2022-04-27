package config

import (
	"os"

	"github.com/aklinker1/project-doctor/cli/log"
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
		log.CheckFatal(err)
		viper.AddConfigPath(Cwd)             // Look in the working directory
		viper.SetConfigType("yaml")          // Prefer yaml files
		viper.SetConfigName("doctor.config") // Default config filename without an extension
	}

	viper.AutomaticEnv() // read in environment variables that match

	// Validate and read in the config
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using project file: %s", viper.ConfigFileUsed())
		projectMap := viper.AllSettings()
		log.Debug("Project as seen by Viper: %+v", projectMap)

		err = validateProject(projectMap)
		log.CheckFatal(err)

		ProjectConfig, err = ParseProjectConfig(projectMap)
		log.CheckFatal(err)
		log.Debug("Received config: %+v", ProjectConfig)
	} else {
		log.CheckFatal(err)
	}
}
