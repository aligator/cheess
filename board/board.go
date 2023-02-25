package board

import (
	"errors"
)

// PiceType encodes the pice and color.
// If it is negative, it is white.
// If it is positive, it is black.
// If it is 0, it is None.
// -7 and 1 encode only the color without a type.
type PiceType int8

const (
	// PiceWhite encodes only the color, not which type it is.
	PiceWhite PiceType = iota - 7
	PiceWhiteKing
	PiceWhiteQueen
	PiceWhiteRook
	PiceWhiteBishop
	PiceWhiteKnight
	PiceWhitePawn

	PiceNone

	// PiceBlack encodes only the color, not which type it is.
	PiceBlack
	PiceBlackKing
	PiceBlackQueen
	PiceBlackRook
	PiceBlackBishop
	PiceBlackKnight
	PiceBlackPawn
)

// Color converts the pice type to the color-only pice type.
func (pt PiceType) Color() PiceType {
	if pt < 0 {
		return PiceWhite
	}
	if pt > 0 {
		return PiceBlack
	}
	return PiceNone
}

func (pt PiceType) IsWhite() bool {
	return pt < 0
}

func (pt PiceType) IsBlack() bool {
	return pt > 0
}

func (pt PiceType) String() string {
	switch pt {
	case PiceNone:
		return "None"
	case PiceWhiteKing:
		return "WhiteKing"
	case PiceWhiteQueen:
		return "WhiteQueen"
	case PiceWhiteRook:
		return "WhiteRook"
	case PiceWhiteBishop:
		return "WhiteBishop"
	case PiceWhiteKnight:
		return "WhiteKnight"
	case PiceWhitePawn:
		return "WhitePawn"
	case PiceBlackKing:
		return "BlackKing"
	case PiceBlackQueen:
		return "BlackQueen"
	case PiceBlackRook:
		return "BlackRook"
	case PiceBlackBishop:
		return "BlackBishop"
	case PiceBlackKnight:
		return "BlackKnight"
	case PiceBlackPawn:
		return "BlackPawn"
	case PiceBlack:
		return "Black"
	case PiceWhite:
		return "White"
	default:
		return "Unknown"
	}
}

type Player struct {
	Color PiceType

	King   BitBoard
	Queen  BitBoard
	Rook   BitBoard
	Bishop BitBoard
	Knight BitBoard
	Pawn   BitBoard
}

func BlackPlayer() Player {
	return Player{
		Color: PiceBlack,

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
		Color: PiceWhite,

		King:   BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00010000),
		Queen:  BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00001000),
		Rook:   BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_10000001),
		Bishop: BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00100100),
		Knight: BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_01000010),
		Pawn:   BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_11111111_00000000),
	}
}

func (p Player) GetType(position Coordinate) PiceType {
	// Ifs are ordered based on what types may be moved more often.
	// To be validated :-)

	var piceType PiceType

	if p.Pawn.Has(position) {
		piceType = PiceWhitePawn
	} else if p.Bishop.Has(position) {
		piceType = PiceWhiteBishop
	} else if p.Knight.Has(position) {
		piceType = PiceWhiteKnight
	} else if p.Queen.Has(position) {
		piceType = PiceWhiteQueen
	} else if p.Rook.Has(position) {
		piceType = PiceWhiteRook
	} else if p.King.Has(position) {
		piceType = PiceWhiteKing
	} else {
		return PiceNone
	}

	if p.Color == PiceWhite {
		return piceType
	}
	return piceType + 7 // Black is the same as white but shifted to > 0
}

func (p Player) Get(pice PiceType) BitBoard {
	switch pice {
	case PiceWhitePawn:
	case PiceBlackPawn:
		return p.Pawn
	case PiceWhiteBishop:
	case PiceBlackBishop:
		return p.Bishop
	case PiceWhiteKnight:
	case PiceBlackKnight:
		return p.Knight
	case PiceWhiteQueen:
	case PiceBlackQueen:
		return p.Queen
	case PiceWhiteRook:
	case PiceBlackRook:
		return p.Rook
	case PiceWhiteKing:
	case PiceBlackKing:
		return p.King
	}

	return BitBoard(0)
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

// NewMove reads all needed data for a move.
// It does not check anything.
func (b Board) NewMove(source, target Coordinate) Move {
	// Just check position in each bit board to find out where it belongs to.
	var playerSource Player
	if b.Black.All().Has(source) {
		playerSource = b.Black
	} else {
		playerSource = b.White
	}

	var playerTarget Player
	if b.Black.All().Has(target) {
		playerTarget = b.Black
	} else {
		playerTarget = b.White
	}

	move := Move{
		Source:     source,
		SourcePice: playerSource.GetType(source),
		Target:     target,
		TargetPice: playerTarget.GetType(target),
	}

	return move
}

func (b Board) CheckMove(move Move) error {
	if move.SourcePice.Color() == move.TargetPice.Color() {
		return errors.New("cannot move to own, occupied square")
	}

	return nil
}

func (b Board) All() BitBoard {
	return b.Black.All() | b.White.All()
}

func (b Board) String() string {
	return b.All().String()
}
