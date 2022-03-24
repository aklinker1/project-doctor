package config

import "path"

var ConfigFile string
var Debug bool
var Cwd string
var ProjectConfig Project

func Dirname() string {
	return path.Base(Cwd)
}
