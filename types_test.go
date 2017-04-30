package dragonpuzzle

import (
	"testing"
)

func TestSock(t *testing.T) {
	names := [...]string{
		"AC", "NC", "RH", "RT", "YH", "YT", "GH", "GT",
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
			if i == AC || j == AC {
				if !r {
					t.Errorf("cannot match %v %v", si, sj)
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
	b := NewBlock(AC,NC,RH,GH)
	if bs := b.String(); bs != "AC,NC,RH,GH" {
		t.Errorf("block String() method is broken: %s", bs);
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
	if bx := b.Turn(1); !bx.EqualTo(NewBlock(GH,AC,NC,RH)) {
		t.Errorf("invalid turn 1: %v", bx)
	}
	if bx := b.Turn(2); !bx.EqualTo(NewBlock(RH,GH,AC,NC)) {
		t.Errorf("invalid turn 2: %v", bx)
	}
	if bx := b.Turn(3); !bx.EqualTo(NewBlock(NC,RH,GH,AC)) {
		t.Errorf("invalid turn 3: %v", bx)
	}
	b2 := NewBlock(RT,GT,AC,NC)
	if !b.Match(b2, N) {
		t.Errorf("cannot match %v %v", b, b2)
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
	b2 = NewBlock(NC,RT,GT,YT)
	if !b.Match(b2, N) {
		t.Errorf("cannot match %v %v", b, b2)
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
