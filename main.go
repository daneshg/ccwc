package main

import (
	ccwc "daneshg/ccwc/cmd"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func printOutput(cc *cobra.Command, wc *ccwc.WCStat, args ...string) {
	flagSet := false
	if cc.Flags().Changed("lines") {
		flagSet = true
		fmt.Printf("%d ", wc.GetLines())
	}
	if cc.Flags().Changed("words") {
		flagSet = true
		fmt.Printf("%d ", wc.GetWords())
	}
	if cc.Flags().Changed("bytes") {
		flagSet = true
		fmt.Printf("%d ", wc.GetBytes())
	}
	if cc.Flags().Changed("multibyte") {
		flagSet = true
		fmt.Printf("%d ", wc.GetChars())
	}

	if !flagSet {
		fmt.Printf("%d %d %d %s", wc.GetLines(), wc.GetWords(), wc.GetBytes(), strings.Join(args, " "))
	}

}

func init() {
	var version string = "0.0.1"
	rootCmd = &cobra.Command{
		Use:     "ccwc",
		Short:   "Custom wc cli tool",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				wcStat, err := ccwc.GetWCFromStdIN()
				if err != nil {
					log.Fatal(err)
				}
				printOutput(cmd, wcStat)
			} else {
				wcStat, err := ccwc.GetWCFromFile(args[0])
				if err != nil {
					log.Fatal(err)
				}
				printOutput(cmd, wcStat, args[0])
			}

		},
	}

	rootCmd.Flags().BoolP("lines", "l", false, "Total number of lines")
	rootCmd.Flags().BoolP("words", "w", false, "Total number of words")
	rootCmd.Flags().BoolP("bytes", "c", false, "Total number of bytes")
	rootCmd.Flags().BoolP("multibyte", "m", false, "Total number of characters")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
