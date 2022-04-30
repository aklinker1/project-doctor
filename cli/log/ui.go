package log

import (
	"fmt"

	"github.com/aklinker1/project-doctor/cli"
)

type TerminalUI struct {
	IsDebug bool
	IsColor bool
}

func NewTerminalUI(isDebug bool, isColor bool) cli.UI {
	return &TerminalUI{
		IsColor: isColor,
		IsDebug: isDebug,
	}
}

func (*TerminalUI) EmptyLine() {
	println()
}

func (ui *TerminalUI) Debug(format string, args ...interface{}) {
	if ui.IsDebug {
		ui.printStyle(Dim, format+"\n", args...)
	}
}

func (ui *TerminalUI) printStyle(style func(string) string, format string, args ...interface{}) {
	internalFormat := format
	if ui.IsColor {
		internalFormat = style(format)
	}
	fmt.Printf(internalFormat, args...)
}

func (ui *TerminalUI) PrintTitle(format string, args ...interface{}) {
	ui.printStyle(Bold, format+"\n", args...)
}

func (ui *TerminalUI) PrintSubtitle(format string, args ...interface{}) {
	ui.printStyle(Dim, format+"\n", args...)
}

func (ui *TerminalUI) PrintSection(format string, args ...interface{}) {
	ui.printStyle(BoldBlue, format+"\n", args...)
}

func (ui *TerminalUI) PrintQuote(spaces string, format string, args ...interface{}) {
	fmtArgs := []interface{}{spaces}
	fmtArgs = append(fmtArgs, args...)
	fmt.Printf(Dim("%s│ ")+format+"\n", fmtArgs...)
}

func (ui *TerminalUI) PrintLabel(format string, args ...interface{}) {
	ui.printStyle(DimItalic, format+"\n", args...)
}

func (ui *TerminalUI) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (ui *TerminalUI) Println(format string, args ...interface{}) {
	ui.Printf(format+"\n", args...)
}

func (ui *TerminalUI) Spinner(status func() string) (stop func(error)) {
	if !ui.IsDebug {
		stop, spin := brailSpinner(status)
		go spin()
		return stop
	}

	ui.Println(" - %s...", status())
	return func(err error) {
		if err == nil {
			ui.Println(" ✔ %s", status())
		} else {
			ui.Println(" ✘ %s", status())
		}
	}
}

func (*TerminalUI) ReadLine() (string, error) {
	panic("Not implemented")
}

func (*TerminalUI) ReadMultiselect(options []string) (string, error) {
	panic("Not implemented")
}
