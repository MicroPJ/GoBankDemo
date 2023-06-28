package gobankdemo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobankdemo",
	Short: "gobankdemo - Golang utility to deploy BankDemo",
	Long:  `Run "gobankdemo -h" for help on any command`,
	PreRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Run 'gobankdemo -h' for help on any command")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
