package cli

var cmdIngest = &Command{
	UsageLine: "ingest",
	Short:     "Complete an ingestion of Media",
	Long: `
Ingest will complete a file system trawl for media files.

For example:
	nottify ingest
`,
}

func init() {
	cmdIngest.Run = ingestRun
}

func ingestRun(args []string) {
	// unimplemented so far
}
