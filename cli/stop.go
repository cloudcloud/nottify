package cli

var cmdStop = &Command{
	UsageLine: "stop",
	Short:     "Stop the Nottify server",
	Long: `
`,
}

func init() {
	cmdStop.Run = stopRun
}

func stopRun(args []string) {
	// stop the server
}
