package dragonpuzzle

import (
	"testing"
)

func TestDir(t *testing.T) {
	chk := func(x int, d Dir, s string, d1 Dir, d3 Dir) {
		if Dir(x) != d {
			t.Errorf("invalid Dir(%d), must be %v", x, d)
		}
		if d.String() != s {
			t.Errorf("invalid d.String() == %q, must be %q", d.String(), s)
		}
		if dx := d.Turn(1); dx != d1 {
			t.Errorf("invalid d(%v).Turn(1) = %v, must be %v", d, dx, d1)
		}
		if dx := d.Turn(3); dx != d3 {
			t.Errorf("invalid d(%v).Turn(3) = %v, must be %v", d, dx, d3)
		}
	}

	chk(0, N, "N", E, W)
	chk(1, E, "E", S, N)
	chk(2, S, "S", W, E)
	chk(3, W, "W", N, S)
}

func TestSock(t *testing.T) {
	names := [...]string{
		"ML", "NC", "RH", "RT", "YH", "YT", "GH", "GT",
	}
	for i := 0; i < 8; i++ {
		s := Side(i)
		if s.String() != names[i] {
			t.Errorf("invalid string for %d %v", i, s)
		}
		expect := 1 << uint(i)
		if x := int(s.bits()); x != expect {
			t.Errorf("invalid s.bits() = %d, must be %d", x, expect)
		}
	}
}

func TestMasks(t *testing.T) {
	match := map[Side]Side{
		NC: NC, RH: RT, YH: YT, GH: GT,
	}
	for i := 0; i < MAXX; i++ {
		for j := 0; j < MAXX; j++ {
			si := Side(i)
			sj := Side(j)
			b := si.bits()
			m := sj.mask()
			r := (b & m) != 0
			if r != si.Match(sj) {
				t.Errorf("inconsistent match %v %v", si, sj)
			}
			if si.Match(sj) != sj.Match(si) {
				t.Errorf("non-communicative match %v %v", si, sj)
			}
			if i == ML || j == ML {
				if r {
					t.Errorf("wrong match %v %v", si, sj)
				}
				continue
			}
			min := si
			max := sj
			if sj < si {
				min = sj
				max = si
			}
			if r {
				if found, ok := match[min]; !ok || found != max {
					t.Errorf("invalid match %v %v => true, false expected", si, sj)
				}
			} else {
				if match[min] == max {
					t.Errorf("cannot match %v %v", si, sj)
				}
			}
		}
	}
}

func TestBlock(t *testing.T) {
	b := NewBlock(ML, NC, RH, GH)
	if bs := b.String(); bs != "ML,NC,RH,GH" {
		t.Errorf("block String() method is broken: %s", bs)
	}
	if !b.EqualTo(b) {
		t.Errorf("block is not equal to itself %v", b)
	}
	if bx := b.Turn(0); !b.EqualTo(bx) {
		t.Errorf("blocks are not equal %v %v", b, bx)
	}
	if bx := b.Turn(1); b.EqualTo(bx) {
		t.Errorf("blocks are wrongly equal %v %v", b, bx)
	}
	if bx := b.Turn(1); !bx.EqualTo(NewBlock(GH, ML, NC, RH)) {
		t.Errorf("invalid turn 1: %v", bx)
	}
	if bx := b.Turn(2); !bx.EqualTo(NewBlock(RH, GH, ML, NC)) {
		t.Errorf("invalid turn 2: %v", bx)
	}
	if bx := b.Turn(3); !bx.EqualTo(NewBlock(NC, RH, GH, ML)) {
		t.Errorf("invalid turn 3: %v", bx)
	}
	b2 := NewBlock(RT, GT, ML, NC)
	if b.Match(b2, N) {
		t.Errorf("wrong match %v %v", b, b2)
	}
	if !b.Match(b2, E) {
		t.Errorf("cannot match %v %v", b, b2)
	}
	if !b.Match(b2, S) {
		t.Errorf("cannot match %v %v", b, b2)
	}
	if !b.Match(b2, W) {
		t.Errorf("cannot match %v %v", b, b2)
	}
	b2 = NewBlock(NC, RT, GT, YT)
	if b.Match(b2, N) {
		t.Errorf("wrong match %v %v", b, b2)
	}
	if b.Match(b2, E) {
		t.Errorf("wrong match %v %v", b, b2)
	}
	if b.Match(b2, S) {
		t.Errorf("cannot match %v %v", b, b2)
	}
	if b.Match(b2, W) {
		t.Errorf("cannot match %v %v", b, b2)
	}
}

func TestTile(t *testing.T) {
	x := NewTile(NC, RH, RT, YH, YT, GH, N)
	if !x.A.EqualTo(NewBlock(NC, RH, ML, GH)) {
		t.Errorf("broken tile ctor #1")
	}
	if !x.B.EqualTo(NewBlock(ML, RT, YH, YT)) {
		t.Errorf("broken tile ctor #2")
	}
	if x.Dir != N {
		t.Errorf("broken tile ctor #3")
	}
	x = x.Turn(2)
	if !x.A.EqualTo(NewBlock(ML, GH, NC, RH)) {
		t.Errorf("broken tile turn #1")
	}
	if !x.B.EqualTo(NewBlock(YH, YT, ML, RT)) {
		t.Errorf("broken tile turn #2")
	}
	if x.Dir != S {
		t.Errorf("broken tile turn #3")
	}
}
