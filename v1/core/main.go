// Package core provides the shared, integral parts within Nottify.
package core

import (
	"io"
	"sync"

	"github.com/cloudcloud/nottify/v1/config"
)

// Nottify gives structure to working with command line actions for
// managing and working with Nottify itself.
type Nottify struct {
	Args    []string
	Command string
	Config  config.Config
	Debug   bool
	Level   int

	fc chan entry
	m  *sync.Map
	wg sync.WaitGroup
}

// Init prepares and sets up a local installation of Nottify.
func (n *Nottify) Init(f string, o io.Writer) *Nottify {
	err := n.setupConfig(f, o)
	if err != nil {
		panic(err)
	}

	return n
}

func (n *Nottify) setupConfig(f string, o io.Writer) error {
	c, err := config.FromFile(f, o, n.Debug)
	c.Level = n.Level
	n.Config = c

	return err
}
