package calm

// BlueprintUploadRequest is the request to upload a blueprint
type BlueprintUploadRequest struct {
	Type string `json:"type"`
}

// BlueprintExportRequest is the export request which only needs a UUID
type BlueprintExportRequest struct{}

// BlueprintExportResponse is the output from the export request
type BlueprintExportResponse struct {
}

// BlueprintRunRequest is the request used run a blueprint
type BlueprintRunRequest struct {
	Description string   `json:"description"`
	Required    []string `json:"required"`
	Type        string   `json:"type"`
	Properties  struct {
		Spec struct {
			Description string `json:"description"`
			Type        string `json:"type"`
			Title       string `json:"title"`
			Properties  struct {
				TargetDetails struct {
					Type                 string `json:"type"`
					Description          string `json:"description"`
					AdditionalProperties bool   `json:"additionalProperties"`
				} `json:"targetDetails"`
				Machine struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"machine"`
				ProviderOperationPayload struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"provider_operation_payload"`
				Attrs struct {
					Type                 string `json:"type"`
					Description          string `json:"description"`
					AdditionalProperties bool   `json:"additionalProperties"`
				} `json:"attrs"`
				Port struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"port"`
				Runid struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"runid"`
			} `json:"properties"`
		} `json:"spec"`
		APIVersion struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"api_version"`
		Metadata struct {
			Description string   `json:"description"`
			Required    []string `json:"required"`
			Type        string   `json:"type"`
			Title       string   `json:"title"`
			Properties  struct {
				LastUpdateTime struct {
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"last_update_time"`
				UseCategoriesMapping struct {
					Default     bool   `json:"default"`
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"use_categories_mapping"`
				Kind struct {
					Default     string   `json:"default"`
					ReadOnly    bool     `json:"readOnly"`
					Type        string   `json:"type"`
					Description string   `json:"description"`
					XNtnxEnum   []string `json:"x-ntnx-enum"`
				} `json:"kind"`
				UUID struct {
					Pattern     string `json:"pattern"`
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"uuid"`
				ProjectReference struct {
					Description string   `json:"description"`
					Required    []string `json:"required"`
					Type        string   `json:"type"`
					Title       string   `json:"title"`
					Properties  struct {
						Kind struct {
							Default     string   `json:"default"`
							ReadOnly    bool     `json:"readOnly"`
							Type        string   `json:"type"`
							Description string   `json:"description"`
							XNtnxEnum   []string `json:"x-ntnx-enum"`
						} `json:"kind"`
						Name struct {
							ReadOnly bool   `json:"readOnly"`
							Type     string `json:"type"`
						} `json:"name"`
						UUID struct {
							Pattern string `json:"pattern"`
							Type    string `json:"type"`
							Format  string `json:"format"`
						} `json:"uuid"`
					} `json:"properties"`
				} `json:"project_reference"`
				CreationTime struct {
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"creation_time"`
				SpecVersion struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"spec_version"`
				SpecHash struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"spec_hash"`
				CategoriesMapping struct {
					Type                 string `json:"type"`
					Description          string `json:"description"`
					AdditionalProperties struct {
						Type  string `json:"type"`
						Items struct {
							Type string `json:"type"`
						} `json:"items"`
					} `json:"additionalProperties"`
				} `json:"categories_mapping"`
				ShouldForceTranslate struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"should_force_translate"`
				OwnerReference struct {
					Description string   `json:"description"`
					Required    []string `json:"required"`
					Type        string   `json:"type"`
					Title       string   `json:"title"`
					Properties  struct {
						Kind struct {
							Default     string   `json:"default"`
							ReadOnly    bool     `json:"readOnly"`
							Type        string   `json:"type"`
							Description string   `json:"description"`
							XNtnxEnum   []string `json:"x-ntnx-enum"`
						} `json:"kind"`
						Name struct {
							ReadOnly bool   `json:"readOnly"`
							Type     string `json:"type"`
						} `json:"name"`
						UUID struct {
							Pattern string `json:"pattern"`
							Type    string `json:"type"`
							Format  string `json:"format"`
						} `json:"uuid"`
					} `json:"properties"`
				} `json:"owner_reference"`
				Categories struct {
					Type                 string `json:"type"`
					Description          string `json:"description"`
					AdditionalProperties struct {
						Type string `json:"type"`
					} `json:"additionalProperties"`
				} `json:"categories"`
				Name struct {
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
				} `json:"name"`
			} `json:"properties"`
		} `json:"metadata"`
	} `json:"properties"`
	Title string `json:"title"`
}

// BlueprintRunResponse is the response from calling Run
type BlueprintRunResponse struct {
	Required   []string `json:"required"`
	Type       string   `json:"type"`
	Properties struct {
		Status struct {
			Required    []string `json:"required"`
			Description string   `json:"description"`
			Properties  struct {
				TrlID struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"trl_id"`
				RequestID struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"request_id"`
				RunID struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"run_id"`
			} `json:"properties"`
		} `json:"status"`
		Metadata struct {
			Description string   `json:"description"`
			Required    []string `json:"required"`
			Type        string   `json:"type"`
			Title       string   `json:"title"`
			Properties  struct {
				LastUpdateTime struct {
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"last_update_time"`
				UseCategoriesMapping struct {
					Default     bool   `json:"default"`
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"use_categories_mapping"`
				Kind struct {
					Default     string   `json:"default"`
					ReadOnly    bool     `json:"readOnly"`
					Type        string   `json:"type"`
					Description string   `json:"description"`
					XNtnxEnum   []string `json:"x-ntnx-enum"`
				} `json:"kind"`
				UUID struct {
					Pattern     string `json:"pattern"`
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"uuid"`
				ProjectReference struct {
					Description string   `json:"description"`
					Required    []string `json:"required"`
					Type        string   `json:"type"`
					Title       string   `json:"title"`
					Properties  struct {
						Kind struct {
							Default     string   `json:"default"`
							ReadOnly    bool     `json:"readOnly"`
							Type        string   `json:"type"`
							Description string   `json:"description"`
							XNtnxEnum   []string `json:"x-ntnx-enum"`
						} `json:"kind"`
						Name struct {
							ReadOnly bool   `json:"readOnly"`
							Type     string `json:"type"`
						} `json:"name"`
						UUID struct {
							Pattern string `json:"pattern"`
							Type    string `json:"type"`
							Format  string `json:"format"`
						} `json:"uuid"`
					} `json:"properties"`
				} `json:"project_reference"`
				CreationTime struct {
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"creation_time"`
				SpecVersion struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"spec_version"`
				SpecHash struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"spec_hash"`
				CategoriesMapping struct {
					Type                 string `json:"type"`
					Description          string `json:"description"`
					AdditionalProperties struct {
						Type  string `json:"type"`
						Items struct {
							Type string `json:"type"`
						} `json:"items"`
					} `json:"additionalProperties"`
				} `json:"categories_mapping"`
				ShouldForceTranslate struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"should_force_translate"`
				OwnerReference struct {
					Description string   `json:"description"`
					Required    []string `json:"required"`
					Type        string   `json:"type"`
					Title       string   `json:"title"`
					Properties  struct {
						Kind struct {
							Default     string   `json:"default"`
							ReadOnly    bool     `json:"readOnly"`
							Type        string   `json:"type"`
							Description string   `json:"description"`
							XNtnxEnum   []string `json:"x-ntnx-enum"`
						} `json:"kind"`
						Name struct {
							ReadOnly bool   `json:"readOnly"`
							Type     string `json:"type"`
						} `json:"name"`
						UUID struct {
							Pattern string `json:"pattern"`
							Type    string `json:"type"`
							Format  string `json:"format"`
						} `json:"uuid"`
					} `json:"properties"`
				} `json:"owner_reference"`
				Categories struct {
					Type                 string `json:"type"`
					Description          string `json:"description"`
					AdditionalProperties struct {
						Type string `json:"type"`
					} `json:"additionalProperties"`
				} `json:"categories"`
				Name struct {
					ReadOnly    bool   `json:"readOnly"`
					Type        string `json:"type"`
					Description string `json:"description"`
					MaxLength   int    `json:"maxLength"`
				} `json:"name"`
			} `json:"properties"`
		} `json:"metadata"`
	} `json:"properties"`
}
