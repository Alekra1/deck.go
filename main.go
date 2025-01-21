package main

import "fmt"

func main() {
	cards := newDeck()

	//hand, remainingCards := deal(cards, 2)
	//hand.print()
	//fmt.Println("-------------------------")
	//remainingCards.print()
	fmt.Println(cards.toString())
	fmt.Println(cards.toByte())
	cards.saveToFile("file.txt")
	cards2 := newDeckFromFile("file.txt")
	fmt.Println(cards2)

	cards.shuffle()
	fmt.Println(cards)
}

