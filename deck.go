package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) print(){
	for _, card := range d{
		fmt.Println(card)
	}
}


func (d deck) toString() string{
	//for _, card := range d{
	//	deckAsOneString += card + ","
	//}
	//deckAsOneString = deckAsOneString[:len(deckAsOneString)-1]
	return strings.Join(d, ",")
}

func (d deck) toByte() []byte{
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) {
	err := os.WriteFile(filename, d.toByte(), 0666)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func newDeckFromFile(filename string) deck{
	content, err := os.ReadFile(filename)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(content), ",")
}


func (d deck) shuffle() deck{
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d{
		newRandomPosition := r.Intn(len(d)-1)
		d[i], d[newRandomPosition] = d[newRandomPosition], d[i]
	}
	return d
}





















