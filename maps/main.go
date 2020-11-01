package main

import "fmt"

func main() {
	colors := map[string]string{
		"black": "#000000",
		"white": "#ffffff",
		"grey":  "#cccccc",
	}

	updateMap(colors, "black", "blackColor")

	printMap(colors)
}

func updateMap(c map[string]string, k string, v string) {
	c[k] = v
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
