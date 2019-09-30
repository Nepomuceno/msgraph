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
// MicrosoftGraphInferenceClassification struct for MicrosoftGraphInferenceClassification
type MicrosoftGraphInferenceClassification struct {
	Id *string `json:"id,omitempty"`

	Overrides *[]MicrosoftGraphInferenceClassificationOverride `json:"overrides,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphInferenceClassification) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphInferenceClassification) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphInferenceClassification) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphInferenceClassification) SetId(v string) {
	o.Id = &v
}

// GetOverrides returns the Overrides field if non-nil, zero value otherwise.
func (o *MicrosoftGraphInferenceClassification) GetOverrides() []MicrosoftGraphInferenceClassificationOverride {
	if o == nil || o.Overrides == nil {
		var ret []MicrosoftGraphInferenceClassificationOverride
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphInferenceClassification) GetOverridesOk() ([]MicrosoftGraphInferenceClassificationOverride, bool) {
	if o == nil || o.Overrides == nil {
		var ret []MicrosoftGraphInferenceClassificationOverride
		return ret, false
	}
	return *o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *MicrosoftGraphInferenceClassification) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []MicrosoftGraphInferenceClassificationOverride and assigns it to the Overrides field.
func (o *MicrosoftGraphInferenceClassification) SetOverrides(v []MicrosoftGraphInferenceClassificationOverride) {
	o.Overrides = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphInferenceClassification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}


