package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/models"
	"github.com/surajsub/terraform-temporal-stack/utils"
	"go.temporal.io/sdk/activity"
)

func TerraformInitEC2Activity(ctx context.Context) (string, error) {
	templog := activity.GetLogger(ctx)
	templog.Info("Starting the ec2 init activity")

	output, err := utils.RunTFInitCommand(utils.EC2_TF_DIRECTORY)
	if err != nil {
		return utils.TF_INIT_FAILED, err
	}
	return output, nil
}

func TerraformApplyEC2Activity(ctx context.Context, subnetid string, sgId string) (string, error) {
	templog := activity.GetLogger(ctx)
	templog.Info(utils.EC2_APPLY, "SubnetNetId", subnetid, "Security Group", sgId)

	output, err := utils.RunTFEC2ApplyCommand(utils.EC2_TF_DIRECTORY, subnetid, sgId)
	if err != nil {
		return "", err
	}
	return output, nil
	if err != nil {
		return utils.TF_APPLY_FAILED, err
	}
	return output, nil
}

func TerraformOutputEC2Activity(ctx context.Context) (map[string]string, error) {
	outputValues, err := utils.RunTFOutputCommand(utils.EC2_TF_DIRECTORY)
	if err != nil {
		return nil, err
	}

	templog := activity.GetLogger(ctx)

	var tfOutput map[string]models.TerraformEC2Output
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	ec2Output := map[string]string{
		"instance_id":        tfOutput[utils.EC2ID].Value,
		"instance_public_ip": tfOutput[utils.EC2PUBLIC].Value,
	}
	templog.Info(utils.EC2_APPLY, "Instance ID", ec2Output["instance_id"], "Instance Public IP", ec2Output["instance_public_ip"])

	return ec2Output, nil
}
