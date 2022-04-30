package cli

import (
	"bytes"
	"fmt"
)

type Error struct {
	ExitCode int
	Message  string
	Op       string
	Err      error
}

const (
	EINTERNAL = 1
	EINVALID  = 2
)

// ExitCode returns the code of the root error, if available. Otherwise returns EINTERNAL.
func ExitCode(err error) int {
	if err == nil {
		return 0
	} else if e, ok := err.(*Error); ok && e.ExitCode != 0 {
		return e.ExitCode
	} else if ok && e.Err != nil {
		return ExitCode(e.Err)
	}
	return EINTERNAL
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return err.Error()
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.ExitCode != 0 {
			fmt.Fprintf(&buf, "<%d> ", e.ExitCode)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}
