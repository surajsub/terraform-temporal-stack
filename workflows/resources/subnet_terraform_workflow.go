package resources

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"go.temporal.io/sdk/workflow"
	"time"
)

func SubnetWorkflow(ctx workflow.Context, vpcID string) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, activities.TerraformInitSubnetActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplySubnetActivity, vpcID).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var subnetOutput map[string]string
	err = workflow.ExecuteActivity(ctx, activities.TerraformOutputSubnetActivity).Get(ctx, &subnetOutput)
	if err != nil {
		return nil, err
	}

	return subnetOutput, nil
}
