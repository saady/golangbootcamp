package main

import (
	"fmt"
)

func main() {
	cities := map[string]int{
		"New York":    8336697,
		"Los Angelos": 3857799,
		"Chicago":     2714856,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants \n", key, value)
	}
}
