package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cmdInit = &Command{
	UsageLine: "init",
	Short:     "Interactive initialisation",
	Long: `
`,
}

var (
	reader *bufio.Reader
)

func init() {
	cmdInit.Run = initRun
}

func initRun(args []string) {
	reader = bufio.NewReader(os.Stdin)

	// init all
	for {
		fmt.Print("Do you wish to setup the database? [Y/n] ")

		t, _ := reader.ReadString('\n')
		t = strings.TrimSpace(t)

		if t == "Y" {
			if !setupConfiguration() {
				// derp
				errorf("Unable to complete Configuration setup")
			} else {
				break
			}
		} else if t == "n" {
			// do nothing
			break
		}
	}

	// run through options and request feedback
	for {
		fmt.Print("Do you wish to add a folder? [Y/n] ")

		t, _ := reader.ReadString('\n')
		t = strings.TrimSpace(t)

		if t == "Y" {
			if !setupDirs() {
				errorf("Unable to add Directory")
			} else {
				break
			}
		} else if t == "n" {
			break
		}
	}

	// done?
	fmt.Println("All done!")
}

func setupConfiguration() bool {
	return true
}

func setupTables() bool {
	userTable := `
`
	_ = userTable

	return true
}

func setupDirs() bool {
	for {
		fmt.Print("Please enter the full path (empty to stop): ")

		p, _ := reader.ReadString('\n')
		p = strings.TrimSpace(p)

		if len(p) < 1 {
			break
		}

		// push to configuration
	}

	return true
}
