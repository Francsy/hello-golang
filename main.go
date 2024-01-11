package main

import "fmt"

var card string

func main() {

	/*
	   One way to declare a variable

	   var card string = "Ace of Spades"

	   Another:

	   card := "Ace of Spades"
	*/
	card := "Ace of Spades"
	card = newCard()
	fmt.Println(card)
}

func newCard() string { // Function that return an string type
	return "Five of Diamonds"
}
