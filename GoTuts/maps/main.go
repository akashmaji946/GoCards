package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello to maps")

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// 	"white": "#00000f",
	// }

	// fmt.Println(colors)

	// colors["white"] = "#000000" //changing
	// colors["black"] = "#ffffff" //appending

	// fmt.Println(colors)

	// var colors map[string]string
	// fmt.Println(colors)
	// colors["white"] = "#000000" //error: cant assign to nil map
	// colors["black"] = "#ffffff" //error
	// fmt.Println(colors)

	colors := make(map[string]string)
	fmt.Println(colors)
	colors["white"] = "#000070"
	colors["black"] = "#ffffff"
	colors["white"] = "#000000"
	fmt.Println(colors)

}
