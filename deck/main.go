package main

// import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFIle("./my_cards.txt")
	// cards := newDeckFromFile("my_cards.txt")
	cards.shuffle()
	cards.print()
	// card1, remindCards1 := deal(cards, 3)
	// card1.print()
	// remindCards1.print()

	// cards.print()
	// fmt.Println(cards.toString())
}
