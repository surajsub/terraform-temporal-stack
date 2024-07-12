package workflows

import (
	"context"
	"github.com/surajsub/terraform-temporal-stack/logger"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.uber.org/zap"
	"log"

	"go.temporal.io/sdk/client"
)

func StartWorkflow(vpc string) {

	tflogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer tflogger.Sync() // flushes buffer, if any

	// Create a custom Temporal logg using Zap
	temporalLogger := logger.NewZapAdapter(tflogger)

	c, err := client.Dial(client.Options{
		Logger: temporalLogger,
	})
	if err != nil {
		log.Fatal("Unable to create Temporal client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "parent-workflow",
		TaskQueue: utils.WORKFLOW_TASK_QUEUE,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, ParentWorkflow, vpc)
	if err != nil {
		log.Fatal("Unable to execute workflow", err)
	}

	log.Printf("Started workflow %s %s", we.GetID(), we.GetRunID())

	var result map[string]interface{}
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatal("Unable to get workflow result", err)
	}
	log.Printf("Workflow result: %v", result)
}
