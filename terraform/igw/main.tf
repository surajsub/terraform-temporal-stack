/* Resource file for internet gateway creation */
resource "aws_internet_gateway" "temporal" {

 vpc_id = var.vpc_id
 tags = {
   Name = "temporal VPC IG"
 }

}

output "igw_id" {
  value = aws_internet_gateway.temporal.id
}

