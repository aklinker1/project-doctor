package config

import (
	_ "embed"
)

//go:embed schema.json
var JSONSchema []byte
