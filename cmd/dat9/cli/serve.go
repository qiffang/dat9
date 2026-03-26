package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mem9-ai/dat9/pkg/backend"
	"github.com/mem9-ai/dat9/pkg/meta"
	"github.com/mem9-ai/dat9/pkg/server"
)

const (
	defaultListenAddr = ":9009"
	defaultDBPath     = "dat9.db"
	defaultBlobDir    = "blobs"
)

// Serve starts the dat9 HTTP server backed by the local SQLite/blob stand-in
// used in P0.
func Serve(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("usage: dat9 serve [listen-addr]")
	}

	addr := envOr("DAT9_LISTEN_ADDR", defaultListenAddr)
	if len(args) == 1 {
		addr = args[0]
	}

	dbPath := envOr("DAT9_DB_PATH", defaultDBPath)
	blobDir := envOr("DAT9_BLOB_DIR", defaultBlobDir)

	if err := ensureParentDir(dbPath); err != nil {
		return err
	}
	if err := os.MkdirAll(blobDir, 0o755); err != nil {
		return fmt.Errorf("create blob dir: %w", err)
	}

	store, err := meta.Open(dbPath)
	if err != nil {
		return fmt.Errorf("open meta store: %w", err)
	}
	defer store.Close()

	b, err := backend.New(store, blobDir)
	if err != nil {
		return fmt.Errorf("create backend: %w", err)
	}

	return server.New(b).ListenAndServe(addr)
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func ensureParentDir(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create db dir: %w", err)
	}
	return nil
}
