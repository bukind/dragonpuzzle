package dragonpuzzle

import (
	"fmt"
)

const (
	BLANK = Color(iota)
	RED
	GREEN
	YELLOW
)

type Color byte

// End is the side description.
type End struct {
	Track byte
	Pos   XY
}

type HT struct {
	Heads byte
	Tails byte
}

// Track is a part of a dragon of certain color.
type Track struct {
	Col   Color
	Count HT
	Ends  byte
}

type TrackList []*Track

// Tile is the tile description.
type Tile struct {
	Ends   []End
	Cells  XYList
	Tracks TrackList
}

var (
	// Head helper
	H = HT{1, 0}
	// Tail helper
	T = HT{0, 1}

	colorToString = map[Color]string{
		BLANK: "blank",
		RED: "red",
		GREEN: "green",
		YELLOW: "yellow",
	}
)

func (c Color) String() string {
	return colorToString[c]
}

func (e End) String() string {
	return fmt.Sprintf("trk#%d@%v", e.Track, e.Pos)
}

func (h HT) String() string {
	return fmt.Sprintf("%.*s%.*s", h.Heads, "HHHHHHHH", h.Tails, "TTTTTTTT")
}

func (h *HT) Merge(hts ...HT) {
	for _, ht := range hts {
		h.Heads += ht.Heads
		h.Tails += ht.Tails
	}
}

func (t *Track) String() string {
	return fmt.Sprintf("Trk(%s, %s%.*s)", t.Col, t.Count, t.Ends, "EEEEEEEE")
}

func (t *Track) Equal(trk *Track) bool {
	if t.Col != trk.Col {
		return false
	}
	if t.Count != trk.Count {
		return false
	}
	if t.Ends != trk.Ends {
		return false
	}
	return true
}

func Green(hts ...HT) *Track {
	return NewTrack(GREEN, hts...)
}
func Red(hts ...HT) *Track {
	return NewTrack(RED, hts...)
}
func Yellow(hts ...HT) *Track {
	return NewTrack(YELLOW, hts...)
}

func NewTrack(c Color, hts ...HT) *Track {
	t := &Track{}
	t.Col = c
	t.Count.Merge(hts...)
	return t
}

func (t TrackList) Clone() TrackList {
	var res TrackList
	for _, x := range t {
		res = append(res, &Track{x.Col, x.Count, x.Ends})
	}
	return res
}

func (t *Tile) String() string {
	return fmt.Sprintf("Tile(ends:%v cells:%v trks:%v)", t.Ends, t.Cells, t.Tracks)
}

func (t *Tile) Turn(steps int) {
	for i := range t.Ends {
		t.Ends[i].Pos.Turn(steps)
	}
	for i := range t.Cells {
		t.Cells[i].Turn(steps)
	}
}

func NewTile(ends []byte, ts ...*Track) (*Tile, error) {
	tile := &Tile{
		Ends: make([]End, 6),
		Cells: append(XYList{}, InitialTileCenters...),
		Tracks: append(TrackList{&Track{}}, ts...),
	}
	if len(ends) != len(tile.Ends) {
		return nil, fmt.Errorf("wrong number of ends %d", len(ends))
	}
	for i, end := range ends {
		if int(end) >= len(tile.Tracks) {
			return nil, fmt.Errorf("too large track index in %d", end)
		}
		tile.Ends[i].Track = end
		tile.Tracks[end].Ends++
		tile.Ends[i].Pos = InitialTileEnds[i]
	}
	return tile, nil
}
