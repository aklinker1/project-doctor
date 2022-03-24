package log

import "github.com/ttacon/chalk"

var TitleStyle = chalk.Bold
var Title = TitleStyle.TextStyle

var SectionHeaderStyle = chalk.Blue.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var SectionHeader = SectionHeaderStyle.Style

var SuccessStyle = chalk.Green.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var Success = SuccessStyle.Style

var ErrorStyle = chalk.Red.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var Error = ErrorStyle.Style

var DimStyle = chalk.Dim
var Dim = DimStyle.TextStyle

var LoadingStyle = chalk.Cyan.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
var Loading = LoadingStyle.Style
