package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#008000",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, key := range c {
		fmt.Println("key is:", color, "and value is:", key)
	}
}

func testOutMaps() {
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#008000",
	}
	delete(colors, "red")

	var colorsEmpty map[string]string

	colorsMake := make(map[string]string)
	colorsMake["white"] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colorsEmpty)
	fmt.Println(colorsMake)
}
