// Package core provides the shared, integral parts within Nottify.
package core

import (
	"github.com/cloudcloud/nottify/v1/config"
)

// CLI will prepare for a CLI session.
func CLI() {
	c, err := config.FromFile("")
	if err != nil {
		panic(err)
	}

	_ = c
}
