// Package config provides all configuration for Nottify, including what may come
// through in the CLI or what can be found on the file system. This also includes
// those options propogated and managed through the database.
package config

// Config gives actions to be taken on all Configuration.
type Config interface {
}

// BaseConfig houses structure for all configuration options.
type BaseConfig struct {
}
