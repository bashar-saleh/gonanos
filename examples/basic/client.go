package basic

import (
	"fmt"
	"gnanos/nanos"
)

// just change the function name to Main and run it from your main package main function
func Main() {

	basicNanosMailBox := NewBasicNanos(8, 2000)

	basicNanos_2MailBox := NewBasicNanos_2(8,1500, basicNanosMailBox)

	basicNanos_2MailBox <- nanos.Message{Content:[]byte(" ")}

	fmt.Scanln()
}
