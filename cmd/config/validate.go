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
	log.Debug(Debug, "Validating config: %v", allConfig)
	bytesToValidate, err := jsoniter.Marshal(allConfig)
	if err != nil {
		return err
	}
	log.Debug(Debug, "Validating config json: %s", string(bytesToValidate))

	err = loadSchema(ctx)
	if err != nil {
		return err
	}

	errs, err := schema.ValidateBytes(ctx, bytesToValidate)
	if err != nil {
		log.Debug(Debug, "Validation failed: %v", err)
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
		log.Debug(Debug, "Validation failed: %v", err)
		return err
	}
	log.Debug(Debug, "Validation succeeded")
	return nil
}

func loadSchema(ctx context.Context) error {
	log.Debug(Debug, "Fetching schema from: %s", schemaUrl)
	if schema != nil {
		log.Debug(Debug, "Schema has already been fetched")
		return nil
	}

	var bytes []byte
	var err error
	if UseLocalSchema {
		bytes, err = getLocalSchema()
	} else {
		bytes, err = getRemoteSchema(ctx)
	}
	if err != nil {
		return err
	}

	log.Debug(Debug, "JSON Schema: %s", strings.TrimSpace(string(bytes)))
	schema = &jsonschema.Schema{}
	return json.Unmarshal(bytes, schema)
}

func getRemoteSchema(ctx context.Context) ([]byte, error) {
	log.Debug(Debug, "Loading schema from: %s", schemaUrl)
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
	log.Debug(Debug, "Loading schema from: file://%s", filename)
	return ioutil.ReadFile(filename)
}
