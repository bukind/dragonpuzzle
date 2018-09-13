package dragonpuzzle

import (
	"testing"
)


func TestTile(t *testing.T) {
	tile, err := NewTile([]byte{0, 1, 2, 0, 3, 1}, Green(), Red(H), Yellow(T))
	t.Log(tile, err)
}
