package cmd

import (
	"errors"
	"flag"
	"strings"
)

var (
	ErrNoVeeamHostFound = errors.New("veeam host not provided")
	ErrNoVeeamUserFound = errors.New("veeam username not provided")
	ErrNoVeeamPassFound = errors.New("veeam password not provided")
)

type cmd struct {
	veeamRaw
	outputType string

	processedFlags Flags
}

type veeamRaw struct {
	host string
	port string
	user string
	pass string
}

func ProvideCMD() Cmd {
	c := &cmd{}

	flag.StringVar(&c.veeamRaw.host, "veeam-host", "", "the Veeam server host")
	flag.StringVar(&c.veeamRaw.port, "veeam-port", "9419", "the Veeam server port")
	flag.StringVar(&c.veeamRaw.user, "veeam-username", "", "veeam server username")
	flag.StringVar(&c.veeamRaw.pass, "veeam-pass", "", "veeam server password")

	flag.StringVar(&c.outputType, "output", "console", "data output type (console, excel)")

	flag.Parse()

	return c
}

func (c *cmd) GetValues() Flags {
	return c.processedFlags
}

func (c *cmd) Parse() (Cmd, error) {

	if c.veeamRaw.host == "" {
		return nil, ErrNoVeeamHostFound
	}

	if c.veeamRaw.user == "" {
		return nil, ErrNoVeeamUserFound
	}

	if c.veeamRaw.pass == "" {
		return nil, ErrNoVeeamPassFound
	}

	c.processedFlags.Veeam.Host = strings.TrimSpace(c.veeamRaw.host)
	c.processedFlags.Veeam.Port = strings.TrimSpace(c.veeamRaw.port)
	c.processedFlags.Veeam.Username = strings.TrimSpace(c.veeamRaw.user)
	c.processedFlags.Veeam.Password = strings.TrimSpace(c.veeamRaw.pass)
	c.processedFlags.OutputType = strings.TrimSpace(c.outputType)

	return c, nil
}
