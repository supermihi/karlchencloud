package db

import (
	"github.com/spf13/cobra"
	u "github.com/supermihi/karlchencloud/server/users"
	"log"
)

var DbCommand = &cobra.Command{
	Use:   "db",
	Short: "Database manipulation tools",
}

var dbConnection string

func getUsersOrFail() *u.SqlUserDatabase {
	db, err := u.NewSqlUserDatabase(dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	DbCommand.PersistentFlags().StringVar(&dbConnection, "database", "users.sqlite", "database location")
}
