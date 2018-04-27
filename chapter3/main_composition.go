package main

import (
	"fmt"
)

type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	User
	GameId int
}

func main() {
	p := Player{
		User{Id: 34, Name: "Yousef", Location: "Casablanca"}, 345,
	}
	fmt.Printf(
		"Id: %d, Name: %s, Location: %s, Game id: %d\n",
		p.Id, p.Name, p.Location, p.GameId)
	//Direcly set a field definied on Player struct
	p.Id = 35
	fmt.Printf("%+v", p)
}
