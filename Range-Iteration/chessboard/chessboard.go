package chessboard

import "fmt"

/* Visualize:

	1 2 3 4 5 6 7 8
A # _ # _ _ _ _ # A
B _ _ _ _ # _ _ _ B
C _ _ # _ _ _ _ _ C
D _ _ _ _ _ _ _ _ D
E _ _ _ _ _ # _ # E
F _ _ _ _ _ _ _ _ F
G _ _ _ # _ _ _ _ G
H # # # # # # _ # H
	1 2 3 4 5 6 7 8

*/

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (num int) {
	RankRows, Exist := cb[rank]
	if !Exist {
		return 0
	}
	for _, row := range RankRows {
		if row {
			num++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (num int) {
	for _, row := range cb {
		if file > len(row) {
			return 0
		}
		if row[file-1] {
			num++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (num int) {
	for i := 1; i <= len(cb); i++ {
		fmt.Println(i, CountInFile(cb, i))
		num += CountInFile(cb, i)
	}
	return
}
