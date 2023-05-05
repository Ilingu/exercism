package cards

import (
	"math"
)

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if isOutOfBound(slice, index) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if isOutOfBound(slice, index) {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	return append(value, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if isOutOfBound(slice, index) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func isNegative(num int) bool {
	return math.Signbit(float64(num))
}

func isOutOfBound(slice []int, index int) bool {
	return isNegative(index) || len(slice) < index+1
}
