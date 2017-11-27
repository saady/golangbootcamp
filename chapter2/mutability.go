package main

import "fmt"

type Artist struct {
    Name, Genre string
    Songs       int
}

func newRelease(a Artist) int {
    a.Songs++
    return a.Songs
}

func main(){
    me := Artist{Name: "SAADY", Genre: "rock", Songs: 1}
    fmt.Printf("%s released their %d'th song \n", me.Name, newRelease(me))
    fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}