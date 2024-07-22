// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestConfig_WriteCreds(t *testing.T) {
	type args struct {
		creds *aws.Credentials
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				creds: &aws.Credentials{
					AccessKeyID:     "ACCESS_KEY_ID",
					SecretAccessKey: "SECRET_ACCESS_KEY",
					SessionToken:    "SESSION_TOKEN",
				},
			},
			want:    "testdata/script.sh",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scriptPath := filepath.Join(t.TempDir(), "script.sh")
			c := &Config{
				ScriptPath: scriptPath,
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
