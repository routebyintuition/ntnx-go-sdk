package pe

import (
	"errors"
	"net/http"
	"net/url"
	"sync"
)

// Service is used to create endpoint specific services to utilize concurrency with timeouts
type Service struct {
	client *Client
}

// AuthCredentials are the username and password received from ENV variables
// never hard code these values into any code!
type AuthCredentials struct {
	Username string
	Password string
}

// Client is an HTTP Client
type Client struct {
	clientMu sync.Mutex
	client   *http.Client

	CalmURL *url.URL

	Login AuthCredentials

	common Service

	Calm *ClustersService
}

// NewClient is used for Nutanix Prism client instantiation
// Required:
//  prismURL: <string value of full URL to Prism Central>
//  httpClient: either nil or an HTTP client with non-default timeouts, etc...
func NewClient(httpClient *http.Client, prismURL string) (*Client, error) {
	if len(prismURL) == 0 {
		return nil, errors.New("URL should be specified in the format: https://{host}:{port}/api/nutanix/v3/")
	}
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(prismURL)
	login, err := GetCredentials()
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: baseURL, Login: login}
	c.common.client = c
	c.VM = (*VMService)(&c.common)
	c.Cluster = (*ClusterService)(&c.common)
	//c.Apps = (*AppsService)(&c.common)

	return c, nil
}
