package main

import (
	"fmt"
	"github.com/supermihi/doppelgopf/game"
)

func main() {
	deck := game.DealCards(1264)

	for i, hand := range deck {
		fmt.Println("Player %i:", game.Player(i))
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
