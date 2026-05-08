package plugin

import (
	"github.com/urfave/cli/v2"
)

var (
	// Flags represents all supported command line interface (CLI) flags for the plugin.
	Flags = []cli.Flag{
		// Plugin Configuration Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_AUDIENCE", "AWS_CREDENTIALS_AUDIENCE"},
			FilePath: "/vela/parameters/aws-credentials/audience,/vela/secrets/aws-credentials/audience",
			Name:     FlagAudience,
			Usage:    "Audience to use for the OIDC provider",
			Value:    "sts.amazonaws.com",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_FORMAT", "AWS_CREDENTIALS_LOG_FORMAT"},
			FilePath: "/vela/parameters/aws-credentials/log_format,/vela/secrets/aws-credentials/log_format",
			Name:     FlagLogFormat,
			Usage:    "set log format - options: (text|json)",
			Value:    "text",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_LEVEL", "AWS_CREDENTIALS_LOG_LEVEL"},
			FilePath: "/vela/parameters/aws-credentials/log_level,/vela/secrets/aws-credentials/log_level",
			Name:     FlagLogLevel,
			Usage:    "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:    "info",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_SCRIPT_PATH", "AWS_CREDENTIALS_SCRIPT_PATH"},
			FilePath: "/vela/parameters/aws-credentials/script_path,/vela/secrets/aws-credentials/script_path",
			Name:     FlagScriptPath,
			Usage:    "path where to write script that contains AWS credentials",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_SCRIPT_FORMAT", "AWS_CREDENTIALS_SCRIPT_FORMAT"},
			FilePath: "/vela/parameters/aws-credentials/script_format,/vela/secrets/aws-credentials/script_format",
			Name:     FlagScriptFormat,
			Usage:    "format of AWS credentials script (shell or credential_file)",
			Value:    ScriptFormatShell,
		},
		&cli.BoolFlag{
			EnvVars: []string{"PARAMETER_SCRIPT_WRITE", "AWS_CREDENTIALS_SCRIPT_WRITE"},
			Name:    FlagScriptWrite,
			Usage:   "if the credentials script should be created",
			Value:   false,
		},
		&cli.BoolFlag{
			EnvVars: []string{"PARAMETER_VERIFY", "AWS_CREDENTIALS_VERIFY"},
			Name:    FlagVerify,
			Usage:   "if the AWS credentials should be validated",
		},

		// AWS Configuration Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_INLINE_SESSION_POLICY", "AWS_CREDENTIALS_INLINE_SESSION_POLICY"},
			FilePath: "/vela/parameters/aws-credentials/inline_session_policy,/vela/secrets/aws-credentials/inline_session_policy",
			Name:     FlagAWSInlineSessionPolicy,
			Usage:    "Inline session policy to use when assuming the role",
		},
		&cli.StringSliceFlag{
			EnvVars:  []string{"PARAMETER_MANAGED_SESSION_POLICIES", "AWS_CREDENTIALS_MANAGED_SESSION_POLICIES"},
			FilePath: "/vela/parameters/aws-credentials/managed_session_policies,/vela/secrets/aws-credentials/managed_session_policies",
			Name:     FlagAWSManagedSessionPolicies,
			Usage:    "list of managed session policies to use when assuming the role",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_REGION", "AWS_CREDENTIALS_REGION"},
			FilePath: "/vela/parameters/aws-credentials/region,/vela/secrets/aws-credentials/region",
			Name:     FlagAWSRegion,
			Usage:    "AWS region to use for assume role",
			Value:    "us-east-1",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_ROLE", "AWS_CREDENTIALS_ROLE"},
			FilePath: "/vela/parameters/aws-credentials/role,/vela/secrets/aws-credentials/role",
			Name:     FlagAWSRole,
			Usage:    "AWS IAM role to assume",
		},
		&cli.IntFlag{
			EnvVars:  []string{"PARAMETER_ROLE_DURATION_SECONDS", "AWS_CREDENTIALS_ROLE_DURATION_SECONDS"},
			FilePath: "/vela/parameters/aws-credentials/role_duration_seconds,/vela/secrets/aws-credentials/role_duration_seconds",
			Name:     FlagAWSRoleDurationSeconds,
			Usage:    "Role duration in seconds",
			Value:    3600,
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_ROLE_SESSION_NAME", "AWS_CREDENTIALS_ROLE_SESSION_NAME"},
			FilePath: "/vela/parameters/aws-credentials/role_session_name,/vela/secrets/aws-credentials/role_session_name",
			Name:     FlagAWSRoleSessionName,
			Usage:    "Role session name",
			Value:    "vela",
		},

		// Vela Configuration Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD_NUMBER", "BUILD_NUMBER"},
			Name:    FlagVelaBuildNumber,
			Usage:   "environment variable reference for reading in build number",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ID_TOKEN_REQUEST_TOKEN"},
			Name:    FlagVelaIDTokenRequestToken,
			Usage:   "environment variable reference for reading in OIDC request token",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ID_TOKEN_REQUEST_URL"},
			Name:    FlagVelaIDTokenRequestURL,
			Usage:   "environment variable reference for reading in OIDC request token URL",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO_ORG", "REPOSITORY_ORG"},
			Name:    FlagVelaOrgName,
			Usage:   "environment variable reference for reading in repository org",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO_NAME", "REPOSITORY_NAME"},
			Name:    FlagVelaRepoName,
			Usage:   "environment variable reference for reading in repository name",
		},
	}
)
