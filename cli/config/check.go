package config

import (
	"errors"
	"fmt"

	"github.com/aklinker1/project-doctor/cli/log"
	"github.com/mitchellh/mapstructure"
)

type Check interface {
	// DisplayName shown in the console
	DisplayName() string
	// Verify that the check is passing
	Verify() error
	// Fix attempts to fix the verification issue
	Fix() error
}

// Parse tool returns a tool that includes operations like validate
func ParseCheck(check map[string]interface{}) Check {
	checkType, hasType := check["type"]
	if !hasType {
		log.CheckFatal(errors.New("Check must have a 'type' field"))
	}

	switch checkType {
	case "tool":
		tool := ToolCheck{}
		err := mapstructure.Decode(check, &tool)
		log.CheckFatal(err)
		return tool
	}

	log.CheckFatal(fmt.Errorf("Unknown check.type = '%s'", checkType))
	return ToolCheck{}
}
