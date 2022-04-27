package config

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/mitchellh/mapstructure"
)

var (
	emptyProject = cli.Project{}
	emptyCommand = cli.Command{}
)

func ParseProject(rawProject map[string]interface{}) (cli.Project, error) {
	checks := []cli.Check{}
	commands := []cli.Command{}

	rawChecks, hasChecks := rawProject["checks"]
	if hasChecks {
		slice, ok := rawChecks.([]interface{})
		if !ok {
			return emptyProject, &cli.Error{
				ExitCode: cli.EINVALID,
				Message:  fmt.Sprintf("The 'checks' field must be a list of objects, got %T", rawChecks),
				Op:       "config.ParseProject",
			}
		}
		for _, item := range slice {
			dic := item.(map[interface{}]interface{})
			parsed, err := ParseCheck(dic)
			if err != nil {
				return emptyProject, err
			}
			checks = append(checks, parsed)
		}
	}

	rawCommands, hasCommands := rawProject["commands"]
	if hasCommands {
		slice, ok := rawCommands.([]interface{})
		if !ok {
			return emptyProject, &cli.Error{
				ExitCode: cli.EINVALID,
				Message:  "The 'commands' field must be a list of objects",
				Op:       "config.ParseProject",
			}
		}
		for _, item := range slice {
			dic := item.(map[interface{}]interface{})
			parsed, err := ParseCommand(dic)
			if err != nil {
				return emptyProject, err
			}
			commands = append(commands, parsed)
		}
	}

	return cli.Project{
		Checks:   checks,
		Commands: commands,
	}, nil
}

func ParseCheck(check map[interface{}]interface{}) (cli.Check, error) {
	checkType, hasType := check["type"]
	if !hasType {
		return nil, &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Check must have 'type' field",
			Op:       "config.ParseCheck",
		}
	}

	switch checkType {
	case "tool":
		tool := ToolCheck{}
		err := mapstructure.Decode(check, &tool)
		return tool, err
	}

	return ToolCheck{}, &cli.Error{
		ExitCode: cli.EINVALID,
		Message:  fmt.Sprintf("Unknown check.type = '%s'", checkType),
		Op:       "config.ParseCheck",
	}
}

func ParseCommand(rawCommand map[interface{}]interface{}) (cli.Command, error) {
	name, hasName := rawCommand["name"]
	if !hasName {
		return emptyCommand, &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Command is missing required field 'name'",
			Op:       "config.ParseCommand",
		}
	}
	nameStr, nameStrOk := name.(string)
	if !nameStrOk {
		return emptyCommand, &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Command 'name' must be a string",
			Op:       "config.ParseCommand",
		}
	}

	run, hasRun := rawCommand["run"]
	if !hasRun {
		return emptyCommand, &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Command is missing required field 'run'",
			Op:       "config.ParseCommand",
		}
	}
	runSlice := []string{}
	if str, ok := run.(string); ok {
		runSlice = append(runSlice, str)
	} else if slice, ok := run.([]interface{}); ok {
		for _, item := range slice {
			runSlice = append(runSlice, item.(string))
		}
	} else {
		return emptyCommand, &cli.Error{
			ExitCode: cli.EINVALID,
			Message: fmt.Sprintf(
				"Command field 'run' must be a string or string array, but was parsed as: %T",
				run,
			),
			Op: "config.ParseCommand",
		}
	}

	return cli.Command{
		Name: nameStr,
		Run:  runSlice,
	}, nil
}
