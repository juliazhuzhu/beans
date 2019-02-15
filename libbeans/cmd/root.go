package cmd

import(
	"github.com/spf13/cobra"
)

type BeansRootCmd struct {
	cobra.Command
	RunCmd        	*cobra.Command
	TestCmd			*cobra.Command		
}

func GenRootCmd() *BeansRootCmd{

	rootCmd := &BeansRootCmd{}
	rootCmd.Use = "beans"
	rootCmd.RunCmd = genRunCmd("beans")
	rootCmd.TestCmd = genTestCmd()
	rootCmd.Run = rootCmd.RunCmd.Run

	rootCmd.AddCommand(rootCmd.RunCmd)
	rootCmd.AddCommand(rootCmd.TestCmd)
	return rootCmd
}