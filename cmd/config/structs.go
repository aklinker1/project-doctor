package config

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/mitchellh/mapstructure"
)

type Project struct {
	Tools    []ToolJSON    `mapstructure:"tools"`
	Commands []CommandJSON `mapstructure:"commands"`
}

// ToolJSON is the raw map that data is loaded into as JSON. Use `ParseTool` to convert this into an
// object you can work with
type ToolJSON map[string]interface{}

type ToolType string

const (
	TOOL_TYPE_BASE   ToolType = "base"
	TOOL_TYPE_PRESET ToolType = "preset"
)

// Parse tool returns a tool that includes operations like validate
func ParseTool(toolJson ToolJSON) Tool {
	typeField, hasTypeField := toolJson["type"]
	if !hasTypeField {
		typeField = TOOL_TYPE_BASE
	}
	typeStr, ok := typeField.(ToolType)
	if !ok {
		log.CheckFatal(fmt.Errorf("Unknown tool.type = %+v", typeField))
	}
	switch typeStr {
	case TOOL_TYPE_BASE:
		raw := InstalledTool{}
		mapstructure.Decode(toolJson, &raw)
		return raw
	default:
		log.CheckFatal(fmt.Errorf("tool.type was not %s (was %s)", TOOL_TYPE_BASE, typeStr))
		return nil
	}
}

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
