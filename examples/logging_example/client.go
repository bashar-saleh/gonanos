package nanos

import (
	"fmt"
	"github.com/bashar-saleh/gonanos/nanos"
)

// just change the function name to Main and run it from your main package main function
func main() {

	// 8 workers = 8 files because we can't make more than one goroutine writing to the same file
	loggingNanos := loggingNanos(8, 1000, 8, "logfile")

	for i := 0; i < 100; i++ {
		loggingNanos <- nanos.Message{Content: []byte("Hello World")}
	}

	// Logging Nanos will return nothing so if the main goroutine finished very quickly the Logging Worker may not finish
	fmt.Scanln()
}
