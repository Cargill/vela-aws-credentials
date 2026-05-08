package plugin

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// FromCLIContext creates and returns a plugin from the urfave/cli context.
func FromCLIContext(ctx *cli.Context, logger *logrus.Entry) *Config {
	return &Config{
		Logger:       logger,
		Audience:     ctx.String(FlagAudience),
		ScriptPath:   ctx.String(FlagScriptPath),
		ScriptFormat: ctx.String(FlagScriptFormat),
		ScriptWrite:  ctx.Bool(FlagScriptWrite),
		Verify:       ctx.Bool(FlagVerify),
		AWS: &AWS{
			Region:                 ctx.String(FlagAWSRegion),
			Role:                   ctx.String(FlagAWSRole),
			RoleDurationSeconds:    ctx.Int(FlagAWSRoleDurationSeconds),
			RoleSessionName:        ctx.String(FlagAWSRoleSessionName),
			InlineSessionPolicy:    ctx.String(FlagAWSInlineSessionPolicy),
			ManagedSessionPolicies: ctx.StringSlice(FlagAWSManagedSessionPolicies),
		},
		Vela: &Vela{
			BuildNumber:     ctx.Int(FlagVelaBuildNumber),
			RepoName:        ctx.String(FlagVelaRepoName),
			OrgName:         ctx.String(FlagVelaOrgName),
			RequestToken:    ctx.String(FlagVelaIDTokenRequestToken),
			RequestTokenURL: ctx.String(FlagVelaIDTokenRequestURL),
		},
	}
}
