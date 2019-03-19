package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gladius-accelerator-cli/env"
	"log"
	"os"
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
		fmt.Println("Domain name set")
	}

	email, _ := cmd.Flags().GetString("email")
	if email != "" {
		env.SetEmail(email)
		fmt.Println("Certbot email set")
	}

	origin, _ := cmd.Flags().GetString("origin")
	if origin != "" {
		env.SetOriginHost(origin)
		fmt.Println("Origin server IP address set")
	}

	if env.VerifyEnvironment() {
		runCommand("Environment teardown", "docker-compose -f build-compose.yaml down --remove-orphans")
		runCommand("Environment build", "docker-compose -f build-compose.yaml up")
	}
}

func serviceStart(*cobra.Command, []string) {
	if env.VerifyEnvironment() {
		runCommand("Masternode teardown", "docker-compose down --remove-orphans")
		runCommand("Masternode start", "docker-compose up -d")
	}
}

func serviceStop(*cobra.Command, []string) {
	runCommand("Environment teardown", "docker-compose -f build-compose.yaml down --remove-orphans")
	runCommand("Masternode teardown", "docker-compose down --remove-orphans")
}

func serviceUpdate(*cobra.Command, []string) {
	runCommand("Environment teardown", "docker-compose -f build-compose.yaml down --remove-orphans")
	runCommand("Masternode teardown", "docker-compose down --remove-orphans")

	if env.VerifyEnvironment() {
		runCommand("Environment update", "docker-compose -f build-compose.yaml pull")
		runCommand("Masternode update", "docker-compose pull")

		runCommand("Environment build", "docker-compose -f build-compose.yaml up")
		runCommand("Masternode start", "docker-compose up -d")
	}
}

func runCommand(message, command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = viper.GetString("DIR_PATH")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("%s failed with %s\n", message, err)
	}
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
