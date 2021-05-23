package db

import (
	"github.com/spf13/cobra"
	"github.com/supermihi/karlchencloud/server/users"
	"log"
)

var DbCommand = &cobra.Command{
	Use:   "db",
	Short: "Database manipulation tools",
}

var dbConnection string

func getUsersOrFail() *users.SqlUserDatabase {
	db, err := users.NewSqlUserDatabase(dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	DbCommand.PersistentFlags().StringVar(&dbConnection, "database", "users.sqlite", "database location")
}
