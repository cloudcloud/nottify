// Package main provides the executable for Nottify. This includes the command line interface, and
// web component management commands.
package main

// default values for incoming variables
const (
	debugModeDefault = false
	formatDefault    = "json"
)

// variables containing provided cli inputs
var (
	debugMode bool
	format    string
)

// prime the environment and settings
func init() {
	// set up command handlers

	// begin flag definitions
}

// Essentially, run the CLI
func main() {
	// get some flags happening
}

// Command provides a structure for individual commands to be run
type Command struct {
	Usage       string `json:"usage"`
	Description string `json:"description"`
	Explanation string `json:"explanation"`

	args []string
}

// Comm defines behaviours that a Command should exhibit
type Comm interface {
	Run() error
	Init([]string) error
}

// LoggedError is a known error for internal usage
type LoggedError struct {
	error
}
