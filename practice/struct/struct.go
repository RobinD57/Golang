package main

import "fmt"

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeedKmh float64
}

func main() {
	a_car := car {
		gasPedal: 22341,
		brakePedal: 0,
		steeringWheel: 12561,
		topSpeedKmh: 225.0,
	}

	fmt.Println(a_car.steeringWheel)
}
