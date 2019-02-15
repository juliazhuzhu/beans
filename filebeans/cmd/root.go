package cmd

import(
	"fmt"
	//"github.com/spf13/cobra"
	"os"
	cmd "xiaoye.com/dory/beans/libbeans/cmd"
)

var RootCmd *cmd.BeansRootCmd

func init() {

	RootCmd = cmd.GenRootCmd()
	RootCmd.AddCommand(genVersionCmd())
}
  
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}