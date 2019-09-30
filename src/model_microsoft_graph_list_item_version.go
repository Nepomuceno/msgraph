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
// MicrosoftGraphListItemVersion struct for MicrosoftGraphListItemVersion
type MicrosoftGraphListItemVersion struct {
	Id *string `json:"id,omitempty"`

	LastModifiedBy *AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	isExplicitNullLastModifiedBy bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	Publication *AnyOfmicrosoftGraphPublicationFacet `json:"publication,omitempty"`
	isExplicitNullPublication bool `json:"-"`
	Fields *AnyOfmicrosoftGraphFieldValueSet `json:"fields,omitempty"`
	isExplicitNullFields bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphListItemVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphListItemVersion) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphListItemVersion) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphListItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphListItemVersion) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphListItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedBy value is set to nil even if false is passed
func (o *MicrosoftGraphListItemVersion) SetLastModifiedByExplicitNull(b bool) {
	o.LastModifiedBy = nil
	o.isExplicitNullLastModifiedBy = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetPublication returns the Publication field if non-nil, zero value otherwise.
func (o *MicrosoftGraphListItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet {
	if o == nil || o.Publication == nil {
		var ret AnyOfmicrosoftGraphPublicationFacet
		return ret
	}
	return *o.Publication
}

// GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphListItemVersion) GetPublicationOk() (AnyOfmicrosoftGraphPublicationFacet, bool) {
	if o == nil || o.Publication == nil {
		var ret AnyOfmicrosoftGraphPublicationFacet
		return ret, false
	}
	return *o.Publication, true
}

// HasPublication returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasPublication() bool {
	if o != nil && o.Publication != nil {
		return true
	}

	return false
}

// SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.
func (o *MicrosoftGraphListItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet) {
	o.Publication = &v
}

// SetPublicationExplicitNull (un)sets Publication to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Publication value is set to nil even if false is passed
func (o *MicrosoftGraphListItemVersion) SetPublicationExplicitNull(b bool) {
	o.Publication = nil
	o.isExplicitNullPublication = b
}
// GetFields returns the Fields field if non-nil, zero value otherwise.
func (o *MicrosoftGraphListItemVersion) GetFields() AnyOfmicrosoftGraphFieldValueSet {
	if o == nil || o.Fields == nil {
		var ret AnyOfmicrosoftGraphFieldValueSet
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphListItemVersion) GetFieldsOk() (AnyOfmicrosoftGraphFieldValueSet, bool) {
	if o == nil || o.Fields == nil {
		var ret AnyOfmicrosoftGraphFieldValueSet
		return ret, false
	}
	return *o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MicrosoftGraphListItemVersion) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given AnyOfmicrosoftGraphFieldValueSet and assigns it to the Fields field.
func (o *MicrosoftGraphListItemVersion) SetFields(v AnyOfmicrosoftGraphFieldValueSet) {
	o.Fields = &v
}

// SetFieldsExplicitNull (un)sets Fields to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Fields value is set to nil even if false is passed
func (o *MicrosoftGraphListItemVersion) SetFieldsExplicitNull(b bool) {
	o.Fields = nil
	o.isExplicitNullFields = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphListItemVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedBy == nil {
		if o.isExplicitNullLastModifiedBy {
			toSerialize["lastModifiedBy"] = o.LastModifiedBy
		}
	} else {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime == nil {
		if o.isExplicitNullLastModifiedDateTime {
			toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
		}
	} else {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Publication == nil {
		if o.isExplicitNullPublication {
			toSerialize["publication"] = o.Publication
		}
	} else {
		toSerialize["publication"] = o.Publication
	}
	if o.Fields == nil {
		if o.isExplicitNullFields {
			toSerialize["fields"] = o.Fields
		}
	} else {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

