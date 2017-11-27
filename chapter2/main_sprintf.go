package main


import "fmt"

func main() {
	name := "Yousef"
	aka := fmt.Sprintf("Yous%df", 3)
	fmt.Printf("%s is know as %s", name, aka)
}
