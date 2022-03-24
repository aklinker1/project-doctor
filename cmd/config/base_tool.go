package config

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	NotInPathError = errors.New("Executable is not in your $PATH")
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
	toolPath, err := tool.getPath()
	if err != nil {
		return err
	}
	if toolPath == "" {
		return NotInPathError
	}
	return nil
}

func (tool BaseTool) DisplayName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return tool.Executable
}

func (tool BaseTool) AttemptInstall() error {
	fmt.Println("    Not installed")
	return errors.New("BaseTool.AttemptInstall not implemented")
}

func (tool BaseTool) getPath() (string, error) {
	out, err := exec.Command("which", tool.Executable).Output()
	if err != nil {
		return "", nil
	}
	return string(out), nil
}
