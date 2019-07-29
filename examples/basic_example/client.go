package basic_example

import (
	"fmt"
	"github.com/bashar-saleh/gonanos/nanos"
)

// just change the function name to Main and run it from your main package main function
func main() {

	basicNanosMailBox := newBasicNanos(8, 2000)

	basicNanos_2MailBox := newBasicNanos_2(8, 1500, basicNanosMailBox)

	basicNanos_2MailBox <- nanos.Message{Content: []byte(" ")}

	fmt.Scanln()
}
