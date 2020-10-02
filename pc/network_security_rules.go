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
	APIVersion string     `json:"api_version"`
	Metadata   Metadata   `json:"metadata"`
	Entities   []Entities `json:"entities"`
}

type NSRImportRequest struct{}
type NSRImportResponse struct{}

type NSRCreateRequest struct{}
type NSRCreateResponse struct{}

type NSRImportDryRunRequest struct{}
type NSRImportDryRunResponse struct{}

type NSRDeleteRequest struct{}
type NSRDeleteResponse struct{}

type NSRGetRequest struct{}
type NSRGetResponse struct{}

type NSRUpdateRequest struct{}
type NSRUpdateResponse struct{}

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
