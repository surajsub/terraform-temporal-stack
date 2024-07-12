package workflows

import (
	"github.com/surajsub/terraform-temporal-stack/utils"
	"github.com/surajsub/terraform-temporal-stack/workflows/resources"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func ParentWorkflow(ctx workflow.Context, vpc_cidr string) (map[string]interface{}, error) {
	cwo := workflow.ChildWorkflowOptions{
		WorkflowExecutionTimeout: time.Hour,
		WorkflowRunTimeout:       time.Minute * 30,
		ParentClosePolicy:        enums.PARENT_CLOSE_POLICY_ABANDON,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	// Start VPC Workflow - Resource 1

	var vpcOutput map[string]string
	err := workflow.ExecuteChildWorkflow(ctx, resources.VPCWorkflow, vpc_cidr).Get(ctx, &vpcOutput)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("VPC created", "vpc_id", vpcOutput["vpc_id"], "and the vpc cidr is ", vpcOutput["vpc_cidr_block"])

	// Start Subnet Workflow - Resource 2
	var subnetOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.SubnetWorkflow, vpcOutput["vpc_id"]).Get(ctx, &subnetOutput)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("Subnet created", "private_subnet_id", subnetOutput["private_subnet_id"])
	workflow.GetLogger(ctx).Info("Public Subnet Created", "public_subnet_id", subnetOutput["public_subnet_id"])

	// Start IGW Workflow - Resource 3

	var igwOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.IGWWorkflow, vpcOutput["vpc_id"]).Get(ctx, &igwOutput)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("IGW Created", "igw_id", igwOutput["igw_id"])

	// Start the NAT Workflow - Resource 4

	var natoutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.NATWorkflow, subnetOutput["public_subnet_id"]).Get(ctx, &natoutput)
	if err != nil {
		workflow.GetLogger(ctx).Info("Failed to create the NAT ", "public subnet id", subnetOutput["public_subnet_id"])
	}
	workflow.GetLogger(ctx).Info("NAT Created", "NAT ID", natoutput["nat_id"])
	workflow.GetLogger(ctx).Info("NAT Gateway ID", "NAT Gateway ID", natoutput["nat_gateway_id"])
	workflow.GetLogger(ctx).Info("NAT Allocation ID", "NAT Allocation ID", natoutput["nat_allocation_id"])

	// Start the Route Table and Association workflow - Resource 5

	var rtOutput map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.RouteTableWorkflow, vpcOutput["vpc_id"], igwOutput[utils.IGWID], natoutput[utils.NATGATEWAYID], subnetOutput[utils.PRIVATE_SUBNET_ID], subnetOutput[utils.PUBLIC_SUBNET_ID]).Get(ctx, &rtOutput)
	if err != nil {
		workflow.GetLogger(ctx).Info(utils.RtError, "igw_id", igwOutput[utils.IGWID])
	}

	// Start the Security Group Workflow - Resource 6

	var sgOutPut map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.SGWorkflow, vpcOutput["vpc_id"], vpcOutput["vpc_cidr_block"]).Get(ctx, &sgOutPut)
	if err != nil {
		log.Fatalln("Failed to execute the ", utils.SG_WorkflowName)
		return nil, err
	}
	workflow.GetLogger(ctx).Info("Security Group created", "sg_id", sgOutPut["sg_id"])

	// Start EC2 Workflow - Resource 7
	var ec2Output map[string]string
	err = workflow.ExecuteChildWorkflow(ctx, resources.EC2Workflow, subnetOutput["public_subnet_id"], sgOutPut["sg_id"]).Get(ctx, &ec2Output)
	if err != nil {
		return nil, err
	}
	workflow.GetLogger(ctx).Info("EC2 instance created", "instance_id", ec2Output["instance_id"], "instance_public_ip", ec2Output["instance_public_ip"])

	// Aggregate results
	results := map[string]interface{}{
		"EC2Workflow": ec2Output,
	}

	return results, nil

}
