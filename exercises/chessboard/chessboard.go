package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	counter := 0
	for _, x := range cb[file] {
		if x == true {
			counter += 1
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	counter := 0
	if rank <= 0 || rank > len(cb) {
		return 0
	}
	for _, x := range cb {
		if x[rank-1] == true {
			counter += 1
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for _, x := range cb {
		counter += len(x)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, x := range cb {
		for _, cell := range x {
			if cell == true {
				counter += 1
			}
		}
	}
	return counter
}
