package main

import (
	"flag"
	"github.com/surajsub/terraform-temporal-stack/worker"
	"github.com/surajsub/terraform-temporal-stack/workflows"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected 'worker' or 'starter' subcommands")
	}

	switch os.Args[1] {
	case "worker":
		worker.RunWorker()
	case "starter":

		starterCmd := flag.NewFlagSet("starter", flag.ExitOnError)
		vpcCdirBlock := starterCmd.String("vpcCdirBlock", "", "CIDR Block for the VPC")

		err := starterCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("failed to parse 'starter' flags")
		}

		// Ensure the required flags are provided
		if *vpcCdirBlock == "" {
			log.Fatal("VPC Cdir block is not provided")
		}

		workflows.StartWorkflow(*vpcCdirBlock)

	default:
		log.Fatal("expected 'worker' or 'starter' subcommands")
	}
}
