package letter

import "strings"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) (out FreqMap) {
	c := make(chan FreqMap)
	go func(ch chan FreqMap) {
		bigstr := strings.Join(l, "")
		c <- Frequency(bigstr)
		close(c)
	}(c)
	return <-c
}
