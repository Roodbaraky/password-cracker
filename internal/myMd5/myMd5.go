package myMd5

import (
	"math"
)

func init() {

}

const Size = 16

const BlockSize = 64

func F(b, c, d uint32) uint32 { return (b & c) | ((^b) & d) }
func G(b, c, d uint32) uint32 { return (b & d) | (c & (^d)) }
func H(b, c, d uint32) uint32 { return b ^ c ^ d }
func I(b, c, d uint32) uint32 { return c ^ (b | (^d)) }

func Sum(data []byte) []byte {
	rounds := [][]uint{
		{7, 12, 17, 22},
		{5, 9, 14, 20},
		{4, 11, 16, 23},
		{6, 10, 15, 21},
	}

	K := make([]uint32, 64)
	for i := 0; i < 64; i++ {
		K[i] = uint32(math.Floor(math.Abs(math.Sin(float64(i+1))) * (1 << 32)))
	}

	L := uint64(len(data) * 8)
	data = append(data, 0x80)
	for len(data)%BlockSize != 56 {
		data = append(data, 0x00)
	}

	s := make([]uint, 64)
	idx := 0
	for r := 0; r < 4; r++ {
		for rep := 0; rep < 4; rep++ {
			for k := 0; k < 4; k++ {
				s[idx] = rounds[r][k]
				idx++
			}
		}
	}

	for i := 0; i < 8; i++ {
		data = append(data, byte(L>>(8*i))&0xFF)
	}

	words := make([]uint32, len(data)/4)
	for i := 0; i < len(words); i++ {
		b0 := data[4*i]
		b1 := data[4*i+1]
		b2 := data[4*i+2]
		b3 := data[4*i+3]
		words[i] = uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24
	}
	A := uint32(0x67452301)
	B := uint32(0xEFCDAB89)
	C := uint32(0x98BADCFE)
	D := uint32(0x10325476)
	for j := 0; j < len(words); j += Size {
		M := words[j : j+Size]

		aa, bb, cc, dd := A, B, C, D

		for i := 0; i < BlockSize; i++ {
			var f uint32
			var g int
			switch {
			case i < Size:
				f = F(B, C, D)
				g = i
			case i < 32:
				f = G(B, C, D)
				g = (5*i + 1) % Size
			case i < 48:
				f = H(B, C, D)
				g = (3*i + 5) % Size
			default:
				f = I(B, C, D)
				g = (7 * i) % Size
			}

			temp := A + f + K[i] + M[g]
			temp = rotl(temp, s[i])
			A, B, C, D = D, B+temp, B, C

		}
		A += aa
		B += bb
		C += cc
		D += dd
	}

	out := make([]byte, 0, Size)
	out = append(out, byte(A), byte(A>>8), byte(A>>16), byte(A>>24))
	out = append(out, byte(B), byte(B>>8), byte(B>>16), byte(B>>24))
	out = append(out, byte(C), byte(C>>8), byte(C>>16), byte(C>>24))
	out = append(out, byte(D), byte(D>>8), byte(D>>16), byte(D>>24))
	return out
}

func rotl(x uint32, n uint) uint32 {
	return (x << n) | (x >> (32 - n))
}
