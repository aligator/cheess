package bit_board

import (
	"fmt"
	"strconv"
)

type Coordinate byte

func At(x, y byte) Coordinate {
	return Coordinate(x + y*8)
}

func (c Coordinate) X() byte {
	return byte(c) % 8
}

func (c Coordinate) Y() byte {
	return byte(c) % 8
}

func (c Coordinate) String() string {
	return "(" + strconv.Itoa(int(c.X())) + "|" + strconv.Itoa(int(c.Y())) + ")"
}

// BitBoard representation with the bit at the coordinate set to 1.
func (c Coordinate) BitBoard() BitBoard {
	return BitBoard(0) | BitBoard(1)<<c
}

// BitBoard represents a chess board using just bits.
//
// The simplest way to map coordinates to bits is to use an 8x8 array of 64 bits,
// with each bit representing a position on the chess board.
// The bits can be numbered from 0-63, with the top left corner being bit 0
// and the bottom right corner being bit 63.
// The bits can then be mapped to the x and y coordinates of the board using
// simple math. For example, bit 0 would map to the coordinate (0,0),
// bit 7 would map to (7,0), bit 56 would map to (0,7),
// and bit 63 would map to (7,7).
//
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	8 |   56   |   57   |   58   |   59   |   60   |   61   |   62   |   63   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	7 |   48   |   49   |   50   |   51   |   52   |   53   |   54   |   55   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	6 |   40   |   41   |   42   |   43   |   44   |   45   |   46   |   47   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	5 |   32   |   33   |   34   |   35   |   36   |   37   |   38   |   39   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	4 |   24   |   25   |   26   |   27   |   28   |   29   |   30   |   31   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	3 |   16   |   17   |   18   |   19   |   20   |   21   |   22   |   23   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	2 |    8   |    9   |   10   |   11   |   12   |   13   |   14   |   15   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	1 |    0   |    1   |    2   |    3   |    4   |    5   |    6   |    7   |
//	  +--------+--------+--------+--------+--------+--------+--------+--------+
//	       a        b        c        d        e        f        g        h
type BitBoard uint64

func (b BitBoard) Has(pos Coordinate) bool {
	return (b>>pos)&1 == 1
}

func (b BitBoard) MoveAll(from, to Coordinate) BitBoard {
	// Before the "from" coordinate, it has to be moved "backwards"
	// and after it, it has to be moved "forward."
	if to >= from {
		b <<= BitBoard(to - from)
	} else {
		b >>= BitBoard(from - to)
	}

	return b
}

func (b BitBoard) String() string {
	bitString := fmt.Sprintf("%064b", b)
	result := ""
	for i, bit := range bitString {
		result += string(bit)
		if i%8 == 7 {
			result += "\n"
		}
	}

	return result
}
