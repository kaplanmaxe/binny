package cmd

import (
	"fmt"
	"os"

	"github.com/kaplanmaxe/binny/calc"
	"github.com/spf13/cobra"
)

var twosComplementBinCmd = &cobra.Command{
	Use:   "twoscomplementbin",
	Short: "Convert a binary string to it's two's complement",
	Long:  "Convert a binary string to it's two's complement",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Only one number can be specified.")
			return
		}
		twosComplement, err := calc.TwosComplementBinary(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("An error occurred converting %s to it's two's complement: %s", args[0], err))
			return
		}
		fmt.Printf("%s two's complement is: %s\n", args[0], twosComplement)
	},
}

func init() {
	rootCmd.AddCommand(twosComplementBinCmd)
}
