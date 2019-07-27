package basic

import (
	"gnanos/nanos"
	"log"
	"time"
)

func NewBasicNanos_2(
	workersMaxCount int,
	taskQueueCapacity int,
	basicNanos chan nanos.Message,

) chan nanos.Message {

	worker := SomeWork_2{
		basicNanos: basicNanos,
	}
	myNanos := nanos.Nanos{
		Worker:            worker,
		TaskQueueCapacity: taskQueueCapacity,
		WorkersMaxCount:   workersMaxCount,
	}

	return myNanos.TasksChannel()
}

type SomeWork_2 struct {
	basicNanos chan nanos.Message
}

func (w SomeWork_2) Work(msg nanos.Message) {

	// preparing the message
	var resTo = make(chan nanos.Message)
	var errorTo = make(chan error)

	log.Println("[SomeWork_2] sending work to [SomeWork] mailbox")
	w.basicNanos <- nanos.Message{
		Content: []byte("Do Work"),
		ResTo:   resTo,
		ErrTo:   errorTo,
	}

	// Waiting for response
	select {
	case _ = <-resTo:
		log.Printf("[SomeWork_2]: work done")
		return
	case err := <-errorTo:
		log.Printf("[SomeWork_2]; error happend: %v", err.Error())
		return
	case <-time.After(time.Second * 4): // timout
		log.Println("[SomeWork_2] timeout")
		return
	}

}
