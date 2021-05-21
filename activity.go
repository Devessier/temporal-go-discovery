package app

import "time"

// SayHello is an activity that simply concatenates "Hi " to the name to greet.
func SayHello(name string) (string, error) {
	greeting := "Hi " + name

	// Simulate a time consuming task
	time.Sleep(2 * time.Second)

	return greeting, nil
}
