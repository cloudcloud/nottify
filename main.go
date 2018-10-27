// Package main provides the executable for Nottify. This includes the command line interface, and
// web component management commands.
package main

import (
	"fmt"
	"os"

	"github.com/cloudcloud/nottify/v1/core"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	c = map[string]interface{}{}
)

// Essentially, run the CLI
func main() {
	args := os.Args
	if len(args) > 0 {
		args = args[1:]
	}

	k := setupKingpin(
		kingpin.New("nottify", "Nottify is your own personal audio streamer."),
	)
	cmd, err := k.Parse(args)
	if err != nil {
		k.Usage(args)
		return
	}

	if *c["debug"].(*bool) {
		fmt.Printf("Command selection is [%s]\n", cmd)
	}

	app(cmd, args, k)
}

func setupKingpin(k *kingpin.Application) *kingpin.Application {
	c["debug"] = k.Flag("debug", "Enable debug mode.").
		Short('d').Default("false").Bool()
	c["dry-run"] = k.Flag("dry-run", "Run the command in dry-run mode.").
		Short('r').Default("false").Bool()
	c["initFile"] = k.Flag("filename", "Filename for the configuration to use.").
		Short('c').Default(".nottify.json").String()

	k.Command("init", "Initialise the local Nottify instance.")

	ingest := k.Command("ingest", "Ingest files from the configured locations.")
	c["ingestPaths"] = ingest.Arg(
		"paths",
		"Paths to include for ingestion. If empty, all configured paths will be used.",
	).Default("").Strings()

	start := k.Command("start", "Bring up the Nottify server.")
	c["start"] = start.Flag("foreground", "Run the server in the foreground.").
		Short('f').Default("false").Bool()

	return k
}

func app(cmd string, args []string, k *kingpin.Application) {
	app := &core.Nottify{
		Args:    args,
		Command: cmd,
	}

	switch cmd {
	case "init":
		app.Init(*c["initFile"].(*string))

	case "ingest":

	case "start":

	default:
		k.Usage(args)
	}
}
