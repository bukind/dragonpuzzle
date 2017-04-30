package dragonpuzzle

// sockets
const (
	ANYS = iota  // does not matter
	NONE
	REDH
	REDT
	YELH
	YELT
	GRNH
	GRNT
	MAXX
)

const (
	N = iota
	E
	W
	S
)

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

type Sock byte

var bits = [...]byte {
	banys, bnone, bredh, bredt, byelh, byelt, bgrnh, bgrnt,
}

var sockNames = [...]string{
	"ANYS", "NONE", "REDH", "REDT", "YELH", "YELT", "GRNH", "GRNT",
}

var masks = map[Sock]byte {
	ANYS: banys | bnone | bredh | bredt | byelh | byelt | bgrnh | bgrnt,
	NONE: banys | bnone,
	REDH: banys | bredt,
	REDT: banys | bredh,
	YELH: banys | byelt,
	YELT: banys | byelh,
	GRNH: banys | bgrnt,
	GRNT: banys | bgrnh,
}

//            0      1      2     3
// block is (north, east, west, south)
type Block struct {
	Sock [4]Sock
}

func (s Sock) String() string {
	return sockNames[s]
}

func (s Sock) Bits() byte {
	return bits[s]
}

func (s Sock) Mask() byte {
	return masks[s]
}

func (s Sock) Match(x Sock) bool {
	return (s.Bits() & x.Mask() != 0)
}

func NewBlock(input ...byte) *Block {
	b := &Block{}
	for i := 0; i < len(b.Sock); i++ {
		b.Sock[i] = Sock(input[i])
	}
	return b
}

func (b *Block) CanJoin(x *Block, dir int) bool {
	if x == nil || b == nil {
		return true
	}
	return b.Sock[dir].Match(x.Sock[S-dir])
}
