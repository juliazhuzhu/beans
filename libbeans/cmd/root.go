package cmd

import(
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"xiaoye.com/dory/beans/libbeans/bean"
)

type BeansRootCmd struct {
	cobra.Command
	RunCmd        	*cobra.Command
	TestCmd			*cobra.Command		
}

func GenRootCmd(name string, beanCreator bean.Creator,runFlags *pflag.FlagSet) *BeansRootCmd{

	rootCmd := &BeansRootCmd{}
	rootCmd.Use = name
	rootCmd.RunCmd = genRunCmd(name,beanCreator,runFlags)
	rootCmd.TestCmd = genTestCmd()
	rootCmd.Run = rootCmd.RunCmd.Run

	rootCmd.Flags().AddFlagSet(rootCmd.RunCmd.Flags())
	rootCmd.AddCommand(rootCmd.RunCmd)
	rootCmd.AddCommand(rootCmd.TestCmd)
	return rootCmd
}