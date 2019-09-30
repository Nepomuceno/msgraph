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
// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	Apps *[]AnyOfmicrosoftGraphManagedMobileApp `json:"apps,omitempty"`

}

// GetApps returns the Apps field if non-nil, zero value otherwise.
func (o *InlineObject2) GetApps() []AnyOfmicrosoftGraphManagedMobileApp {
	if o == nil || o.Apps == nil {
		var ret []AnyOfmicrosoftGraphManagedMobileApp
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetAppsOk() ([]AnyOfmicrosoftGraphManagedMobileApp, bool) {
	if o == nil || o.Apps == nil {
		var ret []AnyOfmicrosoftGraphManagedMobileApp
		return ret, false
	}
	return *o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *InlineObject2) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []AnyOfmicrosoftGraphManagedMobileApp and assigns it to the Apps field.
func (o *InlineObject2) SetApps(v []AnyOfmicrosoftGraphManagedMobileApp) {
	o.Apps = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

