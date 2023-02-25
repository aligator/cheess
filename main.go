package main

import (
	"fmt"

	"github.com/aligator/cheess/board"
)

func main() {
	b := board.New()
	fmt.Println(b.NewMove(board.At(0, 1), board.At(0, 3)))

	fmt.Println(b.NewMove(board.At(0, 6), board.At(0, 3)))
}

/*
ToDo:
- 1 BitBoard for every PiceType & color
- Implement basic BitBoard methods
- Maybe use big.Int as base type?
- Create a Full board representation using 8x8 squares, which gets updated after a successful move
- Add a Move method
- Implement a valid move checking
- Implement a first version of a move generator
- ...
*/
