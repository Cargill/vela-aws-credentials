package plugin

import (
	"flag"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func TestPlugin_FromCLIContext(t *testing.T) {
	// setup types
	flags := flag.NewFlagSet("test", 0)
	flags.String(FlagAudience, "sts.amazonaws.com", "doc")
	flags.String(FlagLogFormat, "json", "doc")
	flags.String(FlagLogLevel, "info", "doc")
	flags.String(FlagScriptFormat, ScriptFormatShell, "doc")
	flags.String(FlagScriptPath, "/path/to/script", "doc")
	flags.Bool(FlagScriptWrite, true, "doc")
	flags.Bool(FlagVerify, true, "doc")

	flags.String(FlagAWSRegion, "us-east-1", "doc")
	flags.String(FlagAWSRole, "testRole", "doc")
	flags.Int(FlagAWSRoleDurationSeconds, 3600, "doc")
	flags.String(FlagAWSRoleSessionName, "testSession", "doc")
	flags.String(FlagAWSInlineSessionPolicy, "{}", "doc")
	flags.String(FlagAWSManagedSessionPolicies, "[arn:aws:iam::aws:policy/ReadOnlyAccess]", "doc")

	flags.Int(FlagVelaBuildNumber, 1234, "doc")
	flags.String(FlagVelaRepoName, "testRepo", "doc")
	flags.String(FlagVelaOrgName, "testOrg", "doc")
	flags.String(FlagVelaIDTokenRequestToken, "testToken", "doc")
	flags.String(FlagVelaIDTokenRequestURL, "http://vela.example.com", "doc")

	// setup tests
	tests := []struct {
		name    string
		context *cli.Context
		want    bool
	}{
		{
			name:    "success",
			context: cli.NewContext(&cli.App{Name: "testing"}, flags, nil),
			want:    true,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FromCLIContext(test.context, logrus.NewEntry(logrus.StandardLogger()))
			if !reflect.DeepEqual(got != nil, test.want) {
				t.Errorf("FromCLIContext for %s is %v want %v", test.name, got != nil, test.want)
			}
		})
	}
}
