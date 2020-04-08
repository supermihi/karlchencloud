package main

import (
	"doppelgopf/doko"
	"fmt"
)

func main() {
	deck := doko.CreateDeck()

	for _, card := range deck {
		fmt.Println(card)
	}
}