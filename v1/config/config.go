// Package config provides all configuration for Nottify, including what may come
// through in the CLI or what can be found on the file system. This also includes
// those options propogated and managed through the database.
//
// Configuration is initially managed through a file in the local directory,
// provided by the ``-c`` command line flag. This filename by default is
// ``.nottify.json`` and will need to be provided for each Nottify execution
// if an alternative is preferred.
package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	Error   = "error"
	Info    = "info"
	Message = "message"
	Warn    = "warn"
)

// Config is a behaviour to provide ease of configuration access.
type Config interface {
	GetDsn() string
	GetFilename() string

	D(string)
	O(string, string)
}

// BaseConfig provides structure around keeping configuration data.
type BaseConfig struct {
	Debug    bool     `json:"debug"`
	Dirs     []string `json:",flow"`
	Database Database `json:"database"`
	Filename string   `json:"-"`

	o io.Writer
}

// Database contains connection information for the database.
type Database struct {
	Database    string `json:"database"`
	Hostname    string `json:"hostname"`
	Password    string `json:"password"`
	Port        int    `json:"port"`
	TablePrefix string `json:"table_prefix"`
	User        string `json:"user"`
}

// GetDsn will provide the standard database connection string.
func (c *BaseConfig) GetDsn() string {
	return fmt.Sprintf(
		"%s:%s@%s:%d/%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Hostname,
		c.Database.Port,
		c.Database.Database,
	)
}

// GetFilename will give the current location for this configuration.
func (c *BaseConfig) GetFilename() string {
	return c.Filename
}

// D will print a debug message.
func (c *BaseConfig) D(m string) {
	if c.Debug {
		fmt.Fprintf(
			c.o,
			"nottify,debug,%v,%s\n",
			time.Now().Format("2006-02-01_03:04:05PM_MST"),
			m,
		)
	}
}

// O will print a message out.
func (c *BaseConfig) O(t, m string) {
	if t == Message {
		fmt.Fprintln(c.o, m)

		return
	}

	fmt.Fprintf(
		c.o,
		"nottify,%s,%v,'%s'\n",
		t,
		time.Now().Format("2006-02-01_03:04:05.000PM_MST"),
		strings.Replace(m, "'", "\\'", -1),
	)
}

// FromFile will use the provided file to instantiate configuration.
func FromFile(f string, out io.Writer) (Config, error) {
	if len(f) < 1 {
		f = fmt.Sprintf(
			"%s%s%s",
			os.Getenv("HOME"),
			string(os.PathSeparator),
			".nottify.json",
		)
	}

	var err error
	if _, err = os.Stat(f); err != nil {
		c := defaultConfig(out)
		c.Filename = f
		err = writeConfig(c)

		c.D("Written base Config")
		return c, err
	}

	h, _ := os.Open(f)
	c := &BaseConfig{}
	err = json.NewDecoder(h).Decode(c)

	c.o = out
	c.D("Loaded base Config")

	return c, err
}

func defaultConfig(out io.Writer) *BaseConfig {
	return &BaseConfig{
		Dirs: []string{"/opt/music/"},
		Database: Database{
			Database: "nottify",
			Hostname: "localhost",
		},
		o: out,
	}
}

func writeConfig(c Config) error {
	// json encode and write to file
	j, err := json.Marshal(c)
	if err != nil {
		return err
	}

	c.D(fmt.Sprintf("Writing data to [%s]: %s", c.GetFilename(), j))
	return ioutil.WriteFile(c.GetFilename(), j, 0755)
}
