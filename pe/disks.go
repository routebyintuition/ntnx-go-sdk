package pe

import (
	"fmt"
	"net/http"
)

// DiskService handles communication to CVM/PE for disk interactions
type DiskService Service

// DiskListRequest provides query paramaters for disk list
type DiskListRequest struct {
	Count               int    `url:"count,omitempty"`
	FilterCriteria      string `url:"filter_criteria,omitempty"`
	SortCriteria        string `url:"sort_criteria,omitempty"`
	SearchString        string `url:"search_string,omitempty"`
	SearchAttributeList string `url:"search_attribute_list,omitempty"`
	Page                int    `url:"page,omitempty"`
	Projection          string `url:"projection,omitempty"`
}

// DiskListResponse provides the response for listing all disk details
type DiskListResponse struct {
	Metadata Metadata   `json:"metadata"`
	Entities []Entities `json:"entities"`
}

// DiskVirtualListRequest provides GET parameters for vdisk list
type DiskVirtualListRequest struct {
	Count               int    `url:"count,omitempty"`
	FilterCriteria      string `url:"filter_criteria,omitempty"`
	SortCriteria        string `url:"sort_criteria,omitempty"`
	SearchString        string `url:"search_string,omitempty"`
	SearchAttributeList string `url:"search_attribute_list,omitempty"`
	Page                int    `url:"page,omitempty"`
	Projection          string `url:"projection,omitempty"`
}

// DiskVirtualListResponse provides details on all virtual disks including NFS paths
type DiskVirtualListResponse struct {
	Metadata Metadata   `json:"metadata"`
	Entities []Entities `json:"entities"`
}

// DiskVirtualGetRequest provides GET parameters for vdisk get by UUID
type DiskVirtualGetRequest struct {
	Query *DiskVirtualGetRequestQuery
}

// DiskVirtualGetRequestQuery provides the UUID for the vdisk get
type DiskVirtualGetRequestQuery struct {
	UUID string `url:"uuid,omitempty"`
}

// DiskVirtualGetResponse provides details on all virtual disks including NFS paths
type DiskVirtualGetResponse struct {
	Metadata Metadata   `json:"metadata"`
	Entities []Entities `json:"entities"`
}

// List makes the call to list hardware disks
func (dsk *DiskService) List(reqdata *DiskListRequest) (*DiskListResponse, *http.Response, error) {

	u := "disks"

	u, err := addOptions(u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	req, err := dsk.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *DiskListResponse
	resp, err := dsk.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

// ListVDisk makes the call to list virtual disks
func (dsk *DiskService) ListVDisk(reqdata *DiskVirtualListRequest) (*DiskVirtualListResponse, *http.Response, error) {

	u := "virtual_disks"

	u, err := addOptions(u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	req, err := dsk.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *DiskVirtualListResponse
	resp, err := dsk.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

// GetVDisk provides vdisk information based upon UUID
func (dsk *DiskService) GetVDisk(reqdata *DiskVirtualGetRequest) (*DiskVirtualGetResponse, *http.Response, error) {

	u := fmt.Sprintf("virtual_disks/%s", reqdata.Query.UUID)

	u, err := addOptions(u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	req, err := dsk.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *DiskVirtualGetResponse
	resp, err := dsk.client.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
