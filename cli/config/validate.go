package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/aklinker1/project-doctor/cli"
	"github.com/aklinker1/project-doctor/cli/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/qri-io/jsonschema"
)

var (
	hasLoadedSchema = false
	schema          *jsonschema.Schema
	schemaUrl       = "https://raw.githubusercontent.com/aklinker1/project-doctor/main/api/schema.json"
)

func validateProject(allConfig map[string]interface{}) error {
	ctx := context.Background()
	log.Debug("Validating config: %v", allConfig)
	bytesToValidate, err := jsoniter.Marshal(allConfig)
	if err != nil {
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Failed to marshal config into JSON string",
			Op:       "config.validateProject",
			Err:      err,
		}
	}
	log.Debug("Validating config json: %s", string(bytesToValidate))

	err = loadSchema(ctx)
	if err != nil {
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  "Failed to load JSON schema",
			Op:       "config.validateProject",
			Err:      err,
		}
	}

	errs, err := schema.ValidateBytes(ctx, bytesToValidate)
	if err != nil {
		log.Debug("Validation failed: %v", err)
		return &cli.Error{
			ExitCode: cli.EINVALID,
			Message:  fmt.Sprintf("Validation failed: %v", err),
			Op:       "config.validateProject",
			Err:      err,
		}
	}
	if len(errs) > 0 {
		fmt.Println(log.Error("Config failed validation:"))
		for _, err := range errs {
			fmt.Println(err)
		}
		return errors.New("Invalid config")
	}

	if err != nil {
		log.Debug("Validation failed: %v", err)
		return err
	}
	log.Debug("Validation succeeded")
	return nil
}

func loadSchema(ctx context.Context) error {
	log.Debug("JSON Schema: %s", strings.TrimSpace(string(JSONSchema)))
	schema = &jsonschema.Schema{}
	return json.Unmarshal(JSONSchema, schema)
}

func getRemoteSchema(ctx context.Context) ([]byte, error) {
	log.Debug("Loading schema from: %s", schemaUrl)
	res, err := http.Get(schemaUrl)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getLocalSchema() ([]byte, error) {
	filename := path.Join(Cwd, "api", "schema.json")
	log.Debug("Loading schema from: file://%s", filename)
	return ioutil.ReadFile(filename)
}
