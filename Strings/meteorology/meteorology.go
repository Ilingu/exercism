package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temp TemperatureUnit) String() string {
	switch temp {
	case 0:
		return "°C"
	default:
		return "°F"
	}
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {
	return fmt.Sprintf("%d %s", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speed SpeedUnit) String() string {
	switch speed {
	case 0:
		return "km/h"
	default:
		return "mph"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (data *MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%s Humidity", data.location, data.temperature, data.windDirection, data.windSpeed, data.humidity, "%")
}
