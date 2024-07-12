package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"log"
)

func TerraformInitRTActivity(ctx context.Context) (string, error) {
	output, err := utils.RunTFInitCommand(utils.RT_TF_DIRECTORY)
	if err != nil {
		return utils.RT_INIT, err
	}
	return output, nil
}

func TerraformApplyRTActivity(ctx context.Context, vpc_id, igw_id, nat_id, private_subnet_id, public_subnet_id string) (string, error) {
	//output, err := utils.RunTFApplyCommand(utils.SUBNET_TF_DIRECTORY)

	log.Println("the values are vpcid ", vpc_id)
	log.Println("the valuare for igw is ", igw_id)
	log.Println("the value for natid is ", nat_id)
	log.Println("the value for private_subnet", private_subnet_id)
	log.Println("the value of public subnet", public_subnet_id)
	output, err := utils.RunTFRTApplyCommand(utils.RT_TF_DIRECTORY, vpc_id, igw_id, nat_id, private_subnet_id, public_subnet_id)

	if err != nil {
		return "", err
	}
	return output, nil
	if err != nil {
		return utils.TF_APPLY_FAILED, err
	}
	return output, nil
}

func TerraformOutputRTActivity(ctx context.Context) (map[string]string, error) {
	outputValues, err := utils.RunTFOutputCommand(utils.RT_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	var tfOutput map[string]models.TerraformRTOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	rtOutput := map[string]string{
		"rt_public_id":  tfOutput[utils.RTPUBLICID].Value,
		"rt_private_id": tfOutput[utils.RTPRIVATEID].Value,
	}

	return rtOutput, nil
}
