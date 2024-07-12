variable "cidr_block" {
  description="CDIR Block for the Subnet"
  default ="10.0.0.0/19"
}

variable "vpc_id" {
  description = "The VPC to create this subnet in"
}

variable "zone" {
  description = "The zone where we want to create the subnet in"
  default = "us-west-2a"
}


variable "private_subnet_cdir" {
  description = "Private Subnet CDIR Block"
  default = "10.0.0.0/19"
}

variable "private_subnet_zone" {
  description = "Private Subnet Zone"
  default ="us-west-2a"
}

variable "public_subnet_cdir" {
  description = "Public Subnet CDIR Block"
  default = "10.0.32.0/19"
}

variable "public_subnet_zone" {
  description = "Public Subnet Zone"
  default = "us-west-2b"
}