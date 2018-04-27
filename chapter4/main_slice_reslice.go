package main

import "fmt"

func main() {

	mySlice := []int{2, 4, 4, 5, 5, 6, 7}
	fmt.Println(mySlice)
	fmt.Println(mySlice[1:4])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[4:])
}
