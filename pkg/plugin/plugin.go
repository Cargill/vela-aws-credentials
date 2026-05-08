// SPDX-License-Identifier: Apache-2.0

package plugin

import (
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
		Logger       *logrus.Entry
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
	c.Logger.Debug("running plugin with provided configuration")

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

	c.Logger.Debug("plugin finished...")

	return nil
}
