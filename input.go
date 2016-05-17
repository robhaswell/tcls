package tcls

import (
	"encoding/json"
	"github.com/adjust/rmq"
)

type ConnectionProducer struct {
	previous []Connection
}

/*
   Report all new connections to the queue.

   Any errors which occur will result in the error being returned. It is the
   responsibility of the caller to log or otherwise handle this error.
*/
func (producer *ConnectionProducer) Report(connections []Connection) error {
	qc := rmq.OpenConnection("tcls", "tcp", "localhost:6379", 0)
	taskQueue := qc.OpenQueue("connections")

	for _, connection := range connections {
		taskBytes, err := json.Marshal(connection)
		if err != nil {
			return err
		}
		taskQueue.PublishBytes(taskBytes)
	}
	return nil
}
