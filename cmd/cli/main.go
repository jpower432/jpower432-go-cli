package main

import (
	"github.com/spf13/cobra"

	"github.com/jpower432/jpower432-go-cli/cmd/cli/commands"
)

func main() {
	rootCmd := commands.NewRootCmd()
	cobra.CheckErr(rootCmd.Execute())
}
