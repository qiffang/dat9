package cli

import (
	"fmt"
	"os"

	"github.com/mem9-ai/dat9/pkg/client"
)

// Cat reads a remote file and writes it to stdout.
//
//	dat9 cat /path/to/file
func Cat(c *client.Client, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: dat9 cat <path>")
	}
	data, err := c.Read(args[0])
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(data)
	return err
}
