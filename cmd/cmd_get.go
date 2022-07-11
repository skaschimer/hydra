package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"
)

func NewGetCmd(root *cobra.Command) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "Get resources",
	}
	cmd.AddCommand(NewGetClientsCmd(root))
	cmdx.RegisterHTTPClientFlags(cmd.PersistentFlags())
	cmdx.RegisterFormatFlags(cmd.PersistentFlags())
	return cmd
}