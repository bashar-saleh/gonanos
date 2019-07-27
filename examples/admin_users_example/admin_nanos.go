package admin_users_example

import (
	"gonanos/nanos"
	"log"
	"time"
)

func newAdminNanos(
	workersMaxCount int,
	taskQueueCapacity int,
	usersMailBox chan nanos.Message,
) chan nanos.Message {

	worker := adminWorker{
		usersMailBox: usersMailBox,
	}

	myNanos := nanos.Nanos{
		WorkersMaxCount:   workersMaxCount,
		TaskQueueCapacity: taskQueueCapacity,
		Worker:            worker,
	}

	return myNanos.TasksChannel()
}

type adminWorker struct {
	usersMailBox chan nanos.Message
}

func (w adminWorker) Work(msg nanos.Message) {

	user_id := "user_1"

	var msgContent = []byte(user_id)
	var resTo = make(chan nanos.Message)
	var errTo = make(chan error)

	// Sending message to Admin Nanos
	log.Println("[AdminNanos] Sending task to UsersNanos id="+user_id)
	w.usersMailBox <- nanos.Message{
		Content: msgContent,
		ResTo:   resTo,
		ErrTo:   errTo,
	}

	select {
	case _ = <-resTo:
		// unmarshal the response
		log.Println("[AdminNanos]: User is Bashar")
		return
	case err := <-errTo:
		log.Println("[AdminNanos]: error happened " + err.Error())
		return
	case <-time.After(time.Second * 1):
		log.Println("Timeout")
		return
	}

}
