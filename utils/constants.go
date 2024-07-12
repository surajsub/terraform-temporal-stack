package utils

import "fmt"

const WORKFLOW_TASK_QUEUE = "MULTI_STACK_QUEUE"

// Terraform file locations

const VPC_RESOURCE = "VPC"
const SUBNET_RESOURCE = "Subnet"
const IGW_RESOURCE = "Internet Gateway"
const SG_RESOURCE = "Security Group"
const EC2_RESOURCE = "EC2 Instance"

const VPC_TF_DIRECTORY = "./terraform/vpc"
const SUBNET_TF_DIRECTORY = "./terraform/subnet"
const IGW_TF_DIRECTORY = "./terraform/igw"
const EC2_TF_DIRECTORY = "./terraform/ec2"
const SG_TF_DIRECTORY = "./terraform/sg"
const NAT_TF_DIRECTORY = "./terraform/nat"
const RT_TF_DIRECTORY = "./terraform/route_table"

var TF_INIT = fmt.Sprintf("terraform", "init", "-input=false")

const TF_INIT_FAILED = "Failed to execute the terraform init command"
const TF_APPLY_FAILED = "Failed to execute the terraform apply command"
const TF_OUTPUT_FAILED = "Failed to execute the terraform output command"

const VPC_WorkflowName = "AWS VPC"
const IGW_WorkflowName = "AWS_Internet_Gateway"
const SUBNET_WorkflowName = "AWS VPC Subnet"
const NAT_WorkflowName = "AWS Nat Service"
const RT_WorkflowName = "AWS Route Table Service"
const SG_WorkflowName = "AWS Security Group"
const EC2_WorkflowName = "AWS EC2 Instance"

const EKS_WorkflowName = "AWS EKS"
const NODE_WorkflowName = "AWS EKS Nodes"

//Define the constants for the variables

const VPC_INIT = "Starting the VPC Init Activity:"
const SUBNET_INIT = "Subnet Init Activity:"
const IGW_INIT = "Internet Gateway Init Activity:"
const NAT_INIT = "NAT Init Activity"
const RT_INIT = "Route Table Init Activity"
const SG_INIT = "Security Group Init Activity:"
const EC2_INIT = "EC2 Init Activity:"
const EKS_INIT = "EKS Init Activity"
const NODE_INIT = "EKS Node Init Activity"

const VPC_APPLY = "VPC Apply Activity:"
const SG_APPLY = "Security Group Apply Activity:"
const SUBNET_APPLY = "AWS Subnet Apply Activity"
const NAT_APPLY = "NAT Apply Activity"
const RT_APPLY = "Route Table Apply Activity"
const EC2_APPLY = "EC2 Apply Activity:"
const EKS_APPLY = "EKS Apply Activity:"
const NODE_APPLY = "EKS Node Apply Activity:"

const VPCID = "vpc_id"
const VPCCIDR = "vpc_cidr_block"

const SUBNETID = "subnet_id"
const SUBNETARN = "subnet_arn"

const PRIVATE_SUBNET_ID = "private_subnet_id"
const PUBLIC_SUBNET_ID = "public_subnet_id"

const IGWID = "igw_id"
const IGWARN = "igw_arn"

const SGID = "sg_id"
const SGARN = "sg_arn"

const NATID = "nat_id"
const NATGATEWAYID = "nat_gateway_id"
const NATALLOCATIONID = "nat_allocation_id"

const RTPRIVATEID = "rt_private_id"
const RTPUBLICID = "rt_public_id"

// These need to map to the output from the tf file
const EC2ID = "instance_id"
const EC2PUBLIC = "instance_public_ip"
