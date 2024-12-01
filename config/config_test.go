package config_test

import (
	"kong-assignment/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockGodotenvLoad() error {
	return nil
}
func TestLoadConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		env      map[string]string
		wantErr  bool
		errMsg   string
		wantConf *config.PostgresDbConfig
	}{
		{
			name:    "valid config",
			env:     map[string]string{"DB_HOST": "localhost", "DB_USER": "user", "DB_PASSWORD": "password", "DB_NAME": "db", "DB_PORT": "5432"},
			wantErr: false,
			wantConf: &config.PostgresDbConfig{
				Host:     "localhost",
				Port:     5432,
				User:     "user",
				Password: "password",
				DbName:   "db",
			},
		},
		{
			name:    "missing env vars",
			env:     map[string]string{"DB_HOST": "localhost", "DB_USER": "user", "DB_PASSWORD": "password", "DB_NAME": "db"},
			wantErr: true,
			errMsg:  "missing required environment variables",
		},
		{
			name:    "invalid port",
			env:     map[string]string{"DB_HOST": "localhost", "DB_USER": "user", "DB_PASSWORD": "password", "DB_NAME": "db", "DB_PORT": "invalid"},
			wantErr: true,
			errMsg:  "invalid port number",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment variables
			os.Clearenv()
			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			// Load config
			conf, err := config.LoadConfig()

			// Check error
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}
			require.NoError(t, err)

			// Check config
			assert.Equal(t, tt.wantConf, conf)
		})
	}
}
