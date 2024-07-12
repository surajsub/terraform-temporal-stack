package models

type TerraformSGOutput struct {
	Value string `json:"value"`
}

type SGApplyOutput struct {
	SubnetId  string `json:"sg_id"`
	SubnetArn string `json:"sg_arn"`
}
