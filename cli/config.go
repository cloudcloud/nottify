package cli

import (
	"fmt"
	"strings"

	"github.com/cloudcloud/nottify/nottify"
)

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
	err error
	val string
)

func configRun(args []string) {
	if len(args) == 0 {
		errorf("No config action provided.\nPlease run 'nottify help config'\n")
	}

	if len(args) > 3 {
		errorf("Too many arguments provided.\nPlease run 'nottify help config'\n")
	}

	conf := nottify.NewConfig()
	t := strings.ToLower(args[0])

	s := strings.Split(strings.ToLower(args[1]), ".")
	if t == "get" {
		val, err = conf.Get(s)
	} else if t == "set" {
		val, err = conf.Set(s, args[2])
	} else if t == "del" {
		val, err = conf.Del(s, args[2])
	} else {
		errorf("Invalid action.\nPlease run 'nottify help config'\n")
	}

	if err != nil {
		errorf("Unable to complete requested action.\n%s\n", err)
	}

	fmt.Printf("%s\n", val)
}
