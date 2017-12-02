package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lon  float64
	Lat  float64
	Date time.Time
}

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(*x == *y)
}
