terraform {

  required_providers {

    aws = {

      source = "hashicorp/aws"

      version = "~> 4.16"

    }

  }

  required_version = ">= 1.2.0"

}

provider "aws" {

  region = "us-west-2"

}


resource "aws_vpc" "temporal" {

  cidr_block = var.cidr_block

  tags = {

    Name = "Temporal-Automated-VPC"

  }

}



output "vpc_id" {
  value = aws_vpc.temporal.id
}

output "vpc_cidr_block" {
  value = aws_vpc.temporal.cidr_block
}