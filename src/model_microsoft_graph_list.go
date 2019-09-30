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
// MicrosoftGraphList struct for MicrosoftGraphList
type MicrosoftGraphList struct {
	Id *string `json:"id,omitempty"`

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
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	List *AnyOfmicrosoftGraphListInfo `json:"list,omitempty"`
	isExplicitNullList bool `json:"-"`
	SharepointIds *AnyOfmicrosoftGraphSharepointIds `json:"sharepointIds,omitempty"`
	isExplicitNullSharepointIds bool `json:"-"`
	System *AnyOfobject `json:"system,omitempty"`
	isExplicitNullSystem bool `json:"-"`
	Columns *[]MicrosoftGraphColumnDefinition `json:"columns,omitempty"`

	ContentTypes *[]MicrosoftGraphContentType `json:"contentTypes,omitempty"`

	Drive *AnyOfmicrosoftGraphDrive `json:"drive,omitempty"`
	isExplicitNullDrive bool `json:"-"`
	Items *[]MicrosoftGraphListItem `json:"items,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphList) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphList) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphList) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphList) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetETag returns the ETag field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetETagOk() (string, bool) {
	if o == nil || o.ETag == nil {
		var ret string
		return ret, false
	}
	return *o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *MicrosoftGraphList) SetETag(v string) {
	o.ETag = &v
}

// SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ETag value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetETagExplicitNull(b bool) {
	o.ETag = nil
	o.isExplicitNullETag = b
}
// GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphList) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedBy value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetLastModifiedByExplicitNull(b bool) {
	o.LastModifiedBy = nil
	o.isExplicitNullLastModifiedBy = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphList) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphList) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetParentReference returns the ParentReference field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret, false
	}
	return *o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *MicrosoftGraphList) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = &v
}

// SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentReference value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetParentReferenceExplicitNull(b bool) {
	o.ParentReference = nil
	o.isExplicitNullParentReference = b
}
// GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetWebUrlOk() (string, bool) {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret, false
	}
	return *o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *MicrosoftGraphList) SetWebUrl(v string) {
	o.WebUrl = &v
}

// SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The WebUrl value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetWebUrlExplicitNull(b bool) {
	o.WebUrl = nil
	o.isExplicitNullWebUrl = b
}
// GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetCreatedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.
func (o *MicrosoftGraphList) SetCreatedByUser(v AnyOfmicrosoftGraphUser) {
	o.CreatedByUser = &v
}

// SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedByUser value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetCreatedByUserExplicitNull(b bool) {
	o.CreatedByUser = nil
	o.isExplicitNullCreatedByUser = b
}
// GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetLastModifiedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.
func (o *MicrosoftGraphList) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser) {
	o.LastModifiedByUser = &v
}

// SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedByUser value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetLastModifiedByUserExplicitNull(b bool) {
	o.LastModifiedByUser = nil
	o.isExplicitNullLastModifiedByUser = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphList) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetList returns the List field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetList() AnyOfmicrosoftGraphListInfo {
	if o == nil || o.List == nil {
		var ret AnyOfmicrosoftGraphListInfo
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetListOk() (AnyOfmicrosoftGraphListInfo, bool) {
	if o == nil || o.List == nil {
		var ret AnyOfmicrosoftGraphListInfo
		return ret, false
	}
	return *o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given AnyOfmicrosoftGraphListInfo and assigns it to the List field.
func (o *MicrosoftGraphList) SetList(v AnyOfmicrosoftGraphListInfo) {
	o.List = &v
}

// SetListExplicitNull (un)sets List to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The List value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetListExplicitNull(b bool) {
	o.List = nil
	o.isExplicitNullList = b
}
// GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds {
	if o == nil || o.SharepointIds == nil {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret
	}
	return *o.SharepointIds
}

// GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool) {
	if o == nil || o.SharepointIds == nil {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret, false
	}
	return *o.SharepointIds, true
}

// HasSharepointIds returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasSharepointIds() bool {
	if o != nil && o.SharepointIds != nil {
		return true
	}

	return false
}

// SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.
func (o *MicrosoftGraphList) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds) {
	o.SharepointIds = &v
}

// SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharepointIds value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetSharepointIdsExplicitNull(b bool) {
	o.SharepointIds = nil
	o.isExplicitNullSharepointIds = b
}
// GetSystem returns the System field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetSystem() AnyOfobject {
	if o == nil || o.System == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetSystemOk() (AnyOfobject, bool) {
	if o == nil || o.System == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.
func (o *MicrosoftGraphList) SetSystem(v AnyOfobject) {
	o.System = &v
}

// SetSystemExplicitNull (un)sets System to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The System value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetSystemExplicitNull(b bool) {
	o.System = nil
	o.isExplicitNullSystem = b
}
// GetColumns returns the Columns field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetColumns() []MicrosoftGraphColumnDefinition {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret, false
	}
	return *o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.
func (o *MicrosoftGraphList) SetColumns(v []MicrosoftGraphColumnDefinition) {
	o.Columns = &v
}

// GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetContentTypes() []MicrosoftGraphContentType {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret
	}
	return *o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetContentTypesOk() ([]MicrosoftGraphContentType, bool) {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret, false
	}
	return *o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasContentTypes() bool {
	if o != nil && o.ContentTypes != nil {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.
func (o *MicrosoftGraphList) SetContentTypes(v []MicrosoftGraphContentType) {
	o.ContentTypes = &v
}

// GetDrive returns the Drive field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetDrive() AnyOfmicrosoftGraphDrive {
	if o == nil || o.Drive == nil {
		var ret AnyOfmicrosoftGraphDrive
		return ret
	}
	return *o.Drive
}

// GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool) {
	if o == nil || o.Drive == nil {
		var ret AnyOfmicrosoftGraphDrive
		return ret, false
	}
	return *o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasDrive() bool {
	if o != nil && o.Drive != nil {
		return true
	}

	return false
}

// SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.
func (o *MicrosoftGraphList) SetDrive(v AnyOfmicrosoftGraphDrive) {
	o.Drive = &v
}

// SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Drive value is set to nil even if false is passed
func (o *MicrosoftGraphList) SetDriveExplicitNull(b bool) {
	o.Drive = nil
	o.isExplicitNullDrive = b
}
// GetItems returns the Items field if non-nil, zero value otherwise.
func (o *MicrosoftGraphList) GetItems() []MicrosoftGraphListItem {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphListItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphList) GetItemsOk() ([]MicrosoftGraphListItem, bool) {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphListItem
		return ret, false
	}
	return *o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MicrosoftGraphList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MicrosoftGraphListItem and assigns it to the Items field.
func (o *MicrosoftGraphList) SetItems(v []MicrosoftGraphListItem) {
	o.Items = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.List == nil {
		if o.isExplicitNullList {
			toSerialize["list"] = o.List
		}
	} else {
		toSerialize["list"] = o.List
	}
	if o.SharepointIds == nil {
		if o.isExplicitNullSharepointIds {
			toSerialize["sharepointIds"] = o.SharepointIds
		}
	} else {
		toSerialize["sharepointIds"] = o.SharepointIds
	}
	if o.System == nil {
		if o.isExplicitNullSystem {
			toSerialize["system"] = o.System
		}
	} else {
		toSerialize["system"] = o.System
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.ContentTypes != nil {
		toSerialize["contentTypes"] = o.ContentTypes
	}
	if o.Drive == nil {
		if o.isExplicitNullDrive {
			toSerialize["drive"] = o.Drive
		}
	} else {
		toSerialize["drive"] = o.Drive
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

