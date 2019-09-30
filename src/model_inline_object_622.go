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
// InlineObject622 struct for InlineObject622
type InlineObject622 struct {
	QuickScan *bool `json:"quickScan,omitempty"`

}

// GetQuickScan returns the QuickScan field if non-nil, zero value otherwise.
func (o *InlineObject622) GetQuickScan() bool {
	if o == nil || o.QuickScan == nil {
		var ret bool
		return ret
	}
	return *o.QuickScan
}

// GetQuickScanOk returns a tuple with the QuickScan field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject622) GetQuickScanOk() (bool, bool) {
	if o == nil || o.QuickScan == nil {
		var ret bool
		return ret, false
	}
	return *o.QuickScan, true
}

// HasQuickScan returns a boolean if a field has been set.
func (o *InlineObject622) HasQuickScan() bool {
	if o != nil && o.QuickScan != nil {
		return true
	}

	return false
}

// SetQuickScan gets a reference to the given bool and assigns it to the QuickScan field.
func (o *InlineObject622) SetQuickScan(v bool) {
	o.QuickScan = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject622) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QuickScan != nil {
		toSerialize["quickScan"] = o.QuickScan
	}
	return json.Marshal(toSerialize)
}

