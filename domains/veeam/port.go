package veeam

type Veeam interface {
	Init() (Veeam, error)
	Logout()
	FetchVmAndObjectData() (Data, error)
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Data struct {
	BackupObjects map[string][]string
	HostObjects   map[string]map[string]bool
}
