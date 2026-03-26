// Package cli implements dat9 CLI commands.
//
// Design follows Plan 9 philosophy: small tools, composable, no flags bloat.
// Each command reads from the client SDK and writes to stdout/stderr.
package cli

import (
	"os"

	"github.com/mem9-ai/dat9/pkg/client"
)

// NewFromEnv creates a dat9 client from environment variables.
func NewFromEnv() *client.Client {
	server := os.Getenv("DAT9_SERVER")
	if server == "" {
		server = "http://localhost:9009"
	}
	apiKey := os.Getenv("DAT9_API_KEY")
	return client.New(server, apiKey)
}
