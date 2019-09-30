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
// MicrosoftGraphOnenoteSection struct for MicrosoftGraphOnenoteSection
type MicrosoftGraphOnenoteSection struct {
	Id *string `json:"id,omitempty"`

	Self *string `json:"self,omitempty"`
	isExplicitNullSelf bool `json:"-"`
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	CreatedBy *AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	isExplicitNullCreatedBy bool `json:"-"`
	LastModifiedBy *AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	isExplicitNullLastModifiedBy bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	IsDefault *bool `json:"isDefault,omitempty"`
	isExplicitNullIsDefault bool `json:"-"`
	Links *AnyOfmicrosoftGraphSectionLinks `json:"links,omitempty"`
	isExplicitNullLinks bool `json:"-"`
	PagesUrl *string `json:"pagesUrl,omitempty"`
	isExplicitNullPagesUrl bool `json:"-"`
	ParentNotebook *AnyOfmicrosoftGraphNotebook `json:"parentNotebook,omitempty"`
	isExplicitNullParentNotebook bool `json:"-"`
	ParentSectionGroup *AnyOfmicrosoftGraphSectionGroup `json:"parentSectionGroup,omitempty"`
	isExplicitNullParentSectionGroup bool `json:"-"`
	Pages *[]MicrosoftGraphOnenotePage `json:"pages,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOnenoteSection) SetId(v string) {
	o.Id = &v
}

// GetSelf returns the Self field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetSelfOk() (string, bool) {
	if o == nil || o.Self == nil {
		var ret string
		return ret, false
	}
	return *o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *MicrosoftGraphOnenoteSection) SetSelf(v string) {
	o.Self = &v
}

// SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Self value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetSelfExplicitNull(b bool) {
	o.Self = nil
	o.isExplicitNullSelf = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOnenoteSection) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphOnenoteSection) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphOnenoteSection) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphOnenoteSection) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedBy value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetLastModifiedByExplicitNull(b bool) {
	o.LastModifiedBy = nil
	o.isExplicitNullLastModifiedBy = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphOnenoteSection) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetIsDefaultOk() (bool, bool) {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret, false
	}
	return *o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *MicrosoftGraphOnenoteSection) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IsDefault value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetIsDefaultExplicitNull(b bool) {
	o.IsDefault = nil
	o.isExplicitNullIsDefault = b
}
// GetLinks returns the Links field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetLinks() AnyOfmicrosoftGraphSectionLinks {
	if o == nil || o.Links == nil {
		var ret AnyOfmicrosoftGraphSectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetLinksOk() (AnyOfmicrosoftGraphSectionLinks, bool) {
	if o == nil || o.Links == nil {
		var ret AnyOfmicrosoftGraphSectionLinks
		return ret, false
	}
	return *o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AnyOfmicrosoftGraphSectionLinks and assigns it to the Links field.
func (o *MicrosoftGraphOnenoteSection) SetLinks(v AnyOfmicrosoftGraphSectionLinks) {
	o.Links = &v
}

// SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Links value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetLinksExplicitNull(b bool) {
	o.Links = nil
	o.isExplicitNullLinks = b
}
// GetPagesUrl returns the PagesUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetPagesUrl() string {
	if o == nil || o.PagesUrl == nil {
		var ret string
		return ret
	}
	return *o.PagesUrl
}

// GetPagesUrlOk returns a tuple with the PagesUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetPagesUrlOk() (string, bool) {
	if o == nil || o.PagesUrl == nil {
		var ret string
		return ret, false
	}
	return *o.PagesUrl, true
}

// HasPagesUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasPagesUrl() bool {
	if o != nil && o.PagesUrl != nil {
		return true
	}

	return false
}

// SetPagesUrl gets a reference to the given string and assigns it to the PagesUrl field.
func (o *MicrosoftGraphOnenoteSection) SetPagesUrl(v string) {
	o.PagesUrl = &v
}

// SetPagesUrlExplicitNull (un)sets PagesUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PagesUrl value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetPagesUrlExplicitNull(b bool) {
	o.PagesUrl = nil
	o.isExplicitNullPagesUrl = b
}
// GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetParentNotebook() AnyOfmicrosoftGraphNotebook {
	if o == nil || o.ParentNotebook == nil {
		var ret AnyOfmicrosoftGraphNotebook
		return ret
	}
	return *o.ParentNotebook
}

// GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool) {
	if o == nil || o.ParentNotebook == nil {
		var ret AnyOfmicrosoftGraphNotebook
		return ret, false
	}
	return *o.ParentNotebook, true
}

// HasParentNotebook returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasParentNotebook() bool {
	if o != nil && o.ParentNotebook != nil {
		return true
	}

	return false
}

// SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.
func (o *MicrosoftGraphOnenoteSection) SetParentNotebook(v AnyOfmicrosoftGraphNotebook) {
	o.ParentNotebook = &v
}

// SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentNotebook value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetParentNotebookExplicitNull(b bool) {
	o.ParentNotebook = nil
	o.isExplicitNullParentNotebook = b
}
// GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup {
	if o == nil || o.ParentSectionGroup == nil {
		var ret AnyOfmicrosoftGraphSectionGroup
		return ret
	}
	return *o.ParentSectionGroup
}

// GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetParentSectionGroupOk() (AnyOfmicrosoftGraphSectionGroup, bool) {
	if o == nil || o.ParentSectionGroup == nil {
		var ret AnyOfmicrosoftGraphSectionGroup
		return ret, false
	}
	return *o.ParentSectionGroup, true
}

// HasParentSectionGroup returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasParentSectionGroup() bool {
	if o != nil && o.ParentSectionGroup != nil {
		return true
	}

	return false
}

// SetParentSectionGroup gets a reference to the given AnyOfmicrosoftGraphSectionGroup and assigns it to the ParentSectionGroup field.
func (o *MicrosoftGraphOnenoteSection) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup) {
	o.ParentSectionGroup = &v
}

// SetParentSectionGroupExplicitNull (un)sets ParentSectionGroup to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentSectionGroup value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteSection) SetParentSectionGroupExplicitNull(b bool) {
	o.ParentSectionGroup = nil
	o.isExplicitNullParentSectionGroup = b
}
// GetPages returns the Pages field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteSection) GetPages() []MicrosoftGraphOnenotePage {
	if o == nil || o.Pages == nil {
		var ret []MicrosoftGraphOnenotePage
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteSection) GetPagesOk() ([]MicrosoftGraphOnenotePage, bool) {
	if o == nil || o.Pages == nil {
		var ret []MicrosoftGraphOnenotePage
		return ret, false
	}
	return *o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteSection) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given []MicrosoftGraphOnenotePage and assigns it to the Pages field.
func (o *MicrosoftGraphOnenoteSection) SetPages(v []MicrosoftGraphOnenotePage) {
	o.Pages = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOnenoteSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Self == nil {
		if o.isExplicitNullSelf {
			toSerialize["self"] = o.Self
		}
	} else {
		toSerialize["self"] = o.Self
	}
	if o.CreatedDateTime == nil {
		if o.isExplicitNullCreatedDateTime {
			toSerialize["createdDateTime"] = o.CreatedDateTime
		}
	} else {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.CreatedBy == nil {
		if o.isExplicitNullCreatedBy {
			toSerialize["createdBy"] = o.CreatedBy
		}
	} else {
		toSerialize["createdBy"] = o.CreatedBy
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
	if o.IsDefault == nil {
		if o.isExplicitNullIsDefault {
			toSerialize["isDefault"] = o.IsDefault
		}
	} else {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.Links == nil {
		if o.isExplicitNullLinks {
			toSerialize["links"] = o.Links
		}
	} else {
		toSerialize["links"] = o.Links
	}
	if o.PagesUrl == nil {
		if o.isExplicitNullPagesUrl {
			toSerialize["pagesUrl"] = o.PagesUrl
		}
	} else {
		toSerialize["pagesUrl"] = o.PagesUrl
	}
	if o.ParentNotebook == nil {
		if o.isExplicitNullParentNotebook {
			toSerialize["parentNotebook"] = o.ParentNotebook
		}
	} else {
		toSerialize["parentNotebook"] = o.ParentNotebook
	}
	if o.ParentSectionGroup == nil {
		if o.isExplicitNullParentSectionGroup {
			toSerialize["parentSectionGroup"] = o.ParentSectionGroup
		}
	} else {
		toSerialize["parentSectionGroup"] = o.ParentSectionGroup
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	return json.Marshal(toSerialize)
}


