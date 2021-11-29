package db

import (
  "github.com/spf13/cobra"
  "log"
)
var (
  password string
)

var passwordCommand = &cobra.Command{
  Use:   "pass [email] <password>",
  Short: "change password",
  Args: cobra.RangeArgs(1, 2),
  Run: func(cmd *cobra.Command, args []string) {
    db := getUsersOrFail()
    user, err := db.FindByEmail(args[0])
    if err != nil {
      log.Fatalf("error finding user: %v", err)
    }
    log.Printf("changing password of user %v with id %v", user.Name, user.Id)
    err = db.ChangePassword(user.Id, args[1])
    if err != nil {
      log.Fatalf("error updating password: %v", err)
    }

  },
}

func init() {
  passwordCommand.Flags().StringVarP(&password, "password", "p", "", "the password")
  DbCommand.AddCommand(passwordCommand)
}
