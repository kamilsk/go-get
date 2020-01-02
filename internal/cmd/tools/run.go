package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewRunCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			tools, err := cfg.Create()
			if err != nil {
				return err
			}
			return tools.Run(ctx, args[0], args[1:]...)
		},
	}
	return cmd
}
