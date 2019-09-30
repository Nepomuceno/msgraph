/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
import (
	"encoding/json"
)
// MicrosoftGraphEdgeSearchEngine struct for MicrosoftGraphEdgeSearchEngine
type MicrosoftGraphEdgeSearchEngine struct {
	// Allows IT admins to set a predefined default search engine for MDM-Controlled devices.
	EdgeSearchEngineType *AnyOfmicrosoftGraphEdgeSearchEngineType `json:"edgeSearchEngineType,omitempty"`

}

// GetEdgeSearchEngineType returns the EdgeSearchEngineType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEdgeSearchEngine) GetEdgeSearchEngineType() AnyOfmicrosoftGraphEdgeSearchEngineType {
	if o == nil || o.EdgeSearchEngineType == nil {
		var ret AnyOfmicrosoftGraphEdgeSearchEngineType
		return ret
	}
	return *o.EdgeSearchEngineType
}

// GetEdgeSearchEngineTypeOk returns a tuple with the EdgeSearchEngineType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEdgeSearchEngine) GetEdgeSearchEngineTypeOk() (AnyOfmicrosoftGraphEdgeSearchEngineType, bool) {
	if o == nil || o.EdgeSearchEngineType == nil {
		var ret AnyOfmicrosoftGraphEdgeSearchEngineType
		return ret, false
	}
	return *o.EdgeSearchEngineType, true
}

// HasEdgeSearchEngineType returns a boolean if a field has been set.
func (o *MicrosoftGraphEdgeSearchEngine) HasEdgeSearchEngineType() bool {
	if o != nil && o.EdgeSearchEngineType != nil {
		return true
	}

	return false
}

// SetEdgeSearchEngineType gets a reference to the given AnyOfmicrosoftGraphEdgeSearchEngineType and assigns it to the EdgeSearchEngineType field.
func (o *MicrosoftGraphEdgeSearchEngine) SetEdgeSearchEngineType(v AnyOfmicrosoftGraphEdgeSearchEngineType) {
	o.EdgeSearchEngineType = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphEdgeSearchEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EdgeSearchEngineType != nil {
		toSerialize["edgeSearchEngineType"] = o.EdgeSearchEngineType
	}
	return json.Marshal(toSerialize)
}


