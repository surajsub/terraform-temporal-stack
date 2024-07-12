variable "aws_instance_type" {
  type        = string
  description = "AWS EC2 Instance type"
  default     = "t2.micro"
}



variable "amiId" {
  description = "ami Id from which instance needs to be created from"
  default     = "ami-01572eda7c4411960"
}

variable "subnet_id" {
  description = "The subnet value to create this instance in"

}

variable "sg_id" {
  description = "The security group to associate this ec2 instance"

}