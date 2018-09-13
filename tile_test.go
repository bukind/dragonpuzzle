package dragonpuzzle

import (
	"testing"
)


/*
func TestTrack(t *testing.T) {
	tests := []struct{
		desc string
		call *Track,
		want *Track,
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if call.Equal(want)
		})
	}
}
*/

func TestTile(t *testing.T) {
	tile, err := NewTile([]byte{0, 1, 2, 0, 3, 1}, Green(), Red(H), Yellow(T))
	t.Log(tile, err)

	tile.Turn(1)
	t.Log(tile)
}
