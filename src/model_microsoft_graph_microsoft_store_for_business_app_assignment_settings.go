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
// MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings struct for MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings
type MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings struct {
	// Whether or not to use device execution context for Microsoft Store for Business mobile app.
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`

}

// GetUseDeviceContext returns the UseDeviceContext field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings) GetUseDeviceContext() bool {
	if o == nil || o.UseDeviceContext == nil {
		var ret bool
		return ret
	}
	return *o.UseDeviceContext
}

// GetUseDeviceContextOk returns a tuple with the UseDeviceContext field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings) GetUseDeviceContextOk() (bool, bool) {
	if o == nil || o.UseDeviceContext == nil {
		var ret bool
		return ret, false
	}
	return *o.UseDeviceContext, true
}

// HasUseDeviceContext returns a boolean if a field has been set.
func (o *MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings) HasUseDeviceContext() bool {
	if o != nil && o.UseDeviceContext != nil {
		return true
	}

	return false
}

// SetUseDeviceContext gets a reference to the given bool and assigns it to the UseDeviceContext field.
func (o *MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings) SetUseDeviceContext(v bool) {
	o.UseDeviceContext = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphMicrosoftStoreForBusinessAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UseDeviceContext != nil {
		toSerialize["useDeviceContext"] = o.UseDeviceContext
	}
	return json.Marshal(toSerialize)
}

