package main

import "fmt"

func main() {
	mySlice := make([]string, 3)
	mySlice[0] = "Agadir"
	mySlice[1] = "Casablanca"
	mySlice[2] = "Rabat"

	fmt.Println("%q", mySlice)
}
