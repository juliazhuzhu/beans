
package cmd

import(
	"github.com/spf13/cobra"
	"fmt"
)

func genRunCmd(name string) *cobra.Command{
	runCmd := cobra.Command{
		Use:   "run",
		Short: "Run " + name,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("RUN")
		},
	}

	return &runCmd
}
