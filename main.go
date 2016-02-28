package main

import (
	"github.com/cloudcloud/nottify/cli"
)

// Essentially, run the CLI
func main() {
	c := cli.New()

	// with the Command, Process it
	c.Process()
}
