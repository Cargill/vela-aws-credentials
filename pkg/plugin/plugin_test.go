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
			},
			wantErr: true,
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
