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
		fmt.Println(i+1, card)
	}

}

func (d deck) shuffle() {

	t := time.Now()      // t is of type 'Time' (time.Time)
	seed := t.UnixNano() // seed is of type int64

	source := rand.NewSource(seed)
	random := rand.New(source)

	for i := range d {
		j := random.Intn(len(d)) //using random created above
		d[i], d[j] = d[j], d[i]
	}

}

func deckToByteArray(d deck) []byte {
	cardsInDeckToStringArray := []string(d)
	cardsToString := strings.Join(cardsInDeckToStringArray, "\n")
	bytearray := []byte(cardsToString)
	return bytearray

}

func byteArrayToDeck(bytearray []byte) deck {
	byteArrayToString := string(bytearray)
	cardsInStringArray := strings.Split(byteArrayToString, "\n")
	deckFromStringArray := deck(cardsInStringArray)
	return deckFromStringArray
}

func (d deck) saveToFile(filename string) error {
	bytearray := deckToByteArray(d)
	return ioutil.WriteFile(filename, bytearray, 0666)
}

func deal(d deck, handsize int) (deck, deck) {
	choosen := d[:handsize]
	remaining := d[handsize:]

	return choosen, remaining
}

func readFromFile(filename string) deck {
	bytearray, err := ioutil.ReadFile(filename)
	if err != nil {
		// Log out error and exit
		fmt.Println("Error: ", err)
		os.Exit(1)

	}
	d := byteArrayToDeck(bytearray)
	return d

}

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := []string{"Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine",
		"Ten", "Ace", "King", "Queen", "Jack"}

	for _, suit := range suits {
		for _, value := range values {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}

	return cards
}
