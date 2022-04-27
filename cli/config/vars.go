package config

import (
	"os"
	"path"

	"github.com/aklinker1/project-doctor/cli"
)

var (
	ConfigFile     string
	Cwd            string
	ProjectConfig  cli.Project
	UseLocalSchema bool
)

func init() {
	UseLocalSchema = os.Getenv("USE_LOCAL_SCHEMA") == "true"
}

func Dirname() string {
	return path.Base(Cwd)
}
