package pc

import (
	"fmt"
	"net/http"
)

// PermissionService is the service for the permissions API in prism central
type PermissionService Service

// PermissionCreateRequest creates a permission based upon data provided
type PermissionCreateRequest struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// PermissionCreateResponse provides a response to the create permission request
type PermissionCreateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// PermissionUpdateRequest updates a permission based on provided UUID and permission data
type PermissionUpdateRequest struct {
	UUID string
	Data PermissionUpdateRequestData
}

// PermissionUpdateRequestData details about the permission to update
type PermissionUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// PermissionUpdateResponse provides response to permission update request
type PermissionUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// PermissionGetRequest provides UUID of permission to get
type PermissionGetRequest struct {
	UUID string
}

// PermissionGetResponse provides details about permission identified by UUID
type PermissionGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// PermissionDeleteRequest provides UUID of permission to delete
type PermissionDeleteRequest struct {
	UUID string
}

// PermissionDeleteResponse is the data of the permission deleted
type PermissionDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// PermissionListRequest provides filter and sort for a permissions list request
type PermissionListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// PermissionListResponse provides permission list results
type PermissionListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// List - lists all permissions
func (ps *PermissionService) List(reqdata *PermissionListRequest) (*PermissionListResponse, *http.Response, error) {

	u := fmt.Sprintf("permissions/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *PermissionListResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Delete - deletes a permission by UUID
func (ps *PermissionService) Delete(reqdata *PermissionDeleteRequest) (*PermissionDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("permissions/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *PermissionDeleteResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get - gets details on a permission
func (ps *PermissionService) Get(reqdata *PermissionGetRequest) (*PermissionGetResponse, *http.Response, error) {

	u := fmt.Sprintf("permissions/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *PermissionGetResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Update - updates a permission
func (ps *PermissionService) Update(reqdata *PermissionUpdateRequest) (*PermissionUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("permissions/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *PermissionUpdateResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Create - Permissions define what a user can do with each type.
// For example, an admin can upload images; a DevOps user can create, edit, or delete a VM; and an operations user can view a VM console.
func (ps *PermissionService) Create(reqdata *PermissionCreateRequest) (*PermissionCreateResponse, *http.Response, error) {

	u := fmt.Sprintf("permissions")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *PermissionCreateResponse
	resp, err := ps.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
