package resources

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func NATWorkflow(ctx workflow.Context, public_subnet_id string) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	templog := workflow.GetLogger(ctx)
	templog.Info(utils.NAT_WorkflowName, "Public SubnetID ", public_subnet_id)

	err := workflow.ExecuteActivity(ctx, activities.TerraformInitNATActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplyNATActivity, public_subnet_id).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var NATOutput map[string]string
	err = workflow.ExecuteActivity(ctx, activities.TerraformOutputNATActivity).Get(ctx, &NATOutput)
	if err != nil {
		return nil, err
	}

	return NATOutput, nil
}
