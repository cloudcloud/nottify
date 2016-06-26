// Package main provides the executable for Nottify. This includes the command line interface, and
// web component management commands.
package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/ogier/pflag"
)

// default values for incoming variables
const (
	version = "0.1.0_a"

	debugModeDefault = false
	debugModeUsage   = "Enable or Disable debug mode for command."

	formatDefault = "json"
	formatUsage   = "Output format for command."
)

// variables containing provided cli inputs
var (
	debugMode bool
	format    string
	commands  map[string]Comm
)

// prime the environment and settings
func init() {
	// set up command handlers
	commands = map[string]Comm{
		"config": gen("config"),
		"ingest": gen("ingest"),
	}

	// begin flag definitions
	flag.StringVarP(&format, "format", "f", formatDefault, formatUsage)
	flag.BoolVarP(&debugMode, "debug", "d", debugModeDefault, debugModeUsage)
}

// Essentially, run the CLI
func main() {
	// get some flags happening
	flag.Parse()

	// pick a winner
	actual, ok := commands[flag.Arg(0)]
	if !ok {
		// now print the commands
		usage(os.Stdout, flag.Arg(0), commands)

		// display full help
		flag.Usage()

		return
	}

	actual.Init(flag.Args(), debugMode, format)
}

func usage(w io.Writer, a string, c map[string]Comm) {
	fmt.Fprintf(w, "Nottify v%s\n\n", version)

	if len(a) > 0 {
		fmt.Fprintf(w, "Command [%s] does not exist!\n\n", a)
	}

	fmt.Fprintf(w, "Available commands:\n")

	for k, v := range c {
		if v == nil {
			continue
		}

		fmt.Fprintf(w, `	%s
		%s
`, k, v.GetShort())
	}

	fmt.Fprintf(w, "\n")
}

func gen(c string) Comm {
	var a Comm

	switch c {
	case "ingest":
		b := new(Ingest)
		b.Description = "Work with import and export of data from filesystem"

		a = b
	case "config":
		b := new(Config)
		b.Description = "Provide interface to view and modify configuration directives"

		a = b
	}

	return a
}

// Command provides a structure for individual commands to be run
type Command struct {
	Description string `json:"description"`
	Explanation string `json:"explanation"`

	args   []string
	debug  bool
	format string
}

// Comm defines behaviours that a Command should exhibit
type Comm interface {
	Run(io.Writer) error
	Init([]string, bool, string) error
	Usage(io.Writer) error
	GetShort() string
}

// LoggedError is a known error for internal usage
type LoggedError struct {
	error
}

// Run provides command execution, dummy method here
func (c *Command) Run(o io.Writer) error {
	return fmt.Errorf("Run() unimplemented for this Command")
}

// Init sets up an environment, effectively storing commands in this parent
func (c *Command) Init(args []string, d bool, format string) error {
	c.args = args
	c.debug = d
	c.format = format

	return nil
}

// Usage provides instructions on usage of this command
func (c *Command) Usage(w io.Writer) error {
	fmt.Fprintf(w, ``)

	return nil
}

// GetShort provides a short description, for usage amongst other things
func (c *Command) GetShort() string {
	return c.Description
}
