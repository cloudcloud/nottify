// Package core provides the shared, integral parts within Nottify.
package core

import (
	"fmt"
	"os"

	"github.com/cloudcloud/nottify/v1/config"
)

// Nottify gives structure to working with command line actions for
// managing and working with Nottify itself.
type Nottify struct {
	Args    []string
	Command string
	Config  config.Config
}

// Init prepares and sets up a local installation of Nottify.
func (n *Nottify) Init(f string) {
	c, err := config.FromFile(f, os.Stdout)
	if err != nil {
		panic(err)
	}

	n.Config = c
	c.O(config.Message, fmt.Sprintf("Setup config at %s!\nEnjoy!", f))
}
