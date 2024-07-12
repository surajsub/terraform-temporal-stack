package models

type TerraformVPCOutput struct {
	Value string `json:"value"`
}

type VPCApplyOutput struct {
	VPCID   string `json:"vpc_id"`
	VPCCIDR string `json:"vpc_cidr_block"`
}

type TerraformCommonOutput struct {
	Value string `json:"value"`
}
