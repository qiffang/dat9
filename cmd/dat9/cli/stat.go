package cli

import (
	"fmt"

	"github.com/mem9-ai/dat9/pkg/client"
)

// Stat shows metadata for a remote path.
//
//	dat9 stat /path/to/file
func Stat(c *client.Client, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: dat9 stat <path>")
	}
	s, err := c.Stat(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("size:     %d\n", s.Size)
	fmt.Printf("isdir:    %v\n", s.IsDir)
	fmt.Printf("revision: %d\n", s.Revision)
	return nil
}
