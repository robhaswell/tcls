package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type signatures struct {
	Signature string `json:"signatureId"`
	Destination struct {
		Name string `json:"name"`
	} `json:"destinationSolarSystem"`
}

func fetch() ([]byte, error) {
	// Fetch the JSON connections from Eve-Scout
	var body []byte
	url := "https://www.eve-scout.com/api/wormholes?systemSearch=Jita&method=shortest&limit=1000&offset=0&order=asc"
	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

func unmarshal(input []byte) ([]Connection, error) {
	// Decode the JSON bytes into a slice of Connection objects
	var connections []Connection
	var result []signatures
	if err := json.Unmarshal(input, &result); err != nil {
		return nil, err
	}
	for _, signature := range(result) {
		connections = append(connections, Connection{
			Dest: System{Name: signature.Destination.Name},
			Sig: Sig{Sig: signature.Signature},
		})
	}
	log.Print(connections)
	return connections, nil
}

func tick() error {
	// This function is called once per INTERVAL
	bytes, err := fetch()
	if err != nil {
		return err
	}
	_, err = unmarshal(bytes)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	for true {
		if err := tick(); err != nil {
			log.Print(err)
		}
		time.Sleep(time.Minute)
	}
}
