package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Error: length{ expected: %d; found: %v }", 52, len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Error: value, {expected 'Two of Spades', found %s }", d[0])
	}

}

func TestSaveToFileAndReadFromFile(t *testing.T) {

	os.Remove("__testingfile.txt")

	d := newDeck()
	d.saveToFile("__testingfile.txt")

	savedDeck := readFromFile("__testingfile.txt")

	if len(savedDeck) != 53 {
		t.Errorf("Error: length {Found: %v; Expected: 53}", len(savedDeck))
	}
}

// if main module is not loaded try usinh this
//go env -w GO111MODULE=auto
