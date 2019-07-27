package admin_users_example

import (
	"fmt"
	"gonanos/nanos"
)

// just change the function name to Main and run it from your main package main function
func main() {

	usersMailBox := newUsersNanos(8,100)
	adminMailBox := newAdminNanos(8,1000, usersMailBox)

	adminMailBox <- nanos.Message{Content: []byte("Start some work")}

	fmt.Scanln()
}
