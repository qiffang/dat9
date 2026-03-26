package cli

import (
	"fmt"

	"github.com/mem9-ai/dat9/pkg/client"
)

// Mv renames or moves a remote file/directory. Metadata-only, zero S3 cost.
//
//	dat9 mv /old/path /new/path
func Mv(c *client.Client, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: dat9 mv <old> <new>")
	}
	return c.Rename(args[0], args[1])
}
