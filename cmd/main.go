package main

import (
	"fmt"
	"github.com/supermihi/karlchencloud/pkg/game/core"
)

func main() {
	deck := core.DealCards(1264)

	for i, hand := range deck {
		fmt.Println("Player %i:", core.Player(i))
		for j, card := range hand {
			fmt.Print(card.String())
			if j < len(hand)-1 {
				fmt.Print(", ")
			} else {
				fmt.Println()
			}
		}
	}
}
