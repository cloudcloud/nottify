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
	"log"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"time"
)

const (
	ErrorLevels = iota
	Error
	Warn
	Info
	Debug
)

var (
	hostname = ""
)

func init() {
	cmd, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal("Unable to determine hostname")
	}

	hostname = strings.TrimSpace(string(cmd))
}

// Config is a behaviour to provide ease of configuration access.
type Config interface {
	GetDirs() []string
	GetDsn() string
	GetFilename() string
	KnownDir(string) bool

	D(string)
	O(int, string)
}

// BaseConfig provides structure around keeping configuration data.
type BaseConfig struct {
	Debug    bool     `json:"debug"`
	Dirs     []string `json:",flow"`
	Database Database `json:"database"`
	Filename string   `json:"-"`
	Level    int      `json:"level"`

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

type message struct {
	Agent string `json:"_agent"`

	Host         string    `json:"host"`
	Level        int       `json:"level"`
	LongMessage  string    `json:"long_message"`
	ShortMessage string    `json:"short_message"`
	Timestamp    time.Time `json:"timestamp"`
	Version      string    `json:"version"`
}

// GetDirs will provide all known directories from configuration.
func (c *BaseConfig) GetDirs() []string {
	return c.Dirs
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

// KnownDir determines if the provided directory string is already known in config.
func (c *BaseConfig) KnownDir(d string) bool {
	for _, x := range c.Dirs {
		if x == d {
			return true
		}
	}

	return false
}

// D will print a debug message.
func (c *BaseConfig) D(m string) {
	if c.Debug {

		c.O(Debug, m)
	}
}

// O will print a message out.
func (c *BaseConfig) O(t int, m string) {
	l := &message{
		Agent:        "nottify",
		Host:         hostname,
		ShortMessage: m,
		LongMessage:  "",
		Level:        t,
		Timestamp:    time.Now(),
		Version:      "1.1",
	}

	if t == Debug {
		l.LongMessage = string(debug.Stack())
	}

	o, err := json.Marshal(l)
	if err != nil {
		fmt.Fprintln(c.o, m)
		return
	}

	if c.Level >= t {
		fmt.Fprintln(c.o, string(o))
	}
}

// FromFile will use the provided file to instantiate configuration.
func FromFile(f string, out io.Writer, d bool) (*BaseConfig, error) {
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
		c.Debug = d
		err = writeConfig(c)

		c.D("Written base Config")
		return c, err
	}

	h, _ := os.Open(f)
	c := &BaseConfig{}
	err = json.NewDecoder(h).Decode(c)

	c.o = out
	c.Debug = d
	c.D("Loaded base Config")
	c.D(fmt.Sprintf("Dirs are %s", c.Dirs))

	return c, err
}

func defaultConfig(out io.Writer) *BaseConfig {
	return &BaseConfig{
		Debug: false,
		Dirs:  []string{"/opt/music/"},
		Database: Database{
			Database: "nottify",
			Hostname: "localhost",
		},
		Level: 2,
		o:     out,
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
