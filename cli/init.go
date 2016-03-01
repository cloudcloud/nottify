package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cloudcloud/nottify/nottify"
)

var cmdInit = &Command{
	UsageLine: "init",
	Short:     "Interactive initialisation",
	Long: `
`,
}

var (
	reader *bufio.Reader
	config *nottify.Config
	db     *nottify.Db
)

func init() {
	cmdInit.Run = initRun

	config = nottify.NewConfig()
}

func initRun(args []string) {
	reader = bufio.NewReader(os.Stdin)

	// init all
	for {
		fmt.Print("Do you wish to setup the database? [Y/n] ")

		t, _ := reader.ReadString('\n')
		t = strings.TrimSpace(t)

		if t == "Y" {
			if err := setupConfiguration(); err != nil {
				// derp
				errorf("Unable to complete Configuration setup [%s]", err.Error())
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

func setupConfiguration() error {
	// prompt for individual sets
	current, _ := config.Get([]string{"database", "hostname"})
	fmt.Printf("Please enter the Database Hostname [%s]: ", current)
	p, _ := reader.ReadString('\n')
	p = strings.TrimSpace(p)
	if len(p) > 0 {
		_, err := config.Set([]string{"database", "hostname"}, p)
		if err != nil {
			return err
		}
	}

	current, _ = config.Get([]string{"database", "user"})
	fmt.Printf("Please enter the Database User [%s]: ", current)
	p, _ = reader.ReadString('\n')
	p = strings.TrimSpace(p)
	if len(p) > 0 {
		_, err := config.Set([]string{"database", "user"}, p)
		if err != nil {
			return err
		}
	}

	// this will be more elegant in the future, of course
	fmt.Print("Please enter the Database Password: ")
	p, _ = reader.ReadString('\n')
	p = strings.TrimSpace(p)
	if len(p) > 0 {
		_, err := config.Set([]string{"database", "password"}, p)
		if err != nil {
			return err
		}
	}

	current, _ = config.Get([]string{"database", "database"})
	fmt.Printf("Please enter the Database Name [%s]: ", current)
	p, _ = reader.ReadString('\n')
	p = strings.TrimSpace(p)
	if len(p) > 0 {
		_, err := config.Set([]string{"database", "database"}, p)
		if err != nil {
			return err
		}
	}

	current, _ = config.Get([]string{"database", "table_prefix"})
	fmt.Printf("Please enter the Database Table Prefix [%s]: ", current)
	p, _ = reader.ReadString('\n')
	p = strings.TrimSpace(p)
	if len(p) > 0 {
		_, err := config.Set([]string{"database", "table_prefix"}, p)
		if err != nil {
			return err
		}
	}

	for {
		fmt.Print("Do you wish to setup the database tables? [Y/n] ")
		p, _ = reader.ReadString('\n')
		p = strings.TrimSpace(p)

		if p == "Y" {
			if err := setupTables(); err != nil {
				return err
			} else {
				break
			}
		} else if p == "n" {
			break
		}
	}

	return nil
}

func setupTables() error {
	db = nottify.NewDb(config)
	if err = db.RunFile("v1_create_tables.sql"); err != nil {
		return err
	}

	return nil
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
		str, err := config.Set([]string{"dirs"}, p)
		if err != nil {
			errorf(err.Error())
		}

		fmt.Println(str)
	}

	return true
}
