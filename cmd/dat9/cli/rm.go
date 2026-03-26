package cli

import (
	"fmt"

	"github.com/mem9-ai/dat9/pkg/client"
)

// Rm removes a remote file or directory.
//
//	dat9 rm /path/to/file
func Rm(c *client.Client, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: dat9 rm <path>")
	}
	return c.Delete(args[0])
}
