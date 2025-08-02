// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/sirupsen/logrus"
)

func (c *Config) AssumeRole(token string) (*aws.Credentials, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(c.AWS.Region))
	if err != nil {
		return nil, err
	}

	// create an STS client
	stsClient := sts.NewFromConfig(cfg)

	var managedPolicies []types.PolicyDescriptorType
	for _, policy := range c.AWS.ManagedSessionPolicies {
		managedPolicies = append(managedPolicies, types.PolicyDescriptorType{Arn: aws.String(policy)})
	}

	input := &sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          aws.String(c.AWS.Role),
		RoleSessionName:  aws.String(c.AWS.RoleSessionName),
		WebIdentityToken: aws.String(token),
		//nolint:gosec // disable G115
		DurationSeconds: aws.Int32(int32(c.AWS.RoleDurationSeconds)),
	}

	if c.AWS.InlineSessionPolicy != "" {
		input.Policy = aws.String(c.AWS.InlineSessionPolicy)
	}

	if len(managedPolicies) > 0 {
		input.PolicyArns = managedPolicies
	}

	// Perform the AssumeRoleWithWebIdentity request
	assumeRoleOutput, err := stsClient.AssumeRoleWithWebIdentity(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to assume role: %w", err)
	}

	creds := aws.Credentials{
		AccessKeyID:     *assumeRoleOutput.Credentials.AccessKeyId,
		SecretAccessKey: *assumeRoleOutput.Credentials.SecretAccessKey,
		SessionToken:    *assumeRoleOutput.Credentials.SessionToken,
	}

	if c.Verify {
		tempCfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(credentials.StaticCredentialsProvider{Value: creds}), config.WithRegion(c.AWS.Region))
		if err != nil {
			return nil, err
		}

		tempClient := sts.NewFromConfig(tempCfg)

		_, err = tempClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
		if err != nil {
			return nil, err
		}

		logrus.Infof("successfully validated credentials for %s", c.AWS.Role)
	}

	return &creds, nil
}

func (c *Config) WriteCreds(creds *aws.Credentials) error {
	switch c.ScriptFormat {
	case "shell":
		return c.writeEnvVarsFile(creds)
	case "credential_file":
		return c.writeCredsFile(creds)
	}

	return fmt.Errorf("unsupported script format: %s", c.ScriptFormat)
}

func (c *Config) writeFile(content string) error {
	if _, err := os.Stat(c.ScriptPath); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(c.ScriptPath), 0700)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(c.ScriptPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) writeEnvVarsFile(creds *aws.Credentials) error {
	content := fmt.Sprintf(
		`#!/bin/sh
export AWS_ACCESS_KEY_ID=%s
export AWS_SECRET_ACCESS_KEY=%s
export AWS_SESSION_TOKEN=%s
export AWS_DEFAULT_REGION=%s`, creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken, c.AWS.Region)
	
	return c.writeFile(content)
}

func (c *Config) writeCredsFile(creds *aws.Credentials) error {
	// Determine profile name
	profileName := "default"
	if c.ProfileName != nil {
		profileName = *c.ProfileName
	}

	// Create new profile content
	newProfileContent := fmt.Sprintf(`[%s]
aws_access_key_id=%s
aws_secret_access_key=%s
aws_session_token=%s`, profileName, creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken)

	// If not appending, just write the new content
	if c.AppendConfig == nil || !*c.AppendConfig {
		return c.writeFile(newProfileContent)
	}

	// Handle append mode
	return c.appendCredentialFile(profileName, newProfileContent)
}

func (c *Config) appendCredentialFile(profileName, newProfileContent string) error {
	// Ensure directory exists
	if _, err := os.Stat(c.ScriptPath); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(c.ScriptPath), 0700)
		if err != nil {
			return err
		}
		// File doesn't exist, create it with the new content
		return c.writeFile(newProfileContent)
	}

	// Read existing file
	existingContent, err := os.ReadFile(c.ScriptPath)
	if err != nil {
		return err
	}

	// Check if profile already exists
	if c.profileExists(string(existingContent), profileName) {
		return fmt.Errorf("profile [%s] already exists in credentials file %s", profileName, c.ScriptPath)
	}

	// Append new profile
	var mergedContent string
	if strings.TrimSpace(string(existingContent)) == "" {
		mergedContent = newProfileContent
	} else {
		mergedContent = strings.TrimRight(string(existingContent), "\n") + "\n" + newProfileContent
	}

	return c.writeFile(mergedContent)
}

func (c *Config) profileExists(content, profileName string) bool {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == fmt.Sprintf("[%s]", profileName) {
			return true
		}
	}
	return false
}
