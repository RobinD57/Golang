package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeedKmh float64
}

func(c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func(c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmh_multiple)
}

func main() {
	a_car := car {
		gasPedal: 22341,
		brakePedal: 0,
		steeringWheel: 12561,
		topSpeedKmh: 225.0,
	}

	fmt.Println(a_car.steeringWheel)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
