package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
)

func TerraformInitSubnetActivity(ctx context.Context) (string, error) {
	output, err := utils.RunTFInitCommand(utils.SUBNET_TF_DIRECTORY)
	if err != nil {
		return utils.TF_INIT_FAILED, err
	}
	return output, nil
}

func TerraformApplySubnetActivity(ctx context.Context, vpcID string) (string, error) {

	output, err := utils.RunTFSubnetApplyCommand(utils.SUBNET_TF_DIRECTORY, vpcID)
	//output, err := runCommand("terraform", "apply", "-input=false", "-auto-approve", "-var", fmt.Sprintf("vpc_id=%s", vpcID), "-chdir=terraform/subnet")
	if err != nil {
		return "", err
	}
	return output, nil
	if err != nil {
		return utils.TF_APPLY_FAILED, err
	}
	return output, nil
}

func TerraformOutputSubnetActivity(ctx context.Context) (map[string]string, error) {
	outputValues, err := utils.RunTFOutputCommand(utils.SUBNET_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	var tfOutput map[string]models.TerraformSubnetOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	subnetOutput := map[string]string{
		"private_subnet_id": tfOutput[utils.PRIVATE_SUBNET_ID].Value,
		"public_subnet_id":  tfOutput[utils.PUBLIC_SUBNET_ID].Value,
	}

	return subnetOutput, nil
}
