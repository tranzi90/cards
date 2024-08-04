package main

import "cards/deck"

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.print()
}
