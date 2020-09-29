package pc

import (
	"errors"
	"fmt"
	"net/http"
)

// UserService handles communication to the VM REST API endpoint
type UserService Service

// UserCreateRequest is the user create request
type UserCreateRequest struct {
	Status     string `json:"status,omitempty"`
	Spec       string `json:"spec,omitempty"`
	APIVersion string `json:"api_version,omitempty"`
	Metadata   int    `json:"metadata,omitempty"`
}

// UserCreateResponse is the user create response
type UserCreateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// UserListRequest is the user list request This operation gets a list of Users, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
type UserListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// UserListResponse is the user list response
type UserListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// UserMeRequest Displays the user currently logged in.
type UserMeRequest struct{}

// UserMeResponse Displays the user currently logged in.
type UserMeResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// UserGetRequest is the request to get details on a user
type UserGetRequest struct {
	UUID string
}

// UserGetResponse is the response details on a user from a get request
type UserGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// UserDeleteRequest is the request to delete a user
type UserDeleteRequest struct {
	UUID string
}

// UserDeleteResponse is the response to the user delete request
type UserDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// UserGetDomainRequest request to get user resources
// Retrieves specified user resource domain information.
type UserGetDomainRequest struct {
	UUID string
}

// UserGetDomainResponse is the response of the user resources
type UserGetDomainResponse struct {
	APIVersion                string                      `json:"api_version,omitempty"`
	ProjectResourceDomainList []ProjectResourceDomainList `json:"project_resource_domain_list,omitempty"`
}

// UserUpdateRequest submits a request to update a existing User based on the input parameters.
type UserUpdateRequest struct {
	UUID string
	Data UserUpdateRequestData
}

// UserUpdateRequestData provides details for user update
type UserUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// UserUpdateResponse provides the response to an update request
type UserUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// Update submits a request to update a existing User based on the input parameters.
func (us *UserService) Update(reqdata *UserUpdateRequest) (*UserUpdateResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("UserUpdateRequest cannot be nil")
	}

	if len(reqdata.UUID) < 8 {
		return nil, nil, errors.New("invalid UUID format in UserUpdateRequest")
	}

	u := fmt.Sprintf("users/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserUpdateResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}

// GetDomain Retrieves specified user resource domain information.
func (us *UserService) GetDomain(reqdata *UserGetDomainRequest) (*UserGetDomainResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("UserGetRequest cannot be nil")
	}

	if len(reqdata.UUID) < 8 {
		return nil, nil, errors.New("invalid UUID format in UserGetRequest")
	}

	u := fmt.Sprintf("users/%s/project_usage_summary", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserGetDomainResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}

// Delete This operation submits a request to delete a existing User.
func (us *UserService) Delete(reqdata *UserDeleteRequest) (*UserDeleteResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("UserDeleteRequest cannot be nil")
	}

	if len(reqdata.UUID) < 8 {
		return nil, nil, errors.New("invalid UUID format in UserDeleteRequest")
	}

	u := fmt.Sprintf("users/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserDeleteResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}

// Get This operation gets a existing User.
func (us *UserService) Get(reqdata *UserGetRequest) (*UserGetResponse, *http.Response, error) {
	if reqdata == nil {
		return nil, nil, errors.New("UserGetRequest cannot be nil")
	}

	if len(reqdata.UUID) < 8 {
		return nil, nil, errors.New("invalid UUID format in UserGetRequest")
	}

	u := fmt.Sprintf("users/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserGetResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}

// Me Displays the user currently logged in.
func (us *UserService) Me(reqdata *UserMeRequest) (*UserMeResponse, *http.Response, error) {

	u := "users/me"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserMeResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}

// Create will clone a virtual machine identified by UUID
func (us *UserService) Create(reqdata *UserCreateRequest) (*UserCreateResponse, *http.Response, error) {

	u := "users"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserCreateResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}

// List operation gets a list of Users, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
func (us *UserService) List(reqdata *UserListRequest) (*UserListResponse, *http.Response, error) {

	u := "users/list"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var userData *UserListResponse
	resp, err := us.client.Do(req, &userData)
	if err != nil {
		return nil, resp, err
	}

	return userData, resp, nil
}
