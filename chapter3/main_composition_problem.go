package main

import "fmt"

type User struct {
	Id       int
	Location string
	Name     string
}

type Player struct {
	Id       int
	Location string
	Name     string
	GameId   int
}

func main() {
	p := Player{}
	p.Id = 42
	p.Location = "Casablanca"
	p.Name = "Youssef"
	p.GameId = 2335
	fmt.Printf("%+v", p)
}
