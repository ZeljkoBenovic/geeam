package app

import (
	"github.com/ZeljkoBenovic/geeam/cmd"
	"github.com/ZeljkoBenovic/geeam/domains/ui"
	"github.com/ZeljkoBenovic/geeam/domains/veeam"
)

type app struct {
	veeam  veeam.Veeam
	config cmd.Cmd
	ui     ui.UI
}

func ProvideApp(
	veeam veeam.Veeam,
	conf cmd.Cmd,
	ui ui.UI,
) App {

	a := &app{
		veeam:  veeam,
		config: conf,
		ui:     ui,
	}

	return a
}

func (a *app) Run() error {
	veeamClient, err := a.veeam.Init()
	if err != nil {
		return err
	}

	defer veeamClient.Logout()

	veeamData, err := veeamClient.FetchVmAndObjectData()
	if err != nil {
		return err
	}

	for _, jobItems := range veeamData.BackupObjects {
		for _, item := range jobItems {
			for _, hostVM := range veeamData.HostObjects {
				if _, ok := hostVM[item]; ok {
					hostVM[item] = true
				}
			}
		}
	}

	a.ui.IngestData(veeamData.HostObjects)
	a.ui.PresentData()

	return nil
}
