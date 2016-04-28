package main

import (
	"testing"
	//"github.com/adjust/rmq"
)

func TestReport(t *testing.T) {
	// A queue can be reported
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

	/*
	connection := rmq.OpenConnection("tcls", "tcp", "localhost:6379", 0)
	taskQueue := connection.OpenQueue("connections")
	*/
}
