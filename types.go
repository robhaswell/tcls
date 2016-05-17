package tcls

type Connection struct {
	Dest System
	Sig  Sig
}

type System string
type SolarSystemId int
type Sig string

/*
func (s *System) ID() SolarSystemId {
	return System2Id[s]
}
*/
