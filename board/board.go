package board

import "github.com/aligator/cheess/bit_board"

type Player struct {
	King   bit_board.BitBoard
	Queen  bit_board.BitBoard
	Rook   bit_board.BitBoard
	Bishop bit_board.BitBoard
	Knight bit_board.BitBoard
	Pawn   bit_board.BitBoard
}

func BlackPlayer() Player {
	// The simplest way to map coordinates to bits is to use an 8x8 array of 64 bits,
	// with each bit representing a position on the chess board.
	// The bits can be numbered from 0-63, with the top left corner being bit 0
	// and the bottom right corner being bit 63.
	// The bits can then be mapped to the x and y coordinates of the board using
	// simple math. For example, bit 0 would map to the coordinate (0,0),
	// bit 7 would map to (7,0), bit 56 would map to (0,7),
	// and bit 63 would map to (7,7).

	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |   0    |   1    |   2    |   3    |   4    |   5    |   6    |   7    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |   8    |   9    |  10    |  11    |  12    |  13    |  14    |  15    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |  16    |  17    |  18    |  19    |  20    |  21    |  22    |  23    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |  24    |  25    |  26    |  27    |  28    |  29    |  30    |  31    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |  32    |  33    |  34    |  35    |  36    |  37    |  38    |  39    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |  40    |  41    |  42    |  43    |  44    |  45    |  46    |  47    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |  48    |  49    |  50    |  51    |  52    |  53    |  54    |  55    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+
	// |  56    |  57    |  58    |  59    |  60    |  61    |  62    |  63    |
	// +--------+--------+--------+--------+--------+--------+--------+--------+

	return Player{
		King:   bit_board.New(0),
		Queen:  bit_board.New(0),
		Rook:   bit_board.New(0),
		Bishop: bit_board.New(0),
		Knight: bit_board.New(0),
		Pawn:   bit_board.New(0),
	}
}

func WhitePlayer() Player {
	return Player{}
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
