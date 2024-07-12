package resources

import (
	"github.com/surajsub/terraform-temporal-stack/activities"
	"go.temporal.io/sdk/workflow"
	"time"
)

const WorkflowName = "AWS_SECURITY_GROUP"

func SGWorkflow(ctx workflow.Context, vpcID string, vpcdir string) (map[string]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, activities.TerraformInitSGActivity).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = workflow.ExecuteActivity(ctx, activities.TerraformApplySGActivity, vpcID, vpcdir).Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	var SGOutput map[string]string
	err = workflow.ExecuteActivity(ctx, activities.TerraformOutputSGActivity).Get(ctx, &SGOutput)
	if err != nil {
		return nil, err
	}

	return SGOutput, nil
}
