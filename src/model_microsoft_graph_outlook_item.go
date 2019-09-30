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
	"time"
	"encoding/json"
)
// MicrosoftGraphOutlookItem struct for MicrosoftGraphOutlookItem
type MicrosoftGraphOutlookItem struct {
	Id *string `json:"id,omitempty"`

	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	ChangeKey *string `json:"changeKey,omitempty"`
	isExplicitNullChangeKey bool `json:"-"`
	Categories *[]string `json:"categories,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookItem) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOutlookItem) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookItem) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOutlookItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookItem) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookItem) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphOutlookItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookItem) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookItem) GetChangeKey() string {
	if o == nil || o.ChangeKey == nil {
		var ret string
		return ret
	}
	return *o.ChangeKey
}

// GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookItem) GetChangeKeyOk() (string, bool) {
	if o == nil || o.ChangeKey == nil {
		var ret string
		return ret, false
	}
	return *o.ChangeKey, true
}

// HasChangeKey returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookItem) HasChangeKey() bool {
	if o != nil && o.ChangeKey != nil {
		return true
	}

	return false
}

// SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.
func (o *MicrosoftGraphOutlookItem) SetChangeKey(v string) {
	o.ChangeKey = &v
}

// SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ChangeKey value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookItem) SetChangeKeyExplicitNull(b bool) {
	o.ChangeKey = nil
	o.isExplicitNullChangeKey = b
}
// GetCategories returns the Categories field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookItem) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookItem) GetCategoriesOk() ([]string, bool) {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret, false
	}
	return *o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookItem) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *MicrosoftGraphOutlookItem) SetCategories(v []string) {
	o.Categories = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOutlookItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime == nil {
		if o.isExplicitNullCreatedDateTime {
			toSerialize["createdDateTime"] = o.CreatedDateTime
		}
	} else {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.LastModifiedDateTime == nil {
		if o.isExplicitNullLastModifiedDateTime {
			toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
		}
	} else {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.ChangeKey == nil {
		if o.isExplicitNullChangeKey {
			toSerialize["changeKey"] = o.ChangeKey
		}
	} else {
		toSerialize["changeKey"] = o.ChangeKey
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

