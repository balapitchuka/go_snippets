package main

import "fmt"

// throws error
// deckSize := 20
var deckSize int

func main() {

	cards_arr := newDeck()
	cards_arr.print()

	deckSize = 10
	fmt.Println(deckSize)
}

func newCard() string {
	return "Five of Diamonds"
}
