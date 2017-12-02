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
	fmt.Println(Bootcamp{Lon: 124545, Lat: 24577, Date: time.Now()})
}
