package cli

var cmdInit = &Command{
	UsageLine: "init",
	Short:     "Interactive initialisation",
	Long: `
`,
}

func init() {
	cmdInit.Run = initRun
}

func initRun(args []string) {
	// init all
}
