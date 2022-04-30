package config

import (
	"os"
	"path"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/validate"
	"github.com/spf13/viper"
)

var (
	// Global flags
	IsColor    bool
	IsDebug    bool
	ConfigFile string

	// Singletons
	cwd     *string
	project *cli.Project
)

func GetCWD() string {
	if cwd != nil {
		return *cwd
	}
	v, err := os.Getwd()
	if err != nil {
		panic(cli.Error{
			ExitCode: cli.EINTERNAL,
			Message:  "Could not get the processes CWD",
			Op:       "config.GetCWD",
			Err:      err,
		})
	}
	cwd = &v
	return v
}

func Dirname() string {
	return path.Base(GetCWD())
}

func GetProject() cli.Project {
	if project != nil {
		return *project
	}
	// ui.Debug("Using project file: %s", viper.ConfigFileUsed())

	projectMap := viper.AllSettings()
	// ui.Debug("Project as seen by Viper: %+v", projectMap)

	err := validate.Project(projectMap)
	if err != nil {
		panic(cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Invalid project config",
			Op:       "config.BeforeExecute",
			Err:      err,
		})
	}

	p, err := ParseProject(projectMap)
	if err != nil {
		panic(cli.Error{
			ExitCode: cli.EINTERNAL,
			Message:  "Failed to parse project from a valid project config file",
			Op:       "config.BeforeExecute",
			Err:      err,
		})
	}
	project = &p

	return p
}
