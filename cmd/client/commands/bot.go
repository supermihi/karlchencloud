package commands

import (
  "context"
  "github.com/spf13/cobra"
  client_implementations "github.com/supermihi/karlchencloud/client/implementations"
  "time"
)

var (
  numBots         int
  invite          string
  firstBotIsOwner bool
  delayInMs       int
)
var botCmd = &cobra.Command{
  Use:   "bot",
  Short: "bot client",
  Args:  cobra.NoArgs,
  Run: func(cmd *cobra.Command, args []string) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    bots := make([]*client_implementations.BotClient, numBots)
    for i := 0; i < numBots; i++ {
      conn := client_implementations.CreateBotLogin(i, serverAddress)
      bots[i] = client_implementations.NewBotClient(conn, i == 0 && firstBotIsOwner, invite, time.Duration(delayInMs)*time.Millisecond)
      go bots[i].Start(ctx)
    }
    <-ctx.Done()
  },
}

func init() {
  botCmd.Flags().IntVarP(&numBots, "num-bots", "n", 1, "number of bots to add")
  botCmd.Flags().IntVarP(&delayInMs, "delay", "d", 0, "delay (in ms) added before any bot action")
  botCmd.Flags().StringVarP(&invite, "invite", "i", "", "invite code")
  botCmd.Flags().BoolVarP(&firstBotIsOwner, "owner", "o", false, "first bot is owner (starts tables)")
  rootCmd.AddCommand(botCmd)
}
