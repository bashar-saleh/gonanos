package nanos

type Nanos struct {
	Worker            Worker
	WorkersMaxCount   int
	TaskQueueCapacity int
}

func (n *Nanos) TasksChannel() chan Message {


	// Creating the KeyCabinets and fill it with keys
	keyCabinets := make(chan int, n.WorkersMaxCount)
	for i := 1; i < n.WorkersMaxCount+1; i++ {
		keyCabinets <- i
	}


	// Creating the Task Channel --> the main door to nanos
	// This is caller will be blocked when the channel is filled --> caller should deal with the block
	var taskChannel = make(chan Message, n.TaskQueueCapacity)

	// Coordinator Start Looping
	go n.coordinate(keyCabinets, taskChannel, n.Worker)

	return taskChannel

}


// Coordinator Loop
func (n *Nanos) coordinate(keyCabinets chan int, taskChannel chan Message, worker Worker) {
	for {

		// Get task from the task channel
		task := <-taskChannel

		// Get Key from key cabinets
		key := <- keyCabinets

		// Assign task to Worker
		go n.assignTaskToWorker(keyCabinets, key, task)

	}
}

// Get the work done and return the key
func (n *Nanos) assignTaskToWorker(keyCabinets chan int, key int, task Message) {
	n.Worker.Work(task)
	keyCabinets <- key
}
