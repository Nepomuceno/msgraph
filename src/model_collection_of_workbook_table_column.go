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
// CollectionOfWorkbookTableColumn struct for CollectionOfWorkbookTableColumn
type CollectionOfWorkbookTableColumn struct {
	Value *[]MicrosoftGraphWorkbookTableColumn `json:"value,omitempty"`

}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *CollectionOfWorkbookTableColumn) GetValue() []MicrosoftGraphWorkbookTableColumn {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphWorkbookTableColumn
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfWorkbookTableColumn) GetValueOk() ([]MicrosoftGraphWorkbookTableColumn, bool) {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphWorkbookTableColumn
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfWorkbookTableColumn) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphWorkbookTableColumn and assigns it to the Value field.
func (o *CollectionOfWorkbookTableColumn) SetValue(v []MicrosoftGraphWorkbookTableColumn) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o CollectionOfWorkbookTableColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

