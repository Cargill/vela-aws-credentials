package plugin

const (
	// FlagAudience represents the name of the flag for setting the OIDC provider audience for the plugin.
	FlagAudience = "audience"
	// FlagLogFormat represents the name of the flag for setting the log format for the plugin.
	FlagLogFormat = "log.format"
	// FlagLogLevel represents the name of the flag for setting the log level for the plugin.
	FlagLogLevel = "log.level"
	// FlagScriptFormat represents the name of the flag for setting the format of the AWS credentials script for the plugin.
	FlagScriptFormat = "script_format"
	// FlagScriptPath represents the name of the flag for setting the path to write the AWS credentials script for the plugin.
	FlagScriptPath = "script_path"
	// FlagScriptWrite represents the name of the flag for setting whether to write the AWS credentials script for the plugin.
	FlagScriptWrite = "script_write"
	// FlagVerify represents the name of the flag for setting whether to validate the AWS credentials for the plugin.
	FlagVerify = "verify"

	// AWS Configuration Flags

	// FlagAWSInlineSessionPolicy represents the name of the flag for setting the AWS inline session policy for the plugin.
	FlagAWSInlineSessionPolicy = "aws.inline_session_policy"
	// FlagAWSManagedSessionPolicies represents the name of the flag for setting the AWS managed session policies for the plugin.
	FlagAWSManagedSessionPolicies = "aws.managed_session_policies"
	// FlagAWSRegion represents the name of the flag for setting the AWS region for the plugin.
	FlagAWSRegion = "aws.region"
	// FlagAWSRole represents the name of the flag for setting the AWS IAM role to assume for the plugin.
	FlagAWSRole = "aws.role"
	// FlagAWSRoleDurationSeconds represents the name of the flag for setting the duration in seconds for assuming the AWS IAM role for the plugin.
	FlagAWSRoleDurationSeconds = "aws.role_duration_seconds"
	// FlagAWSRoleSessionName represents the name of the flag for setting the session name when assuming the AWS IAM role for the plugin.
	FlagAWSRoleSessionName = "aws.role_session_name"

	// Vela Configuration Flags

	// FlagVelaBuildNumber represents the name of the flag for capturing the build number from Vela for the plugin.
	FlagVelaBuildNumber = "vela.build_number"
	// FlagVelaIDTokenRequestToken represents the name of the flag for capturing the OIDC request token from Vela for the plugin.
	FlagVelaIDTokenRequestToken = "vela.id_token_request_token"
	// FlagVelaIDTokenRequestURL represents the name of the flag for capturing the OIDC request token URL from Vela for the plugin.
	FlagVelaIDTokenRequestURL = "vela.id_token_request_url"
	// FlagVelaOrgName represents the name of the flag for capturing the organization name from Vela for the plugin.
	FlagVelaOrgName = "vela.org_name"
	// FlagVelaRepoName represents the name of the flag for capturing the repository name from Vela for the plugin.
	FlagVelaRepoName = "vela.repo_name"
)
