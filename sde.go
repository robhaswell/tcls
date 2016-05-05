// Build useable Go datastructures from the EVE SDE and write them to data.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Solarsystem struct {
	SolarSystemId SolarSystemId `yaml:"solarSystemID"`
	Stargates map[Stargate]struct {
		Destination Stargate
	}
}

type Stargate int

// Map system names to their IDs
var system2id = make(map[System]SolarSystemId)

// Map system IDs to their names
var id2system = make(map[SolarSystemId]System)

// Map system IDs with which stargates they contain
var stargates = make(map[SolarSystemId][]Stargate)

// Map stargates to their resident system IDs
var stargate2systemId = make(map[Stargate]SolarSystemId)

// Map system IDs with which system IDs they connect to via stargates
var jumps = make(map[SolarSystemId][]SolarSystemId)

func walkFunc(path string, info os.FileInfo, err error) error {
	dir, filename := filepath.Split(path)
	if filename != "solarsystem.staticdata" {
		return nil
	}
	var solarsystem Solarsystem

	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &solarsystem); err != nil {
		return err
	}

	system := System(filepath.Base(dir))
	solarSystemId := solarsystem.SolarSystemId

	fmt.Println(path, system, solarSystemId)

	system2id[system] = solarSystemId
	id2system[solarSystemId] = system

	for stargate, destination := range(solarsystem.Stargates) {
		stargates[solarSystemId] = append(stargates[solarSystemId], destination.Destination)
		stargate2systemId[stargate] = solarSystemId
	}
	return nil
}

func main() {
	if err := filepath.Walk("sde/fsd/universe/eve/", walkFunc); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Normalise our system->stargate->stargate->system data into just
	// system->system
	for solarSystemId, residentStargates := range(stargates) {
		for _, stargate := range(residentStargates) {
			destSolarSystemId := stargate2systemId[stargate]
			jumps[solarSystemId] = append(jumps[solarSystemId], destSolarSystemId)
		}
	}

	fp, err := os.OpenFile("data.go", os.O_WRONLY | os.O_CREATE, 0600)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer fp.Close()
	_, err = fmt.Fprintf(fp, `// Generated data from 'go run sde.go types.go'
package main

var System2Id = %#v

var Id2System = %#v

var Jumps = %#v
`, system2id, id2system, jumps)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
