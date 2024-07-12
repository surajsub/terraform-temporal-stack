

resource "aws_subnet" "private_temporal" {
  vpc_id            = var.vpc_id
  cidr_block        = var.private_subnet_cdir
  availability_zone = var.private_subnet_zone

  tags = {
    "Name"                            = "temporal-private-us-west-2a"
    "kubernetes.io/role/internal-elb" = "1"
    "kubernetes.io/cluster/demo"      = "owned"

  }
}


resource "aws_subnet" "public_temporal" {
  vpc_id            = var.vpc_id
  cidr_block        = var.public_subnet_cdir
  availability_zone = var.public_subnet_zone


  tags = {
    "Name"                            = "temporal-public-us-west-2b"
    "kubernetes.io/role/internal-elb" = "1"
    "kubernetes.io/cluster/demo"      = "owned"
  }
}

output "private_subnet_id" {
  value = aws_subnet.private_temporal.id
}

output "private_subnet_arn" {
  value = aws_subnet.private_temporal.arn
}

output "public_subnet_id" {
  value = aws_subnet.public_temporal.id
}

output "public_subnet_arn" {
  value = aws_subnet.public_temporal.arn
}