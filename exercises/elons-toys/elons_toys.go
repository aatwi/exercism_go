package elon

import "strconv"

// Drive drives the car.
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.battery = car.battery - car.batteryDrain
		car.distance = car.distance + car.speed
	}
}

// DisplayDistance returns the distance traveled.
func (car *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// DisplayBattery displays the remaining battery.
func (car *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// CanFinish returns true if the car can finish the track
func (car *Car) CanFinish(trackDistance int) bool {
	timeRequired := float64(trackDistance) / float64(car.speed)
	batteryConsumed := float64(car.batteryDrain) * timeRequired
	return float64(car.battery) >= batteryConsumed
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
