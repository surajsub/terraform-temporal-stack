package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
)

func TerraformInitNATActivity(ctx context.Context) (string, error) {
	output, err := utils.RunTFInitCommand(utils.NAT_TF_DIRECTORY)
	if err != nil {
		return utils.NAT_INIT, err
	}
	return output, nil
}

func TerraformApplyNATActivity(ctx context.Context, public_subnet_id string) (string, error) {
	//output, err := utils.RunTFApplyCommand(utils.SUBNET_TF_DIRECTORY)
	output, err := utils.RunTFNATApplyCommand(utils.NAT_TF_DIRECTORY, public_subnet_id)

	if err != nil {
		return "", err
	}
	return output, nil
	if err != nil {
		return utils.TF_APPLY_FAILED, err
	}
	return output, nil
}

func TerraformOutputNATActivity(ctx context.Context) (map[string]string, error) {
	outputValues, err := utils.RunTFOutputCommand(utils.NAT_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	var tfOutput map[string]models.TerraformNATOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	natOutput := map[string]string{
		"nat_id":            tfOutput[utils.NATID].Value,
		"nat_gateway_id":    tfOutput[utils.NATGATEWAYID].Value,
		"nat_allocation_id": tfOutput[utils.NATALLOCATIONID].Value,
	}

	return natOutput, nil
}
