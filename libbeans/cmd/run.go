
package cmd

import(
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	//"context"
	//svc "xiaoye.com/dory/beans/libbeans/service"
	"xiaoye.com/dory/beans/libbeans/bean"
	"xiaoye.com/dory/beans/libbeans/cmd/instance"
	"fmt"
)



func genRunCmd(name string,bt bean.Creator,runFlags *pflag.FlagSet) *cobra.Command{
	runCmd := cobra.Command{
		Use:   "run",
		Short: "Run " + name,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("RUN")
			instance.Run(bt)
		},
	}

	if runFlags != nil {
		runCmd.Flags().AddFlagSet(runFlags)
	}

	//_, cancel := context.WithCancel(context.Background())
	
	//bean, err := bt(&b.Beat)
	//svc.HandleSignals(bean.Stop, cancel)

	return &runCmd
}
