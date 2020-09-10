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
	Type  string `json:"type"`
	Items Items  `json:"items"`
}

// RegistryCreateRequest is the request needed to create registries
type RegistryCreateRequest struct {
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties Properties `json:"properties"`
}

// RegistryCreateResponse is the request needed to create registries
type RegistryCreateResponse struct {
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties Properties `json:"properties"`
}

// RegistryGetRequest is the request needed to get registries
type RegistryGetRequest struct {
	Name string `json:"name"`
}

// RegistryGetResponse is the request needed to get registries
type RegistryGetResponse struct {
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties Properties `json:"properties"`
}

// RegistryDeleteRequest is the request needed to delete registries
type RegistryDeleteRequest struct {
	Name string `json:"name"`
}

// RegistryDeleteResponse is the request needed to delete registries
type RegistryDeleteResponse struct {
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties Properties `json:"properties"`
}

// Delete will delete the virtual machine associated with the provided UUID
func (rs *RegistryService) Delete(reqdata *RegistryDeleteRequest) (*RegistryDeleteResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("RegistryDeleteRequest cannot be nil")
	}

	if len(reqdata.Name) < 3 {
		return nil, nil, errors.New("invalid name format in RegistryDeleteRequest")
	}

	u := fmt.Sprintf("registries/%s", reqdata.Name)

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

	u := fmt.Sprintf("registries/%s", reqdata.Name)

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

	u := fmt.Sprintf("registries")

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
	u := "registries"

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
