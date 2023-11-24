package statusTableUI

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/pfe-manager/pkg/models"
	"github.com/rodaine/table"
)

var tbl table.Table

func UpdateTable(services []models.Service) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	tbl = table.New("Service name", "Host", "Status")
	for _, service := range services {
		tbl.AddRow(service.Name, service.Host, service.Status)
	}
	tbl.Print()
}
