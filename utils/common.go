package utils

import (
	"errors"
	"fmt"
	"github.com/surajsub/terraform-temporal-stack/logger"
	"go.uber.org/zap"
	"net"
	"os/exec"
)

// This will contain a list of common functions that can be used through the code

func runTFCommand(name, directory string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = VPC_TF_DIRECTORY
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", name, args, err, string(output))
	}
	return string(output), nil
}

func RunTFInitCommand(directory string) (string, error) {

	cmdArgs := []string{"init", "-input=false"}
	cmd := exec.Command("terraform", cmdArgs...)

	//cmd := exec.Command("terraform", "init", "-input=false")
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "init -input=false", err, string(output))
	}
	return string(output), nil
}

func RunTFApplyCommand(directory string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve")
	//terraform", "apply", "-input=false", "-auto-approve"
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "apply -input=false -auto-approve", err, string(output))
	}
	return string(output), nil
}

func RunTFOutputCommand(directory string) (string, error) {
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "output -json", err, string(output))
	}
	return string(output), nil
}
func RunTFVPCApplyCommand(directory, vpc_cidr_block string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve", "-var", fmt.Sprintf("cidr_block=%s", vpc_cidr_block))
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

func RunTFSubnetApplyCommand(directory, vpcid string) (string, error) {

	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve", //nolint:gosec
		"-var", fmt.Sprintf("vpc_id=%s", vpcid))

	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

func RunTFIGWApplyCommand(directory, vpcid string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve", "-var", fmt.Sprintf("vpc_id=%s", vpcid))
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

func RunTFNATApplyCommand(directory, public_subnet_id string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve", "-var", fmt.Sprintf("subnet_id=%s", public_subnet_id))
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

// Start the route table apply

func RunTFRTApplyCommand(directory, vpc_id, igw_id, nat_id, private_subnet_id, public_subnet_id string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve",
		"-var", fmt.Sprintf("vpc_id=%s", vpc_id),
		"-var", fmt.Sprintf("igw_id=%s", igw_id),
		"-var", fmt.Sprintf("nat_id=%s", nat_id),
		"-var", fmt.Sprintf("private_subnet_id=%s", private_subnet_id),
		"-var", fmt.Sprintf("public_subnet_id=%s", public_subnet_id))

	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

// start the SG Apply
func RunTFSGApplyCommand(directory, vpcid, vpcdir string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve", "-var", fmt.Sprintf("vpc_id=%s", vpcid), "-var", fmt.Sprintf("vpc_cidr_block=%s", vpcdir))
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(SG_WorkflowName+"error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

func RunTFEC2ApplyCommand(directory, subnetid, sgid string) (string, error) {
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve", "-var", fmt.Sprintf("subnet_id=%s", subnetid), "-var", fmt.Sprintf("sg_id=%s", sgid))
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(EC2_WorkflowName+"error running command %s %v: %v - output: %s", "terraform", "apply -json", err, string(output))
	}
	return string(output), nil
}

func GetTemporalZap() *logger.ZapAdapter {
	tflogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer tflogger.Sync()

	temporalLogger := logger.NewZapAdapter(tflogger)
	return temporalLogger
}

func CarveSubnets(cidr string, count int) ([]string, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, errors.New("invalid CIDR block")
	}

	ones, bits := ipNet.Mask.Size()
	subnetSize := bits - (ones + 1)
	requiredSize := ones + count

	if requiredSize > bits {
		return nil, errors.New("count is too large for the given CIDR block")
	}

	var subnets []string
	currentIP := ipNet.IP

	for i := 0; i < count; i++ {
		subnet := &net.IPNet{
			IP:   currentIP,
			Mask: net.CIDRMask(ones+1, bits),
		}
		subnets = append(subnets, subnet.String())
		currentIP = incrementIP(currentIP, 1<<(subnetSize))
	}

	return subnets, nil
}

func incrementIP(ip net.IP, inc int) net.IP {
	ip = ip.To4()
	if ip == nil {
		return nil
	}
	newIP := make(net.IP, len(ip))
	copy(newIP, ip)
	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i] += byte(inc)
		if newIP[i] != 0 {
			break
		}
		inc >>= 8
	}
	return newIP
}
