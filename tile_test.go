package dragonpuzzle

import (
	"testing"
)


func TestTrack(t *testing.T) {
	tests := []struct{
		desc string
		call *Track
		want *Track
	}{
		{
			desc: "nil track",
			call: &Track{},
			want: &Track{Col: BLANK},
		},
		{
			desc: "red no hts",
			call: Red(),
			want: &Track{Col: RED, Count: HT{}, Ends: 0},
		},
		{
			desc: "green no hts",
			call: Green(),
			want: &Track{Col: GREEN, Count: HT{}, Ends: 0},
		},
		{
			desc: "yellow no hts",
			call: Yellow(),
			want: &Track{Col:YELLOW, Count: HT{}, Ends: 0},
		},
		{
			desc: "red with 2 heads",
			call: Red(H, H),
			want: &Track{Col: RED, Count: HT{2, 0}, Ends: 0},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if !test.call.Equal(test.want) {
				t.Errorf("tracks not equal: got %s, want %s", test.call, test.want)
			}
		})
	}
}

func TestTile(t *testing.T) {
	tile, err := NewTile([]byte{0, 1, 2, 0, 3, 1}, Green(), Red(H), Yellow(T))
	t.Log(tile, err)

	tile.Turn(1)
	t.Log(tile)
}
