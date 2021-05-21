package db

import (
	"github.com/spf13/cobra"
	"log"
)

var listUsersCommand = &cobra.Command{
	Use:   "list",
	Short: "list users",
	Run: func(cmd *cobra.Command, args []string) {
		db := getUsersOrFail()
		ids, err := db.ListIds()
		if err != nil {
			log.Fatalf("error listing users: %v", err)
		}
		for _, id := range ids {
			user, err := db.GetName(id)
			if err != nil {
				log.Fatalf("error retrieving user name: %v", err)
			}
			log.Print(user)
		}

	},
}

func init() {
	DbCommand.AddCommand(listUsersCommand)
}
