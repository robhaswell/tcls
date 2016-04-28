package main

type ConnectionProducer struct {
	previous []Connection
}

func (producer *ConnectionProducer) Report(connections []Connection) error {
	// Calculate new connections and report them.
	return nil
}
