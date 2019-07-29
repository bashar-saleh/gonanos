package basic_example

import (
	"github.com/bashar-saleh/gonanos/nanos"
	"log"
	"time"
)

func newBasicNanos(
	workersMaxCount int,
	taskQueueCapacity int,

) chan nanos.Message {

	worker := someWork{}
	myNanos := nanos.Nanos{
		Worker:            worker,
		TaskQueueCapacity: taskQueueCapacity,
		WorkersMaxCount:   workersMaxCount,
	}

	return myNanos.TasksChannel()
}

type someWork struct{}

func (w someWork) Work(msg nanos.Message) {

	// Do The Work
	log.Println("[someWork] working on task")
	time.Sleep(time.Second * 1)

	select {
	case msg.ResTo <- nanos.Message{Content: []byte("SomeResult")}:
		return
	default:
		return
	}

}
