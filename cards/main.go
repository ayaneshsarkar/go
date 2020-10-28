package main

//import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Seven of Diamonds")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
