package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "command",
	Short: "cmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is command")
	},
}

func main() {
	if err := Command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
