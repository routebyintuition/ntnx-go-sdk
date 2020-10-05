package pc

import (
	"fmt"
	"net/http"
)

// BatchService is the service for the batch API in prism central
type BatchService Service

// BatchSubmitRequest batch item list
type BatchSubmitRequest struct {
	APIVersion      string           `json:"api_version,omitempty"`
	APIRequestList  []APIRequestList `json:"api_request_list,omitempty"`
	ExecutionOrder  string           `json:"execution_order,omitempty"`
	ActionOnFailure string           `json:"action_on_failure,omitempty"`
}

// BatchSubmitResponse is the response
type BatchSubmitResponse struct {
	APIResponseList []APIResponseList `json:"api_response_list,omitempty"`
}

// Submit - Batching allows for instructions for several operations to be sent using a single HTTP request.
// Depending on the batch parameters, the Nutanix v3 gateway processes each independent operation sequentially or in parallel. Once all operations in the batch have been completed, a consolidated response is returned and the HTTP connection is closed.
// The batch API takes an array of logical HTTP requests represented as JSON arrays.
// Maximum size of the array should not exceed 60. Each request comprises the following:
//    A method (corresponding to HTTP methods such as GET, PUT, and POST)
//    A relative URL (relative_url)
//    (Optional) A body (for POST and PUT requests). The batch API returns an array of logical HTTP responses represented as JSON arrays containing the following:
//    A status code
//    (Optional) A body represented as a JSON-encoded string
func (bs *BatchService) Submit(reqdata *BatchSubmitRequest) (*BatchSubmitResponse, *http.Response, error) {

	u := fmt.Sprintf("batch")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *BatchSubmitResponse
	resp, err := bs.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
