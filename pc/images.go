package pc

import (
	"fmt"
	"net/http"
)

// ImageService is the image-specific service
type ImageService Service

// ImageGetFileRequest Downloads the image based on the UUID specified.
type ImageGetFileRequest struct {
	UUID string
}

// ImageGetFileResponse provides the disk contents to save in string binary format
type ImageGetFileResponse string

/*
type ImageUploadRequest struct{
	UUID string
	Data ImageUploadRequestData
}
type ImageUploadResponse struct{}
*/

// ImageCreateRequest provides am image configuration to create
type ImageCreateRequest struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

// ImageCreateResponse is the response from an image create
type ImageCreateResponse struct {
	APIVersion string   `json:"api_version"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

// ImageListRequest This operation gets a list of images, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
type ImageListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// ImageListResponse response of all images in system
type ImageListResponse struct {
	APIVersion string     `json:"api_version"`
	Metadata   Metadata   `json:"metadata"`
	Entities   []Entities `json:"entities"`
}

// ImageDeleteRequest provides UUID of image to delete
type ImageDeleteRequest struct {
	UUID string
}

// ImageDeleteResponse response of delete image call on image UUID
type ImageDeleteResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// ImageGetRequest is the UUID of the image to retrieve
type ImageGetRequest struct {
	UUID string
}

// ImageGetResponse details about requested image
type ImageGetResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
	Status     Status   `json:"status,omitempty"`
}

// ImageUpdateRequest contains the UUID of the image to update and the data to be updated
type ImageUpdateRequest struct {
	UUID string
	Data ImageUpdateRequestData
}

// ImageUpdateRequestData image data provided during update request
type ImageUpdateRequestData struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// ImageUpdateResponse response to an image update request
type ImageUpdateResponse struct {
	APIVersion string   `json:"api_version,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

// ImageMigrateRequest requests an image migration, PC and PE
type ImageMigrateRequest struct {
	ClusterReference   ClusterReference     `json:"cluster_reference,omitempty"`
	ImageReferenceList []ImageReferenceList `json:"image_reference_list,omitempty"`
}

// ImageMigrateResponse migration task UUID response
type ImageMigrateResponse struct {
	TaskUUID string `json:"task_uuid"`
}

// GetFile Downloads the image based on the UUID specified.
func (is *ImageService) GetFile(reqdata *ImageGetFileRequest) (*ImageGetFileResponse, *http.Response, error) {
	u := fmt.Sprintf("images/%s/file", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageGetFileResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Create Images are raw ISO, QCOW2, or VMDK files that are uploaded by a user can be attached to a VM.
// An ISO image is attached as a virtual CD-ROM drive, and QCOW2 and VMDK files are attached as SCSI disks.
// An image has to be explicitly added to the self-service catalog before users can create VMs from it.
func (is *ImageService) Create(reqdata *ImageCreateRequest) (*ImageCreateResponse, *http.Response, error) {
	u := fmt.Sprintf("images")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageCreateResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// List This operation gets a list of images, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
func (is *ImageService) List(reqdata *ImageListRequest) (*ImageListResponse, *http.Response, error) {
	u := fmt.Sprintf("images/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageListResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Delete submits a request to delete an existing image
func (is *ImageService) Delete(reqdata *ImageDeleteRequest) (*ImageDeleteResponse, *http.Response, error) {
	u := fmt.Sprintf("images/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageDeleteResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get gets an existing image
func (is *ImageService) Get(reqdata *ImageGetRequest) (*ImageGetResponse, *http.Response, error) {
	u := fmt.Sprintf("images/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageGetResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Update updates the image with provided UUID with provided DATA config
func (is *ImageService) Update(reqdata *ImageUpdateRequest) (*ImageUpdateResponse, *http.Response, error) {
	u := fmt.Sprintf("images/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("PUT", u, reqdata.Data)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageUpdateResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Migrate migrates an image to cluster resources
func (is *ImageService) Migrate(reqdata *ImageMigrateRequest) (*ImageMigrateResponse, *http.Response, error) {
	u := fmt.Sprintf("images/migrate")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *ImageMigrateResponse

	resp, err := is.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
