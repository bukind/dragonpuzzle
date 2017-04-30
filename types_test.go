package dragonpuzzle

import (
	"testing"
)

func TestSock(t *testing.T) {
	names := [...]string{
		"ANYS", "NONE", "REDH", "REDT", "YELH", "YELT", "GRNH", "GRNT",
	}
	for i := 0; i < 8; i++ {
		s := Sock(i)
		if s.String() != names[i] {
			t.Errorf("invalid string for %d %v", i, s)
		}
		expect := 1 << uint(i)
		if x := int(s.Bits()); x != expect {
			t.Errorf("invalid s.Bits() = %d, must be %d", x, expect)
		}
	}
}

func TestMasks(t *testing.T) {
	match := map[Sock]Sock{
		NONE: NONE, REDH: REDT, YELH: YELT, GRNH: GRNT,
	}
	for i := 0; i < MAXX; i++ {
		for j := 0; j < MAXX; j++ {
			si := Sock(i)
			sj := Sock(j)
			b := si.Bits()
			m := sj.Mask()
			r := (b & m) != 0
			if i == ANYS || j == ANYS {
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
