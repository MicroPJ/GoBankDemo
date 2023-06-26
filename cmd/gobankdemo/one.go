package gobankdemo

import (
	"fmt"

	"github.com/micropj/gobankdemo/pkg/gobankdemo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(oneCmd)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gobankdemo",
	Long:  `All software has versions. This is gobankdemo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GoBankDemo v0.1\n")
	},
}

var oneCmd = &cobra.Command{
	Use:     "one",
	Aliases: []string{"o"},
	Short:   "Option one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		res := gobankdemo.One(args)
		fmt.Printf(res)
	},
}
