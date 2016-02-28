package cli

import "fmt"

var cmdConfig = &Command{
	UsageLine: "config [action] [name] [value]",
	Short:     "Work a specific config item",
	Long: `
Config will allow for interaction of a specific Configuration option.

For example:
	nottify config set nottify.base_path /media/files/music/
	nottify config get nottify.base_path
`,
}

func init() {
	cmdConfig.Run = configRun
}

var (
	basePath string
	pinCode  int
	pqDsn    string
)

func configRun(args []string) {
	if len(args) == 0 {
		errorf("No config action provided.\nPlease run 'nottify help config'\n")
	}

	if len(args) > 3 {
		errorf("Too many arguments provided.\nPlease run 'nottify help config'\n")
	}

	if args[0] == "get" {
		fmt.Println("Get the variable")
	} else if args[0] == "set" {
		fmt.Println("Set the variable")
	} else {
		errorf("Invalid action.\nPlease run 'nottify help config'\n")
	}
}
