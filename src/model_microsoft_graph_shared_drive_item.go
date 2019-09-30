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
// MicrosoftGraphSharedDriveItem struct for MicrosoftGraphSharedDriveItem
type MicrosoftGraphSharedDriveItem struct {
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
	Owner *AnyOfmicrosoftGraphIdentitySet `json:"owner,omitempty"`
	isExplicitNullOwner bool `json:"-"`
	DriveItem *AnyOfmicrosoftGraphDriveItem `json:"driveItem,omitempty"`
	isExplicitNullDriveItem bool `json:"-"`
	Items *[]MicrosoftGraphDriveItem `json:"items,omitempty"`

	List *AnyOfmicrosoftGraphList `json:"list,omitempty"`
	isExplicitNullList bool `json:"-"`
	ListItem *AnyOfmicrosoftGraphListItem `json:"listItem,omitempty"`
	isExplicitNullListItem bool `json:"-"`
	Root *AnyOfmicrosoftGraphDriveItem `json:"root,omitempty"`
	isExplicitNullRoot bool `json:"-"`
	Site *AnyOfmicrosoftGraphSite `json:"site,omitempty"`
	isExplicitNullSite bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSharedDriveItem) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphSharedDriveItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphSharedDriveItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphSharedDriveItem) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetETag returns the ETag field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetETagOk() (string, bool) {
	if o == nil || o.ETag == nil {
		var ret string
		return ret, false
	}
	return *o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *MicrosoftGraphSharedDriveItem) SetETag(v string) {
	o.ETag = &v
}

// SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ETag value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetETagExplicitNull(b bool) {
	o.ETag = nil
	o.isExplicitNullETag = b
}
// GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedByExplicitNull(b bool) {
	o.LastModifiedBy = nil
	o.isExplicitNullLastModifiedBy = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphSharedDriveItem) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetParentReference returns the ParentReference field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret, false
	}
	return *o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *MicrosoftGraphSharedDriveItem) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = &v
}

// SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentReference value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetParentReferenceExplicitNull(b bool) {
	o.ParentReference = nil
	o.isExplicitNullParentReference = b
}
// GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetWebUrlOk() (string, bool) {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret, false
	}
	return *o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *MicrosoftGraphSharedDriveItem) SetWebUrl(v string) {
	o.WebUrl = &v
}

// SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The WebUrl value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetWebUrlExplicitNull(b bool) {
	o.WebUrl = nil
	o.isExplicitNullWebUrl = b
}
// GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetCreatedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.
func (o *MicrosoftGraphSharedDriveItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser) {
	o.CreatedByUser = &v
}

// SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedByUser value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetCreatedByUserExplicitNull(b bool) {
	o.CreatedByUser = nil
	o.isExplicitNullCreatedByUser = b
}
// GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.
func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser) {
	o.LastModifiedByUser = &v
}

// SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedByUser value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedByUserExplicitNull(b bool) {
	o.LastModifiedByUser = nil
	o.isExplicitNullLastModifiedByUser = b
}
// GetOwner returns the Owner field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetOwner() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.Owner == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetOwnerOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Owner == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Owner field.
func (o *MicrosoftGraphSharedDriveItem) SetOwner(v AnyOfmicrosoftGraphIdentitySet) {
	o.Owner = &v
}

// SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Owner value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetOwnerExplicitNull(b bool) {
	o.Owner = nil
	o.isExplicitNullOwner = b
}
// GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetDriveItem() AnyOfmicrosoftGraphDriveItem {
	if o == nil || o.DriveItem == nil {
		var ret AnyOfmicrosoftGraphDriveItem
		return ret
	}
	return *o.DriveItem
}

// GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetDriveItemOk() (AnyOfmicrosoftGraphDriveItem, bool) {
	if o == nil || o.DriveItem == nil {
		var ret AnyOfmicrosoftGraphDriveItem
		return ret, false
	}
	return *o.DriveItem, true
}

// HasDriveItem returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasDriveItem() bool {
	if o != nil && o.DriveItem != nil {
		return true
	}

	return false
}

// SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.
func (o *MicrosoftGraphSharedDriveItem) SetDriveItem(v AnyOfmicrosoftGraphDriveItem) {
	o.DriveItem = &v
}

// SetDriveItemExplicitNull (un)sets DriveItem to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DriveItem value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetDriveItemExplicitNull(b bool) {
	o.DriveItem = nil
	o.isExplicitNullDriveItem = b
}
// GetItems returns the Items field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetItems() []MicrosoftGraphDriveItem {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphDriveItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetItemsOk() ([]MicrosoftGraphDriveItem, bool) {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphDriveItem
		return ret, false
	}
	return *o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Items field.
func (o *MicrosoftGraphSharedDriveItem) SetItems(v []MicrosoftGraphDriveItem) {
	o.Items = &v
}

// GetList returns the List field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetList() AnyOfmicrosoftGraphList {
	if o == nil || o.List == nil {
		var ret AnyOfmicrosoftGraphList
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetListOk() (AnyOfmicrosoftGraphList, bool) {
	if o == nil || o.List == nil {
		var ret AnyOfmicrosoftGraphList
		return ret, false
	}
	return *o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given AnyOfmicrosoftGraphList and assigns it to the List field.
func (o *MicrosoftGraphSharedDriveItem) SetList(v AnyOfmicrosoftGraphList) {
	o.List = &v
}

// SetListExplicitNull (un)sets List to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The List value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetListExplicitNull(b bool) {
	o.List = nil
	o.isExplicitNullList = b
}
// GetListItem returns the ListItem field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetListItem() AnyOfmicrosoftGraphListItem {
	if o == nil || o.ListItem == nil {
		var ret AnyOfmicrosoftGraphListItem
		return ret
	}
	return *o.ListItem
}

// GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetListItemOk() (AnyOfmicrosoftGraphListItem, bool) {
	if o == nil || o.ListItem == nil {
		var ret AnyOfmicrosoftGraphListItem
		return ret, false
	}
	return *o.ListItem, true
}

// HasListItem returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasListItem() bool {
	if o != nil && o.ListItem != nil {
		return true
	}

	return false
}

// SetListItem gets a reference to the given AnyOfmicrosoftGraphListItem and assigns it to the ListItem field.
func (o *MicrosoftGraphSharedDriveItem) SetListItem(v AnyOfmicrosoftGraphListItem) {
	o.ListItem = &v
}

// SetListItemExplicitNull (un)sets ListItem to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ListItem value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetListItemExplicitNull(b bool) {
	o.ListItem = nil
	o.isExplicitNullListItem = b
}
// GetRoot returns the Root field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetRoot() AnyOfmicrosoftGraphDriveItem {
	if o == nil || o.Root == nil {
		var ret AnyOfmicrosoftGraphDriveItem
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetRootOk() (AnyOfmicrosoftGraphDriveItem, bool) {
	if o == nil || o.Root == nil {
		var ret AnyOfmicrosoftGraphDriveItem
		return ret, false
	}
	return *o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the Root field.
func (o *MicrosoftGraphSharedDriveItem) SetRoot(v AnyOfmicrosoftGraphDriveItem) {
	o.Root = &v
}

// SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Root value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetRootExplicitNull(b bool) {
	o.Root = nil
	o.isExplicitNullRoot = b
}
// GetSite returns the Site field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharedDriveItem) GetSite() AnyOfmicrosoftGraphSite {
	if o == nil || o.Site == nil {
		var ret AnyOfmicrosoftGraphSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharedDriveItem) GetSiteOk() (AnyOfmicrosoftGraphSite, bool) {
	if o == nil || o.Site == nil {
		var ret AnyOfmicrosoftGraphSite
		return ret, false
	}
	return *o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *MicrosoftGraphSharedDriveItem) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given AnyOfmicrosoftGraphSite and assigns it to the Site field.
func (o *MicrosoftGraphSharedDriveItem) SetSite(v AnyOfmicrosoftGraphSite) {
	o.Site = &v
}

// SetSiteExplicitNull (un)sets Site to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Site value is set to nil even if false is passed
func (o *MicrosoftGraphSharedDriveItem) SetSiteExplicitNull(b bool) {
	o.Site = nil
	o.isExplicitNullSite = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSharedDriveItem) MarshalJSON() ([]byte, error) {
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
	if o.Owner == nil {
		if o.isExplicitNullOwner {
			toSerialize["owner"] = o.Owner
		}
	} else {
		toSerialize["owner"] = o.Owner
	}
	if o.DriveItem == nil {
		if o.isExplicitNullDriveItem {
			toSerialize["driveItem"] = o.DriveItem
		}
	} else {
		toSerialize["driveItem"] = o.DriveItem
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.List == nil {
		if o.isExplicitNullList {
			toSerialize["list"] = o.List
		}
	} else {
		toSerialize["list"] = o.List
	}
	if o.ListItem == nil {
		if o.isExplicitNullListItem {
			toSerialize["listItem"] = o.ListItem
		}
	} else {
		toSerialize["listItem"] = o.ListItem
	}
	if o.Root == nil {
		if o.isExplicitNullRoot {
			toSerialize["root"] = o.Root
		}
	} else {
		toSerialize["root"] = o.Root
	}
	if o.Site == nil {
		if o.isExplicitNullSite {
			toSerialize["site"] = o.Site
		}
	} else {
		toSerialize["site"] = o.Site
	}
	return json.Marshal(toSerialize)
}

