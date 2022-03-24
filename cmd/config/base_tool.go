package config

import (
	"errors"
	"regexp"

	"github.com/aklinker1/project-doctor/cmd/exec"
	"github.com/aklinker1/project-doctor/cmd/log"
)

var (
	NotInPathError = errors.New("Executable is not in your $PATH")
)

type BaseTool struct {
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

func (tool BaseTool) Verify() error {
	// Check installation
	toolPath := tool.getPath()
	if toolPath == "" {
		return NotInPathError
	}

	// Check version
	if tool.VersionRegex != "" {
		installedVersion, err := tool.getVersion(toolPath)
		log.Debug("%s's version: %s", tool.Executable, installedVersion)
		if err != nil {
			return err
		}
		versionRegex, err := regexp.Compile(tool.VersionRegex)
		if err != nil {
			return err
		}
		log.Debug("Comparing %s to /%s/", installedVersion, versionRegex)
		if !versionRegex.MatchString(installedVersion) {
			log.Debug("Version mismatch: %s vs /%s/", installedVersion, versionRegex)
			return NewWrongVersionError(tool, installedVersion)
		}
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
	return nil
}

func (tool BaseTool) getPath() string {
	return exec.Which(exec.Shell(), tool.Executable)
}

func (tool BaseTool) getVersion(toolPath string) (string, error) {
	output, err := exec.Command(exec.Shell(), "-c", tool.GetVersion)
	if err != nil {
		return "", err
	}
	return output, nil
}
