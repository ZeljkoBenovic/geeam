package engine

import (
	"fmt"

	"github.com/ZeljkoBenovic/geeam/cmd"
	"github.com/ZeljkoBenovic/geeam/core/app"
	ui2 "github.com/ZeljkoBenovic/geeam/domains/ui"
	"github.com/ZeljkoBenovic/geeam/domains/veeam"
)

var uiOption = map[string]ui2.AvailableUIs{
	"console": ui2.Console,
	"excel":   ui2.ExcelTable,
}

func BootstrapCoreApp() (app.App, error) {
	// parse flags
	conf, err := cmd.ProvideCMD().Parse()
	if err != nil {
		return nil, fmt.Errorf("could not parse the provided flags: %w", err)
	}

	// create UI instance
	uiInstance, err := ui2.ProvideUI(uiOption[conf.GetValues().OutputType])
	if err != nil {
		return nil, fmt.Errorf("could not create a UI instance: %w", err)
	}

	// create domains
	veeamDomain := veeam.ProvideVeeam(
		veeam.Config{
			Host:     conf.GetValues().Veeam.Host,
			Port:     conf.GetValues().Veeam.Port,
			Username: conf.GetValues().Veeam.Username,
			Password: conf.GetValues().Veeam.Password,
		})

	// create core app
	return app.ProvideApp(veeamDomain, conf, uiInstance), nil
}
