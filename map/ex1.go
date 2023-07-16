package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	person := make(map[int]string)
	person[1] = "1"
	colors["white"] = "#ffffff"
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	delete(colors, "white")
	fmt.Println(colors)
	fmt.Println(person)
}
