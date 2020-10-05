package pc

import (
	"fmt"
	"net/http"
)

// NetworkSecurityRuleService provides api function for network security rules
type NetworkSecurityRuleService Service

// NSRExportRequest empty GET request to receive export of rules
type NSRExportRequest struct{}

// NSRExportResponse 200 response on success
type NSRExportResponse struct{}

// NSRListRequest is the sort and filter options for a nsr list
type NSRListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// NSRListResponse provides the list of network security rules
type NSRListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// NSRImportRequest - Imports previously exported network security rules
type NSRImportRequest struct{}

// NSRImportResponse provides the UUID of the task performing the import
type NSRImportResponse struct {
	TaskUUID string `json:"task_uuid,omitempty"`
}

// NSRCreateRequest details on the security rule to create
type NSRCreateRequest struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       string   `json:"spec,omitempty"`
}

// NSRCreateResponse details on the rule created from request
type NSRCreateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       string   `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// NSRImportDryRunRequest empty request to dry run an import
type NSRImportDryRunRequest struct{}

// NSRImportDryRunResponse provides entity details
type NSRImportDryRunResponse struct {
	EntityList []EntityList `json:"entity_list,omitempty"`
}

// NSRDeleteRequest provides UUID of NSR to delete
type NSRDeleteRequest struct {
	UUID string
}

// NSRDeleteResponse provides details of deleted NSR
type NSRDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       string   `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// NSRGetRequest gets an NSR by provided UUID
type NSRGetRequest struct {
	UUID string
}

// NSRGetResponse provides details on an NSR from provided UUID
type NSRGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       string   `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// NSRUpdateRequest updates a NSR identified by provided UUID with provided Data element details
type NSRUpdateRequest struct {
	UUID string
	Data NSRUpdateRequestData
}

// NSRUpdateRequestData data to use to update NSR identified by UUID
type NSRUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       string   `json:"spec,omitempty"`
}

// NSRUpdateResponse provides update details on NSR
type NSRUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       string   `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// Update - This operation submits a request to update a existing Network security rule based on the input parameters.
func (nsrs *NetworkSecurityRuleService) Update(reqdata *NSRUpdateRequest) (*NSRUpdateResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("GET", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRUpdateResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get - This operation gets a existing Network security rule.
func (nsrs *NetworkSecurityRuleService) Get(reqdata *NSRGetRequest) (*NSRGetResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRGetResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Delete - This operation submits a request to delete a existing Network security rule.
func (nsrs *NetworkSecurityRuleService) Delete(reqdata *NSRDeleteRequest) (*NSRDeleteResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRDeleteResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// ImportDryRun - Generates a report on the impact of importing the policy data
func (nsrs *NetworkSecurityRuleService) ImportDryRun(reqdata *NSRImportDryRunRequest) (*NSRImportDryRunResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/import/dry_run")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRImportDryRunResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Create - This operation submits a request to create a new Network security rule based on the input parameters.
func (nsrs *NetworkSecurityRuleService) Create(reqdata *NSRCreateRequest) (*NSRCreateResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRCreateResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Import - Imports previously exported network security rules
func (nsrs *NetworkSecurityRuleService) Import(reqdata *NSRImportRequest) (*NSRImportResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/import/apply")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRImportResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// List - This operation gets a list of Network security rules, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
func (nsrs *NetworkSecurityRuleService) List(reqdata *NSRListRequest) (*NSRListResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRListResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Export - Export all network security rules to save and for subsequent import
func (nsrs *NetworkSecurityRuleService) Export(reqdata *NSRExportRequest) (*NSRExportResponse, *http.Response, error) {

	u := fmt.Sprintf("network_security_rules/export")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := nsrs.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *NSRExportResponse
	resp, err := nsrs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
