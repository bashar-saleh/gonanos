package nanos

import (
	"fmt"
)


type Worker interface {
	Work(task Message)
}


// SelfHealing Decorator for recovering from panics situation
type SelfHealingWorker struct {
	Worker    Worker
	PanicsMax int
	Name      string
	Logger    chan Message
}

func (w *SelfHealingWorker) Work(task Message) {

	defer func() {
		if err := recover(); err != nil {
			if w.Logger != nil {
				w.Logger <- Message{Content: []byte(fmt.Sprintf("[%v]: Panic %v", w.Name, err))}
			}
			if w.PanicsMax == 0 {
				if w.Logger != nil {
					w.Logger <- Message{Content: []byte(fmt.Sprintf("[%v]: Exceeded tha allowd count of panics, System will stop", w.Name))}
				}
				panic(fmt.Sprintf("[%v]: System Stop because of panic %v", w.Name, err))
			} else {
				w.PanicsMax--
				if w.Logger != nil {
					w.Logger <- Message{Content: []byte(fmt.Sprintf("[%v]: Healing from panic number %v", w.Name,w.PanicsMax ))}
				}
				go w.Worker.Work(task)
			}
		}
	}()

	w.Worker.Work(task)
}
