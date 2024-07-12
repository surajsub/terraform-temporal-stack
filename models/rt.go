package models

type TerraformRTOutput struct {
	Value string `json:"value"`
}

type RTApplyOutput struct {
	NatID           string `json:"nat_id"`
	NatGateway      string `json:"nat_gateway_id"`
	NatAllocationID string `json:"nat_allocation_id"`
}
