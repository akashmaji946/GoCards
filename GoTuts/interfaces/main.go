package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type commonBot interface {
	// all types which associate themselves
	// with these exact functions
	// become honorary of type 'commonBot'
	getGreet() string
}

func main() {
	fmt.Println("Hello World! Welcome to interfaces.")

	eb := englishBot{}
	sb := spanishBot{}

	fmt.Println(eb)
	fmt.Println(sb)

	printGreet(eb)
	printGreet(sb)

}

// ant type that is honorary of commonBot can use this
func printGreet(cb commonBot) {
	fmt.Println(cb.getGreet())
}

// func printGreet(eb englishBot) {
// 	fmt.Println(eb.getGreet())
// }

// we cant have this in same file
// func printGreet(sb spanishBot) {
// 	fmt.Println(sb.getGreet())
// }

func (eb englishBot) getGreet() string {
	//logic
	return "Hello"

}

func (sb spanishBot) getGreet() string {
	//logic
	return "Hola"
}
