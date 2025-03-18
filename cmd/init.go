package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "add",
	Short: "short init",
	Long:  `long init`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init cmd run begin")
		//打印flag
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("config").Value,
			cmd.Flags().Lookup("license").Value,
			cmd.Parent().Flags().Lookup("source").Value,
		)
		fmt.Println("init cmd run end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
