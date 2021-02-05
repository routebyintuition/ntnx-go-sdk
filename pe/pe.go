package pe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sync"

	"github.com/google/go-querystring/query"
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
	Disk    *DiskService
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
	c.Disk = (*DiskService)(&c.common)
	//c.Apps = (*AppsService)(&c.common)

	return c, nil
}

// NewRequest creates an API request
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	u, err := c.PEURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	encodedAuth := FormatCredentials(c.Login)
	basicAuthHeader := fmt.Sprintf("Basic %s", encodedAuth)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", basicAuthHeader)

	return req, nil
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

// FormatCredentials takes in an AuthCredential type and returns a base64 encoded string
// in the proper format for use by the REST API
func FormatCredentials(a AuthCredentials) string {
	authString := fmt.Sprintf("%s:%s", a.Username, a.Password)

	encoded := base64.StdEncoding.EncodeToString([]byte(authString))

	return encoded
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

// CheckResponse looks for common response errors
// Need to fill this in at some point
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return errors.New(fmt.Sprintln("Response error: ", r.StatusCode, r.Status))
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Received an error on request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
