// Build the star map and calculate distances
package tcls

import (
	"encoding/gob"
	"fmt"
	"os"

	astar "github.com/beefsack/go-astar"
)

func LoadStarMap() error {
	fp1, err := os.Open("data/system2id.gob")
	if err != nil {
		return err
	}
	dec1 := gob.NewDecoder(fp1)
	if err := dec1.Decode(&System2id); err != nil {
		return err
	}

	fp2, err := os.Open("data/id2system.gob")
	if err != nil {
		return err
	}
	dec2 := gob.NewDecoder(fp2)
	if err := dec2.Decode(&Id2system); err != nil {
		return err
	}

	fp3, err := os.Open("data/jumps.gob")
	if err != nil {
		return err
	}
	dec3 := gob.NewDecoder(fp3)
	if err := dec3.Decode(&Jumps); err != nil {
		return err
	}

	return nil
}

func (s SolarSystemId) PathNeighbors() []astar.Pather {
	var paths []astar.Pather
	for _, adjacentSolarSystemId := range(Jumps[s]) {
		paths = append(paths, astar.Pather(adjacentSolarSystemId))
	}
	return paths
}

func (s SolarSystemId) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

func (s SolarSystemId) PathEstimatedCost(to astar.Pather) float64 {
	return 1.0
}

func (s *SolarSystemId) JumpsTo(to SolarSystemId) (int, error) {
	_, distance, found := astar.Path(s, to)
	if !found {
		return 0, fmt.Errorf("No path found from %s to %s", s, to)
	}
	return int(distance), nil
}

func (s *System) JumpsTo(to System) (int, error) {
	id := s.ID()
	return id.JumpsTo(to.ID())
}
