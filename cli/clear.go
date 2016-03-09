package cli

var cmdClear = &Command{
	UsageLine: "clear [type] [options]",
	Short:     "Clear out data from the current environment",
	Long: `
Clear provides a simple way to remove data. Current cached items, data within the database, or other generated data can be removed.

For example:
	nottify clear database prefixDir=Mastodon
`,
}

func init() {
	cmdClear.Run = clearRun
}

func clearRun(args []string) {
	if len(args) == 0 {
		// could allow for full data emptying
	}

	//
}
