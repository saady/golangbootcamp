package main

import "fmt"

type Point struct {
	x, y int
}

var (
	p = Point{1, 3}
	q = &Point{1, 3}
	r = Point{x: 1}
	s = Point{}
)

func main() {
	fmt.Println(p, q, r, s)
}
