package karbon

// Cert is the certificate data
type Cert struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Format      string `json:"format"`
}

// Name is the karbon naming type for response parsing
type Name struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	MaxLength   int    `json:"maxLength"`
	MinLength   int    `json:"minLength"`
	Pattern     string `json:"pattern"`
	Example     string `json:"example"`
}

// Port are the port details from karbon
type Port struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// URL are the url values for karbon
type URL struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Properties are the overall data type for various response properties
type Properties struct {
	Cert         Cert         `json:"cert"`
	Name         Name         `json:"name"`
	Port         Port         `json:"port"`
	URL          URL          `json:"url"`
	Endpoint     Endpoint     `json:"endpoint"`
	UUID         UUID         `json:"uuid"`
	RegistryName RegistryName `json:"registry_name"`
	BuildDate    BuildDate    `json:"build_date"`
	GitCommit    GitCommit    `json:"git_commit"`
	Version      Version      `json:"version"`
}

// Endpoint are the endpoint response details
type Endpoint struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// UUID is the UUID format in karbon responses
type UUID struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Items are the individual items in karbon response
type Items struct {
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties Properties `json:"properties"`
}

// RegistryName is the naming format for the registry responses
type RegistryName struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type BuildDate struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}
type GitCommit struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}
type Version struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}
