package dragonpuzzle

// XY is the position coordinates.
type XY struct {
	X byte
	Y byte
}

var (
	// Original tile positions.
	TilePos = []XY{
		XY{0, 1},
		XY{1, 2},
		XY{3, 2},
		XY{4, 1},
		XY{3, 0},
		XY{1, 0},
	}
)
