package main

import (
	"fmt"

	"github.com/aligator/cheess/board"
)

func main() {
	b := board.New()

	fmt.Println(b.White.Rook.Has(board.At(7, 0)), board.At(1, 0), b.White.Rook)
}

/*
ToDo:
- 1 BitBoard for every PieceType & color
- Implement basic BitBoard methods
- Maybe use big.Int as base type?
- Create a Full board representation using 8x8 squares, which gets updated after a successful move
- Add a Move method
- Implement a valid move checking
- Implement a first version of a move generator
- ...
*/
