package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
	app "temporal-helloworld"
	wf "temporal-helloworld/workflows"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GREETINGTASKQUEUE,
	}
	name := "Joel & humana"
	we, err := c.ExecuteWorkflow(context.Background(), options, wf.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
