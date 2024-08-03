package main

import "cards/deck"

func main() {
	cards := deck.NewDeck()
	cards.shuffle()
	cards.print()
}
