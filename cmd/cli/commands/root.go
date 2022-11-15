package commands

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/jpower432/jpower432-go-cli/cmd/cli/commands/options"
)

// NewRootCmd creates a new cobra.Command for the command root.
func NewRootCmd() *cobra.Command {
	o := options.Common{}

	o.IOStreams = options.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
	cmd := &cobra.Command{
		Use:           filepath.Base(os.Args[0]),
		Short:         "A go client",
		Long:          "A long sentence about this CLI",
		SilenceErrors: false,
		SilenceUsage:  false,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(NewOtherCmd(&o))
	cmd.AddCommand(NewVersionCmd(&o))

	return cmd
}
