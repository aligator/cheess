package main

import (
	"fmt"

	"github.com/aligator/cheess/board"
)

func main() {
	b := board.New()
	fmt.Println(b)
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
