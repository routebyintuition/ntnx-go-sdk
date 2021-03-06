package karbon

import (
	"net/http"
)

// MetaService handles communication to the clusters REST API endpoint
type MetaService Service

// MetaGetRequest gets a detailed list of version data
type MetaGetRequest struct{}

// MetaGetResponse are the version details
type MetaGetResponse struct {
	BuildDate string `json:"build_date"`
	GitCommit string `json:"git_commit"`
	Version   string `json:"version"`
}

// GetVersion Get the Karbon controller version of the k8 cluster.
func (ms *MetaService) GetVersion(reqdata *MetaGetRequest) (*MetaGetResponse, *http.Response, error) {

	u := "v1-alpha.1/version"

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ms.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var meta *MetaGetResponse
	resp, err := ms.client.Do(req, &meta)
	if err != nil {
		return nil, resp, err
	}

	return meta, resp, nil
}
