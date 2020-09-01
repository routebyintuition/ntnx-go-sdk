package calm

import (
	"fmt"
	"net/http"
)

// BlueprintService handles communication to the Calm BPs REST API endpoint
type BlueprintService Service

type BlueprintCloneRequest struct{}
type BlueprintGetRuntimeVariablesRequest struct{}
type BlueprintCreateRequest struct{}
type BlueprintAbortRequest struct{}
type BlueprintGetPendingLaunchRequest struct{}
type BlueprintLaunchRequest struct{}
type BlueprintLaunchMarketplaceRequest struct{}
type BlueprintImportRequest struct{}
type BlueprintDownloadRequest struct{}

type BlueprintUploadResponse struct{}
type BlueprintCloneResponse struct{}
type BlueprintGetRuntimeVariablesResponse struct{}
type BlueprintCreateResponse struct{}
type BlueprintAbortResponse struct{}
type BlueprintGetPendingLaunchResponse struct{}
type BlueprintLaunchResponse struct{}
type BlueprintLaunchMarketplaceResponse struct{}
type BlueprintImportResponse struct{}
type BlueprintDownloadResponse struct{}

// Run makes the call to cluster run
func (bps *BlueprintService) Run(reqdata *BlueprintRunRequest, uuid string) (*BlueprintRunResponse, *http.Response, error) {

	u := fmt.Sprintf("blueprints/%s/run_script", uuid)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bps.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var calmResp *BlueprintRunResponse
	resp, err := bps.client.Do(req, &calmResp)
	if err != nil {
		return nil, resp, err
	}

	return calmResp, resp, nil
}

// Export will export the identified blueprint
func (bps *BlueprintService) Export(reqdata *BlueprintExportRequest, uuid string) (*BlueprintExportResponse, *http.Response, error) {

	u := fmt.Sprintf("blueprints/%s/export_json", uuid)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bps.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var calmResp *BlueprintExportResponse
	resp, err := bps.client.Do(req, &calmResp)
	if err != nil {
		return nil, resp, err
	}

	return calmResp, resp, nil
}

// Upload will update the cluster defined by the UUID with the provided information
func (bps *BlueprintService) Upload(reqdata *BlueprintUploadRequest, uuid string) (*BlueprintUploadResponse, *http.Response, error) {

	u := fmt.Sprintf("clusters/%s", uuid)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bps.client.NewRequest("PUT", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var calmResp *BlueprintUploadResponse
	resp, err := bps.client.Do(req, &calmResp)
	if err != nil {
		return nil, resp, err
	}

	return calmResp, resp, nil
}
