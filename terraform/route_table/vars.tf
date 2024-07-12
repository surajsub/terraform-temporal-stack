variable vpc_id {
  description = "VPC ID"
}


variable "nat_id" {
  description = "NAT Id to be used for the route"
}

variable "igw_id" {
  description = "Internet Gateway"
}


variable "public_subnet_id" {
  description = "THe public subnet id created for this resource"

}

variable "private_subnet_id" {
  description = "The private subnet id created for this resource"
}
