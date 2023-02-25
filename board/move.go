package board

type Move struct {
	Source     Coordinate
	SourcePice PiceType
	Target     Coordinate
	TargetPice PiceType

	PromotedPice PiceType

	Capture          bool
	DoublePush       bool
	EnPassantCapture bool
	Castling         bool
}
