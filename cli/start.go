package cli

var cmdStart = &Command{
	UsageLine: "start",
	Short:     "Begin the Nottify server",
	Long: `
`,
}

func init() {
	cmdStart.Run = startRun
}

func startRun(args []string) {
	// spin up the server
}
