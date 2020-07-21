package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"'
	// card := newCard()
	cards := newDeck()

	handCard, remainingDeck := deal(cards, 5)
	handCard.print()
	remainingDeck.print()
	// cards.print()
}
