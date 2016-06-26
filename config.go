package main

import (
	"io"

	"github.com/cloudcloud/nottify/config"
)

// Config provides structure for working with configuration
type Config struct {
	Command

	Config   *config.File `json:"config" yaml:"config"`
	Filename string       `json:"filename" yaml:"filename"`
}

// Run executes the command processing
func (c *Config) Run(o io.Writer) error {
	return nil
}

// Init will use provided arguments to prepare for a Run
func (c *Config) Init(args []string, d bool, format string) error {
	c.args = args
	c.debug = d
	c.format = format

	return nil
}

// GetShort returns the simple short description
func (c *Config) GetShort() string {
	return c.Description
}
