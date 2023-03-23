package ui

type UI interface {
	IngestData(map[string]map[string]bool)
	PresentData()
}

type AvailableUIs int

const (
	Console AvailableUIs = iota
	ExcelTable
)
