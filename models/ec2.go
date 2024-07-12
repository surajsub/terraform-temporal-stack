package models

type TerraformEC2Output struct {
	Value string `json:"value"`
}

type EC2ApplyOutput struct {
	//EC2 Apply Output Structure.
	InstanceID       string `json:"instance_id"`
	InstancePublicIP string `json:"instance_public_ip"`
}
