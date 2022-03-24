package config

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"

	"github.com/aklinker1/project-doctor/cmd/log"
)

var (
	NotInPathError    = errors.New("Executable is not in your $PATH")
	WrongVersionError = errors.New("Executable is installed, but does not match the required version")
)

type BaseTool struct {
	Name string `mapstructure:"name"`

	Executable     string            `mapstructure:"executable"`
	GetVersionArgs []string          `mapstructure:"getVersionArgs"`
	VersionRegex   string            `mapstructure:"versionRegex"`
	ChangeVersions map[string]string `mapstructure:"changeVersions"`

	InstallUrl        string `mapstructure:"installUrl"`
	UnixInstallUrl    string `mapstructure:"unixInstallUrl"`
	MacInstallUrl     string `mapstructure:"macInstallUrl"`
	WindowsInstallUrl string `mapstructure:"windowsInstallUrl"`

	PackageManagers map[string]string `mapstructure:"packageManagers"`
}

func (tool BaseTool) Verify() (string, error) {
	// Check installation
	toolPath, err := tool.getPath()
	if err != nil {
		return "", err
	}
	if toolPath == "" {
		return "", NotInPathError
	}

	// Check version
	var installedVersion string
	if tool.VersionRegex != "" {
		installedVersion, err = tool.getVersion(toolPath)
		log.Debug(Debug, "%s's version: %s", tool.Executable, installedVersion)
		if err != nil {
			return "", err
		}
		versionRegex, err := regexp.Compile(tool.VersionRegex)
		if err != nil {
			return "", err
		}
		log.Debug(Debug, "Comparing %s to /%s/", installedVersion, versionRegex)
		if !versionRegex.MatchString(installedVersion) {
			log.Debug(Debug, "Version mismatch: %s vs /%s/", installedVersion, versionRegex)
			return installedVersion, WrongVersionError
		}
	}

	return installedVersion, nil
}

func (tool BaseTool) DisplayName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return tool.Executable
}

func (tool BaseTool) AttemptInstall() error {
	return NotInPathError
}

func (tool BaseTool) getPath() (string, error) {
	log.Debug(Debug, "Executing: which %s", tool.Executable)
	out, err := exec.Command("which", tool.Executable).Output()
	log.Debug(Debug, "Output: %s", out)
	if err != nil {
		// Assume errors mean it's not installed
		return "", nil
	}
	return strings.TrimSpace(string(out)), nil
}

func (tool BaseTool) getVersion(toolPath string) (string, error) {
	log.Debug(Debug, "Executing: %s %v", tool.Executable, tool.GetVersionArgs)
	out, err := exec.Command(tool.Executable, tool.GetVersionArgs...).Output()
	log.Debug(Debug, "Output: %s", out)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
