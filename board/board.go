package board

import "fmt"

type Color rune

const (
	White Color = 'w'
	Black Color = 'b'
)

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return "Unknown"
	}
}

type PieceType byte

const (
	None PieceType = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func (pt PieceType) String() string {
	switch pt {
	case None:
		return "None"
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	default:
		return "Unknown"
	}
}

type Player struct {
	King   BitBoard
	Queen  BitBoard
	Rook   BitBoard
	Bishop BitBoard
	Knight BitBoard
	Pawn   BitBoard
}

func BlackPlayer() Player {
	return Player{
		King:   BitBoard(0b00010000_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Queen:  BitBoard(0b00001000_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Rook:   BitBoard(0b10000001_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Bishop: BitBoard(0b00100100_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Knight: BitBoard(0b01000010_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Pawn:   BitBoard(0b00000000_11111111_00000000_00000000_00000000_00000000_00000000_00000000),
	}
}

func WhitePlayer() Player {
	return Player{
		King:   BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00010000),
		Queen:  BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00001000),
		Rook:   BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_10000001),
		Bishop: BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00100100),
		Knight: BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_01000010),
		Pawn:   BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_11111111_00000000),
	}
}

func (p Player) Get(position Coordinate) PieceType {
	// Ifs are ordered based on what types may be moved more often.
	// To be validated :-)

	if p.Pawn.Has(position) {
		return Pawn
	}
	if p.Bishop.Has(position) {
		return Bishop
	}
	if p.Knight.Has(position) {
		return Knight
	}
	if p.Queen.Has(position) {
		return Queen
	}
	if p.Rook.Has(position) {
		return Rook
	}
	if p.King.Has(position) {
		return King
	}

	return None
}

func (p Player) All() BitBoard {
	return (p.King | p.Queen | p.Rook | p.Bishop | p.Knight | p.Pawn)
}

func (p Player) String() string {
	return p.All().String()
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

// Move a piece from position1 to position2.
func (b Board) Move(position1, position2 Coordinate) {
	// Just check position in each bit board to find out where it belongs to.
	var color Color
	var piece PieceType
	if b.Black.All().Has(position1) {
		color = Black
		piece = b.Black.Get(position1)
	} else {
		color = White
		piece = b.White.Get(position1)
	}

	fmt.Printf("found %v %v\n", color, piece)
}

func (b Board) All() BitBoard {
	return b.Black.All() | b.White.All()
}

func (b Board) String() string {
	return b.All().String()
}
