package nutanix

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/routebyintuition/ntnx-go-sdk/calm"
	"github.com/routebyintuition/ntnx-go-sdk/karbon"
	"github.com/routebyintuition/ntnx-go-sdk/pc"
	"github.com/routebyintuition/ntnx-go-sdk/pe"
)

// Client is an HTTP Client
type Client struct {
	config   *Config
	clientMu sync.Mutex
	client   *http.Client

	common Service

	PC *pc.Client
	// PE *pe.Service
	Calm   *calm.Client
	Karbon *karbon.Client
	// VM      *pc.VMService
}

// Service is used to create endpoint specific services to utilize concurrency with timeouts
type Service struct {
	client *Client
}

// Config is used to initialize the required client configs like URL, auth, etc....
type Config struct {
	PrismCentral       *pc.ServiceConfig     `json:"prism_central,omitempty"`
	PrismElement       *pe.ServiceConfig     `json:"prism_element,omitempty"`
	Calm               *calm.ServiceConfig   `json:"calm,omitempty"`
	Karbon             *karbon.ServiceConfig `json:"karbon,omitempty"`
	InsecureSkipVerify bool                  `json:"insecure,omitempty"`
}

// PrismElement is the config for PE requests
type PrismElement struct {
	URL  *string
	User *string
	Pass *string
}

// NewClient is used to initial
func NewClient(httpClient *http.Client, conf *Config) (*Client, error) {
	if (Config{}) == *conf {
		return nil, errors.New("no service configurations provided")
	}
	if httpClient == nil {
		if conf.InsecureSkipVerify {
			httpClient = &http.Client{Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}}
		} else {
			httpClient = &http.Client{}
		}
	}

	c := &Client{client: httpClient, config: conf}

	c.common.client = c
	var pcError, calmError, karbonError error
	if conf.PrismCentral != nil {
		c.PC, pcError = pc.NewClient(httpClient, conf.PrismCentral)
		if pcError != nil {
			fmt.Println("error in configuring Prism Central connector: ", pcError)
			os.Exit(1)
		}
	}

	if conf.Calm != nil {
		c.Calm, calmError = calm.NewClient(httpClient, conf.Calm)
		if calmError != nil {
			fmt.Println("error in configuring Calm connector: ", calmError)
			os.Exit(1)
		}
	}

	if conf.Karbon != nil {
		c.Karbon, karbonError = karbon.NewClient(httpClient, conf.Karbon)
		if karbonError != nil {
			fmt.Println("error in configuring Karbon connector: ", karbonError)
			os.Exit(1)
		}
	}

	return c, nil
}

// String returns the pointer for a string
func String(str string) *string {
	return &str
}
