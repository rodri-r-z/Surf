package logger

import (
	"os"
	"surf/ansi"
	"surf/code"
)

var logLevel = ansi.Colorize("cyan_bright_bold", "LOG |")
var errorLevel = ansi.Colorize("red_bright_bold", "ERROR |")
var helpLevel = ansi.Colorize("green_bright_bold", "HELP |")
var helpPrefix = ansi.Colorize("black_bright", "  |>")

// Log prints a set of messages to the console
// in O(n) time
func Log(message ...string) {
	for _, m := range message {
		println(logLevel, m)
	}
}

// Error prints a set of error messages to the console
// in O(n) time
func Error(message ...string) {
	for _, m := range message {
		println(errorLevel, ansi.Colorize("red_bright", m))
	}
}

// Help prints a set of help messages to the console
// in O(n) time
func Help(message ...string) {
	for _, m := range message {
		println(helpPrefix, helpLevel, m)
	}
}

// TokenError prints an error message related to the given token
// with its trace and help messages to the console
// and then exits the program with code 1
// in O(n) time
func TokenError(
	token code.Token,
	message string,
	help ...string,
) {
	Error(message)
	Log("Full context:", token.GetTrace(), token.GetTraceContext(), token.GetTraceIndicator())

	for _, h := range help {
		Help(h)
	}

	os.Exit(1)
}