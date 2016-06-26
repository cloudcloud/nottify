package main

import "io"

// Ingest structures detail for the ingest Command
type Ingest struct {
	Command
}

// Run will process the full Ingest command
func (i *Ingest) Run(o io.Writer) error {
	return nil
}

// Init will use provided arguments to prepare for a Run
func (i *Ingest) Init(args []string, d bool, format string) error {
	i.args = args
	i.debug = d
	i.format = format

	return nil
}
