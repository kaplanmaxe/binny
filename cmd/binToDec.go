package cmd

import (
	"fmt"
	"os"

	"github.com/kaplanmaxe/binny/calc"
	"github.com/spf13/cobra"
)

var signed bool

var binToDecCmd = &cobra.Command{
	Use:   "bintodec",
	Short: "Convert a decimal number to a binary number",
	Long:  "Convert a decimal number to a binary number",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Only one number can be specified.")
			return
		}
		binNum := args[0]
		decNum, err := calc.BinToDec(binNum, signed)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("An error occurred converting %d to binary: %s", decNum, err))
			return
		}
		fmt.Printf("%s in decimal is: %d\n", binNum, decNum)
	},
}

func init() {
	rootCmd.AddCommand(binToDecCmd)
	binToDecCmd.Flags().BoolVarP(&signed, "signed", "s", false, "Specifies a signed number")
}
