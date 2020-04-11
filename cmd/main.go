package main

import (
	"fmt"
	"github.com/supermihi/doppelgopf/game"
)

func main() {
	deck := game.Deck()

	for _, card := range deck {
		fmt.Println(card)
	}
}
