package board

type Color = rune

const (
	White Color = 'w'
	Black Color = 'b'
)

type PieceType byte

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

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

// Move a piece from position 1 to position 2.
// Slower than MoveExact but no information about the piece to move is needed.
func (b Board) Move(x1, y1, x2, y2 byte) {
	// Just check position in each bit board -> call MoveExact with the correct info.
}

// MoveExact does the same as Move, but it doesn't need to find out the type and color of the pice.
// It is faster to use.
// Incorrect piece or color has undefined behavior.
func (b Board) MoveExact(piece PieceType, color Color, x1, y1, x2, y2 byte) {

}

func (b Board) All() BitBoard {
	return b.Black.All() | b.White.All()
}

func (b Board) String() string {
	return b.All().String()
}
