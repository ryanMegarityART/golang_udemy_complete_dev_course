package main

import "fmt"

func main() {
	cards := newDeckFromFile("cards.txt")
	cards.print()
	hand, remainingCards := cards.deal(3)
	hand.print()
	remainingCards.print()
	fmt.Println(hand.toString())
	cards.saveToFile("cards.txt")
}
