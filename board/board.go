package board

import (
	"github.com/aligator/cheess/bit_board"
)

type Player struct {
	King   bit_board.BitBoard
	Queen  bit_board.BitBoard
	Rook   bit_board.BitBoard
	Bishop bit_board.BitBoard
	Knight bit_board.BitBoard
	Pawn   bit_board.BitBoard
}

func (p Player) String() string {
	return p.All().String()
}

func (p Player) All() bit_board.BitBoard {
	return (p.King | p.Queen | p.Rook | p.Bishop | p.Knight | p.Pawn)
}

func BlackPlayer() Player {
	return Player{
		King:   bit_board.New(0b00010000_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Queen:  bit_board.New(0b00001000_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Rook:   bit_board.New(0b10000001_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Bishop: bit_board.New(0b00100100_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Knight: bit_board.New(0b01000010_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Pawn:   bit_board.New(0b00000000_11111111_00000000_00000000_00000000_00000000_00000000_00000000),
	}
}

func WhitePlayer() Player {
	return Player{
		King:   bit_board.New(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00010000),
		Queen:  bit_board.New(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00001000),
		Rook:   bit_board.New(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_10000001),
		Bishop: bit_board.New(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00100100),
		Knight: bit_board.New(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_01000010),
		Pawn:   bit_board.New(0b00000000_00000000_00000000_00000000_00000000_00000000_11111111_00000000),
	}
}

type Board struct {
	White Player
	Black Player
}

func New() Board {
	return Board{
		Black: BlackPlayer(),
		White: WhitePlayer(),
	}
}
