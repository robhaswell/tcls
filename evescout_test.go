package tcls

import (
	"testing"
)

func TestGetsConnections(t *testing.T) {
	connections, err := readConnections()
	if err != nil {
		t.Fatal(err)
	}
	if len(connections) == 0 {
		t.Fatal("No connections found")
	}
}
