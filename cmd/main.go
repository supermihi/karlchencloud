package main

import (
	"fmt"
	"github.com/supermihi/doppelgopf/doko"
)

func main() {
	deck := doko.Deck()

	for _, card := range deck {
		fmt.Println(card)
	}
}
