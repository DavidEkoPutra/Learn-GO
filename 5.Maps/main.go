package main

import (
	statuscode "3.Assignment/5.Maps"
	"fmt"
)

func main() {

	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{

		"red":   "#ff0000",
		"green": "#42134",
		"white": "#ffffff",
	}

	fmt.Println(statuscode.StatusText(003))
	//colors["white"] = "#fffff"

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex Code for", color, "is", hex)
	}
}
