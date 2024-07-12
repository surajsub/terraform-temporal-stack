package models

type VPC struct {
	Name      string `json:"name"`
	CdirBlock string `json:"cdir_block"`
	Tags      []Tags `json:"tags"`
}

type Subnet struct {
	Name             string `json:"subnet"`
	SubnetBlock      string `json:"subnet_block"`
	AvailabilityZone string `json:"availability_zone"`
	Tags             []Tags `json:"tags"`
}

type IGW struct {
	Name string `json:"name"`
	Tags []Tags `json:"tags"`
}

type SecurityGroup struct {
	Name string `json:"name"`
	Rule []Rule `json:"rule"`
	Tags []Tags `json:"tags"`
}

type Rule struct {
	RuleType string `json:"rule_type"`
}

type EC2Instance struct {
	Tags []Tags `json:"tags"`
}

type AWSTemporalResponse struct{}

type Resources struct {
	VPC           VPC           `json:"vpc"`
	Subnet        Subnet        `json:"subnet"`
	IGW           IGW           `json:"igw"`
	SecurityGroup SecurityGroup `json:"security_group"`
	EC2Instance   EC2Instance   `json:"ec2instance"`
}

type AwsTemporalRequest struct {
	Region    string    `json:"region"`
	Prefix    string    `json:"prefix"`
	Resources Resources `json:"resources"`
}

type Tags struct {
	Key   string `json:"name"`
	Value string `json:"value"`
}
