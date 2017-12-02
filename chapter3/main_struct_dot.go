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
	event := Bootcamp{Lon: 23.34555, Lat: -134.2653}

	event.Date = time.Now()
	fmt.Printf("event on %s, location (%f, %f)", event.Date, event.Lat, event.Lon)

}
