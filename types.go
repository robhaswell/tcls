package tcls

type Connection struct {
	Dest System
	Sig  Sig
}

type System string
type SolarSystemId int
type Sig string

func (s System) ID() SolarSystemId {
	return System2id[s]
}

// Map system names to their IDs
var System2id = make(map[System]SolarSystemId)

// Map system IDs to their names
var Id2system = make(map[SolarSystemId]System)

// Map system IDs with which system IDs they connect to via stargates
var Jumps = make(map[SolarSystemId][]SolarSystemId)

