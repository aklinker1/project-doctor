package cli

type Check interface {
	// DisplayName shown in the console
	DisplayName() string
	// Verify that the check is passing
	Verify(ui UI) error
	// Fix attempts to fix the verification issue
	Fix(ui UI) error
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

// ChecksValidator is used to validate all checks and attempt to fix any failing checks
type ChecksValidator interface {
	Validate(checks []Check) []error
}

// UI defines how to show the user information and get their input
type UI interface {
	EmptyLine()
	Debug(format string, args ...interface{})
	PrintTitle(format string, args ...interface{})
	PrintSubtitle(format string, args ...interface{})
	PrintSection(format string, args ...interface{})
	PrintQuote(spaces string, format string, args ...interface{})
	PrintLabel(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Println(format string, args ...interface{})
	Spinner(status func() string) (stop func(error))
	ReadLine() (string, error)
	ReadMultiselect(options []string) (string, error)
}
