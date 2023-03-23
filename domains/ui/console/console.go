package console

import (
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type ConsoleUI interface {
	IngestData(map[string]map[string]bool)
	PresentData()
}

type console struct {
	table table.Table
}

func NewUI() (ConsoleUI, error) {
	return &console{}, nil
}

func (c *console) IngestData(data map[string]map[string]bool) {

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	c.table = table.New("VM Name", "Backup", "Host")
	c.table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for hostName, hostObjects := range data {
		for objectName, hasBackup := range hostObjects {
			if !hasBackup && !strings.Contains(objectName, "vCLS-") {
				c.table.AddRow(objectName, "NO", hostName)
			}
		}
	}
}

func (c *console) PresentData() {
	c.table.Print()
}
