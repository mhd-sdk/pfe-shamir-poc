package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/pfe-service/config"
	"github.com/pfe-service/pkg/api"
)

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
	if len(os.Args) > 2 {
		config.Override(config.Config{
			Name: os.Args[1],
			Host: os.Args[2],
		})
	}
	if err != nil {
		panic(err)
	}
	go api.InitStatusWS()
	api.Init()
}
