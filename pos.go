package dragonpuzzle

import "fmt"

// XY is the position coordinates.
type XY struct {
	X int8
	Y int8
}

func (x *XY) Turn(steps int) {
	switch steps {
	case 0:
		return
	case 1:
		*x = XY{-x.Y, x.X}
	case 2:
		*x = XY{-x.X, -x.Y}
	case 3:
		*x = XY{x.Y, -x.X}
	default:
		panic(fmt.Sprintf("wrong rotation %d", steps))
	}
}

var (
	// Y
	//
	// ^
	// |
	// 2     1   2
	// 1   0 a . b 3
	// 0     5   4
	// |
	// +-- 0 1 2 3 4 --> X

	// Initial tile ends positions.
	InitialTileEnds = []XY{
		XY{0, 1},
		XY{1, 2},
		XY{3, 2},
		XY{4, 1},
		XY{3, 0},
		XY{1, 0},
	}
	// Initial tile centers.
	InitialTileCenters = []XY{
		XY{1, 1},
		XY{3, 1},
	}
)
