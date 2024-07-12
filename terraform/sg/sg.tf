
variable "vpc_id" {}
variable "vpc_cidr_block" {}

resource "aws_security_group" "temporal" {
  name        = "allow_ssh"
  description = "sg for to allow ssh"
  vpc_id      = var.vpc_id

  tags = {
    Name = "Temporal SG"
  }
}


resource "aws_vpc_security_group_ingress_rule" "temporal_ssh" {
  security_group_id = aws_security_group.temporal.id
  cidr_ipv4         = "0.0.0.0/0"
  description       = "inbound rule to port 22"
  from_port         = 22
  ip_protocol       = "tcp"
  to_port           = 22
}

resource "aws_vpc_security_group_ingress_rule" "temporal_web" {
  security_group_id = aws_security_group.temporal.id
  cidr_ipv4         = "0.0.0.0/0"
  description       = "inbound rule to port 80"
  from_port         = 80
  ip_protocol       = "tcp"
  to_port           = 80
}


resource "aws_vpc_security_group_egress_rule" "temporal" {
  security_group_id = aws_security_group.temporal.id
  cidr_ipv4         = "0.0.0.0/0"
  ip_protocol       = "-1" # semantically equivalent to all ports
}

output "sg_id" {
  value = aws_security_group.temporal.id
}

output "sg_arn" {
  value = aws_security_group.temporal.arn
}

