package main

// Ingest structures detail for the ingest Command
type Ingest struct {
	Command
}

// Run will process the full Ingest command
func (c *Ingest) Run() error {
	return nil
}

// Init will use provided arguments to prepare for a Run
func (c *Ingest) Init(args []string) error {
	c.args = args

	return nil
}
