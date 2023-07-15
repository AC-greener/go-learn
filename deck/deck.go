package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert a deck to a string
func (d deck) toString() string {
	return strings.Join(d, ";")
}

// Save deck to a file
func (d deck) saveToFIle(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ";")
	return deck(s)
}

//	func (d deck) shuffle() {
//		for i := range d {
//			newPosition := rand.Intn(len(d) - 1)
//			d[i], d[newPosition] = d[newPosition], d[i]
//		}
//	}
//
// Shuffle the deck in place
func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
