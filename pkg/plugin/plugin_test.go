// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "all fields are populated",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "shell",
			},
			wantErr: false,
		},
		{
			name: "AWS Role field is empty",
			config: &Config{
				AWS: &AWS{
					Role:                "",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "shell",
			},
			wantErr: true,
		},
		{
			name: "AWS RoleDurationSeconds field is 0",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 0,
				},
				Vela: &Vela{
					RequestToken: "testToken",
				},
				ScriptFormat: "shell",
			},
			wantErr: true,
		},
		{
			name: "Vela RequestTokenURL field is empty",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken: "testToken",
				},
				ScriptFormat: "shell",
			},
			wantErr: true,
		},
		{
			name: "Vela RequestToken field is empty",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "shell",
			},
			wantErr: true,
		},
		{
			name: "AppendConfig set with shell format should error",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "shell",
				AppendConfig: boolPtr(true),
			},
			wantErr: true,
		},
		{
			name: "ProfileName set with shell format should error",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "shell",
				ProfileName:  stringPtr("custom"),
			},
			wantErr: true,
		},
		{
			name: "AppendConfig and ProfileName set with credential_file format should pass",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "credential_file",
				AppendConfig: boolPtr(true),
				ProfileName:  stringPtr("custom"),
			},
			wantErr: false,
		},
		{
			name: "AppendConfig false with credential_file should pass",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "credential_file",
				AppendConfig: boolPtr(false),
			},
			wantErr: false,
		},
		{
			name: "AppendConfig and ProfileName nil should pass with any format",
			config: &Config{
				AWS: &AWS{
					Role:                "testRole",
					RoleDurationSeconds: 3600,
				},
				Vela: &Vela{
					RequestToken:    "testToken",
					RequestTokenURL: "http://127.0.0.1",
				},
				ScriptFormat: "shell",
				AppendConfig: nil,
				ProfileName:  nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				assert.Error(t, err, "An error was expected")
			} else {
				assert.NoError(t, err, "No error was expected")
			}
		})
	}
}
