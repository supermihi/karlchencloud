package db

import (
	"github.com/spf13/cobra"
	users2 "github.com/supermihi/karlchencloud/room/users"
	"log"
)

var DbCommand = &cobra.Command{
	Use:   "db",
	Short: "Database manipulation tools",
}

var dbConnection string

func getUsersOrFail() *users2.SqlUserDatabase {
	db, err := users2.NewSqlUserDatabase(dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	DbCommand.PersistentFlags().StringVar(&dbConnection, "database", "users.sqlite", "database location")
}
