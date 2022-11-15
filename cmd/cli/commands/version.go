package commands

import (
	"github.com/spf13/cobra"

	"github.com/jpower432/jpower432-go-cli/cmd/cli/commands/options"
	"github.com/jpower432/jpower432-go-cli/version"
)

// NewVersionCmd creates a new cobra.Command for the version subcommand.
func NewVersionCmd(common *options.Common) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return version.WriteVersion(common.IOStreams.Out)
		},
	}
}
