// +build !windows

package app

import (
	"github.com/Benbentwo-Sandbox/test-go-bin-generic-make-mine-0/pkg/cmd"
	"os"
)

// Run runs the command, if args are not nil they will be set on the command
func Run(args []string) error {
	cmd := cmd.NewMainCmd(os.Stdin, os.Stdout, os.Stderr, nil)
	if args != nil {
		args = args[1:]
		cmd.SetArgs(args)
	}
	return cmd.Execute()
}
