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
// OutlookItem struct for OutlookItem
type OutlookItem struct {
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	ChangeKey *string `json:"changeKey,omitempty"`
	isExplicitNullChangeKey bool `json:"-"`
	Categories *[]string `json:"categories,omitempty"`

}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *OutlookItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutlookItem) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *OutlookItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *OutlookItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *OutlookItem) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *OutlookItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutlookItem) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *OutlookItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *OutlookItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *OutlookItem) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.
func (o *OutlookItem) GetChangeKey() string {
	if o == nil || o.ChangeKey == nil {
		var ret string
		return ret
	}
	return *o.ChangeKey
}

// GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutlookItem) GetChangeKeyOk() (string, bool) {
	if o == nil || o.ChangeKey == nil {
		var ret string
		return ret, false
	}
	return *o.ChangeKey, true
}

// HasChangeKey returns a boolean if a field has been set.
func (o *OutlookItem) HasChangeKey() bool {
	if o != nil && o.ChangeKey != nil {
		return true
	}

	return false
}

// SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.
func (o *OutlookItem) SetChangeKey(v string) {
	o.ChangeKey = &v
}

// SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ChangeKey value is set to nil even if false is passed
func (o *OutlookItem) SetChangeKeyExplicitNull(b bool) {
	o.ChangeKey = nil
	o.isExplicitNullChangeKey = b
}
// GetCategories returns the Categories field if non-nil, zero value otherwise.
func (o *OutlookItem) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutlookItem) GetCategoriesOk() ([]string, bool) {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret, false
	}
	return *o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *OutlookItem) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *OutlookItem) SetCategories(v []string) {
	o.Categories = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o OutlookItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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


