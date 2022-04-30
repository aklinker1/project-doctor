package validate

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/qri-io/jsonschema"
)

//go:embed project-schema.json
var projectSchema []byte

var (
	schema *jsonschema.Schema
)

func Project(allConfig map[string]interface{}) error {
	ctx := context.Background()
	// ui.Debug("Validating config: %v", allConfig)
	bytesToValidate, err := jsoniter.Marshal(allConfig)
	if err != nil {
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Failed to marshal config into JSON string",
			Op:       "config.ValidateProject",
			Err:      err,
		}
	}
	// ui.Debug("Validating config json: %s", string(bytesToValidate))

	err = loadSchema(ctx)
	if err != nil {
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Failed to load JSON schema",
			Op:       "config.ValidateProject",
			Err:      err,
		}
	}

	errs, err := schema.ValidateBytes(ctx, bytesToValidate)
	if err != nil {
		// ui.Debug("Validation failed: %v", err)
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  fmt.Sprintf("Validation failed: %v", err),
			Op:       "config.ValidateProject",
			Err:      err,
		}
	}
	if len(errs) > 0 {
		fmt.Println(log.BoldRed("Config failed validation:"))
		for _, err := range errs {
			fmt.Println(err)
		}
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Project config failed validation",
			Op:       "config.validateProject",
			Err:      errs[0],
		}
	}

	if err != nil {
		// ui.Debug("Validation failed: %v", err)
		return err
	}
	// ui.Debug("Validation succeeded")
	return nil
}

func loadSchema(ctx context.Context) error {
	// ui.Debug("JSON Schema: %s", strings.TrimSpace(string(JSONSchema)))
	schema = &jsonschema.Schema{}
	return json.Unmarshal(projectSchema, schema)
}
