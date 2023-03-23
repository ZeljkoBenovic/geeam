package ui

import (
	"errors"

	"github.com/ZeljkoBenovic/geeam/domains/ui/console"
	"github.com/ZeljkoBenovic/geeam/domains/ui/excel"
)

var (
	ErrUINotSupported = errors.New("ui not supported")
)

var uiTypesFactory = map[AvailableUIs]func() (UI, error){
	Console: func() (UI, error) {
		return console.NewUI()
	},
	ExcelTable: func() (UI, error) {
		return excel.NewUI()
	},
}

func ProvideUI(uiType AvailableUIs) (UI, error) {
	if UIInstance, ok := uiTypesFactory[uiType]; ok {
		return UIInstance()
	}

	return nil, ErrUINotSupported
}
