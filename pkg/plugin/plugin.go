// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"slices"

	"github.com/sirupsen/logrus"
)

type (
	// Config struct represents fields user can present to plugin.
	Config struct {
		Audience     string
		Verify       bool
		ScriptPath   string
		ScriptFormat string
		ScriptWrite  bool
		AWS          *AWS
		Vela         *Vela
	}

	// AWS struct represents the config for the AWS role assumption.
	AWS struct {
		Region                 string
		Role                   string
		RoleDurationSeconds    int
		RoleSessionName        string
		InlineSessionPolicy    string
		ManagedSessionPolicies []string
	}

	// Vela struct represents the config for the Vela API calls.
	Vela struct {
		BuildNumber     int
		RepoName        string
		OrgName         string
		RequestToken    string
		RequestTokenURL string
	}
)

// Exec generates a set of temporary AWS credentials for later usage.
func (c *Config) Exec() error {
	logrus.Debug("running plugin with provided configuration")

	token, err := c.GenerateVelaToken()
	if err != nil {
		return err
	}

	creds, err := c.AssumeRole(token)
	if err != nil {
		return err
	}

	if c.ScriptWrite {
		err = c.WriteCreds(creds)
		if err != nil {
			return err
		}
	}

	logrus.Debug("plugin finished...")

	return nil
}

// Validate function to validate plugin configuration.
func (c *Config) Validate() error {
	logrus.Debug("validating plugin configuration")

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
			c.ScriptFormat = "/vela/secrets/aws/setup.sh"
		case "credential_file":
			c.ScriptFormat = "/vela/secrets/aws/creds"
		}
	}

	if c.Vela.RequestToken == "" {
		return fmt.Errorf("no request token provided - make sure you have set `id_request: yes` in the step")
	}

	return nil
}
