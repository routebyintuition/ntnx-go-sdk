package pe

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
)

// NUTANIX_PE_USER is the environment variable name for the PE username
const NUTANIX_PE_USER = "NUTANIX_PE_USER"

// NUTANIX_PE_PASS is the environment variables for the PE password
const NUTANIX_PE_PASS = "NUTANIX_PE_PASS"

// NUTANIX_PE_URL is the environment variables for the PE URL
const NUTANIX_PE_URL = "NUTANIX_PE_URL"

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

	PEURL *url.URL

	Login AuthCredentials

	common Service

	Cluster *ClusterService
	// User    *UserService
	// VM      *VMService
}

// ServiceConfig is the configuration to use PC
type ServiceConfig struct {
	URL  *string
	User *string
	Pass *string
}

// NewClient is used for Nutanix Prism client instantiation
// Required:
//  prismURL: <string value of full URL to Prism Central>
//  httpClient: either nil or an HTTP client with non-default timeouts, etc...
func NewClient(httpClient *http.Client, conf *ServiceConfig) (*Client, error) {
	err := CheckConfig(conf)
	if err != nil {
		fmt.Println("minimal amount of information not available for Prism Central config: ", err)
		return nil, err
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(*conf.URL)
	login, err := GetCredentials(conf)
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, PEURL: baseURL, Login: login}
	c.common.client = c
	// c.VM = (*VMService)(&c.common)
	c.Cluster = (*ClusterService)(&c.common)
	//c.Apps = (*AppsService)(&c.common)

	return c, nil
}

// GetCredentials gets the username and password from the ENV
// variables for use in the application
func GetCredentials(conf *ServiceConfig) (AuthCredentials, error) {
	a := AuthCredentials{}
	if len(*conf.User) < 3 {
		a.Username = os.Getenv("NUTANIX_PE_USER")
	} else {
		a.Username = *conf.User
	}
	if len(*conf.Pass) < 8 {
		a.Password = os.Getenv("NUTANIX_PE_PASS")
	} else {
		a.Password = *conf.Pass
	}

	if len(a.Username) == 0 {
		return AuthCredentials{}, errors.New("No username found in environment variable NutanixUser or in provided config")
	}
	if len(a.Password) == 0 {
		return AuthCredentials{}, errors.New("No password found in environment variable NutanixPass or in provided config")
	}

	return a, nil
}

// CheckConfig will ensure that the minimal amount of information is present in the config
func CheckConfig(conf *ServiceConfig) error {
	errOut := errors.New("insufficient data for Prism Central config")
	if conf == nil {
		return errOut
	}
	if conf.URL == nil {
		URL := os.Getenv(NUTANIX_PE_URL)
		if len(URL) < 3 {
			return errors.New("insufficient data for Prism Element config: no URL")
		}
		conf.URL = &URL
	}
	if conf.User == nil {
		Username := os.Getenv(NUTANIX_PE_USER)
		if len(Username) < 3 {
			return errors.New("insufficient data for Prism Element config: no username")
		}
		conf.User = &Username
	}
	if conf.Pass == nil {
		Password := os.Getenv(NUTANIX_PE_PASS)
		if len(Password) < 8 {
			return errors.New("insufficient data for Prism Element config: no password")
		}
		conf.Pass = &Password
	}
	if len(*conf.URL) < 8 || len(*conf.User) < 3 || len(*conf.Pass) < 8 {
		return errOut
	}

	return nil
}
