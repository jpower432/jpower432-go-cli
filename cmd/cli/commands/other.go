package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jpower432/jpower432-go-cli/cmd/cli/commands/options"
)

type OtherOptions struct {
	*options.Common
}

// NewOtherCmd creates a new cobra.Command for a subcommand.
func NewOtherCmd(options *options.Common) *cobra.Command {
	o := OtherOptions{
		Common: options,
	}

	cmd := &cobra.Command{
		Use:           "other",
		SilenceErrors: false,
		SilenceUsage:  false,
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(o.Complete(args))
			cobra.CheckErr(o.Validate())
			cobra.CheckErr(o.Run(cmd.Context()))
		},
	}

	return cmd
}

func (o *OtherOptions) Complete(_ []string) error {
	// Do completion things
	return nil
}

func (o *OtherOptions) Validate() error {
	// Do validation things
	return nil
}

func (o *OtherOptions) Run(_ context.Context) error {
	fmt.Println(" I do other things")
	return nil
}
