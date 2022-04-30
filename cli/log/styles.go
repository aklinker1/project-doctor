package log

import "github.com/ttacon/chalk"

var BoldStyle = chalk.Bold
var Bold = BoldStyle.TextStyle

var BoldBlueStyle = chalk.Blue.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var BoldBlue = BoldBlueStyle.Style

var BoldGreenStyle = chalk.Green.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var BoldGreen = BoldGreenStyle.Style

var BoldRedStyle = chalk.Red.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var BoldRed = BoldRedStyle.Style

var DimStyle = chalk.Dim
var Dim = DimStyle.TextStyle

var ItalicStyle = chalk.Italic
var Italic = ItalicStyle.TextStyle

var DimItalic = func(color string) string {
	return Italic(Dim(color))
}

var BoldCyanStyle = chalk.Cyan.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var BoldCyan = BoldCyanStyle.Style
