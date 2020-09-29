package pc

import (
	"fmt"
	"net/http"
)

// CategoryService service for operations against the category prism central API
type CategoryService Service

// CategoryDeleteKeyRequest provides the category key to delete
type CategoryDeleteKeyRequest struct {
	Name string `json:"name,omitempty"`
}

// CategoryDeleteKeyResponse responds with just a 200 success HTTP code
type CategoryDeleteKeyResponse struct{}

// CategoryGetKeyRequest request to get the category key by name
type CategoryGetKeyRequest struct {
	Name string `json:"name,omitempty"`
}

// CategoryGetKeyResponse response for a GetKey
type CategoryGetKeyResponse struct {
	APIVersion    string       `json:"api_version,omitempty"`
	Name          string       `json:"name,omitempty"`
	Capabilities  Capabilities `json:"capabilities,omitempty"`
	Description   string       `json:"description,omitempty"`
	SystemDefined bool         `json:"system_defined,omitempty"`
}

// CategoryUpdateKeyRequest - Create or Update a category Key.
type CategoryUpdateKeyRequest struct {
	Name string `json:"name,omitempty"`
	Data CategoryUpdateKeyRequestData
}

// CategoryUpdateKeyRequestData data for create or update category key
type CategoryUpdateKeyRequestData struct {
	APIVersion   string       `json:"api_version,omitempty"`
	Name         string       `json:"name,omitempty"`
	Description  string       `json:"description,omitempty"`
	Capabilities Capabilities `json:"capabilities,omitempty"`
}

// CategoryUpdateKeyResponse response to creation or update of a key
type CategoryUpdateKeyResponse struct {
	APIVersion    string       `json:"api_version,omitempty"`
	Name          string       `json:"name,omitempty"`
	Capabilities  Capabilities `json:"capabilities,omitempty"`
	Description   string       `json:"description,omitempty"`
	SystemDefined bool         `json:"system_defined,omitempty"`
}

// CategoryListKeysRequest list keys request with filter and sort options
type CategoryListKeysRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// CategoryListKeysResponse provides the list of category keys
type CategoryListKeysResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// CategoryQueryRequest are the category query parameters
type CategoryQueryRequest struct {
	APIVersion        string         `json:"api_version,omitempty"`
	GroupMemberCount  int            `json:"group_member_count,omitempty"`
	GroupMemberOffset int            `json:"group_member_offset,omitempty"`
	UsageType         string         `json:"usage_type,omitempty"`
	CategoryFilter    CategoryFilter `json:"category_filter,omitempty"`
}

// CategoryQueryResponse category query results
type CategoryQueryResponse struct {
	APIVersion string    `json:"api_version,omitempty"`
	Metadata   Metadata  `json:"metadata,omitempty"`
	Results    []Results `json:"results,omitempty"`
}

// CategoryDeleteValueRequest passes a key and value to delete for the category
type CategoryDeleteValueRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// CategoryDeleteValueResponse returns a 200 HTTP code on success
type CategoryDeleteValueResponse struct{}

// CategoryGetValueRequest returns a 200 HTTP code on success
type CategoryGetValueRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// CategoryGetValueResponse are the results from getting a category value
type CategoryGetValueResponse struct {
	APIVersion     string         `json:"api_version,omitempty"`
	Name           string         `json:"name,omitempty"`
	Value          string         `json:"value,omitempty"`
	AssignmentRule AssignmentRule `json:"assignment_rule,omitempty"`
	Description    string         `json:"description,omitempty"`
	SystemDefined  bool           `json:"system_defined,omitempty"`
}

// CategoryUpdateValueRequest create or update category value by name and value with provided DATA
type CategoryUpdateValueRequest struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Data  CategoryUpdateValueRequestData
}

// CategoryUpdateValueRequestData is the data to update or create for the category value
type CategoryUpdateValueRequestData struct {
	APIVersion     string         `json:"api_version,omitempty"`
	Value          string         `json:"value,omitempty"`
	Description    string         `json:"description"`
	AssignmentRule AssignmentRule `json:"assignment_rule"`
}

// CategoryUpdateValueResponse is the response upon update or creation of a category value
type CategoryUpdateValueResponse struct {
	APIVersion     string         `json:"api_version"`
	Name           string         `json:"name"`
	Value          string         `json:"value"`
	AssignmentRule AssignmentRule `json:"assignment_rule"`
	Description    string         `json:"description"`
	SystemDefined  bool           `json:"system_defined"`
}

// CategoryListValuesRequest is the category name and data associated to list values
type CategoryListValuesRequest struct {
	Name string `json:"name,omitempty"`
	Data CategoryListValuesRequestData
}

// CategoryListValuesRequestData is the filter and sort for the category value list
type CategoryListValuesRequestData struct {
	Kind          string `json:"kind"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Filter        string `json:"filter"`
	SortOrder     string `json:"sort_order"`
	SortAttribute string `json:"sort_attribute"`
}

// CategoryListValuesResponse provides the list of values associated to a category key
type CategoryListValuesResponse struct {
	APIVersion string     `json:"api_version"`
	Metadata   Metadata   `json:"metadata"`
	Entities   []Entities `json:"entities"`
}

// ListValues - List the values for a specified key.
func (cas *CategoryService) ListValues(reqdata *CategoryListValuesRequest) (*CategoryListValuesResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s/list", reqdata.Name)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("POST", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryListValuesResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// UpdateValue - Create or Update a category value. Creates when value doesn’t exist.
func (cas *CategoryService) UpdateValue(reqdata *CategoryUpdateValueRequest) (*CategoryUpdateValueResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s/%s", reqdata.Name, reqdata.Value)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryUpdateValueResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetValue - Get a category value.
func (cas *CategoryService) GetValue(reqdata *CategoryGetValueRequest) (*CategoryGetValueResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s/%s", reqdata.Name, reqdata.Value)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryGetValueResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteValue - Delete a category value.
func (cas *CategoryService) DeleteValue(reqdata *CategoryDeleteValueRequest) (*CategoryDeleteValueResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s/%s", reqdata.Name, reqdata.Value)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryDeleteValueResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Query - Get list of entities attached to categories or policies in which categories are used as defined by the filter criteria.
func (cas *CategoryService) Query(reqdata *CategoryQueryRequest) (*CategoryQueryResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/query")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryQueryResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// ListKeys - List the category keys.
func (cas *CategoryService) ListKeys(reqdata *CategoryListKeysRequest) (*CategoryListKeysResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryListKeysResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// UpdateKey - Create or Update a category Key.
func (cas *CategoryService) UpdateKey(reqdata *CategoryUpdateKeyRequest) (*CategoryUpdateKeyResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s", reqdata.Name)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryUpdateKeyResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetKey - Get a category key.
func (cas *CategoryService) GetKey(reqdata *CategoryGetKeyRequest) (*CategoryGetKeyResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s", reqdata.Name)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryGetKeyResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteKey - Delete a category Key.
func (cas *CategoryService) DeleteKey(reqdata *CategoryDeleteKeyRequest) (*CategoryDeleteKeyResponse, *http.Response, error) {

	u := fmt.Sprintf("categories/%s", reqdata.Name)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cas.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *CategoryDeleteKeyResponse
	resp, err := cas.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
