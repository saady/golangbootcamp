package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var continent string
	
	switch city {
	case "Agadir", "Taroudant", "Ouled teima":
		region, continent = "Souss massa", "Africa"
	case "New york", "NYC":
		region, continent = "New York", "North america"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}


func main() {
	region, continent := location("Agadir")
	fmt.Printf("SAADY lives in %s, %s", region, continent)
}