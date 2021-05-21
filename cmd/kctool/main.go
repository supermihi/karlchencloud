package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/supermihi/karlchencloud/cmd/kctool/db"
	"os"
)

func main() {
	rootCmd.AddCommand(db.DbCommand)
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "kctool",
	Short: "Karlchencloud management tools",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
