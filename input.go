package main

import (
	"encoding/json"
	"github.com/adjust/rmq"
)

type ConnectionProducer struct {
	previous []Connection
}

func (producer *ConnectionProducer) Report(connections []Connection) error {
	/*
		Report all new connections to the queue.

		Any errors which occur will result in the error being returned. It is the
		responsibility of the caller to log or otherwise handle this error.
	*/
	connection := rmq.OpenConnection("tcls", "tcp", "localhost:6379", 0)
	taskQueue := connection.OpenQueue("connections")

	for connection := range connections {
		taskBytes, err := json.Marshal(connection)
		if err != nil {
			return err
		}
		taskQueue.PublishBytes(taskBytes)
	}
	return nil
}
