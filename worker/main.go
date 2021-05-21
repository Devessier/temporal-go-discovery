package main

import (
	"log"
	"temporal-go-discovery/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// Launch a worker to process workflows from `app.GreetingTaskQueue` queue.
func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("error during client creation", err)
	}
	defer c.Close()

	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})

	w.RegisterWorkflow(app.SayHelloWorkflow)
	w.RegisterActivity(app.SayHello)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("error while trying to launch worker", err)
	}
}
