package beer

import (
	"errors"
	"fmt"
)

func Song() string {
	song, _ := Verses(99, 0)
	return song
}

func Verses(start, stop int) (string, error) {
	if stop >= start {
		return "", errors.New("start should be > stop")
	}

	var verses string
	for i := start; i >= stop; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}

		verses += verse + "\n"
	}
	return verses, nil
}

func Verse(n int) (string, error) {
	switch {
	case n > 99, n < 0:
		return "", errors.New("out of range")
	case n == 2:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", n, n, n-1), nil
	case n == 1:
		return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", n, n), nil
	case n == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}
