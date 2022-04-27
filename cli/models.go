package cli

type Check interface {
	// DisplayName shown in the console
	DisplayName() string
	// Verify that the check is passing
	Verify() error
	// Fix attempts to fix the verification issue
	Fix() error
}

// Command is a helpful command that describes how to consume/run a project
type Command struct {
	// Name is a description of what the commands will do
	Name string
	// Run lists ways to execute the command from the CLI
	Run []string
}

// Project stores the project's config that will be checked by the CLI tool. This is the final,
// resolved project containing all the checks and all the commands
type Project struct {
	Checks   []Check
	Commands []Command
}
