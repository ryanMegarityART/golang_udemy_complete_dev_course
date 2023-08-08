package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.txt")
	cards.shuffle()
	cards.saveToFile("cardsShuffled.txt")
	cards.print()
}
