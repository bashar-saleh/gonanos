package basic

import (
	"gnanos/nanos"
	"log"
	"time"
)


func NewBasicNanos(
	workersMaxCount int,
	taskQueueCapacity int,

	)  chan nanos.Message{

	worker := SomeWork{}
	myNanos := nanos.Nanos{
		Worker:worker,
		TaskQueueCapacity:taskQueueCapacity,
		WorkersMaxCount:workersMaxCount,
	}

	return  myNanos.TasksChannel()
}




type SomeWork struct {}

func (w SomeWork) Work(msg nanos.Message)  {

	// Do The Work
	log.Println("[SomeWork] working on task")
	time.Sleep(time.Second * 1)




	select {
	case msg.ResTo <- nanos.Message{Content: []byte("SomeResult")}:
	default:
		return
	}

}
