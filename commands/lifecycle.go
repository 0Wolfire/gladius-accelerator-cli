package commands

import (
	"github.com/spf13/cobra"
	"gladius-accelerator-cli/env"
	"os/exec"
)

var commandConfig = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run:   serviceConfig,
}

var commandStart = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  "",
	Run:   serviceStart,
}

var commandStop = &cobra.Command{
	Use:   "stop",
	Short: "",
	Long:  "",
	Run:   serviceStop,
}

var commandUpdate = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  "",
	Run:   serviceUpdate,
}

func serviceConfig(cmd *cobra.Command, args []string) {
	domain, _ := cmd.Flags().GetString("domain")
	if domain != "" {
		env.SetDomain(domain)
	}

	email, _ := cmd.Flags().GetString("email")
	if email != "" {
		env.SetEmail(email)
	}

	origin, _ := cmd.Flags().GetString("origin")
	if origin != "" {
		env.SetOriginHost(origin)
	}

	if env.VerifyEnvironment() {
		//exec.Command("/bin/sh", "-c", "docker-compose -f build-compose.yaml up").Run()
	}
}

func serviceStart(*cobra.Command, []string) {
	exec.Command("/bin/sh", "-c", "docker-compose up -d").Run()
}

func serviceStop(*cobra.Command, []string) {
	exec.Command("/bin/sh", "-c", "docker-compose -f build-compose.yaml down").Run()
	exec.Command("/bin/sh", "-c", "docker-compose down").Run()
}

func serviceUpdate(*cobra.Command, []string) {
	exec.Command("/bin/sh", "-c", "docker-compose -f build-compose.yaml down").Run()
	exec.Command("/bin/sh", "-c", "docker-compose down").Run()

	exec.Command("/bin/sh", "-c", "docker-compose -f build-compose.yaml pull").Run()
	exec.Command("/bin/sh", "-c", "docker-compose pull").Run()

	exec.Command("/bin/sh", "-c", "docker-compose -f build-compose.yaml up").Run()
	exec.Command("/bin/sh", "-c", "docker-compose up -d").Run()
}

func init() {
	rootCommand.AddCommand(commandConfig)
	rootCommand.AddCommand(commandStart)
	rootCommand.AddCommand(commandStop)
	rootCommand.AddCommand(commandUpdate)

	setupFlags()
}

func setupFlags() {
	commandConfig.Flags().StringP("domain", "d", "", "Request certificates and initialize accelerator for this domain")
	commandConfig.Flags().StringP("email", "e", "", "Email to register automatically with CertBot")
	commandConfig.Flags().StringP("origin", "o", "", "Origin server IP address")
}
