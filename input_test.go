package main

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/adjust/rmq"
)

type TaskConsumer struct{}

var ch = make(chan Connection)

func (consumer *TaskConsumer) Consume(delivery rmq.Delivery) {
	var connection Connection
	payload := []byte(delivery.Payload())

	if err := json.Unmarshal(payload, &connection); err != nil {
		log.Print(err)
		delivery.Reject()
		return
	}
	ch <- connection
	delivery.Ack()
}

// A queue can be reported
// TODO: Refactor this using the actual consumer logic
func TestReportIntegration(t *testing.T) {
	connections := []Connection{
		Connection{Dest: System("Aaa"), Sig: Sig("AAA")},
		Connection{Dest: System("Bbb"), Sig: Sig("BBB")},
		Connection{Dest: System("Ccc"), Sig: Sig("CCC")},
	}
	producer := new(ConnectionProducer)
	err := producer.Report(connections)
	if err != nil {
		t.Fatal(err)
	}

	qc := rmq.OpenConnection("tcls", "tcp", "localhost:6379", 0)
	taskQueue := qc.OpenQueue("connections")
	taskQueue.StartConsuming(10, time.Second)

	taskConsumer := &TaskConsumer{}
	taskQueue.AddConsumer("task consumer", taskConsumer)

	var result []Connection
	for {
		connection := <-ch
		result = append(result, connection)
		if reflect.DeepEqual(result, connections) {
			return
		}
	}
}
