
package cmd

import(
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"fmt"
)

func genRunCmd(name string,runFlags *pflag.FlagSet) *cobra.Command{
	runCmd := cobra.Command{
		Use:   "run",
		Short: "Run " + name,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("RUN")
		},
	}

	if runFlags != nil {
		runCmd.Flags().AddFlagSet(runFlags)
	}

	return &runCmd
}
