// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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
		DurationSeconds:  aws.Int32(int32(c.AWS.RoleDurationSeconds)),
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
	content := fmt.Sprintf(
		`#!/bin/sh
export AWS_ACCESS_KEY_ID=%s
export AWS_SECRET_ACCESS_KEY=%s
export AWS_SESSION_TOKEN=%s
export AWS_DEFAULT_REGION=%s`, creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken, c.AWS.Region)

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
