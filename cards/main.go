package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	// var card string = "Ace of Spades"
	// fmt.Println(card)

	// var card2 = "This is also another card"
	// fmt.Println("This is a card", card2)

	// //var joker string  = "I am a joker"
	// joker := "I am a joker"
	// fmt.Println(joker)

	// //reassign
	// joker = "Joker ?"
	// fmt.Println(joker)

	// card := newCard() //calling a fn.
	// fmt.Println(card)

	// card2 := newCard2() //calling
	// fmt.Println(card2)

	// cards := []string{
	// 	"Hello I am jack;",
	// 	"Ace of Spades;",
	// 	newCard(), //comma is necessary when } is in nextline
	// }

	// cards = append(cards, "Jack of sparrows")

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, " -> ", card)
	// }

	// cards := deck{
	// 	newCard(),
	// 	"Ace of spades",
	// 	"12 of Diamond",
	// 	"King of panam",
	// }

	// fmt.Println(cards)
	// //using normal loop
	// for k, v := range cards {
	// 	fmt.Println(k, v)
	// }
	// //using receiver fn
	// cards.print()

	// cards := newDeck()
	// cards.print()

	// choosen, remaining := deal(cards, 7)
	// fmt.Println("Choosen")
	// choosen.print()
	// fmt.Println("Remaining")
	// remaining.print()

	// cards := newDeck()
	// cards.saveToFile("sample.txt")

	// savedcards := readFromFile("sample.txt")
	// savedcards.print()

	// cards := newDeck()
	// fmt.Println("Before Shuffling")
	// cards.print()

	// // shuffle
	// cards.shuffle()

	// fmt.Println("After Shuffling")
	// cards.print()

}

func newCard() string {
	return "Hello I am a new diamond card"

}

func newCard2() int {
	return 123
}
