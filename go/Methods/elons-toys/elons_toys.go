package elon

import "fmt"

func (car *Car) Drive() {
	if car.batteryDrain < car.battery {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%s", car.battery, "%")
}

func (car Car) CanFinish(trackDistance int) bool {
	max := (car.battery / car.batteryDrain) * car.speed
	return max >= trackDistance
}
