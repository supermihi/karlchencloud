package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd.AddCommand(interactiveCmd)
	Execute()
}

var serverAddress string
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Karlchencloud console clients",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&serverAddress, "server", "localhost:50501", "server location")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
