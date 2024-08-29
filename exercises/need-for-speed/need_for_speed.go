package speed

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		battery:      100,
		batteryDrain: batteryDrain,
		distance:     0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	battery := car.battery - car.batteryDrain
	if battery < 0 {
		return car
	}
	return Car{
		speed:        car.speed,
		battery:      battery,
		distance:     car.distance + car.speed,
		batteryDrain: car.batteryDrain,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	neededTime := track.distance / car.speed
	neededBattery := car.batteryDrain * neededTime
	if neededBattery > car.battery {
		return false
	}
	return true
}
