package plugin

import (
	"fmt"
	"slices"
)

// Validate function to validate plugin configuration.
func (c *Config) Validate() error {
	c.Logger.Debug("validating plugin configuration")

	// validate that a webhook was supplied
	if len(c.AWS.Role) == 0 {
		return fmt.Errorf("no role provided")
	}

	if c.AWS.RoleDurationSeconds == 0 {
		return fmt.Errorf("no role duration provided")
	}

	if c.Vela.RequestTokenURL == "" {
		return fmt.Errorf("no request token url provided")
	}

	supportedFormats := []string{"shell", "credential_file"}
	if !slices.Contains(supportedFormats, c.ScriptFormat) {
		return fmt.Errorf("only script formats of %s are supported", supportedFormats)
	}

	if c.ScriptPath == "" {
		switch c.ScriptFormat {
		case "shell":
			c.ScriptPath = "/vela/secrets/aws/setup.sh"
		case "credential_file":
			c.ScriptPath = "/vela/secrets/aws/creds"
		}
	}

	if c.Vela.RequestToken == "" {
		return fmt.Errorf("no request token provided - make sure you have set `id_request: yes` in the step")
	}

	return nil
}
