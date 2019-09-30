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
// BaseItem struct for BaseItem
type BaseItem struct {
	CreatedBy *AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	isExplicitNullCreatedBy bool `json:"-"`
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`

	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	ETag *string `json:"eTag,omitempty"`
	isExplicitNullETag bool `json:"-"`
	LastModifiedBy *AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	isExplicitNullLastModifiedBy bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	ParentReference *AnyOfmicrosoftGraphItemReference `json:"parentReference,omitempty"`
	isExplicitNullParentReference bool `json:"-"`
	WebUrl *string `json:"webUrl,omitempty"`
	isExplicitNullWebUrl bool `json:"-"`
	CreatedByUser *AnyOfmicrosoftGraphUser `json:"createdByUser,omitempty"`
	isExplicitNullCreatedByUser bool `json:"-"`
	LastModifiedByUser *AnyOfmicrosoftGraphUser `json:"lastModifiedByUser,omitempty"`
	isExplicitNullLastModifiedByUser bool `json:"-"`
}

// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *BaseItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BaseItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *BaseItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *BaseItem) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *BaseItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *BaseItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *BaseItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *BaseItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseItem) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *BaseItem) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetETag returns the ETag field if non-nil, zero value otherwise.
func (o *BaseItem) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetETagOk() (string, bool) {
	if o == nil || o.ETag == nil {
		var ret string
		return ret, false
	}
	return *o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *BaseItem) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *BaseItem) SetETag(v string) {
	o.ETag = &v
}

// SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ETag value is set to nil even if false is passed
func (o *BaseItem) SetETagExplicitNull(b bool) {
	o.ETag = nil
	o.isExplicitNullETag = b
}
// GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.
func (o *BaseItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *BaseItem) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *BaseItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedBy value is set to nil even if false is passed
func (o *BaseItem) SetLastModifiedByExplicitNull(b bool) {
	o.LastModifiedBy = nil
	o.isExplicitNullLastModifiedBy = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *BaseItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *BaseItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *BaseItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *BaseItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseItem) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *BaseItem) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetParentReference returns the ParentReference field if non-nil, zero value otherwise.
func (o *BaseItem) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret, false
	}
	return *o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *BaseItem) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *BaseItem) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = &v
}

// SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentReference value is set to nil even if false is passed
func (o *BaseItem) SetParentReferenceExplicitNull(b bool) {
	o.ParentReference = nil
	o.isExplicitNullParentReference = b
}
// GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.
func (o *BaseItem) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetWebUrlOk() (string, bool) {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret, false
	}
	return *o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *BaseItem) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *BaseItem) SetWebUrl(v string) {
	o.WebUrl = &v
}

// SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The WebUrl value is set to nil even if false is passed
func (o *BaseItem) SetWebUrlExplicitNull(b bool) {
	o.WebUrl = nil
	o.isExplicitNullWebUrl = b
}
// GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.
func (o *BaseItem) GetCreatedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *BaseItem) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.
func (o *BaseItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser) {
	o.CreatedByUser = &v
}

// SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedByUser value is set to nil even if false is passed
func (o *BaseItem) SetCreatedByUserExplicitNull(b bool) {
	o.CreatedByUser = nil
	o.isExplicitNullCreatedByUser = b
}
// GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.
func (o *BaseItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *BaseItem) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.
func (o *BaseItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser) {
	o.LastModifiedByUser = &v
}

// SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedByUser value is set to nil even if false is passed
func (o *BaseItem) SetLastModifiedByUserExplicitNull(b bool) {
	o.LastModifiedByUser = nil
	o.isExplicitNullLastModifiedByUser = b
}

// MarshalJSON returns the JSON representation of the model.
func (o BaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy == nil {
		if o.isExplicitNullCreatedBy {
			toSerialize["createdBy"] = o.CreatedBy
		}
	} else {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.ETag == nil {
		if o.isExplicitNullETag {
			toSerialize["eTag"] = o.ETag
		}
	} else {
		toSerialize["eTag"] = o.ETag
	}
	if o.LastModifiedBy == nil {
		if o.isExplicitNullLastModifiedBy {
			toSerialize["lastModifiedBy"] = o.LastModifiedBy
		}
	} else {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.ParentReference == nil {
		if o.isExplicitNullParentReference {
			toSerialize["parentReference"] = o.ParentReference
		}
	} else {
		toSerialize["parentReference"] = o.ParentReference
	}
	if o.WebUrl == nil {
		if o.isExplicitNullWebUrl {
			toSerialize["webUrl"] = o.WebUrl
		}
	} else {
		toSerialize["webUrl"] = o.WebUrl
	}
	if o.CreatedByUser == nil {
		if o.isExplicitNullCreatedByUser {
			toSerialize["createdByUser"] = o.CreatedByUser
		}
	} else {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.LastModifiedByUser == nil {
		if o.isExplicitNullLastModifiedByUser {
			toSerialize["lastModifiedByUser"] = o.LastModifiedByUser
		}
	} else {
		toSerialize["lastModifiedByUser"] = o.LastModifiedByUser
	}
	return json.Marshal(toSerialize)
}

