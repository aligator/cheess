package main

import (
	"fmt"

	"github.com/aligator/cheess/board"
	"github.com/aligator/cheess/board/bit_board"
	"github.com/aligator/cheess/board/lookup"
)

func main() {
	b := board.Board{
		White: board.Player{
			Color: board.PiceWhite,
			Queen: bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001),
		},
		Black: board.Player{
			Color: board.PiceBlack,
		},
	}

	fmt.Println(lookup.KingMove)

	for i := 0; i < 64; i++ {
		b.White.King = 1 << i
		m := b.NewMove(bit_board.Coordinate(i), 0)
		err := b.CheckMove(m)
		if err != nil {
			panic(err)
		}
	}
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
