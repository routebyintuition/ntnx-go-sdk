package pc

import (
	"fmt"
	"net/http"
)

// TaskService handles communication to the tasks REST API endpoint
type TaskService Service

// TaskGetRequest get task request by UUID
type TaskGetRequest struct {
	UUID string
}

// TaskGetResponse provides task details
type TaskGetResponse struct {
	APIVersion           string                 `json:"api_version,omitempty"`
	UUID                 string                 `json:"uuid,omitempty"`
	LogicalTimestamp     int                    `json:"logical_timestamp,omitempty"`
	ProgressMessage      string                 `json:"progress_message,omitempty"`
	OperationType        string                 `json:"operation_type,omitempty"`
	PercentageComplete   int                    `json:"percentage_complete,omitempty"`
	Status               string                 `json:"status,omitempty"`
	ParentTaskReference  ParentTaskReference    `json:"parent_task_reference,omitempty"`
	SubtaskReferenceList []SubtaskReferenceList `json:"subtask_reference_list,omitempty"`
	ClusterReference     ClusterReference       `json:"cluster_reference,omitempty"`
	CreationTime         string                 `json:"creation_time,omitempty"`
	CreationTimeUsecs    int                    `json:"creation_time_usecs,omitempty"`
	LastUpdateTime       string                 `json:"last_update_time,omitempty"`
	StartTime            string                 `json:"start_time,omitempty"`
	StartTimeUsecs       int                    `json:"start_time_usecs,omitempty"`
	CompletionTime       string                 `json:"completion_time,omitempty"`
	CompletionTimeUsecs  int                    `json:"completion_time_usecs,omitempty"`
	EntityReferenceList  []EntityReferenceList  `json:"entity_reference_list,omitempty"`
	ErrorCode            string                 `json:"error_code,omitempty"`
	ErrorDetail          string                 `json:"error_detail,omitempty"`
}

// TaskListRequest provides filter and sort for task list request
type TaskListRequest struct {
	Kind          string `json:"kind,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Filter        string `json:"filter,omitempty"`
	SortOrder     string `json:"sort_order,omitempty"`
	SortAttribute string `json:"sort_attribute,omitempty"`
}

// TaskListResponse provides task list results
type TaskListResponse struct {
	APIVersion string     `json:"api_version,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`
	Entities   []Entities `json:"entities,omitempty"`
}

// List - This operation gets a list of Tasks, allowing for sorting and pagination. Note: Entities that have not been created successfully are not listed.
func (ts *TaskService) List(reqdata *TaskListRequest) (*TaskListResponse, *http.Response, error) {

	u := fmt.Sprintf("tasks/list")

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ts.client.NewRequest("POST", u, reqdata)
	if err != nil {
		return nil, nil, err
	}

	var result *TaskListResponse
	resp, err := ts.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Get - This operation gets a existing Task.
func (ts *TaskService) Get(reqdata *TaskGetRequest) (*TaskGetResponse, *http.Response, error) {

	u := fmt.Sprintf("tasks/%s", reqdata.UUID)

	u, err := addOptions(u, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ts.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var result *TaskGetResponse
	resp, err := ts.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
