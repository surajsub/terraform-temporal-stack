package worker

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"github.com/surajsub/terraform-temporal-stack/logger"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"github.com/surajsub/terraform-temporal-stack/workflows"
	"github.com/surajsub/terraform-temporal-stack/workflows/resources"
	"go.uber.org/zap"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func RunWorker() {

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
		log.Panic("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, utils.WORKFLOW_TASK_QUEUE, worker.Options{})

	// Register workflows
	w.RegisterWorkflow(workflows.ParentWorkflow)
	w.RegisterWorkflow(resources.VPCWorkflow)
	w.RegisterWorkflow(resources.SubnetWorkflow)
	w.RegisterWorkflow(resources.IGWWorkflow)
	w.RegisterWorkflow(resources.NATWorkflow)
	w.RegisterWorkflow(resources.RouteTableWorkflow)
	w.RegisterWorkflow(resources.SGWorkflow)
	w.RegisterWorkflow(resources.EC2Workflow)

	// Register activities
	w.RegisterActivity(activities.TerraformInitVPCActivity)
	w.RegisterActivity(activities.TerraformApplyVPCActivity)
	w.RegisterActivity(activities.TerraformOutputVPCActivity)

	// Register the Subnet work
	w.RegisterActivity(activities.TerraformInitSubnetActivity)
	w.RegisterActivity(activities.TerraformApplySubnetActivity)
	w.RegisterActivity(activities.TerraformOutputSubnetActivity)

	// Register the IGW work
	w.RegisterActivity(activities.TerraformInitIGWActivity)
	w.RegisterActivity(activities.TerraformApplyIGWActivity)
	w.RegisterActivity(activities.TerraformOutputIGWActivity)

	// Register the NAT Work

	w.RegisterActivity(activities.TerraformInitNATActivity)
	w.RegisterActivity(activities.TerraformApplyNATActivity)
	w.RegisterActivity(activities.TerraformOutputNATActivity)

	// Register the RT Work

	w.RegisterActivity(activities.TerraformInitRTActivity)
	w.RegisterActivity(activities.TerraformApplyRTActivity)
	w.RegisterActivity(activities.TerraformOutputRTActivity)

	// Register the SG work
	w.RegisterActivity(activities.TerraformInitSGActivity)
	w.RegisterActivity(activities.TerraformApplySGActivity)
	w.RegisterActivity(activities.TerraformOutputSGActivity)

	// Register the EC2 work
	w.RegisterActivity(activities.TerraformInitEC2Activity)
	w.RegisterActivity(activities.TerraformApplyEC2Activity)
	w.RegisterActivity(activities.TerraformOutputEC2Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("Unable to start worker", err)
	}
}
