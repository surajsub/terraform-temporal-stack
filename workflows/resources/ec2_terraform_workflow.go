package resources

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"log"
	"time"
)

func EC2Workflow(ctx workflow.Context, subnetID string, sgid string) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	templog := workflow.GetLogger(ctx)
	templog.Info(utils.EC2_WorkflowName, "Subnet id ", subnetID, "Security Group id ", sgid)

	err := workflow.ExecuteActivity(ctx, activities.TerraformInitEC2Activity).Get(ctx, nil)
	if err != nil {
		log.Println("Unable to execute the terraform initc2activity", zap.Error(err))
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplyEC2Activity, subnetID, sgid).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var ec2Output map[string]string
	err = workflow.ExecuteActivity(ctx, activities.TerraformOutputEC2Activity).Get(ctx, &ec2Output)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.EC2_WorkflowName, "Post Provisioning getting the data  ", ec2Output["instance_id"], "Instance ip of the ec2 instance ", ec2Output["instance_public_ip"])

	return ec2Output, nil
}
