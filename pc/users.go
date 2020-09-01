package pc

import "net/http"

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
type UserCreateResponse struct{}

// UserListRequest is the user list request
type UserListRequest struct{}

// UserListResponse is the user list response
type UserListResponse struct{}

// UserActiveRequest is the list of users logged in request
type UserActiveRequest struct{}

// UserActiveResponse is the list of users logged in response
type UserActiveResponse struct{}

// UserGetRequest is the request to get details on a user
type UserGetRequest struct{}

// UserGetResponse is the response details on a user from a get request
type UserGetResponse struct{}

// UserDeleteRequest is the request to delete a user
type UserDeleteRequest struct{}

// UserDeleteResponse is the response to the user delete request
type UserDeleteResponse struct{}

// UserResourcesRequest request to get user resources
type UserResourcesRequest struct{}

// UserResourcesResponse is the response of the user resources
type UserResourcesResponse struct{}

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
