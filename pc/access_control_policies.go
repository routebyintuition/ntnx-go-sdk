package pc

// AccessControlPolicies handles communication to the AccessControlPolicies REST API endpoint
type AccessControlPolicies Service

// CreateACPRequest is the request data to create a new access control policy
type CreateACPRequest struct {
	APIVersion string      `json:"api_version"`
	Metadata   ACPMetadata `json:"metadata"`
	Spec       ACPSpec     `json:"spec"`
}
