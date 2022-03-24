package config

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cmd/log"
	"github.com/mitchellh/mapstructure"
)

type Project struct {
	Tools []ToolJson `mapstructure:"tools"`
}

type Tool interface {
	Verify() error
	DisplayName() string
	AttemptInstall() error
}

// ToolJson is the raw map that data is loaded into as JSON. Use `ParseTool` to convert this into an
// object you can work with
type ToolJson map[string]interface{}

type ToolType string

const (
	TOOL_TYPE_BASE   ToolType = "base"
	TOOL_TYPE_PRESET ToolType = "preset"
)

// Parse tool returns a tool that includes operations like validate
func ParseTool(toolJson ToolJson) Tool {
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
		raw := BaseTool{}
		mapstructure.Decode(toolJson, &raw)
		return raw
	default:
		log.CheckFatal(fmt.Errorf("tool.type was not %s (was %s)", TOOL_TYPE_BASE, typeStr))
		return nil
	}
}
