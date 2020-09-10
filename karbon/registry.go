package karbon

import (
	"errors"
	"fmt"
	"net/http"
)

// RegistryService handles communication to the clusters REST API endpoint
type RegistryService Service

// RegistryListRequest is the request needed to get a list of registries
type RegistryListRequest struct{}

// RegistryListResponse is the request needed to get a list of registries
type RegistryListResponse struct {
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}

// RegistryCreateRequest is the request needed to create registries
type RegistryCreateRequest struct {
	Cert string
	Name string
	Port int
	URL  string
}

// RegistryCreateResponse is the request needed to create registries
type RegistryCreateResponse struct {
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}

// RegistryGetRequest is the request needed to get registries
type RegistryGetRequest struct {
	Name string `json:"name"`
}

// RegistryGetResponse is the request needed to get registries
type RegistryGetResponse struct {
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}

// RegistryDeleteRequest is the request needed to delete registries
type RegistryDeleteRequest struct {
	Name string `json:"name"`
}

// RegistryDeleteResponse is the request needed to delete registries
type RegistryDeleteResponse struct {
	RegistryName string `json:"registry_name"`
}

// Delete will delete the virtual machine associated with the provided UUID
func (rs *RegistryService) Delete(reqdata *RegistryDeleteRequest) (*RegistryDeleteResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("RegistryDeleteRequest cannot be nil")
	}

	if len(reqdata.Name) < 3 {
		return nil, nil, errors.New("invalid name format in RegistryDeleteRequest")
	}

	u := fmt.Sprintf("v1-alpha.1/registries/%s", reqdata.Name)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var registry *RegistryDeleteResponse
	resp, err := rs.client.Do(req, &registry)
	if err != nil {
		return nil, resp, err
	}

	return registry, resp, nil
}

// Get passes as name of a specific registry to get information about that registry
func (rs *RegistryService) Get(reqdata *RegistryGetRequest) (*RegistryGetResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("RegistryGetRequest cannot be nil")
	}

	if len(reqdata.Name) < 3 {
		return nil, nil, errors.New("invalid name format in RegistryGetRequest")
	}

	u := fmt.Sprintf("v1-alpha.1/registries/%s", reqdata.Name)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var registry *RegistryGetResponse
	resp, err := rs.client.Do(req, &registry)
	if err != nil {
		return nil, resp, err
	}

	return registry, resp, nil
}

// List returns a list of all registries
func (rs *RegistryService) List(reqdata *RegistryGetRequest) (*RegistryGetResponse, *http.Response, error) {

	u := fmt.Sprintf("v1-alpha.1/registries")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var registry *RegistryGetResponse
	resp, err := rs.client.Do(req, &registry)
	if err != nil {
		return nil, resp, err
	}

	return registry, resp, nil
}

// Create makes the call to cluster list
func (rs *RegistryService) Create(reqdata *RegistryCreateRequest) (*RegistryCreateResponse, *http.Response, error) {
	u := "v1-alpha.1/registries"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var registry *RegistryCreateResponse
	resp, err := rs.client.Do(req, &registry)
	if err != nil {
		return nil, resp, err
	}

	return registry, resp, nil
}
