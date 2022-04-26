package config

type ICheck interface {
	// DisplayName shown in the console
	DisplayName() string
	// Verify that the check is passing
	Verify() error
	// Fix attempts to fix the verification issue
	Fix() error
}

type Check struct {
	Type string `mapstructure:"type"`
	InstalledTool
	// Preset
}
