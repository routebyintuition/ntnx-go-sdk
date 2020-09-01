package nutanix

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/routebyintuition/ntnx-go-sdk/calm"
	"github.com/routebyintuition/ntnx-go-sdk/pc"
)

// Client is an HTTP Client
type Client struct {
	config   *Config
	clientMu sync.Mutex
	client   *http.Client

	common Service

	PC *pc.Client
	// PE *pe.Service
	Calm *calm.Client
	// VM      *pc.VMService
}

// Service is used to create endpoint specific services to utilize concurrency with timeouts
type Service struct {
	client *Client
}

// Config is used to initialize the required client configs like URL, auth, etc....
type Config struct {
	PrismCentral *pc.ServiceConfig   `json:"prism_central,omitempty"`
	PrismElement *PrismElement       `json:"prism_element,omitempty"`
	Calm         *calm.ServiceConfig `json:"calm,omitempty"`
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
		httpClient = &http.Client{}
	}

	c := &Client{client: httpClient, config: conf}

	c.common.client = c
	var pcError, calmError error
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

	// c.Calm = (*calm.Service)(&c.common)
	//c.Apps = (*AppsService)(&c.common)

	return c, nil
}

// String returns the pointer for a string
func String(str string) *string {
	return &str
}
