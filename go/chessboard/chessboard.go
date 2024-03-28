package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.

// Define a struct named "File" to represent a file in a chessboard.
type File []bool

// Define a struct named "Chessboard" to represent a chessboard.
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	var c int
	for _, v := range cb[file] {
		if v {
			c++
		}
	}
	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var c int
	if rank >= 1 && rank <= 8 {
		for _, v := range cb {
			if v[rank-1] { c++ }
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var c int
	for _, f := range cb {
		for range f {
			c++
		}
	}
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var c int
	for _, f := range cb {
		for _, r := range f {
			if r { c++ }
		}
	}
	return c
}
