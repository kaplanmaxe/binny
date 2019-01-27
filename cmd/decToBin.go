package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kaplanmaxe/binny/calc"
	"github.com/spf13/cobra"
)

var decToBinCmd = &cobra.Command{
	Use:   "dectobin",
	Short: "Convert a decimal number to a binary number",
	Long:  "Convert a decimal number to a binary number",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Only one number can be specified.")
			return
		}
		decNum, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "You must enter a number")
			return
		}
		binNum, err := calc.DecToBin(int(decNum))
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("An error occurred converting %d to binary: %s", decNum, err))
			return
		}
		fmt.Printf("%d in binary is: %s\n", decNum, binNum)
	},
}

func init() {
	rootCmd.AddCommand(decToBinCmd)
}
