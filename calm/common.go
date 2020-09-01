package calm

type TargetDetails struct {
	Type                 string `json:"type"`
	Description          string `json:"description"`
	AdditionalProperties bool   `json:"additionalProperties"`
}
type Machine struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ProviderOperationPayload struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Attrs struct {
	Type                 string `json:"type"`
	Description          string `json:"description"`
	AdditionalProperties bool   `json:"additionalProperties"`
}
type Port struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Runid struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Spec struct {
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Properties  *Properties `json:"properties"`
}
type APIVersion struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type LastUpdateTime struct {
	ReadOnly    bool   `json:"readOnly"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Format      string `json:"format"`
}
type UseCategoriesMapping struct {
	Default     bool   `json:"default"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Kind struct {
	Default     string   `json:"default"`
	ReadOnly    bool     `json:"readOnly"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	XNtnxEnum   []string `json:"x-ntnx-enum"`
}
type UUID struct {
	Pattern     string `json:"pattern"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Format      string `json:"format"`
}

type ProjectReference struct {
	Description string      `json:"description"`
	Required    []string    `json:"required"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Properties  *Properties `json:"properties"`
}
type CreationTime struct {
	ReadOnly    bool   `json:"readOnly"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Format      string `json:"format"`
}
type SpecVersion struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type SpecHash struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Items struct {
	Type string `json:"type"`
}
type AdditionalProperties struct {
	Type  string `json:"type"`
	Items Items  `json:"items"`
}
type CategoriesMapping struct {
	Type                 string                `json:"type"`
	Description          string                `json:"description"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties"`
}
type ShouldForceTranslate struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type OwnerReference struct {
	Description string      `json:"description"`
	Required    []string    `json:"required"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Properties  *Properties `json:"properties"`
}

type Categories struct {
	Type                 string                `json:"type"`
	Description          string                `json:"description"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties"`
}
type Name struct {
	ReadOnly    bool   `json:"readOnly"`
	Type        string `json:"type"`
	Description string `json:"description"`
	MaxLength   int    `json:"maxLength"`
}
type Properties struct {
	LastUpdateTime           *LastUpdateTime           `json:"last_update_time"`
	UseCategoriesMapping     *UseCategoriesMapping     `json:"use_categories_mapping"`
	Kind                     *Kind                     `json:"kind"`
	UUID                     *UUID                     `json:"uuid"`
	ProjectReference         *ProjectReference         `json:"project_reference"`
	CreationTime             *CreationTime             `json:"creation_time"`
	SpecVersion              *SpecVersion              `json:"spec_version"`
	SpecHash                 *SpecHash                 `json:"spec_hash"`
	CategoriesMapping        *CategoriesMapping        `json:"categories_mapping"`
	ShouldForceTranslate     *ShouldForceTranslate     `json:"should_force_translate"`
	OwnerReference           *OwnerReference           `json:"owner_reference"`
	Categories               *Categories               `json:"categories"`
	Name                     *Name                     `json:"name"`
	Spec                     *Spec                     `json:"spec"`
	APIVersion               *APIVersion               `json:"api_version"`
	Metadata                 *Metadata                 `json:"metadata"`
	TargetDetails            *TargetDetails            `json:"targetDetails"`
	Machine                  *Machine                  `json:"machine"`
	ProviderOperationPayload *ProviderOperationPayload `json:"provider_operation_payload"`
	Attrs                    *Attrs                    `json:"attrs"`
	Port                     *Port                     `json:"port"`
	Runid                    *Runid                    `json:"runid"`
}
type Metadata struct {
	Description string      `json:"description"`
	Required    []string    `json:"required"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Properties  *Properties `json:"properties"`
}
