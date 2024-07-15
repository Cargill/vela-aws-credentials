// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"net/url"

	"github.com/go-vela/sdk-go/vela"
)

func (c *Config) GenerateVelaToken() (string, error) {
	tokenURL, err := url.Parse(c.Vela.RequestTokenURL)
	if err != nil {
		return "", err
	}

	client, err := vela.NewClient(fmt.Sprintf("https://%s", tokenURL.Hostname()), "vela", nil)
	if err != nil {
		return "", err
	}

	client.Authentication.SetTokenAuth(c.Vela.RequestToken)

	opt := &vela.IDTokenOptions{Audience: []string{c.Audience}}

	token, _, err := client.Build.GetIDToken(c.Vela.OrgName, c.Vela.RepoName, c.Vela.BuildNumber, opt)
	if err != nil {
		return "", err
	}

	return token.GetToken(), nil
}
