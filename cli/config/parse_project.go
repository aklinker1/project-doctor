package config

import "github.com/mitchellh/mapstructure"

var (
	EmptyProject = Project{}
)

func ParseProjectConfig(allConfig map[string]interface{}) (project Project, err error) {
	err = mapstructure.Decode(allConfig, &project)
	return
}
