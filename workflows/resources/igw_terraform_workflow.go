package resources

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/workflow"
	"time"
)

func IGWWorkflow(ctx workflow.Context, vpcID string) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	templog := workflow.GetLogger(ctx)
	templog.Info(utils.IGW_WorkflowName, "VPCID ", vpcID)

	err := workflow.ExecuteActivity(ctx, activities.TerraformInitIGWActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplyIGWActivity, vpcID).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var iGWOutput map[string]string
	err = workflow.ExecuteActivity(ctx, activities.TerraformOutputIGWActivity).Get(ctx, &iGWOutput)
	if err != nil {
		return nil, err
	}

	return iGWOutput, nil
}
