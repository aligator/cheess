package board

import "github.com/aligator/cheess/board/bit_board"

type Move struct {
	SourcePlayer *Player
	Source       bit_board.Coordinate
	SourcePice   PiceType

	TargetPlayer *Player
	Target       bit_board.Coordinate
	TargetPice   PiceType

	PromotedPice PiceType

	Capture          bool
	DoublePush       bool
	EnPassantCapture bool
	Castling         bool
}
