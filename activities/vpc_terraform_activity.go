package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/activity"
)

func TerraformInitVPCActivity(ctx context.Context) (string, error) {
	activity.GetLogger(ctx).Info("Initiating the Terraform VPC init Activity")
	output, err := utils.RunTFInitCommand(utils.VPC_TF_DIRECTORY)
	if err != nil {
		return utils.TF_INIT_FAILED, err
	}
	return output, nil
}

func TerraformApplyVPCActivity(ctx context.Context, cidrBlock string) (string, error) {
	activity.GetLogger(ctx).Info("Initiating the Terraform VPC Apply Activity")
	output, err := utils.RunTFVPCApplyCommand(utils.VPC_TF_DIRECTORY, cidrBlock)
	if err != nil {
		return utils.TF_APPLY_FAILED, err
	}
	return output, nil
}

func TerraformOutputVPCActivity(ctx context.Context) (map[string]string, error) {
	activity.GetLogger(ctx).Info("Initiating the Terraform VPC Output Activity")
	outputValues, err := utils.RunTFOutputCommand(utils.VPC_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	var tfOutput map[string]models.TerraformCommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	vpcOutput := map[string]string{
		"vpc_id":         tfOutput[utils.VPCID].Value,
		"vpc_cidr_block": tfOutput[utils.VPCCIDR].Value,
	}

	return vpcOutput, nil
}
