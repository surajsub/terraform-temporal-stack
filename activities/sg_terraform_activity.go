package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/activity"
)

func TerraformInitSGActivity(ctx context.Context) (string, error) {
	output, err := utils.RunTFInitCommand(utils.SG_TF_DIRECTORY)
	if err != nil {
		return utils.TF_INIT_FAILED, err
	}
	return output, nil
}

func TerraformApplySGActivity(ctx context.Context, vpcID string, vpc_cdir_block string) (string, error) {
	templog := utils.GetTemporalZap()
	templog.Info(utils.SG_APPLY, "VPC ID", vpcID, "VPC Cdir Block", vpc_cdir_block)
	fmt.Printf("the vpc cdir is set to %s\n and the vpcid is %s", vpc_cdir_block, vpcID)

	output, err := utils.RunTFSGApplyCommand(utils.SG_TF_DIRECTORY, vpcID, vpc_cdir_block)
	if err != nil {
		templog.Error(utils.EC2_APPLY, "Failed to perform the Apply for Security Group")
		return "", err
	}
	return output, nil

}

func TerraformOutputSGActivity(ctx context.Context) (map[string]string, error) {
	templog := activity.GetLogger(ctx)
	outputValues, err := utils.RunTFOutputCommand(utils.SG_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	var tfOutput map[string]models.TerraformSGOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	sgOutput := map[string]string{
		"sg_id":  tfOutput[utils.SGID].Value,
		"sg_arn": tfOutput[utils.SGARN].Value,
	}

	templog.Info(utils.SG_APPLY, "Security Group ID", sgOutput[utils.SGID], "Security Group ARN ", sgOutput[utils.SGARN])
	//fmt.Println("SG ACTIVITY :  the value is %s", sgOutput["sg_id"])

	return sgOutput, nil
}
