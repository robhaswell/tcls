// Build the star map and calculate distances
package tcls
/*

import (
	astar "github.com/beefsack/go-astar"
)

func (s *SolarSystemId) PathNeighbours() []astar.Pather {
	var paths []astar.Pather
	for _, adjacentSolarSystemId := range(Jumps[s]) {
		paths = append(paths, astar.Pather(adjacentSolarSystemId))
	}
	return paths
}

func (s *SolarSystemId) PathNeighbourCost(to astar.Pather) float64 {
	return 1.0
}

func (s *SolarSystemId) PathEstimatedCost(to astar.Pather) float64 {
	return 1.0
}

func (s *SolarSystemId) JumpsTo(to SolarSystemId) (int, error) {
	path, distance, found := astar.Path(s, to)
	if !found {
		return nil, error("No path found from", s, "to", to)
	}
	return int(distance), 0
}

func (s *System) JumpsTo(to System) (int, error) {
	return s.ID().JumpsTo(to.ID())
}
*/
