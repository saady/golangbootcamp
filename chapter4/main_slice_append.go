package main

import "fmt"

func main() {
	cities := []string{}
	cities = append(cities, "Casablanca", "Agadir")
	fmt.Println(cities)

	otherCities := []string{"Paris", "marseille"}
	cities = append(cities, otherCities...)
	fmt.Println(cities)
}
