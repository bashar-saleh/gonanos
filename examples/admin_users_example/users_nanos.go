package admin_users_example

import (
	"github.com/bashar-saleh/gonanos/nanos"
	"log"
	"time"
)

func newUsersNanos(
	workersMaxCount int,
	taskQueueCapacity int,
) chan nanos.Message {

	worker := usersWorker{}
	myNanos := nanos.Nanos{
		Worker:            worker,
		TaskQueueCapacity: taskQueueCapacity,
		WorkersMaxCount:   workersMaxCount,
	}
	return myNanos.TasksChannel()
}

type usersWorker struct{}

func (w usersWorker) Work(msg nanos.Message) {

	// unmarshal msg
	user_id := "user_1"

	// Do the work
	log.Println("[UsersNanos] Searching for user with id " + user_id)
	time.After(time.Millisecond * 10)

	select {
	case msg.ResTo <- nanos.Message{Content: []byte("User name is Bashar")}:
		return
	default:
		return
	}
}
