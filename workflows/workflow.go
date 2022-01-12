package workflows

import (
	"go.temporal.io/sdk/workflow"
	act "temporal-helloworld/activities"
	"time"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result string
	err := workflow.ExecuteActivity(ctx, act.ComposeGreeting, name).Get(ctx, &result)
	return result, err
}
