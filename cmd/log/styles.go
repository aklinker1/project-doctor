package log

import "github.com/ttacon/chalk"

var Title = chalk.Bold.TextStyle
var Success = chalk.Green.NewStyle().WithTextStyle(chalk.Bold).Style
var Error = chalk.Red.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor).Style
var Dim = chalk.Dim.TextStyle
