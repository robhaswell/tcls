// Build useable Go datastructures from the EVE SDE and write them to data.go
package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
	tcls "../.."
)

// Internal types
type SolarSystem struct {
	SolarSystemId tcls.SolarSystemId `yaml:"solarSystemID"`
	Stargates     map[Stargate]struct {
		Destination Stargate
	}
}

type Stargate int

// Map system IDs with which stargates they contain
var stargates = make(map[tcls.SolarSystemId][]Stargate)

// Map stargates to their resident system IDs
var stargate2systemId = make(map[Stargate]tcls.SolarSystemId)

func walkFunc(path string, info os.FileInfo, err error) error {
	dir, filename := filepath.Split(path)
	if filename != "solarsystem.staticdata" {
		return nil
	}
	var solarsystem SolarSystem

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

	system := tcls.System(filepath.Base(dir))
	solarSystemId := solarsystem.SolarSystemId

	fmt.Println(path, system, solarSystemId)

	tcls.System2id[system] = solarSystemId
	tcls.Id2system[solarSystemId] = system

	for stargate, destination := range solarsystem.Stargates {
		stargates[solarSystemId] = append(stargates[solarSystemId], destination.Destination)
		stargate2systemId[stargate] = solarSystemId
	}
	return nil
}

func main() {
	if err := filepath.Walk("data/sde/fsd/universe/eve/", walkFunc); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Normalise our system->stargate->stargate->system data into just
	// system->system
	for solarSystemId, residentStargates := range stargates {
		for _, stargate := range residentStargates {
			destSolarSystemId := stargate2systemId[stargate]
			tcls.Jumps[solarSystemId] = append(tcls.Jumps[solarSystemId], destSolarSystemId)
		}
	}

	// Write some gobs of data
	fp1, err := os.OpenFile("data/tcls.System2id.gob", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer fp1.Close()
	enc1 := gob.NewEncoder(fp1)
	err = enc1.Encode(tcls.System2id)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fp2, err := os.OpenFile("data/tcls.Id2system.gob", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer fp2.Close()
	enc2 := gob.NewEncoder(fp2)
	err = enc2.Encode(tcls.Id2system)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fp3, err := os.OpenFile("data/jumps.gob", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer fp3.Close()
	enc3 := gob.NewEncoder(fp3)
	err = enc3.Encode(tcls.Jumps)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
