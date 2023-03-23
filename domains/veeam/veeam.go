package veeam

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/veeamhub/veeam-vbr-sdk-go/client"
)

type veeam struct {
	host, username, password string
	ctx                      context.Context
	client                   *client.APIClient
	data                     Data
}

const (
	apiVersion       = "1.0-rev1"       // default API version (1.0-rev1)
	apiVersion11Rev0 = "1.1-rev0"       // newer API version (1.1-rev0)
	skipTls          = true             // skip TLS certificate verification
	timeout          = 30 * time.Second // 30 seconds
)

func ProvideVeeam(conf Config) Veeam {
	v := &veeam{
		host:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		username: conf.Username,
		password: conf.Password,
		data: Data{
			BackupObjects: make(map[string][]string),
			HostObjects:   make(map[string]map[string]bool),
		},
	}

	return v
}

func (v *veeam) Init() (Veeam, error) {
	// create tls client
	tlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTls,
			},
		},
		Timeout: timeout,
	}

	// setup config
	conf := client.NewConfiguration()
	conf.HTTPClient = tlsClient
	conf.Host = v.host
	conf.Scheme = "https"

	// create new veeam api client
	v.client = client.NewAPIClient(conf)

	// new login
	login, httpResp, err := v.client.LoginApi.CreateToken(context.Background()).
		XApiVersion(apiVersion).
		GrantType("password").
		Username(v.username).
		Password(v.password).
		Execute()
	if err != nil {
		fmt.Printf("HTTP Response: %v\n", httpResp)

		return nil, fmt.Errorf("could not authenticate to VBR API: %w", err)
	}

	// set auth context
	v.ctx = context.WithValue(context.Background(), client.ContextAccessToken, login.AccessToken)

	return v, nil
}

func (v *veeam) Logout() {
	_, _, _ = v.client.LoginApi.Logout(v.ctx).XApiVersion(apiVersion).Execute()
}

func (v *veeam) FetchVmAndObjectData() (Data, error) {
	if err := v.getJobObjects(); err != nil {
		return Data{}, err
	}

	if err := v.getHostObjects(); err != nil {
		return Data{}, err
	}

	return v.data, nil
}

func (v *veeam) getJobObjects() error {
	// try default api version first
	jobs, _, err := v.client.AutomationApi.ExportJobs(v.ctx).XApiVersion(apiVersion).Execute()
	if err != nil {
		// try newer api version
		jobs, _, err = v.client.AutomationApi.ExportJobs(v.ctx).XApiVersion(apiVersion11Rev0).Execute()
		if err != nil {
			return fmt.Errorf("could not retreive job objects: %w", err)
		}

	}

	for _, j := range jobs.GetJobs() {
		for _, vm := range j.GetVirtualMachines().Includes {
			v.data.BackupObjects[j.GetName()] = append(v.data.BackupObjects[j.GetName()], vm.Name)
		}
	}

	return nil
}

func (v *veeam) getHostObjects() error {
	// fetch all connected hosts
	inventoryHosts, r, err := v.client.InventoryBrowserApi.GetAllInventoryVmwareHosts(v.ctx).XApiVersion(apiVersion).Execute()
	if err != nil {
		fmt.Printf("HTTP Response: %#v", r)

		return fmt.Errorf("could not fetch VMWare hosts inventory: %w", err)
	}

	// create a map with a host name as key
	for _, h := range inventoryHosts.GetData() {
		v.data.HostObjects = map[string]map[string]bool{
			h.InventoryObject.GetName(): {},
		}
	}

	// get all object in a host and append to hosts array
	for hostName := range v.data.HostObjects {
		hostObjects, hresp, hostErr := v.client.InventoryBrowserApi.GetVmwareHostObject(v.ctx, hostName).XApiVersion(apiVersion).Execute()
		if hostErr != nil {
			fmt.Printf("HTTP Response: %#v", hresp)

			return fmt.Errorf("could not fetch host inventory objects: %w", err)
		}

		for _, h := range hostObjects.GetData() {
			if h.GetInventoryObject().Type == "VirtualMachine" {
				v.data.HostObjects[hostName][h.GetInventoryObject().Name] = false
			}
		}
	}

	return nil
}
