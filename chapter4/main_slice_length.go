package main

import "fmt"

func main() {
	cities := []string{
		"Casablanca",
		"Rabat",
		"Agadir",
	}
	fmt.Println(len(cities))
	countries := make([]string, 42)
	fmt.Println(len(countries))
}
