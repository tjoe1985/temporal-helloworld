package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	app "temporal-helloworld"
	"temporal-helloworld/activities"
	"temporal-helloworld/workflows"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Println("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Worker and Activity functions
	w := worker.New(c, app.GREETINGTASKQUEUE, worker.Options{})
	// register the workflows and activities
	w.RegisterWorkflow(workflows.GreetingWorkflow)
	w.RegisterActivity(activities.ComposeGreeting)
	// start listening to the task queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Println(" unable to start worker: ", err)
	}
}
