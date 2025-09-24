package pbits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitsAndBytes(t *testing.T) {
	cases := []struct {
		m     Mask
		bits  int
		bytes int
	}{
		{NoMask, 0, 0},
		{Mask1, 1, 1},
		{Mask8, 8, 1},
		{Mask9, 9, 2},
		{Mask16, 16, 2},
		{Mask24, 24, 3},
		{Mask32, 32, 4},
		{Mask33, 33, 5},
		{Mask64, 64, 8},
	}
	for _, c := range cases {
		assert.Equal(t, c.bits, c.m.BitsCount(), "BitsCount mismatch for %v", c.m)
		assert.Equal(t, c.bytes, c.m.BytesCount(), "BytesCount mismatch for %v", c.m)
	}
}

func Test24BitsInto32(t *testing.T) {
	mvIn24Bits := uint32(Mask24.MaxValue())
	mvIn8Bits := uint32(Mask8.MaxValue())
	// 24+8 bits require 32 bits (uint32)
	v1 := Pack[uint32](mvIn24Bits, NoMask, Mask24)
	v2 := Pack[uint32](mvIn8Bits, Mask24, Mask8)
	valueOn32Bits := v1 | v2
	uv1 := Unpack[uint32](valueOn32Bits, NoMask, Mask24)
	uv2 := Unpack[uint32](valueOn32Bits, Mask24, Mask8)
	assert.Equal(t, mvIn24Bits, uv1, "unpacked v1 values not match")
	assert.Equal(t, mvIn8Bits, uv2, "unpacked v2 values not match")
}
func TestPackMultipleValue(t *testing.T) {
	// We pack the example values specified by MaxValue for a given mask,
	// we want to compress all these values into one 64-bit number.
	// The total number of bits needed to compress these numbers is just 64 bits
	// storing values in bigendian format,
	// which allows you to use the data stored in this way to create, for example, sorted keys in keyvalue databases
	packed := Pack[uint64](Mask3.MaxValue(), NoMask, Mask3) |
		Pack[uint64](Mask5.MaxValue(), Mask3, Mask5) |
		Pack[uint64](Mask2.MaxValue(), Mask5, Mask2) |
		Pack[uint64](Mask8.MaxValue(), Mask2, Mask8) |
		Pack[uint64](Mask16.MaxValue(), Mask8, Mask16) |
		Pack[uint64](Mask16.MaxValue(), Mask16, Mask13) |
		Pack[uint64](Mask16.MaxValue(), Mask13, Mask4) |
		Pack[uint64](Mask16.MaxValue(), Mask4, Mask12)

	// unpack in the same order as we packed
	assert.Equal(t, Mask3.MaxValue(), Unpack[uint64](packed, NoMask, Mask3))
	assert.Equal(t, Mask5.MaxValue(), Unpack[uint64](packed, Mask3, Mask5))
	assert.Equal(t, Mask2.MaxValue(), Unpack[uint64](packed, Mask5, Mask2))
	assert.Equal(t, Mask8.MaxValue(), Unpack[uint64](packed, Mask2, Mask8))
	assert.Equal(t, Mask16.MaxValue(), Unpack[uint64](packed, Mask8, Mask16))
	assert.Equal(t, Mask13.MaxValue(), Unpack[uint64](packed, Mask16, Mask13))
	assert.Equal(t, Mask4.MaxValue(), Unpack[uint64](packed, Mask13, Mask4))
	assert.Equal(t, Mask12.MaxValue(), Unpack[uint64](packed, Mask4, Mask12))

	// another try
	packed = Pack[uint64](Mask6.MaxValue(), NoMask, Mask6) |
		Pack[uint64](Mask7.MaxValue(), Mask6, Mask7) |
		Pack[uint64](Mask10.MaxValue(), Mask7, Mask10) |
		Pack[uint64](Mask11.MaxValue(), Mask10, Mask11) |
		Pack[uint64](Mask14.MaxValue(), Mask11, Mask14) |
		Pack[uint64](Mask15.MaxValue(), Mask14, Mask15) |
		Pack[uint64](Mask12.MaxValue(), Mask15, Mask12)

	assert.Equal(t, Mask6.MaxValue(), Unpack[uint64](packed, NoMask, Mask6))
	assert.Equal(t, Mask7.MaxValue(), Unpack[uint64](packed, Mask6, Mask7))
	assert.Equal(t, Mask10.MaxValue(), Unpack[uint64](packed, Mask7, Mask10))
	assert.Equal(t, Mask11.MaxValue(), Unpack[uint64](packed, Mask10, Mask11))
	assert.Equal(t, Mask14.MaxValue(), Unpack[uint64](packed, Mask11, Mask14))
	assert.Equal(t, Mask15.MaxValue(), Unpack[uint64](packed, Mask14, Mask15))
	assert.Equal(t, Mask12.MaxValue(), Unpack[uint64](packed, Mask15, Mask12))

	// another try
	packed = Pack[uint64](Mask9.MaxValue(), NoMask, Mask9) |
		Pack[uint64](Mask17.MaxValue(), Mask9, Mask17) |
		Pack[uint64](Mask18.MaxValue(), Mask17, Mask18) |
		Pack[uint64](Mask19.MaxValue(), Mask18, Mask19)

	assert.Equal(t, Mask9.MaxValue(), Unpack[uint64](packed, NoMask, Mask9))
	assert.Equal(t, Mask17.MaxValue(), Unpack[uint64](packed, Mask9, Mask17))
	assert.Equal(t, Mask18.MaxValue(), Unpack[uint64](packed, Mask17, Mask18))
	assert.Equal(t, Mask19.MaxValue(), Unpack[uint64](packed, Mask18, Mask19))

	// another try
	packed = Pack[uint64](Mask20.MaxValue(), NoMask, Mask20) |
		Pack[uint64](Mask21.MaxValue(), Mask20, Mask21) |
		Pack[uint64](Mask22.MaxValue(), Mask21, Mask22)

	assert.Equal(t, Mask20.MaxValue(), Unpack[uint64](packed, NoMask, Mask20))
	assert.Equal(t, Mask21.MaxValue(), Unpack[uint64](packed, Mask20, Mask21))
	assert.Equal(t, Mask22.MaxValue(), Unpack[uint64](packed, Mask21, Mask22))

	// another try
	packed = Pack[uint64](Mask23.MaxValue(), NoMask, Mask23) |
		Pack[uint64](Mask25.MaxValue(), Mask23, Mask25)

	assert.Equal(t, Mask23.MaxValue(), Unpack[uint64](packed, NoMask, Mask23))
	assert.Equal(t, Mask25.MaxValue(), Unpack[uint64](packed, Mask23, Mask25))

	// another try
	packed = Pack[uint64](Mask28.MaxValue(), NoMask, Mask28) |
		Pack[uint64](Mask29.MaxValue(), Mask28, Mask29)

	assert.Equal(t, Mask28.MaxValue(), Unpack[uint64](packed, NoMask, Mask28))
	assert.Equal(t, Mask29.MaxValue(), Unpack[uint64](packed, Mask28, Mask29))

	// another try
	packed = Pack[uint64](Mask26.MaxValue(), NoMask, Mask26) |
		Pack[uint64](Mask27.MaxValue(), Mask26, Mask27)

	assert.Equal(t, Mask26.MaxValue(), Unpack[uint64](packed, NoMask, Mask26))
	assert.Equal(t, Mask27.MaxValue(), Unpack[uint64](packed, Mask26, Mask27))

	// another try
	packed = Pack[uint64](Mask30.MaxValue(), NoMask, Mask30) |
		Pack[uint64](Mask32.MaxValue(), Mask30, Mask32)

	assert.Equal(t, Mask30.MaxValue(), Unpack[uint64](packed, NoMask, Mask30))
	assert.Equal(t, Mask32.MaxValue(), Unpack[uint64](packed, Mask30, Mask32))

	// another try
	packed = Pack[uint64](Mask62.MaxValue(), NoMask, Mask62) |
		Pack[uint64](Mask1.MaxValue(), Mask62, Mask1)

	assert.Equal(t, Mask62.MaxValue(), Unpack[uint64](packed, NoMask, Mask62))
	assert.Equal(t, Mask1.MaxValue(), Unpack[uint64](packed, Mask62, Mask1))

	// another try
	packed = Pack[uint64](Mask63.MaxValue(), NoMask, Mask63) |
		Pack[uint64](Mask1.MaxValue(), Mask63, Mask1)

	assert.Equal(t, Mask63.MaxValue(), Unpack[uint64](packed, NoMask, Mask63))
	assert.Equal(t, Mask1.MaxValue(), Unpack[uint64](packed, Mask63, Mask1))

	// another try
	packed = Pack[uint64](Mask63.MaxValue(), NoMask, Mask63) |
		Pack[uint64](Mask2.MaxValue(), Mask63, Mask2)

	assert.Equal(t, Mask63.MaxValue(), Unpack[uint64](packed, NoMask, Mask63))
	// overflow
	assert.NotEqual(t, Mask2.MaxValue(), Unpack[uint64](packed, Mask63, Mask2))

	// another try
	packed = Pack[uint64](Mask62.MaxValue(), NoMask, Mask62) |
		Pack[uint64](Mask2.MaxValue(), Mask62, Mask2)

	assert.Equal(t, Mask62.MaxValue(), Unpack[uint64](packed, NoMask, Mask62))
	assert.Equal(t, Mask2.MaxValue(), Unpack[uint64](packed, Mask62, Mask2))

	// another try
	packed = Pack[uint64](Mask33.MaxValue(), NoMask, Mask33) |
		Pack[uint64](Mask1.MaxValue(), Mask33, Mask1)

	assert.Equal(t, Mask33.MaxValue(), Unpack[uint64](packed, NoMask, Mask33))
	assert.Equal(t, Mask1.MaxValue(), Unpack[uint64](packed, Mask33, Mask1))

	// another try
	packed = Pack[uint64](Mask32.MaxValue(), NoMask, Mask32) |
		Pack[uint64](Mask32.MaxValue(), Mask32, Mask32)

	assert.Equal(t, Mask32.MaxValue(), Unpack[uint64](packed, NoMask, Mask32))
	assert.Equal(t, Mask32.MaxValue(), Unpack[uint64](packed, Mask32, Mask32))

	// another try
	packed = Pack[uint64](Mask34.MaxValue(), NoMask, Mask34) |
		Pack[uint64](Mask30.MaxValue(), Mask34, Mask30)

	assert.Equal(t, Mask34.MaxValue(), Unpack[uint64](packed, NoMask, Mask34))
	assert.Equal(t, Mask30.MaxValue(), Unpack[uint64](packed, Mask34, Mask30))

	// max uints
	assert.Equal(t, Mask64.MaxValue(), Unpack[uint64](^uint64(0), NoMask, Mask64))
	assert.Equal(t, Mask32.MaxValue(), Unpack[uint64](uint64(^uint32(0)), NoMask, Mask32))
	assert.Equal(t, Mask16.MaxValue(), Unpack[uint64](uint64(^uint16(0)), NoMask, Mask16))

	// another try
	packed = Pack[uint64](Mask48.MaxValue(), NoMask, Mask48)

	assert.Equal(t, Mask48.MaxValue(), Unpack[uint64](packed, NoMask, Mask48))
	// rest bits are zero, so value is also zero
	assert.Equal(t, NoMask.MaxValue(), Unpack[uint64](packed, Mask48, Mask16))

}

// overflow test
func TestMask_Protect(t *testing.T) {
	for i, tm := range tableOfMasks {

		testname := fmt.Sprintf("MaskProtect,%d", tm.bitsCount)
		t.Run(testname, func(t *testing.T) {
			m := Mask(i)
			assert.Equal(t, m.MaxValue(), m.Protect(m.MaxValue()))
			if i == 0 {
				assert.Equal(t, m.MaxValue(), m.Protect(m.MaxValue()+1))
			} else {
				assert.NotEqual(t, m.MaxValue(), m.Protect(m.MaxValue()+1))
			}
		})
	}
}
