# dragonpuzzle
A puzzle from the game SNAP

There are tiles with dragon parts of three colors.  Parts are
exposed to the sides of tiles with their endings.

A tile has 6 sides.

Endings have the following attributes:
- Color
- Connection number (to match other endings in the same tile).
- Position (the side) on the tile.

The tile also has the following mapping:
- Each connection maps to the tuple of how many heads and tails are in
the tile.

Tile{
  Ends: Blank, Red{1}, Red{2}, Blank, Blank, Blank,
  Map: map[int]Count{1: Head(1), 2: Tail(1)},
}
