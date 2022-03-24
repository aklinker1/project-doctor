package config

import (
	"errors"
	"time"
)

type BaseTool struct {
	Name         string `mapstructure:"name"`
	Executable   string `mapstructure:"executable"`
	VersionRegex string `mapstructure:"versionRegex"`

	InstallUrl        string `mapstructure:"installUrl"`
	UnixInstallUrl    string `mapstructure:"unixInstallUrl"`
	MacInstallUrl     string `mapstructure:"macInstallUrl"`
	WindowsInstallUrl string `mapstructure:"windowsInstallUrl"`

	PackageManagers map[string]string `mapstructure:"packageManagers"`
}

func (tool BaseTool) Verify() error {
	time.Sleep(2 * time.Second)
	return errors.New("BaseTool.Verify not implemented")
}

func (tool BaseTool) DisplayName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return tool.Executable
}
