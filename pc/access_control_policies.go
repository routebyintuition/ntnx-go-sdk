package pc

import (
	"fmt"
	"net/http"
)

// AccessControlPoliciesService handles communication to the AccessControlPolicies REST API endpoint
type AccessControlPoliciesService Service

// CreateACPRequest is the request data to create a new access control policy
type CreateACPRequest struct {
	APIVersion string      `json:"api_version"`
	Metadata   ACPMetadata `json:"metadata"`
	Spec       ACPSpec     `json:"spec"`
}

type CreateACPResponse struct {
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}

type ListACPRequest struct {
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Properties  Properties `json:"properties"`
}

type ListACPResponse struct {
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}

type DeleteACPRequest struct {
	UUID string `json:"uuid,omitempty"`
}

type DeleteACPResponse struct {
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}

type GetACPRequest struct {
	UUID string `json:"uuid,omitempty"`
}

type GetACPResponse struct {
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}

type UpdateACPRequest struct {
	UUID        string     `json:"uuid"`
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}

type UpdateACPResponse struct {
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}

// Create makes the call to acp list
func (acp *AccessControlPoliciesService) Create(reqdata *CreateACPRequest) (*CreateACPResponse, *http.Response, error) {
	u := "access_control_policies"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := acp.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var acpResp *CreateACPResponse
	resp, err := acp.client.Do(req, &acpResp)
	if err != nil {
		return nil, resp, err
	}

	return acpResp, resp, nil
}

// List makes the call to acp list
func (acp *AccessControlPoliciesService) List(reqdata *ListACPRequest) (*ListACPResponse, *http.Response, error) {
	u := "access_control_policies/list"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := acp.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var acpResp *ListACPResponse
	resp, err := acp.client.Do(req, &acpResp)
	if err != nil {
		return nil, resp, err
	}

	return acpResp, resp, nil
}

// Update makes the call to acp delete
func (acp *AccessControlPoliciesService) Update(reqdata *UpdateACPRequest) (*UpdateACPResponse, *http.Response, error) {
	u := fmt.Sprintf("access_control_policies/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := acp.client.NewRequest("PUT", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var acpResp *UpdateACPResponse
	resp, err := acp.client.Do(req, &acpResp)
	if err != nil {
		return nil, resp, err
	}

	return acpResp, resp, nil
}

// Delete makes the call to acp delete
func (acp *AccessControlPoliciesService) Delete(reqdata *DeleteACPRequest) (*DeleteACPResponse, *http.Response, error) {
	u := fmt.Sprintf("access_control_policies/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := acp.client.NewRequest("DELETE", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var acpResp *DeleteACPResponse
	resp, err := acp.client.Do(req, &acpResp)
	if err != nil {
		return nil, resp, err
	}

	return acpResp, resp, nil
}

// Get makes the call to acp delete
func (acp *AccessControlPoliciesService) Get(reqdata *GetACPRequest) (*GetACPResponse, *http.Response, error) {
	u := fmt.Sprintf("access_control_policies/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := acp.client.NewRequest("GET", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var acpResp *GetACPResponse
	resp, err := acp.client.Do(req, &acpResp)
	if err != nil {
		return nil, resp, err
	}

	return acpResp, resp, nil
}
