package board

import (
	"fmt"

	"github.com/aligator/cheess/board/bit_board"
	"github.com/aligator/cheess/board/lookup"
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

	King   bit_board.BitBoard
	Queen  bit_board.BitBoard
	Rook   bit_board.BitBoard
	Bishop bit_board.BitBoard
	Knight bit_board.BitBoard
	Pawn   bit_board.BitBoard
}

func BlackPlayer() Player {
	return Player{
		Color: PiceBlack,

		King:   bit_board.BitBoard(0b00010000_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Queen:  bit_board.BitBoard(0b00001000_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Rook:   bit_board.BitBoard(0b10000001_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Bishop: bit_board.BitBoard(0b00100100_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Knight: bit_board.BitBoard(0b01000010_00000000_00000000_00000000_00000000_00000000_00000000_00000000),
		Pawn:   bit_board.BitBoard(0b00000000_11111111_00000000_00000000_00000000_00000000_00000000_00000000),
	}
}

func WhitePlayer() Player {
	return Player{
		Color: PiceWhite,

		King:   bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00010000),
		Queen:  bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00001000),
		Rook:   bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_10000001),
		Bishop: bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00100100),
		Knight: bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_01000010),
		Pawn:   bit_board.BitBoard(0b00000000_00000000_00000000_00000000_00000000_00000000_11111111_00000000),
	}
}

func (p Player) GetType(position bit_board.Coordinate) PiceType {
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

func (p Player) Get(pice PiceType) bit_board.BitBoard {
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

	return bit_board.BitBoard(0)
}

func (p Player) All() bit_board.BitBoard {
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
func (b Board) NewMove(source, target bit_board.Coordinate) Move {
	// Just check position in each bit board to find out where it belongs to.
	var playerSource *Player
	if b.Black.All().Has(source) {
		playerSource = &b.Black
	} else {
		playerSource = &b.White
	}

	var playerTarget *Player
	if b.Black.All().Has(target) {
		playerTarget = &b.Black
	} else {
		playerTarget = &b.White
	}

	move := Move{
		SourcePlayer: playerSource,
		Source:       source,
		SourcePice:   playerSource.GetType(source),

		TargetPlayer: playerTarget,
		Target:       target,
		TargetPice:   playerTarget.GetType(target),
	}

	return move
}

func (b Board) CheckMove(move Move) error {
	// if move.SourcePice.Color() == move.TargetPice.Color() {
	// 	return errors.New("cannot move to own, occupied square")
	// }

	var canMoveTo bit_board.BitBoard
	switch move.SourcePice {
	case PiceBlackKing:
	case PiceWhiteKing:
		canMoveTo = b.computeKingIncomplete(move)
	case PiceBlackKnight:
	case PiceWhiteKnight:
		canMoveTo = b.computeKnightIncomplete(move)
	}

	fmt.Println(canMoveTo)
	return nil
}

func (b Board) computeKingIncomplete(move Move) bit_board.BitBoard {
	// 1. Move the KingMove LOT to the position of the king.
	var kingMoves bit_board.BitBoard = lookup.KingMove.MoveAll(lookup.MoveCenter, move.Source)

	// 2. Clip the fileA or fileH due to overflow on left and right side.
	x := move.Source.X()
	if x == 0 {
		kingMoves &= lookup.ClearFile[0]
	} else if x == 7 {
		kingMoves &= lookup.ClearFile[7]
	}

	// 3. Remove all own occupied positions.
	return kingMoves & ^move.SourcePlayer.All()
}

func (b Board) computeKnightIncomplete(move Move) bit_board.BitBoard {
	// 1. Move the KingMove LOT to the position of the king.
	var knightMoves bit_board.BitBoard = lookup.KnightMove.MoveAll(lookup.MoveCenter, move.Source)

	// 2. Clip the fileA/B or fileH/G due to overflow on left and right side.
	x := move.Source.X()
	if x <= 1 {
		knightMoves &= lookup.ClearFile[0] & lookup.ClearFile[1]
	} else if x >= 6 {
		knightMoves &= lookup.ClearFile[7] & lookup.ClearFile[6]
	}

	// 3. Remove all own occupied positions.
	return knightMoves & ^move.SourcePlayer.All()
}

func (b Board) All() bit_board.BitBoard {
	return b.Black.All() | b.White.All()
}

func (b Board) String() string {
	return b.All().String()
}
