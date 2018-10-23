// Package config provides all configuration for Nottify, including what may come
// through in the CLI or what can be found on the file system. This also includes
// those options propogated and managed through the database.
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is a behaviour to provide ease of configuration access.
type Config interface {
	GetDsn() string
}

// BaseConfig provides structure around keeping configuration data.
type BaseConfig struct {
	Dirs     []string `json:",flow"`
	Database struct {
		Database    string `json:"database"`
		Hostname    string `json:"hostname"`
		Password    string `json:"password"`
		Port        int    `json:"port"`
		TablePrefix string `json:"table_prefix"`
		User        string `json:"user"`
	} `json:"database"`
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

// FromFile will use the provided file to instantiate configuration.
func FromFile(f string) (Config, error) {
	if len(f) < 1 {
		f = "~/.nottify"
	}

	h, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	c := &BaseConfig{}
	d := json.NewDecoder(h)
	err = d.Decode(c)

	return c, err
}
