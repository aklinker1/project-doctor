package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aklinker1/project-doctor/cmd/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/qri-io/jsonschema"
)

var (
	schemaUrl       = "https://raw.githubusercontent.com/aklinker1/project-doctor/main/api/schema.json"
	hasLoadedSchema = false
	schema          *jsonschema.Schema
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
		fmt.Println(log.Error("Invalid config:"))
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
	res, err := http.Get(schemaUrl)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Debug(Debug, "JSON Schema: %s", strings.TrimSpace(string(bytes)))
	schema = &jsonschema.Schema{}
	if err = json.Unmarshal(bytes, schema); err != nil {
		return err
	}
	return nil
}
