package pc

import (
	"fmt"
	"net/http"
)

// AccessControlPoliciesService handles communication to the AccessControlPolicies REST API endpoint
type AccessControlPoliciesService Service

// CreateACPRequest is the request data to create a new access control policy
type CreateACPRequest struct {
	APIVersion string      `json:"api_version,omitempty"`
	Metadata   ACPMetadata `json:"metadata,omitempty"`
	Spec       ACPSpec     `json:"spec,omitempty"`
}

type CreateACPResponse struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    []string   `json:"required,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
}

type ListACPRequest struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
}

type ListACPResponse struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    []string   `json:"required,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
}

type DeleteACPRequest struct {
	UUID string `json:"uuid,omitempty"`
}

type DeleteACPResponse struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    []string   `json:"required,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
}

type GetACPRequest struct {
	UUID string `json:"uuid,omitempty"`
}

type GetACPResponse struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    []string   `json:"required,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
}

type UpdateACPRequest struct {
	UUID        string     `json:"uuid,omitempty"`
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    []string   `json:"required,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
}

type UpdateACPResponse struct {
	Type        string     `json:"type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    []string   `json:"required,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
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
