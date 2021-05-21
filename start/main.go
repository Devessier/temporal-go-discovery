package main

import (
	"context"
	"log"
	"temporal-go-discovery/app"

	"go.temporal.io/sdk/client"
)

// Execute `app.SayHelloWorkflow` workflow with the ID `greeting-workflow`.
// Wait for the workflow to complete and get its result into greeting variable.
func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("error during client creation", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}
	nameToGreet := "Baptiste Devessier"
	workflowExecution, err := c.ExecuteWorkflow(context.Background(), workflowOptions, app.SayHelloWorkflow, nameToGreet)
	if err != nil {
		log.Fatalln("error while executing workflow", err)
	}

	var greeting string

	err = workflowExecution.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("error while getting workflow result", err)
	}

	log.Printf("workflow id: %s with run id: %s gave result: %s\n", workflowExecution.GetID(), workflowExecution.GetRunID(), greeting)
}
