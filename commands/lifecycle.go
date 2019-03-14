package commands

import (
	"github.com/spf13/cobra"
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

func serviceConfig(*cobra.Command, []string) {
	exec.Command("/bin/sh", "-c", "docker-compose -f build-compose.yaml up").Run()
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
}
