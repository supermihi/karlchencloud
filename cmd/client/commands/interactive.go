package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/client/implementations"
	"log"
)

var (
	name     string
	email    string
	userId   string
	password string
)

func askForMissingClientData() {
	for name == "" {
		log.Printf("Input your display name:")
		name = implementations.UserInputString()
	}
	for email == "" {
		log.Printf("Input your email address:")
		email = implementations.UserInputString()
	}
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive command-line client",
	Run: func(cmd *cobra.Command, args []string) {
		askForMissingClientData()
		conn := client.LoginData{
			Email:                 email,
			Name:                  name,
			UserId:                userId,
			Password:              password,
			ServerAddress:         serverAddress,
			RegisterIfEmptyUserId: true,
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		cliHandler := implementations.CliHandler{}
		karlchenClient := client.NewKarlchenClient(conn, &cliHandler)
		go karlchenClient.Start(ctx)
		<-ctx.Done()
	},
}

func init() {
	interactiveCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
	interactiveCmd.Flags().StringVarP(&email, "email", "e", "", "Your email address")
	interactiveCmd.Flags().StringVarP(&userId, "id", "i", "", "User ID")
	interactiveCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	rootCmd.AddCommand(interactiveCmd)
}
