package models

type TerraformIGWOutput struct {
	Value string `json:"value"`
}

type IGWApplyOutput struct {
	IGWId  string `json:"igw_id"`
	IGWArn string `json:"igw_arn"`
}
