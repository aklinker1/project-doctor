package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
)

type Project struct {
	Tools []Tool `mapstructure:"tools"`
}

type Tool map[string]interface{}

type ToolType string

const (
	TOOL_TYPE_RAW    ToolType = "raw"
	TOOL_TYPE_PRESET ToolType = "preset"
)

type RawToolType struct {
	Type ToolType `mapstructure:"type"`
}

type PresetToolType struct {
	Type   ToolType `mapstructure:"type"`
	Preset string   `mapstructure:"preset"`
}

func AsRaw(tool Tool) (result RawToolType) {
	typeField, hasTypeField := tool["type"]
	if !hasTypeField {
		typeField = TOOL_TYPE_RAW
	}
	typeStr, ok := typeField.(ToolType)
	if !ok {
		cobra.CheckErr(fmt.Errorf("Unknown tool.type = %+v", typeField))
	}
	if typeStr == TOOL_TYPE_RAW {
		mapstructure.Decode(tool, &result)
		return result
	} else {
		cobra.CheckErr(fmt.Errorf("tool.type was not %s (was %s)", TOOL_TYPE_RAW, typeStr))
		return result
	}
}
