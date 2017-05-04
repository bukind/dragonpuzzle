package dragonpuzzle

import (
	"fmt"
)

// Side colors
const (
	ML = iota // no match at all (middle line)
	NC        // no color, match only itself
	RH        // red head
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
	S
	W
)

// internal bits
const (
	bno = 1 << iota
	bnc
	brh
	brt
	byh
	byt
	bgh
	bgt
)

// Dir is direction such as N, E, S, W
type Dir byte

// String is rep for direction.
func (d Dir) String() string {
	return "NESW"[d : d+1]
}

// Turn is to turn direction n times 90 degrees clockwise.
func (d Dir) Turn(n int) Dir {
	return Dir((int(d) + n) & 0x3)
}

// Return the number of rotations to do to turn x into d.
func (d Dir) Diff(x Dir) int {
	return int(d) + 4 - int(x)
}

// Side is the state of the side of the block
type Side byte

var bits = [...]byte{
	bno, bnc, brh, brt, byh, byt, bgh, bgt,
}

var sockNames = [...]string{
	"ML", "NC", "RH", "RT", "YH", "YT", "GH", "GT",
}

var masks = map[Side]byte{
	ML: 0,
	NC: bnc,
	RH: brt,
	RT: brh,
	YH: byt,
	YT: byh,
	GH: bgt,
	GT: bgh,
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
	return (s.bits()&x.mask() != 0)
}

// block is (north, east, west, south)
type Block struct {
	Sides [4]Side
}

// NewBlock creates a new block from N, E, S, W
func NewBlock(n, e, s, w Side) *Block {
	b := &Block{}
	b.Sides[N] = n
	b.Sides[E] = e
	b.Sides[S] = s
	b.Sides[W] = w
	return b
}

func (b *Block) String() string {
	return fmt.Sprintf("%s,%s,%s,%s", b.Sides[N], b.Sides[E], b.Sides[S], b.Sides[W])
}

func (b *Block) EqualTo(x *Block) bool {
	return b.Sides == x.Sides
}

// Match is to test if block matches another on direction dir
func (b *Block) Match(x *Block, dir Dir) bool {
	if x == nil || b == nil {
		return true
	}
	return b.Sides[dir].Match(x.Sides[dir.Turn(2)])
}

// Turn is to create a block, turned n times 90 degrees clockwise.
func (b *Block) Turn(n int) *Block {
	x := &Block{}
	for i := 0; i < 4; i++ {
		d := Dir(i)
		x.Sides[d.Turn(n)] = b.Sides[d]
	}
	return x
}

// Tile is a tile containing two blocks.
// The tile may be in 4 orientations given by Dir from A to B.
type Tile struct {
	A   *Block
	B   *Block
	Dir Dir // direction from b to a
}

// NewTile creates a tile.
func NewTile(n, e1, e2, s, w2, w1 Side, dir Dir) *Tile {
	t := &Tile{
		A:   NewBlock(n, e1, ML, w1),
		B:   NewBlock(ML, e2, s, w2),
		Dir: N,
	}
	return t.Turn(dir.Diff(t.Dir))
}

// Turn turns the tile.
func (t *Tile) Turn(n int) *Tile {
	return &Tile{t.A.Turn(n), t.B.Turn(n), t.Dir.Turn(n)}
}
