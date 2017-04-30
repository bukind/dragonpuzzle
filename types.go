package dragonpuzzle

import (
	"fmt"
)

// Side colors
const (
	AC = iota  // any color, match anything
	NC         // no color, match only itself
	RH         // red head
	RT
	YH
	YT
	GH
	GT
	MAXX
)

// Directions
const (
	N = iota
	E
	W
	S
)

// internal bits
const (
	banys = 1 << iota
	bnone
	bredh
	bredt
	byelh
	byelt
	bgrnh
	bgrnt
)

// Side is the state of the side of the block
type Side byte

var bits = [...]byte {
	banys, bnone, bredh, bredt, byelh, byelt, bgrnh, bgrnt,
}

var sockNames = [...]string{
	"AC", "NC", "RH", "RT", "YH", "YT", "GH", "GT",
}

var masks = map[Side]byte {
	AC: banys | bnone | bredh | bredt | byelh | byelt | bgrnh | bgrnt,
	NC: banys | bnone,
	RH: banys | bredt,
	RT: banys | bredh,
	YH: banys | byelt,
	YT: banys | byelh,
	GH: banys | bgrnt,
	GT: banys | bgrnh,
}

// String is the string representation of Side.
func (s Side) String() string {
	return sockNames[s]
}

// bits is bit value of the side.
func (s Side) bits() byte {
	return bits[s]
}

// mask is bitmask which is used to check for side matching.
func (s Side) mask() byte {
	return masks[s]
}

// Match is to check if a side match another.
func (s Side) Match(x Side) bool {
	return (s.bits() & x.mask() != 0)
}

// block is (north, east, west, south)
type Block struct {
	Sides [4]Side
}

// NewBlock creates a new block from N, E, S, W
func NewBlock(n, e, s, w Side) *Block {
	b := &Block{}
	b.Sides[0] = n
	b.Sides[1] = e
	b.Sides[2] = s
	b.Sides[3] = w
	return b
}

func (b *Block) String() string {
	return fmt.Sprintf("%s,%s,%s,%s", b.Sides[0], b.Sides[1], b.Sides[2], b.Sides[3])
}

func (b *Block) EqualTo(x *Block) bool {
	return b.Sides == x.Sides
}

// Match is to test if block matches another on direction dir
func (b *Block) Match(x *Block, dir int) bool {
	if x == nil || b == nil {
		return true
	}
	return b.Sides[dir].Match(x.Sides[(dir + 2) & 0x3])
}

// Turn is to create a block, turned 90 \grad n times clockwise.
func (b *Block) Turn(times int) *Block {
	x := &Block{}
	for i := 0; i < 4; i++ {
		x.Sides[(i+times) & 3] = b.Sides[i]
	}
	return x
}
