package config

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cli/log"
)

type Project struct {
	Tools    []map[string]interface{} `mapstructure:"tools"`
	Commands []CommandJSON            `mapstructure:"commands"`
}

// ToolJSON is the raw map that data is loaded into as JSON. Use `ParseTool` to convert this into an
// object you can work with
type ToolJSON map[string]interface{}

type CommandJSON struct {
	Name    string      `mapstructure:"name"`
	Command interface{} `mapstructure:"command"`
}

type Command struct {
	Name    string
	Command []string
}

// Parse commmand returns a tool with fields resolved
func ParseCommand(cmdJSON CommandJSON) Command {
	if str, ok := cmdJSON.Command.(string); ok {
		return Command{
			Name:    cmdJSON.Name,
			Command: []string{str},
		}
	} else if array, ok := cmdJSON.Command.([]interface{}); ok {
		strings := []string{}
		for _, item := range array {
			strings = append(strings, item.(string))
		}
		return Command{
			Name:    cmdJSON.Name,
			Command: strings,
		}
	}
	log.CheckFatal(fmt.Errorf("command.[*].type needs to be a string or []string, not %T", cmdJSON.Command))
	return Command{}
}
