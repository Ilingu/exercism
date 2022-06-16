package twelve

import (
	"fmt"
	"strings"
)

var OrderedGave = [12]string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
var DigitToWord = map[int]string{1: "first", 2: "second", 3: "third", 4: "fourth", 5: "fifth", 6: "sixth", 7: "seventh", 8: "eighth", 9: "ninth", 10: "tenth", 11: "eleventh", 12: "twelfth"}

func Reverse(input []string) []string {
	var output []string
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}

func Verse(i int) string {
	sliceOfGaves := Reverse(OrderedGave[:i])
	if i > 1 {
		sliceOfGaves[len(sliceOfGaves)-1] = "and " + sliceOfGaves[len(sliceOfGaves)-1]
	}
	Gave := strings.Join(sliceOfGaves, ", ")

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", DigitToWord[i], Gave)
}

func Song() (out string) {
	song := make([]string, 0)
	for day := 1; day <= 12; day++ {
		song = append(song, Verse(day))
	}
	return strings.Join(song, "\n")
}
