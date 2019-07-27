package nanos

import (
	"gnanos/nanos"
	"log"
	"os"
	"strconv"
	"time"
)

func loggingNanos(
	workersMaxCount int,
	taskQueueCapacity,
	panicsMax int,
	filename string,
	) chan nanos.Message {

	worker := &nanos.SelfHealingWorker{
		Worker:    loggingWorker(filename, workersMaxCount),
		PanicsMax: panicsMax,
	}

	myNanos := nanos.Nanos{
		Worker:          worker,
		WorkersMaxCount: workersMaxCount,
		TaskQueueCapacity:taskQueueCapacity,
	}

	return myNanos.TasksChannel()

}



func loggingWorker(filename string, workerMaxCount int) nanos.Worker {

	var keyCabinets = make(chan int, workerMaxCount)
	for i := 1; i < workerMaxCount+1; i++ {
		keyCabinets <- i
	}

	worker := logging{filename: filename, keyCabinets: keyCabinets}
	return worker
}

type logging struct {
	filename    string
	keyCabinets chan int
}

func (w logging) Work(msg nanos.Message) {

	var key = <-w.keyCabinets

	filename := w.filename + "_" + strconv.Itoa(key)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var logger = new(log.Logger)
	logger.SetOutput(file)

	logger.Printf("%v -- %v", time.Now().String(), string(msg.Content))

	w.keyCabinets <- key

}
