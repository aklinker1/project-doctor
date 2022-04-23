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

	"github.com/aklinker1/project-doctor/cmd/log"
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
		return err
	}
	log.Debug("Validating config json: %s", string(bytesToValidate))

	err = loadSchema(ctx)
	if err != nil {
		return err
	}

	errs, err := schema.ValidateBytes(ctx, bytesToValidate)
	if err != nil {
		log.Debug("Validation failed: %v", err)
		return err
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
	log.Debug("Fetching schema from: %s", schemaUrl)
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
