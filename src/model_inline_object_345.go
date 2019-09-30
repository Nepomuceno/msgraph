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
// InlineObject345 struct for InlineObject345
type InlineObject345 struct {
	Id *string `json:"id,omitempty"`
	isExplicitNullId bool `json:"-"`
	GroupId *string `json:"groupId,omitempty"`
	isExplicitNullGroupId bool `json:"-"`
	RenameAs *string `json:"renameAs,omitempty"`
	isExplicitNullRenameAs bool `json:"-"`
	SiteCollectionId *string `json:"siteCollectionId,omitempty"`
	isExplicitNullSiteCollectionId bool `json:"-"`
	SiteId *string `json:"siteId,omitempty"`
	isExplicitNullSiteId bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *InlineObject345) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject345) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject345) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject345) SetId(v string) {
	o.Id = &v
}

// SetIdExplicitNull (un)sets Id to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Id value is set to nil even if false is passed
func (o *InlineObject345) SetIdExplicitNull(b bool) {
	o.Id = nil
	o.isExplicitNullId = b
}
// GetGroupId returns the GroupId field if non-nil, zero value otherwise.
func (o *InlineObject345) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject345) GetGroupIdOk() (string, bool) {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret, false
	}
	return *o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineObject345) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *InlineObject345) SetGroupId(v string) {
	o.GroupId = &v
}

// SetGroupIdExplicitNull (un)sets GroupId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The GroupId value is set to nil even if false is passed
func (o *InlineObject345) SetGroupIdExplicitNull(b bool) {
	o.GroupId = nil
	o.isExplicitNullGroupId = b
}
// GetRenameAs returns the RenameAs field if non-nil, zero value otherwise.
func (o *InlineObject345) GetRenameAs() string {
	if o == nil || o.RenameAs == nil {
		var ret string
		return ret
	}
	return *o.RenameAs
}

// GetRenameAsOk returns a tuple with the RenameAs field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject345) GetRenameAsOk() (string, bool) {
	if o == nil || o.RenameAs == nil {
		var ret string
		return ret, false
	}
	return *o.RenameAs, true
}

// HasRenameAs returns a boolean if a field has been set.
func (o *InlineObject345) HasRenameAs() bool {
	if o != nil && o.RenameAs != nil {
		return true
	}

	return false
}

// SetRenameAs gets a reference to the given string and assigns it to the RenameAs field.
func (o *InlineObject345) SetRenameAs(v string) {
	o.RenameAs = &v
}

// SetRenameAsExplicitNull (un)sets RenameAs to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The RenameAs value is set to nil even if false is passed
func (o *InlineObject345) SetRenameAsExplicitNull(b bool) {
	o.RenameAs = nil
	o.isExplicitNullRenameAs = b
}
// GetSiteCollectionId returns the SiteCollectionId field if non-nil, zero value otherwise.
func (o *InlineObject345) GetSiteCollectionId() string {
	if o == nil || o.SiteCollectionId == nil {
		var ret string
		return ret
	}
	return *o.SiteCollectionId
}

// GetSiteCollectionIdOk returns a tuple with the SiteCollectionId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject345) GetSiteCollectionIdOk() (string, bool) {
	if o == nil || o.SiteCollectionId == nil {
		var ret string
		return ret, false
	}
	return *o.SiteCollectionId, true
}

// HasSiteCollectionId returns a boolean if a field has been set.
func (o *InlineObject345) HasSiteCollectionId() bool {
	if o != nil && o.SiteCollectionId != nil {
		return true
	}

	return false
}

// SetSiteCollectionId gets a reference to the given string and assigns it to the SiteCollectionId field.
func (o *InlineObject345) SetSiteCollectionId(v string) {
	o.SiteCollectionId = &v
}

// SetSiteCollectionIdExplicitNull (un)sets SiteCollectionId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SiteCollectionId value is set to nil even if false is passed
func (o *InlineObject345) SetSiteCollectionIdExplicitNull(b bool) {
	o.SiteCollectionId = nil
	o.isExplicitNullSiteCollectionId = b
}
// GetSiteId returns the SiteId field if non-nil, zero value otherwise.
func (o *InlineObject345) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject345) GetSiteIdOk() (string, bool) {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret, false
	}
	return *o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *InlineObject345) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *InlineObject345) SetSiteId(v string) {
	o.SiteId = &v
}

// SetSiteIdExplicitNull (un)sets SiteId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SiteId value is set to nil even if false is passed
func (o *InlineObject345) SetSiteIdExplicitNull(b bool) {
	o.SiteId = nil
	o.isExplicitNullSiteId = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject345) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id == nil {
		if o.isExplicitNullId {
			toSerialize["id"] = o.Id
		}
	} else {
		toSerialize["id"] = o.Id
	}
	if o.GroupId == nil {
		if o.isExplicitNullGroupId {
			toSerialize["groupId"] = o.GroupId
		}
	} else {
		toSerialize["groupId"] = o.GroupId
	}
	if o.RenameAs == nil {
		if o.isExplicitNullRenameAs {
			toSerialize["renameAs"] = o.RenameAs
		}
	} else {
		toSerialize["renameAs"] = o.RenameAs
	}
	if o.SiteCollectionId == nil {
		if o.isExplicitNullSiteCollectionId {
			toSerialize["siteCollectionId"] = o.SiteCollectionId
		}
	} else {
		toSerialize["siteCollectionId"] = o.SiteCollectionId
	}
	if o.SiteId == nil {
		if o.isExplicitNullSiteId {
			toSerialize["siteId"] = o.SiteId
		}
	} else {
		toSerialize["siteId"] = o.SiteId
	}
	return json.Marshal(toSerialize)
}


