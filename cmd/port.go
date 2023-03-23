package cmd

type Cmd interface {
	Parse() (Cmd, error)
	GetValues() Flags
}

type Flags struct {
	Veeam
	OutputType       string
	VCenterServerURL string
}

type Veeam struct {
	Host     string
	Port     string
	Username string
	Password string
}
