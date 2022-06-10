package booking

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func Digits(str string) string {
	val, err := strconv.Atoi(str)
	if err != nil {
		return str
	}
	if string(str[0]) == "0" {
		return str
	}
	if val > 0 && val < 10 {
		return "0" + str
	}
	return str
}

func mapDigits(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		arr[i] = Digits(arr[i])
	}
	return arr
}

func formatStringToDate(StrToTransform string) string {
	Layer1Date := strings.Split(StrToTransform, " ")

	Layer2Date := mapDigits(reverseArray(strings.Split(Layer1Date[0], "/")))
	Layer3Date := reverseArray(Layer2Date[1:])
	Layer2Date = append(Layer2Date[:1], Layer3Date...)

	LayerR2Date := mapDigits(strings.Split(Layer1Date[1], ":"))

	Layer1Date[0] = strings.Join(Layer2Date, "-")
	Layer1Date[1] = strings.Join(LayerR2Date, ":")
	return strings.Join(Layer1Date, " ")
}

func parseDate(layout string, date string) time.Time {
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	date = formatStringToDate(date)
	parsedDate := parseDate("2006-01-02 15:04:05", date)
	return parsedDate
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layer1 := strings.Split(date, ",")
	layer2 := strings.Split(layer1[0], " ")

	layer2[1] = Digits(layer2[1])
	layer1[0] = strings.Join(layer2, " ")

	date = strings.Join(layer1, ",")
	parsedDate := parseDate("January 02, 2006 15:04:05", date)
	return time.Now().After(parsedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layer1 := strings.Split(date, ",")
	layer2 := strings.Split(layer1[1], " ")

	layer2[2] = Digits(layer2[2])
	layer1[1] = strings.Join(layer2, " ")

	date = strings.Join(layer1, ",")

	parsedDate := parseDate("Monday, January 02, 2006 15:04:05", date)
	HourInTheDay := parsedDate.Hour()
	return HourInTheDay >= 12 && HourInTheDay < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	parsedDate := Schedule(date)
	Day, Month, Date, Year, Hour, Min := parsedDate.Weekday(), parsedDate.Month(), parsedDate.Day(), parsedDate.Year(), parsedDate.Hour(), parsedDate.Minute()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", Day, Month, Date, Year, Hour, Min)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
