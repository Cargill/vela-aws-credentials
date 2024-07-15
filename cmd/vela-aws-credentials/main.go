// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"time"

	"github.com/Cargill/vela-aws-credentials/pkg/plugin"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

var (
	version = "dev"
)

func main() {
	// create new CLI application
	app := cli.NewApp()

	// Config Information

	app.Name = "vela-aws-credentials"
	app.HelpName = "vela-aws-credentials"
	app.Usage = "Vela AWS credentials plugin for temporary AWS credentials."

	// Config Metadata

	app.Action = run
	app.Compiled = time.Now()
	app.Version = version

	// Config Flags

	app.Flags = []cli.Flag{

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_LEVEL", "AWS_CREDENTIALS_LOG_LEVEL"},
			FilePath: "/vela/parameters/aws-credentials/log_level,/vela/secrets/aws-credentials/log_level",
			Name:     "log.level",
			Usage:    "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:    "info",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_REGION", "AWS_CREDENTIALS_REGION"},
			FilePath: "/vela/parameters/aws-credentials/region,/vela/secrets/aws-credentials/region",
			Name:     "aws.region",
			Usage:    "AWS region to use for assume role",
			Value:    "us-east-1",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_ROLE", "AWS_CREDENTIALS_ROLE"},
			FilePath: "/vela/parameters/aws-credentials/role,/vela/secrets/aws-credentials/role",
			Name:     "aws.role",
			Usage:    "AWS IAM role to assume",
		},
		&cli.IntFlag{
			EnvVars:  []string{"PARAMETER_ROLE_DURATION_SECONDS", "AWS_CREDENTIALS_ROLE_DURATION_SECONDS"},
			FilePath: "/vela/parameters/aws-credentials/role_duration_seconds,/vela/secrets/aws-credentials/role_duration_seconds",
			Name:     "aws.role_duration_seconds",
			Usage:    "Role duration in seconds",
			Value:    3600,
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_ROLE_SESSION_NAME", "AWS_CREDENTIALS_ROLE_SESSION_NAME"},
			FilePath: "/vela/parameters/aws-credentials/role_session_name,/vela/secrets/aws-credentials/role_session_name",
			Name:     "aws.role_session_name",
			Usage:    "Role session name",
			Value:    "vela",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_AUDIENCE", "AWS_CREDENTIALS_AUDIENCE"},
			FilePath: "/vela/parameters/aws-credentials/audience,/vela/secrets/aws-credentials/audience",
			Name:     "audience",
			Usage:    "Audience to use for the OIDC provider",
			Value:    "sts.amazonaws.com",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_INLINE_SESSION_POLICY", "AWS_CREDENTIALS_INLINE_SESSION_POLICY"},
			FilePath: "/vela/parameters/aws-credentials/inline_session_policy,/vela/secrets/aws-credentials/inline_session_policy",
			Name:     "aws.inline_session_policy",
			Usage:    "Inline session policy to use when assuming the role",
		},
		&cli.StringSliceFlag{
			EnvVars:  []string{"PARAMETER_MANAGED_SESSION_POLICIES", "AWS_CREDENTIALS_MANAGED_SESSION_POLICIES"},
			FilePath: "/vela/parameters/aws-credentials/managed_session_policies,/vela/secrets/aws-credentials/managed_session_policies",
			Name:     "aws.managed_session_policies",
			Usage:    "list of managed session policies to use when assuming the role",
		},

		// vela flags
		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD_NUMBER", "BUILD_NUMBER"},
			Name:    "vela.build_number",
			Usage:   "environment variable reference for reading in build number",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO_NAME", "REPOSITORY_NAME"},
			Name:    "vela.repo_name",
			Usage:   "environment variable reference for reading in repository name",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO_ORG", "REPOSITORY_ORG"},
			Name:    "vela.org_name",
			Usage:   "environment variable reference for reading in repository org",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ID_TOKEN_REQUEST_TOKEN"},
			Name:    "vela.id_token_request_token",
			Usage:   "environment variable reference for reading in OIDC request token",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ID_TOKEN_REQUEST_URL"},
			Name:    "vela.id_token_request_url",
			Usage:   "environment variable reference for reading in OIDC request token URL",
		},

		&cli.BoolFlag{
			EnvVars: []string{"PARAMETER_VERIFY", "AWS_CREDENTIALS_VERIFY"},
			Name:    "verify",
			Usage:   "if the AWS credentials should be validated",
		},
		&cli.StringFlag{
			EnvVars: []string{"PARAMETER_SCRIPT_PATH", "AWS_CREDENTIALS_SCRIPT_PATH"},
			Name:    "script_path",
			Usage:   "path where to write script that contains AWS credentials",
			Value:   "/vela/secrets/aws/setup.sh",
		},
		&cli.BoolFlag{
			EnvVars: []string{"PARAMETER_SCRIPT_WRITE", "AWS_CREDENTIALS_SCRIPT_WRITE"},
			Name:    "script_write",
			Usage:   "if the credentials script should be created",
			Value:   false,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(c *cli.Context) error {
	// set the log level for the plugin
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.WithFields(logrus.Fields{
		"code":     "https://github.com/Cargill/vela-aws-credentials",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/pipeline/aws-credentials",
		"registry": "https://hub.docker.com/r/cargill/vela-aws-credentials",
	}).Info("Vela AWS Credentials Config")

	// create the plugin
	p := &plugin.Config{
		Audience:    c.String("audience"),
		Verify:      c.Bool("verify"),
		ScriptPath:  c.String("script_path"),
		ScriptWrite: c.Bool("script_write"),
		AWS: &plugin.AWS{
			Region:                 c.String("aws.region"),
			Role:                   c.String("aws.role"),
			RoleDurationSeconds:    c.Int("aws.role_duration_seconds"),
			RoleSessionName:        c.String("aws.role_session_name"),
			InlineSessionPolicy:    c.String("aws.inline_session_policy"),
			ManagedSessionPolicies: c.StringSlice("aws.managed_session_policies"),
		},
		Vela: &plugin.Vela{
			BuildNumber:     c.Int("vela.build_number"),
			RepoName:        c.String("vela.repo_name"),
			OrgName:         c.String("vela.org_name"),
			RequestToken:    c.String("vela.id_token_request_token"),
			RequestTokenURL: c.String("vela.id_token_request_url"),
		},
	}

	// validate the plugin
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
