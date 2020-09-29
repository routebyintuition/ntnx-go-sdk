package pc

import (
	"fmt"
	"net/http"
)

// ProjectService works with the prism central projects API
type ProjectService Service

// ProjectCreateRequest contains project details to create
type ProjectCreateRequest struct {
	Data ProjectCreateRequestData
}

// ProjectCreateRequestData is data associate with the project create request
type ProjectCreateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// ProjectCreateResponse is the project response and details
type ProjectCreateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// ProjectDeleteRequest UUID of project to delete
type ProjectDeleteRequest struct {
	UUID string
}

// ProjectDeleteResponse response from project delete request
type ProjectDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// ProjectGetRequest UUID of project to get details on
type ProjectGetRequest struct {
	UUID string
}

// ProjectGetResponse details of project identified by UUID
type ProjectGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// ProjectUpdateRequest is the UUID and data to update for a project
type ProjectUpdateRequest struct {
	UUID string
	Data ProjectUpdateRequestData
}

// ProjectUpdateRequestData is the data for the project update identified by UUID
type ProjectUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// ProjectUpdateResponse response from project update request
type ProjectUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// ProjectListRequest provides filter and sort options for a list of projects request
type ProjectListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// ProjectListResponse provides the response to a project list request
type ProjectListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// List - This operation gets a list of Projects, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
func (ps *CategoryService) List(reqdata *ProjectListRequest) (*ProjectListResponse, *http.Response, error) {

	u := fmt.Sprintf("projects/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *ProjectListResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Update - This operation submits a request to update a existing Project based on the input parameters.
func (ps *CategoryService) Update(reqdata *ProjectUpdateRequest) (*ProjectUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("projects/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *ProjectUpdateResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get - This operation gets a existing Project.
func (ps *CategoryService) Get(reqdata *ProjectGetRequest) (*ProjectGetResponse, *http.Response, error) {

	u := fmt.Sprintf("projects/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *ProjectGetResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Delete - This operation submits a request to delete a existing Project.
func (ps *CategoryService) Delete(reqdata *ProjectDeleteRequest) (*ProjectDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("projects/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *ProjectDeleteResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Create - This operation submits a request to create a new Project based on the input parameters.
func (ps *CategoryService) Create(reqdata *ProjectCreateRequest) (*ProjectCreateResponse, *http.Response, error) {

	u := fmt.Sprintf("projects")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("POST", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *ProjectCreateResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
