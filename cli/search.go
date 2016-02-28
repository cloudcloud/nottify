package cli

var cmdSearch = &Command{
	UsageLine: "search [options]",
	Short:     "Find data within the Database",
	Long: `
`,
}

func init() {
	cmdSearch.Run = searchRun
}

func searchRun(args []string) {
	// get some queries together
	errorf("Unimplemented!")
}
