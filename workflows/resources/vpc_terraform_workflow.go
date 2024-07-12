package resources

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func VPCWorkflow(ctx workflow.Context, vpc string) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	templog := workflow.GetLogger(ctx)
	templog.Info(utils.VPC_WorkflowName, "VPC Value is ", vpc)
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, activities.TerraformInitVPCActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplyVPCActivity, vpc).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var vpcOutput map[string]string
	err = workflow.ExecuteActivity(ctx, activities.TerraformOutputVPCActivity).Get(ctx, &vpcOutput)
	if err != nil {
		return nil, err
	}

	templog.Info(utils.VPC_WorkflowName, "VPC Value is ", vpcOutput["vpc_id"])
	return vpcOutput, nil
}
