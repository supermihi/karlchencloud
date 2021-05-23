package commands

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/supermihi/karlchencloud/client"
	"github.com/supermihi/karlchencloud/client/implementations"
)

var (
	numBots int
	invite  string
)
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "bot client",
	Run: func(cmd *cobra.Command, args []string) {
		conn := implementations.CreateBotLogin(1, serverAddress)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		botHandler := implementations.NewBotHandler(false, invite)
		karlchenClient := client.NewClientImplementation(conn, botHandler)
		go karlchenClient.Start(ctx)
		<-ctx.Done()
	},
}

func init() {
	botCmd.Flags().IntVarP(&numBots, "num-bots", "n", 1, "number of bots to add")
	botCmd.Flags().StringVarP(&invite, "invite", "i", "", "invite code")
	rootCmd.AddCommand(botCmd)
}
