package excel

import (
	"fmt"
)

type ExcelUI interface {
	IngestData(map[string]map[string]bool)
	PresentData()
}

type excel struct {
}

func NewUI() (ExcelUI, error) {
	return &excel{}, nil
}

// TODO: implement excel UI
func (e *excel) IngestData(map[string]map[string]bool) {
	fmt.Println("Creating excel table")
}

func (e *excel) PresentData() {

}
