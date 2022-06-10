package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}
type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, batteryDrain: batteryDrain, speed: speed, distance: 0}
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time, the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for i := 0; i < track.distance; i++ {
		res := Drive(car)
		if res.distance == track.distance {
			return true
		}
		if res.battery == car.battery {
			return false
		}
		car = res
	}
	// OR: max := (car.battery / car.batteryDrain) * car.speed
	// how many times does battery contains batteryDrain, each time = Un + speed, so time speed
	return true
}
