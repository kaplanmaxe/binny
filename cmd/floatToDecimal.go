package cmd

import (
	"fmt"
	"os"

	"github.com/kaplanmaxe/binny/calc"
	"github.com/spf13/cobra"
)

var floatToDecCmd = &cobra.Command{
	Use:   "floattodec",
	Short: "Convert a float binary string to decimal",
	Long:  "Convert a float binary string to decimal",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Only one number can be specified.")
			return
		}
		binNum := args[0]
		decNum, err := calc.FloatToDecimal(binNum)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("An error occurred converting %s to decimal: %s", binNum, err))
			return
		}
		fmt.Printf("%s decimal is: %s\n", binNum, decNum)
	},
}

func init() {
	rootCmd.AddCommand(floatToDecCmd)
}
