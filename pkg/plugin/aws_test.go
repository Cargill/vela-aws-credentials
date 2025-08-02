// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestConfig_WriteCreds(t *testing.T) {
	type args struct {
		creds        *aws.Credentials
		scriptFormat string
		appendConfig *bool
		profileName  *string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "shell",
			args: args{
				creds: &aws.Credentials{
					AccessKeyID:     "ACCESS_KEY_ID",
					SecretAccessKey: "SECRET_ACCESS_KEY",
					SessionToken:    "SESSION_TOKEN",
				},
				scriptFormat: "shell",
			},
			want:    "testdata/script.shell",
			wantErr: false,
		},
		{
			name: "credential_file",
			args: args{
				creds: &aws.Credentials{
					AccessKeyID:     "ACCESS_KEY_ID",
					SecretAccessKey: "SECRET_ACCESS_KEY",
					SessionToken:    "SESSION_TOKEN",
				},
				scriptFormat: "credential_file",
			},
			want:    "testdata/script.credential_file",
			wantErr: false,
		},
		{
			name: "credential_file_with_custom_profile",
			args: args{
				creds: &aws.Credentials{
					AccessKeyID:     "ACCESS_KEY_ID",
					SecretAccessKey: "SECRET_ACCESS_KEY",
					SessionToken:    "SESSION_TOKEN",
				},
				scriptFormat: "credential_file",
				profileName:  stringPtr("production"),
			},
			want:    "testdata/script.credential_file_custom_profile",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scriptPath := filepath.Join(t.TempDir(), fmt.Sprintf("script.%s", tt.args.scriptFormat))
			c := &Config{
				ScriptPath:   scriptPath,
				ScriptFormat: tt.args.scriptFormat,
				AppendConfig: tt.args.appendConfig,
				ProfileName:  tt.args.profileName,
				AWS: &AWS{
					Region: "us-east-1",
				},
			}

			err := c.WriteCreds(tt.args.creds)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteCreds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			expected, err := os.ReadFile(tt.want)
			assert.NoError(t, err)
			assert.NotEqual(t, 0, len(expected), "expected golden script.sh file")

			got, err := os.ReadFile(scriptPath)
			assert.NoError(t, err)

			if diff := cmp.Diff(string(expected), string(got)); diff != "" {
				t.Errorf("WriteCreds() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConfig_WriteCreds_AppendConfig(t *testing.T) {
	tests := []struct {
		name         string
		existingFile string
		profileName  string
		expectedFile string
		wantErr      bool
	}{
		{
			name:         "append_to_nonexistent_file",
			existingFile: "",
			profileName:  "default",
			expectedFile: "testdata/script.credential_file",
			wantErr:      false,
		},
		{
			name: "append_new_profile",
			existingFile: `[existing]
aws_access_key_id=EXISTING_KEY
aws_secret_access_key=EXISTING_SECRET
aws_session_token=EXISTING_TOKEN`,
			profileName:  "default",
			expectedFile: "testdata/script.credential_file_append_new",
			wantErr:      false,
		},
		{
			name: "fail_when_profile_exists",
			existingFile: `[default]
aws_access_key_id=OLD_KEY
aws_secret_access_key=OLD_SECRET
aws_session_token=OLD_TOKEN`,
			profileName: "default",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			scriptPath := filepath.Join(tempDir, "credentials")

			// Create existing file if specified
			if tt.existingFile != "" {
				err := os.WriteFile(scriptPath, []byte(tt.existingFile), 0600)
				assert.NoError(t, err)
			}

			c := &Config{
				ScriptPath:   scriptPath,
				ScriptFormat: "credential_file",
				AppendConfig: boolPtr(true),
				ProfileName:  stringPtr(tt.profileName),
				AWS: &AWS{
					Region: "us-east-1",
				},
			}

			creds := &aws.Credentials{
				AccessKeyID:     "ACCESS_KEY_ID",
				SecretAccessKey: "SECRET_ACCESS_KEY",
				SessionToken:    "SESSION_TOKEN",
			}

			err := c.WriteCreds(creds)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			// Read expected content
			expected, err := os.ReadFile(tt.expectedFile)
			assert.NoError(t, err)

			// Read actual content
			got, err := os.ReadFile(scriptPath)
			assert.NoError(t, err)

			if diff := cmp.Diff(string(expected), string(got)); diff != "" {
				t.Errorf("WriteCreds() append mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
