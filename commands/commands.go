package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "accelerator",
	Short: "Gladius Website Accelerator",
	Long:  "Gladius Website Accelerator. Upgrade your old website with scalable resources, certificate management, and service worker enhancements.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Gladius Accelerator")
		fmt.Println("Let's work on upgrading your existing website")
		fmt.Println("")
		fmt.Println("Run a few simple commands to get started:")
		fmt.Println("")
		fmt.Println("accelerator config -d domain.com -e certbot@email.com -o 255.255.255.255")
		fmt.Println("")
		fmt.Println("Set up your website and domain, run only when initially setting up or changing servers")
		fmt.Println("")
		fmt.Println("-d | Domain name to request certs via certbot, limit of 5 requests per week per domain")
		fmt.Println("   | Be sure this domain is registered in Digital Ocean and is pointing to this droplet")
		fmt.Println("")
		fmt.Println("-e | Email to register for acme via CertBot")
		fmt.Println("")
		fmt.Println("-o | Origin server ip address you wish to redirect content from. Must be serving content from port 80")
		fmt.Println("")
		fmt.Println("accelerator start")
		fmt.Println("Start our service to serve content")
		fmt.Println("")
		fmt.Println("accelerator stop")
		fmt.Println("Stop our service")
		fmt.Println("")
		fmt.Println("accelerator update")
		fmt.Println("Restart, update, and renew certificates that are coming up for renewal")
	},
}

// Execute main entry point to run the CLI
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
