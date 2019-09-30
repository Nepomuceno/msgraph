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
// ItemAttachment struct for ItemAttachment
type ItemAttachment struct {
	Item *AnyOfmicrosoftGraphOutlookItem `json:"item,omitempty"`
	isExplicitNullItem bool `json:"-"`
}

// GetItem returns the Item field if non-nil, zero value otherwise.
func (o *ItemAttachment) GetItem() AnyOfmicrosoftGraphOutlookItem {
	if o == nil || o.Item == nil {
		var ret AnyOfmicrosoftGraphOutlookItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemAttachment) GetItemOk() (AnyOfmicrosoftGraphOutlookItem, bool) {
	if o == nil || o.Item == nil {
		var ret AnyOfmicrosoftGraphOutlookItem
		return ret, false
	}
	return *o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *ItemAttachment) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given AnyOfmicrosoftGraphOutlookItem and assigns it to the Item field.
func (o *ItemAttachment) SetItem(v AnyOfmicrosoftGraphOutlookItem) {
	o.Item = &v
}

// SetItemExplicitNull (un)sets Item to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Item value is set to nil even if false is passed
func (o *ItemAttachment) SetItemExplicitNull(b bool) {
	o.Item = nil
	o.isExplicitNullItem = b
}

// MarshalJSON returns the JSON representation of the model.
func (o ItemAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Item == nil {
		if o.isExplicitNullItem {
			toSerialize["item"] = o.Item
		}
	} else {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}


