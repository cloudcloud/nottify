package nottify

import (
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config defines the YAML file format content
type Config struct {
	Dirs     []string `yaml:",flow"`
	Hostname string   `yaml:"hostname"`
	Port     int      `yaml:"port"`
	Database struct {
		Hostname    string `yaml:"hostname"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Database    string `yaml:"database"`
		TablePrefix string `yaml:"table_prefix"`
	}
}

var (
	filename string
	err      error
)

// NewConfig provisions an instance of Config
func NewConfig() *Config {
	c := new(Config)

	if len(filename) < 1 {
		filename, err = discoverConfigFile()
		if err != nil {
			// this needs elegance, pls
			panic(err)
		}
	}

	// load / parse the yaml from filename
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// more elegance
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &c)

	if err != nil {
		// again, elegance
		panic(err)
	}

	return c
}

// Get will retrieve a value from the Configuration
func (c *Config) Get(s []string) (string, error) {
	// first section
	main := s[0]

	// manual time, reflection is nuts
	switch main {
	case "dirs":
		return strings.Join(c.Dirs, ", "), nil
	case "hostname":
		return c.Hostname, nil
	case "port":
		return strconv.FormatInt(int64(c.Port), 10), nil
	case "database":
		if len(s) < 2 {
			return "", errors.New("Database requires a Sub Index to get")
		}

		// database has sub items
		sub := s[1]
		switch sub {
		case "hostname":
			return c.Database.Hostname, nil
		case "user":
			return c.Database.User, nil
		case "password":
			return c.Database.Password, nil
		case "database":
			return c.Database.Database, nil
		case "table_prefix", "tableprefix":
			return c.Database.TablePrefix, nil
		}
	}

	return "", errors.New("Unable to get configuration entry")
}

// Set will change the value of an existing Configuration Directive
func (c *Config) Set(s []string, v interface{}) (string, error) {
	// outer section
	main := s[0]

	// perhaps reflection in the future
	switch main {
	case "dirs":
		c.Dirs = append(c.Dirs, v.(string))
	case "hostname":
		c.Hostname = v.(string)
	case "port":
		c.Port = v.(int)
	case "database":
		if len(s) < 2 {
			return "", errors.New("Database requires a Sub Index to set")
		}

		// now the subs
		sub := s[1]
		switch sub {
		case "hostname":
			c.Database.Hostname = v.(string)
		case "user":
			c.Database.User = v.(string)
		case "password":
			c.Database.Password = v.(string)
		case "database":
			c.Database.Database = v.(string)
		case "table_prefix", "tableprefix":
			c.Database.TablePrefix = v.(string)
		default:
			return "", errors.New("Unknown Database Sub Index for Setting")
		}
	default:
		return "", errors.New("Unknown configuration directive to set")
	}

	return c.push()
}

// Del will remove an entry from a list, or reset the default
func (c *Config) Del(s []string, d string) (string, error) {
	// outer section
	main := s[0]

	// again, need to investigate reflection for something pragmatic
	switch main {
	case "dirs":
		// trimFromList()
		var n []string
		for i := 0; i < len(c.Dirs); i++ {
			if c.Dirs[i] != d {
				n = append(n, c.Dirs[i])
			}
		}
		c.Dirs = n
	case "hostname":
		c.Hostname = "localhost"
	case "port":
		c.Port = 8080
	case "database":
		if len(s) < 2 {
			return "", errors.New("Database requires a Sub Item to Delete")
		}

		// and the subs
		sub := s[1]
		switch sub {
		case "hostname":
			c.Database.Hostname = "localhost"
		case "user":
			c.Database.User = "nottify"
		case "password":
			c.Database.Password = ""
		case "database":
			c.Database.Database = "nottify"
		case "table_prefix", "tableprefix":
			c.Database.TablePrefix = "nott"
		default:
			return "", errors.New("Unknown Database Sub Index for Delete")
		}
	default:
		return "", errors.New("Unknown configuration directive to Delete")
	}

	return c.push()
}

func (c *Config) push() (string, error) {
	// transform back to yaml
	y, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}

	// push back to the file
	err = ioutil.WriteFile(filename, y, 0755)
	return fmt.Sprintf("Written to file %s\n", filename), nil
}

func discoverConfigFile() (string, error) {
	// do some dir traversal
	gopath := build.Default.GOPATH
	filename = path.Join(gopath, "src/github.com/cloudcloud/nottify/conf.yml")
	tmpFile := ""

	if _, err := os.Stat(filename); err != nil {
		// grab the example and copy it across
		tmpFile = filename + ".example"
		if _, err := os.Stat(tmpFile); err != nil {
			return "", errors.New("Unable to discover config.yml file")
		}

		// using tmpFile as the configuration file now
		if err = os.Link(tmpFile, filename); err != nil {
			return "", err
		}
	}

	return filename, nil
}
