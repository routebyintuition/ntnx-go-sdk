package pc

import (
	"fmt"
	"net/http"
)

// RoleService handles communication to the roles REST API endpoint
type RoleService Service

// RoleCreateRequest provides role details for creation
type RoleCreateRequest struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

// RoleCreateResponse proides the response from role create
type RoleCreateResponse struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

// RoleDeleteRequest provides UUID of role to delete
type RoleDeleteRequest struct {
	UUID string
}

// RoleDeleteResponse is the response and data for the role delete request
type RoleDeleteResponse struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

// RoleGetRequest gets details on a role by UUID
type RoleGetRequest struct {
	UUID string
}

// RoleGetResponse provides details on the role requested by UUID
type RoleGetResponse struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

// RoleUpdateRequest updates the role identified by UUID with the data provided
type RoleUpdateRequest struct {
	UUID string
	Data RoleUpdateRequestData
}

// RoleUpdateRequestData is the data for the role to be updated
type RoleUpdateRequestData struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

// RoleUpdateResponse is role data post update
type RoleUpdateResponse struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

// RoleListRequest is the filter and sort for the list request
type RoleListRequest struct {
	Kind          string `json:"kind"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Filter        string `json:"filter"`
	SortOrder     string `json:"sort_order"`
	SortAttribute string `json:"sort_attribute"`
}

// RoleListResponse is the list of roles
type RoleListResponse struct {
	APIVersion string     `json:"api_version"`
	Metadata   Metadata   `json:"metadata"`
	Entities   []Entities `json:"entities"`
}

// List - request a list of roles
func (rs *RoleService) List(reqdata *RoleListRequest) (*RoleListResponse, *http.Response, error) {

	u := fmt.Sprintf("roles/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *RoleListResponse
	resp, err := rs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Update - updates a role with the provided data by UUID
func (rs *RoleService) Update(reqdata *RoleUpdateRequest) (*RoleUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("roles/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *RoleUpdateResponse
	resp, err := rs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get - gets a role by UUID
func (rs *RoleService) Get(reqdata *RoleGetRequest) (*RoleGetResponse, *http.Response, error) {

	u := fmt.Sprintf("roles/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *RoleGetResponse
	resp, err := rs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Delete - deletes a role
func (rs *RoleService) Delete(reqdata *RoleDeleteRequest) (*RoleDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("roles/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *RoleDeleteResponse
	resp, err := rs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Create - A role is a collection of permissions defined for one or more kinds. A kind represents the type of an entity (such as VM).
// Roles are defined by users who have permission to create roles and assign roles to projects. All users in a project inherit the role.
func (rs *RoleService) Create(reqdata *RoleCreateRequest) (*RoleCreateResponse, *http.Response, error) {

	u := fmt.Sprintf("roles")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *RoleCreateResponse
	resp, err := rs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
