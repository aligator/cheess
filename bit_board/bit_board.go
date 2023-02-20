package bit_board

import (
	"math/big"
	"unsafe"
)

type BitBoard big.Int

func New(values uint64) BitBoard {
	// Cast without changing the bit-representation.
	intValues := *(*int64)(unsafe.Pointer(&values))
	return BitBoard(*big.NewInt(intValues))
}
