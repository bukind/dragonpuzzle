package dragonpuzzle

import (
	"fmt"
)


type BlankT int

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
	Heads byte
	Tails byte
	Ends  byte
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

var H = HT{1, 0}
var T = HT{0, 1}

func NewTrack(c Color, hts ...HT) *Track {
	t := &Track{}
	t.Col = c
	for _, x := range hts {
		t.Heads += x.Heads
		t.Tails += x.Tails
	}
	return t
}

// Tile is the tile description.
type Tile struct {
	Ends []End
	Tracks []*Track
}

func NewTile(ends []byte, ts ...*Track) (*Tile, error) {
	tile := &Tile{
		Ends: make([]End, 6),
		Tracks: append([]*Track{&Track{Col:BLANK}}, ts...),
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
		tile.Ends[i].Pos = TilePos[i]
	}
	return tile, nil
}
