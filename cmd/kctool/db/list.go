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
			user, err := db.GetData(id)
			if err != nil {
				log.Fatalf("error retrieving user name: %v", err)
			}
			log.Printf("[%s] %s (%s)", user.Id, user.Name, user.Email)
		}

	},
}

func init() {
	DbCommand.AddCommand(listUsersCommand)
}
