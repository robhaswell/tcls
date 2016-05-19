package tcls

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	LoadStarMap()
	flag.Parse()
	os.Exit(m.Run())
}

// The distance between adjacent systems is 1
func TestAdjacent(t *testing.T) {
	s1 := System("Asoutar")
	s2 := System("Rayl")
	jumps, err := s1.JumpsTo(s2)

	if err != nil {
		t.Error(err)
	}

	if jumps != 1 {
		t.Error("Unexpected number of jumps:", jumps)
	}
}

// The distance between Jita and Amarr is 9 jumps
func TestJitaAmarr(t *testing.T) {
	s1 := System("Jita")
	s2 := System("Amarr")
	jumps, err := s1.JumpsTo(s2)

	if err != nil {
		t.Error(err)
	}

	if jumps != 9 {
		t.Error("Unexpected number of jumps:", jumps)
	}
}

// The distance between regional gates is 1 jump
func TestRegional(t *testing.T) {
	s1 := System("HED-GP")
	s2 := System("Keberz")
	jumps, err := s1.JumpsTo(s2)

	if err != nil {
		t.Error(err)
	}

	if jumps != 1 {
		t.Error("Unexpected number of jumps:", jumps)
	}
}
