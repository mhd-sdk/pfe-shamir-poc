package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/pfe-manager/config"
	"github.com/pfe-manager/pkg/api"
	"github.com/pfe-manager/pkg/servicesManager"
)

type ShamirPart struct {
	Key byte
	Val []byte
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	// clear console
	clearConsole()
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	go servicesManager.Init()
	api.Init()
}
