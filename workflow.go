package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// Our workflow that calls SayHello activity.
func SayHelloWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var greeting string

	err := workflow.ExecuteActivity(ctx, SayHello, name).Get(ctx, &greeting)

	return greeting, err
}
