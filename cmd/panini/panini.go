package panini

import (
	"context"

	"github.com/going/gclone/operations"
	"github.com/rclone/rclone/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
}

var commandDefinition = &cobra.Command{
	Use: "panini remote:path",

	Run: func(command *cobra.Command, args []string) {
		//cmd.CheckArgs(1, 2, command, args)
		fdst := cmd.NewFsSrc(args)
		cmd.Run(false, false, command, func() error {
			return operations.Panini(context.Background(), fdst)
		})
	},
}
