package config

import (
	"os"
	"path"
)

var (
	ConfigFile     string
	Debug          bool
	Cwd            string
	ProjectConfig  Project
	UseLocalSchema bool
)

func init() {
	UseLocalSchema = os.Getenv("USE_LOCAL_SCHEMA") == "true"
}

func Dirname() string {
	return path.Base(Cwd)
}
