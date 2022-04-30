package config

import (
	"errors"
	"regexp"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/exec"
)

var (
	NotInPathError = errors.New("Executable is not in your $PATH")
)

type ToolCheck struct {
	Name string `mapstructure:"name"`

	Executable     string            `mapstructure:"executable"`
	GetVersion     string            `mapstructure:"getVersion"`
	VersionRegex   string            `mapstructure:"versionRegex"`
	ChangeVersions map[string]string `mapstructure:"changeVersions"`

	InstallUrl        string `mapstructure:"installUrl"`
	UnixInstallUrl    string `mapstructure:"unixInstallUrl"`
	MacInstallUrl     string `mapstructure:"macInstallUrl"`
	WindowsInstallUrl string `mapstructure:"windowsInstallUrl"`

	PackageManagers map[string]string `mapstructure:"packageManagers"`
}

func (tool ToolCheck) Verify(ui cli.UI) error {
	defaultShell, err := exec.Shell()
	if err != nil {
		return err
	}

	return tool.verifyShell(ui, defaultShell)
}

func (tool ToolCheck) verifyShell(ui cli.UI, shell string) error {
	// Check installation
	toolPath, _ := exec.Which(ui, shell, tool.Executable)
	if toolPath == "" {
		return NotInPathError
	}

	// Check version
	if tool.VersionRegex != "" {
		installedVersion, err := exec.Command(ui, shell, tool.GetVersion)
		ui.Debug("%s's version: %s", tool.Executable, installedVersion)
		if err != nil {
			return err
		}
		versionRegex, err := regexp.Compile(tool.VersionRegex)
		if err != nil {
			return err
		}
		ui.Debug("Comparing %s to /%s/", installedVersion, versionRegex)
		if !versionRegex.MatchString(installedVersion) {
			ui.Debug("Version mismatch: %s vs /%s/", installedVersion, versionRegex)
			return NewWrongVersionError(tool, installedVersion)
		}
	}

	return nil
}

func (tool ToolCheck) DisplayName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return tool.Executable
}

func (tool ToolCheck) Fix(ui cli.UI) error {
	return nil
}
