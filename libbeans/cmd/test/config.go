package test

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GenTestConfigCmd() *cobra.Command {
	configTestCmd := cobra.Command{
		Use:   "config",
		Short: "Test configuration settings",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("test config....")
		},
	}

	return &configTestCmd
}