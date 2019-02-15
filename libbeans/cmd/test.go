
package cmd

import (
	"github.com/spf13/cobra"
	"xiaoye.com/dory/beans/libbeans/cmd/test"
)

func genTestCmd() *cobra.Command {
	exportCmd := &cobra.Command{
		Use:   "test",
		Short: "Test config",
	}

	exportCmd.AddCommand(test.GenTestConfigCmd())
	return exportCmd
}
