// Package cli provides all methods for Command Line usage of Nottify. Each individual action point
// is placed within its own file to enhance readability of source.
package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

// LoggedError is a localised wrapper of error
type LoggedError struct {
	error
}

// Command is a generic struct for defining Commands to be made available.
type Command struct {
	Run                    func(args []string)
	UsageLine, Short, Long string
}

// Name will use the particular Command instance to generate a display name.
func (cmd *Command) Name() string {
	name := cmd.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

var commands = []*Command{
	cmdClear,
	cmdConfig,
	cmdIngest,
	cmdInit,
	cmdSearch,
	cmdStart,
	cmdStop,
}

// New will begin the Object init for a specific Command.
func New() *Command {
	c := new(Command)

	return c
}

// Process is the runner for handling the full command entered.
func (cmd *Command) Process() *Command {
	flag.Usage = func() { usage(1) }
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || args[0] == "help" {
		if len(args) == 1 {
			usage(0)
		} else if len(args) > 1 {
			for _, cmd := range commands {
				if cmd.Name() == args[1] {
					tmpl(os.Stdout, helpTemplate, cmd)
					return nil
				}
			}
		}
		usage(2)
	}

	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(LoggedError); !ok {
				panic(err)
			}
			os.Exit(1)
		}
	}()

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Run(args[1:])
			return nil
		}
	}

	errorf("Unknown command [%q]\nRun 'nottify help' for usage.\n", args[0])
	return nil
}

func errorf(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprintf(os.Stderr, format, args...)
	panic(LoggedError{})
}

func usage(exitCode int) {
	tmpl(os.Stderr, usageTemplate, commands)
	os.Exit(exitCode)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

const usageTemplate = `usage: nottify command [arguments]

The commands are:
{{range .}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "nottify help [command]" for more information.
`

var helpTemplate = `usage: nottify {{.UsageLine}}
{{.Long}}
`
