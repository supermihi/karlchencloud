package db

import (
	"github.com/spf13/cobra"
	"log"
)

var addUserCommand = &cobra.Command{
	Use:   "add [email] [name] [password]",
	Short: "add users to database",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		db := getUsersOrFail()
		id, err := db.Add(args[0], args[2], args[1])
		if err != nil {
			log.Fatalf("error adding user: %v", err)
		}
		log.Printf("created user with id %v", id)
	},
}

func init() {
	DbCommand.AddCommand(addUserCommand)
}
