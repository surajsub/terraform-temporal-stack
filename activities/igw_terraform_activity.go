package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
)

func TerraformInitIGWActivity(ctx context.Context) (string, error) {
	output, err := utils.RunTFInitCommand(utils.IGW_TF_DIRECTORY)
	if err != nil {
		return utils.TF_INIT_FAILED, err
	}
	return output, nil
}

func TerraformApplyIGWActivity(ctx context.Context, vpcID string) (string, error) {
	//output, err := utils.RunTFApplyCommand(utils.SUBNET_TF_DIRECTORY)
	output, err := utils.RunTFIGWApplyCommand(utils.IGW_TF_DIRECTORY, vpcID)

	if err != nil {
		return "", err
	}
	return output, nil
	if err != nil {
		return utils.TF_APPLY_FAILED, err
	}
	return output, nil
}

func TerraformOutputIGWActivity(ctx context.Context) (map[string]string, error) {
	outputValues, err := utils.RunTFOutputCommand(utils.IGW_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	var tfOutput map[string]models.TerraformIGWOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	igwOutput := map[string]string{
		"igw_id":  tfOutput["igw_id"].Value,
		"igw_arn": tfOutput["igw_arn"].Value,
	}

	return igwOutput, nil
}
