
provider "aws" {
  region = "us-west-2"
}


resource "aws_key_pair" "temporal" {
  key_name   = "temporal-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvpFyZo8aFbXeUBr7osSCJNgvavWbM/06niWrOvYX2xwWdhXmXSrbX8ZbabVohBK41 email@example.com"
}

resource "aws_instance" "temporal" {


  ami                    = var.amiId
  instance_type          = var.aws_instance_type
  subnet_id              = var.subnet_id
  key_name               = aws_key_pair.temporal.key_name
  vpc_security_group_ids = [var.sg_id]


  tags = {

    Name = "suraj-temporal"

  }

}


output "instance_id" {
  value = aws_instance.temporal.id
}

output "instance_public_ip" {
  value = aws_instance.temporal.public_ip
}