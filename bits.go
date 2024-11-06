package pbits

import (
	"fmt"
	"math"
)

// Mask defines the maximum number that can be stored in bit packing,
// This is quite problematic, so it is better to use constant arrays
// e.g. to encode the number 1939 using the smallest possible number of bits in a 64-bit number you need to use this:
// x := 1939 << 53 & Mask11.MaxValue() so encoding 1939 requires Mask11 because it stores numbers from 0-2047
// the value 53 is taken from this,  that we want to put the result in a 64-bit number,
// i.e. 64bits-Mask11 gives us that we start saving the first bit from position
// no. 53 In the BigEndian format and so the last bits are the first ones,
// so we should write the sorting value, i.e. e.g. timstamp from the 64th bit down,
// because after writing to disk they will be the first
type Mask uint

const (
	// NoMask max value: 0, bitsCount: 0, memory use: 0 byte(s) size: 0
	NoMask Mask = iota

	// Mask1 max value: 1, bitsCount: 1, memory use: 1 byte(s) size: 8
	Mask1

	// Mask2 max value: 3, bitsCount: 2, memory use: 1 byte(s) size: 8
	Mask2

	// Mask3 max value: 7, bitsCount: 3, memory use: 1 byte(s) size: 8
	Mask3

	// Mask4 max value: 15, bitsCount: 4, memory use: 1 byte(s) size: 8
	Mask4

	// Mask5 max value: 31, bitsCount: 5, memory use: 1 byte(s) size: 8
	Mask5

	// Mask6 max value: 63, bitsCount: 6, memory use: 1 byte(s) size: 8
	Mask6

	// Mask7 max value: 127, bitsCount: 7, memory use: 1 byte(s) size: 8
	Mask7

	// Mask8 max value: 255, bitsCount: 8, memory use: 1 byte(s) size: 8
	Mask8

	// Mask9 max value: 511, bitsCount: 9, memory use: 2 byte(s) size: 16
	Mask9

	// Mask10 max value: 1023, bitsCount: 10, memory use: 2 byte(s) size: 16
	Mask10

	// Mask11 max value: 2047, bitsCount: 11, memory use: 2 byte(s) size: 16
	Mask11

	// Mask12 max value: 4095, bitsCount: 12, memory use: 2 byte(s) size: 16
	Mask12

	// Mask13 max value: 8191, bitsCount: 13, memory use: 2 byte(s) size: 16
	Mask13

	// Mask14 max value: 16383, bitsCount: 14, memory use: 2 byte(s) size: 16
	Mask14

	// Mask15 max value: 32767, bitsCount: 15, memory use: 2 byte(s) size: 16
	Mask15

	// Mask16 max value: 65535, bitsCount: 16, memory use: 2 byte(s) size: 16
	Mask16

	// Mask17 max value: 131071, bitsCount: 17, memory use: 3 byte(s) size: 32
	Mask17

	// Mask18 max value: 262143, bitsCount: 18, memory use: 3 byte(s) size: 32
	Mask18

	// Mask19 max value: 524287, bitsCount: 19, memory use: 3 byte(s) size: 32
	Mask19

	// Mask20 max value: 1048575, bitsCount: 20, memory use: 3 byte(s) size: 32
	Mask20

	// Mask21 max value: 2097151, bitsCount: 21, memory use: 3 byte(s) size: 32
	Mask21

	// Mask22 max value: 4194303, bitsCount: 22, memory use: 3 byte(s) size: 32
	Mask22

	// Mask23 max value: 8388607, bitsCount: 23, memory use: 3 byte(s) size: 32
	Mask23

	// Mask24 max value: 16777215, bitsCount: 24, memory use: 3 byte(s) size: 32
	Mask24

	// Mask25 max value: 33554431, bitsCount: 25, memory use: 4 byte(s) size: 32
	Mask25

	// Mask26 max value: 67108863, bitsCount: 26, memory use: 4 byte(s) size: 32
	Mask26

	// Mask27 max value: 134217727, bitsCount: 27, memory use: 4 byte(s) size: 32
	Mask27

	// Mask28 max value: 268435455, bitsCount: 28, memory use: 4 byte(s) size: 32
	Mask28

	// Mask29 max value: 536870911, bitsCount: 29, memory use: 4 byte(s) size: 32
	Mask29

	// Mask30 max value: 1073741823, bitsCount: 30, memory use: 4 byte(s) size: 32
	Mask30

	// Mask31 max value: 2147483647, bitsCount: 31, memory use: 4 byte(s) size: 32
	Mask31

	// Mask32 max value: 4294967295, bitsCount: 32, memory use: 4 byte(s) size: 32
	Mask32

	// Mask33 max value: 8589934591, bitsCount: 33, memory use: 5 byte(s) size: 64
	Mask33

	// Mask34 max value: 17179869183, bitsCount: 34, memory use: 5 byte(s) size: 64
	Mask34

	// Mask35 max value: 34359738367, bitsCount: 35, memory use: 5 byte(s) size: 64
	Mask35

	// Mask36 max value: 68719476735, bitsCount: 36, memory use: 5 byte(s) size: 64
	Mask36

	// Mask37 max value: 137438953471, bitsCount: 37, memory use: 5 byte(s) size: 64
	Mask37

	// Mask38 max value: 274877906943, bitsCount: 38, memory use: 5 byte(s) size: 64
	Mask38

	// Mask39 max value: 549755813887, bitsCount: 39, memory use: 5 byte(s) size: 64
	Mask39

	// Mask40 max value: 1099511627775, bitsCount: 40, memory use: 5 byte(s) size: 64
	Mask40

	// Mask41 max value: 2199023255551, bitsCount: 41, memory use: 6 byte(s) size: 64
	Mask41

	// Mask42 max value: 4398046511103, bitsCount: 42, memory use: 6 byte(s) size: 64
	Mask42

	// Mask43 max value: 8796093022207, bitsCount: 43, memory use: 6 byte(s) size: 64
	Mask43

	// Mask44 max value: 17592186044415, bitsCount: 44, memory use: 6 byte(s) size: 64
	Mask44

	// Mask45 max value: 35184372088831, bitsCount: 45, memory use: 6 byte(s) size: 64
	Mask45

	// Mask46 max value: 70368744177663, bitsCount: 46, memory use: 6 byte(s) size: 64
	Mask46

	// Mask47 max value: 140737488355327, bitsCount: 47, memory use: 6 byte(s) size: 64
	Mask47

	// Mask48 max value: 281474976710655, bitsCount: 48, memory use: 6 byte(s) size: 64
	Mask48

	// Mask49 max value: 562949953421311, bitsCount: 49, memory use: 7 byte(s) size: 64
	Mask49

	// Mask50 max value: 1125899906842623, bitsCount: 50, memory use: 7 byte(s) size: 64
	Mask50

	// Mask51 max value: 2251799813685247, bitsCount: 51, memory use: 7 byte(s) size: 64
	Mask51

	// Mask52 max value: 4503599627370495, bitsCount: 52, memory use: 7 byte(s) size: 64
	Mask52

	// Mask53 max value: 9007199254740991, bitsCount: 53, memory use: 7 byte(s) size: 64
	Mask53

	// Mask54 max value: 18014398509481983, bitsCount: 54, memory use: 7 byte(s) size: 64
	Mask54

	// Mask55 max value: 36028797018963967, bitsCount: 55, memory use: 7 byte(s) size: 64
	Mask55

	// Mask56 max value: 72057594037927935, bitsCount: 56, memory use: 8 byte(s) size: 64
	Mask56

	// Mask57 max value: 144115188075855871, bitsCount: 57, memory use: 8 byte(s) size: 64
	Mask57

	// Mask58 max value: 288230376151711743, bitsCount: 58, memory use: 8 byte(s) size: 64
	Mask58

	// Mask59 max value: 576460752303423487, bitsCount: 59, memory use: 8 byte(s) size: 64
	Mask59

	// Mask60 max value: 1152921504606846975, bitsCount: 60, memory use: 8 byte(s) size: 64
	Mask60

	// Mask61 max value: 2305843009213693951, bitsCount: 61, memory use: 8 byte(s) size: 64
	Mask61

	// Mask62 max value: 4611686018427387903, bitsCount: 62, memory use: 8 byte(s) size: 64
	Mask62

	// Mask63 max value: 9223372036854775807, bitsCount: 63, memory use: 8 byte(s) size: 64
	Mask63

	// Mask64 max value: 18446744073709551615, bitsCount: 64, memory use: 8 byte(s) size: 64
	Mask64
)

// const MaskSignedBit = 0b1111111111111111111111111111111111111111111111111111111111111111

// MaxValue max value that can be stored in this mask
func (self Mask) MaxValue() uint64 {
	return tableOfMasks[self].value
}

func (self Mask) BytesCount() int {
	return tableOfMasks[self].bytesCount
}

// Size int size type
func (self Mask) Size() int {
	return tableOfMasks[self].size
}

// BitsCount how many bitsCount you need sto store value
// BitsCount is eault to Mask value
func (self Mask) BitsCount() int {
	return tableOfMasks[self].size
}

// Protect guarantees to return a number in the mask range if the higher supply returns zero
func (self Mask) Protect(val uint64) uint64 {
	return val & uint64(self)
}

var tableOfMasks = [65]item{}

type item struct {
	// mask value
	value uint64
	// bitsCount how many bitsCount you need sto store value
	bitsCount uint
	// bytesCount how many bytes you need to store value
	bytesCount int
	// size int size type
	size int
}

func Value(m Mask) uint64 {
	return tableOfMasks[m].value
}

func init() {
	// initialize all masks
	for i := range tableOfMasks {
		vv := powUint64(2, i) - 1
		bc := int(math.Log2(float64(vv)) / 8)
		if i > 0 {
			bc++
		} else {
			bc = 0
		}

		var x int
		switch bc {
		case 0:
			x = 0
			break
		case 1:
			x = 8
			break
		case 2:
			x = 16
			break
		case 3:
			x = 32
			break
		case 4:
			x = 32
			break
		case 5:
			x = 64
			break
		case 6:
			x = 64
			break
		case 7:
			x = 64
			break
		case 8:
			x = 64
			break
		default:
			panic("not supported int size")
		}
		if i == 64 {
			vv = ^uint64(0)
		}
		tableOfMasks[i] = item{value: uint64(vv),
			bitsCount:  uint(i),
			bytesCount: bc,
			size:       x,
		}
		// simple tool to generate code
		fmt.Printf("\t// Mask%d max value: %d, bitsCount: %d, memory use: %d byte(s) size: %d\n\tMask%d\n\n",
			i,
			tableOfMasks[i].value,
			tableOfMasks[i].bitsCount,
			tableOfMasks[i].bytesCount,
			tableOfMasks[i].size,
			i,
		)
	}

}

// Pack @value of max value defined by @mask
// if it is first value in packed integer @prevMask should be equal NoMask
func Pack[T uint | uint8 | uint16 | uint32 | uint64](value T, prevMask Mask, mask Mask) T {
	return T(uint64(value) & mask.MaxValue() << uint64(prevMask))
}

// Unpack uint from @value, which max value is defined by @mask
// if it is first unpacking value @prevMask should be equal NoMask
func Unpack[T uint | uint8 | uint16 | uint32 | uint64](value T, prevMask Mask, mask Mask) T {
	if prevMask == NoMask {
		return value & T(mask.MaxValue())
	}
	return T(uint64(value) >> prevMask & mask.MaxValue())
}

func powUint64(x, y int) uint64 {
	return uint64(math.Pow(float64(x), float64(y)))
}
