package cmd

import(
	"fmt"
	"github.com/spf13/pflag"
	"flag"
	"os"
	cmd "xiaoye.com/dory/beans/libbeans/cmd"
	"xiaoye.com/dory/beans/filebeans/beaner"
)

var Name = "beans"
var RootCmd *cmd.BeansRootCmd

func init() {
	var runFlags = pflag.NewFlagSet(Name, pflag.ExitOnError)
	f := flag.CommandLine.Lookup("once")
	if f != nil {
		runFlags.AddGoFlag(f)
	}
	f = flag.CommandLine.Lookup("modules")
	if f != nil {
		runFlags.AddGoFlag(f)
	}

	RootCmd = cmd.GenRootCmd(Name, beaner.New,runFlags)
	RootCmd.AddCommand(genVersionCmd())
}
  
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}