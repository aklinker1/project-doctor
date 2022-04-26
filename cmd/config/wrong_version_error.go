package config

import (
	"fmt"
)

type wrongVersionError struct {
	executable           string
	expectedVersionRegex string
	InstalledVersion     string
}

var WrongVersionError = &wrongVersionError{}

func NewWrongVersionError(tool ToolCheck, installedVersion string) error {
	return &wrongVersionError{
		executable:           tool.Executable,
		expectedVersionRegex: tool.VersionRegex,
		InstalledVersion:     installedVersion,
	}
}

func AsWrongVersionError(err error) *wrongVersionError {
	return err.(*wrongVersionError)
}

func (e *wrongVersionError) Error() string {
	return fmt.Sprintf("%s requires version to match /%s/, but got '%s'", e.executable, e.expectedVersionRegex, e.InstalledVersion)
}

func (e *wrongVersionError) Is(tgt error) bool {
	_, ok := tgt.(*wrongVersionError)
	return ok
}
