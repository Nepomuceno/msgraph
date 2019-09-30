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
// MicrosoftGraphSite struct for MicrosoftGraphSite
type MicrosoftGraphSite struct {
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
	Root *AnyOfobject `json:"root,omitempty"`
	isExplicitNullRoot bool `json:"-"`
	SharepointIds *AnyOfmicrosoftGraphSharepointIds `json:"sharepointIds,omitempty"`
	isExplicitNullSharepointIds bool `json:"-"`
	SiteCollection *AnyOfmicrosoftGraphSiteCollection `json:"siteCollection,omitempty"`
	isExplicitNullSiteCollection bool `json:"-"`
	Analytics *AnyOfmicrosoftGraphItemAnalytics `json:"analytics,omitempty"`
	isExplicitNullAnalytics bool `json:"-"`
	Columns *[]MicrosoftGraphColumnDefinition `json:"columns,omitempty"`

	ContentTypes *[]MicrosoftGraphContentType `json:"contentTypes,omitempty"`

	Drive *AnyOfmicrosoftGraphDrive `json:"drive,omitempty"`
	isExplicitNullDrive bool `json:"-"`
	Drives *[]MicrosoftGraphDrive `json:"drives,omitempty"`

	Items *[]MicrosoftGraphBaseItem `json:"items,omitempty"`

	Lists *[]MicrosoftGraphList `json:"lists,omitempty"`

	Sites *[]MicrosoftGraphSite `json:"sites,omitempty"`

	Onenote *AnyOfmicrosoftGraphOnenote `json:"onenote,omitempty"`
	isExplicitNullOnenote bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSite) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphSite) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphSite) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphSite) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetETag returns the ETag field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetETagOk() (string, bool) {
	if o == nil || o.ETag == nil {
		var ret string
		return ret, false
	}
	return *o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *MicrosoftGraphSite) SetETag(v string) {
	o.ETag = &v
}

// SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ETag value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetETagExplicitNull(b bool) {
	o.ETag = nil
	o.isExplicitNullETag = b
}
// GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphSite) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetLastModifiedByExplicitNull(b bool) {
	o.LastModifiedBy = nil
	o.isExplicitNullLastModifiedBy = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphSite) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphSite) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetParentReference returns the ParentReference field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		var ret AnyOfmicrosoftGraphItemReference
		return ret, false
	}
	return *o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *MicrosoftGraphSite) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = &v
}

// SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentReference value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetParentReferenceExplicitNull(b bool) {
	o.ParentReference = nil
	o.isExplicitNullParentReference = b
}
// GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetWebUrlOk() (string, bool) {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret, false
	}
	return *o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *MicrosoftGraphSite) SetWebUrl(v string) {
	o.WebUrl = &v
}

// SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The WebUrl value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetWebUrlExplicitNull(b bool) {
	o.WebUrl = nil
	o.isExplicitNullWebUrl = b
}
// GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetCreatedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.
func (o *MicrosoftGraphSite) SetCreatedByUser(v AnyOfmicrosoftGraphUser) {
	o.CreatedByUser = &v
}

// SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedByUser value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetCreatedByUserExplicitNull(b bool) {
	o.CreatedByUser = nil
	o.isExplicitNullCreatedByUser = b
}
// GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetLastModifiedByUser() AnyOfmicrosoftGraphUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		var ret AnyOfmicrosoftGraphUser
		return ret, false
	}
	return *o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.
func (o *MicrosoftGraphSite) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser) {
	o.LastModifiedByUser = &v
}

// SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedByUser value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetLastModifiedByUserExplicitNull(b bool) {
	o.LastModifiedByUser = nil
	o.isExplicitNullLastModifiedByUser = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphSite) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetRoot returns the Root field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetRoot() AnyOfobject {
	if o == nil || o.Root == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetRootOk() (AnyOfobject, bool) {
	if o == nil || o.Root == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given AnyOfobject and assigns it to the Root field.
func (o *MicrosoftGraphSite) SetRoot(v AnyOfobject) {
	o.Root = &v
}

// SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Root value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetRootExplicitNull(b bool) {
	o.Root = nil
	o.isExplicitNullRoot = b
}
// GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds {
	if o == nil || o.SharepointIds == nil {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret
	}
	return *o.SharepointIds
}

// GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool) {
	if o == nil || o.SharepointIds == nil {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret, false
	}
	return *o.SharepointIds, true
}

// HasSharepointIds returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasSharepointIds() bool {
	if o != nil && o.SharepointIds != nil {
		return true
	}

	return false
}

// SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.
func (o *MicrosoftGraphSite) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds) {
	o.SharepointIds = &v
}

// SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharepointIds value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetSharepointIdsExplicitNull(b bool) {
	o.SharepointIds = nil
	o.isExplicitNullSharepointIds = b
}
// GetSiteCollection returns the SiteCollection field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetSiteCollection() AnyOfmicrosoftGraphSiteCollection {
	if o == nil || o.SiteCollection == nil {
		var ret AnyOfmicrosoftGraphSiteCollection
		return ret
	}
	return *o.SiteCollection
}

// GetSiteCollectionOk returns a tuple with the SiteCollection field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetSiteCollectionOk() (AnyOfmicrosoftGraphSiteCollection, bool) {
	if o == nil || o.SiteCollection == nil {
		var ret AnyOfmicrosoftGraphSiteCollection
		return ret, false
	}
	return *o.SiteCollection, true
}

// HasSiteCollection returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasSiteCollection() bool {
	if o != nil && o.SiteCollection != nil {
		return true
	}

	return false
}

// SetSiteCollection gets a reference to the given AnyOfmicrosoftGraphSiteCollection and assigns it to the SiteCollection field.
func (o *MicrosoftGraphSite) SetSiteCollection(v AnyOfmicrosoftGraphSiteCollection) {
	o.SiteCollection = &v
}

// SetSiteCollectionExplicitNull (un)sets SiteCollection to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SiteCollection value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetSiteCollectionExplicitNull(b bool) {
	o.SiteCollection = nil
	o.isExplicitNullSiteCollection = b
}
// GetAnalytics returns the Analytics field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics {
	if o == nil || o.Analytics == nil {
		var ret AnyOfmicrosoftGraphItemAnalytics
		return ret
	}
	return *o.Analytics
}

// GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetAnalyticsOk() (AnyOfmicrosoftGraphItemAnalytics, bool) {
	if o == nil || o.Analytics == nil {
		var ret AnyOfmicrosoftGraphItemAnalytics
		return ret, false
	}
	return *o.Analytics, true
}

// HasAnalytics returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasAnalytics() bool {
	if o != nil && o.Analytics != nil {
		return true
	}

	return false
}

// SetAnalytics gets a reference to the given AnyOfmicrosoftGraphItemAnalytics and assigns it to the Analytics field.
func (o *MicrosoftGraphSite) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics) {
	o.Analytics = &v
}

// SetAnalyticsExplicitNull (un)sets Analytics to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Analytics value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetAnalyticsExplicitNull(b bool) {
	o.Analytics = nil
	o.isExplicitNullAnalytics = b
}
// GetColumns returns the Columns field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetColumns() []MicrosoftGraphColumnDefinition {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret, false
	}
	return *o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.
func (o *MicrosoftGraphSite) SetColumns(v []MicrosoftGraphColumnDefinition) {
	o.Columns = &v
}

// GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetContentTypes() []MicrosoftGraphContentType {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret
	}
	return *o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetContentTypesOk() ([]MicrosoftGraphContentType, bool) {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret, false
	}
	return *o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasContentTypes() bool {
	if o != nil && o.ContentTypes != nil {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.
func (o *MicrosoftGraphSite) SetContentTypes(v []MicrosoftGraphContentType) {
	o.ContentTypes = &v
}

// GetDrive returns the Drive field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetDrive() AnyOfmicrosoftGraphDrive {
	if o == nil || o.Drive == nil {
		var ret AnyOfmicrosoftGraphDrive
		return ret
	}
	return *o.Drive
}

// GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool) {
	if o == nil || o.Drive == nil {
		var ret AnyOfmicrosoftGraphDrive
		return ret, false
	}
	return *o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasDrive() bool {
	if o != nil && o.Drive != nil {
		return true
	}

	return false
}

// SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.
func (o *MicrosoftGraphSite) SetDrive(v AnyOfmicrosoftGraphDrive) {
	o.Drive = &v
}

// SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Drive value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetDriveExplicitNull(b bool) {
	o.Drive = nil
	o.isExplicitNullDrive = b
}
// GetDrives returns the Drives field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetDrives() []MicrosoftGraphDrive {
	if o == nil || o.Drives == nil {
		var ret []MicrosoftGraphDrive
		return ret
	}
	return *o.Drives
}

// GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetDrivesOk() ([]MicrosoftGraphDrive, bool) {
	if o == nil || o.Drives == nil {
		var ret []MicrosoftGraphDrive
		return ret, false
	}
	return *o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasDrives() bool {
	if o != nil && o.Drives != nil {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.
func (o *MicrosoftGraphSite) SetDrives(v []MicrosoftGraphDrive) {
	o.Drives = &v
}

// GetItems returns the Items field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetItems() []MicrosoftGraphBaseItem {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphBaseItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetItemsOk() ([]MicrosoftGraphBaseItem, bool) {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphBaseItem
		return ret, false
	}
	return *o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MicrosoftGraphBaseItem and assigns it to the Items field.
func (o *MicrosoftGraphSite) SetItems(v []MicrosoftGraphBaseItem) {
	o.Items = &v
}

// GetLists returns the Lists field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetLists() []MicrosoftGraphList {
	if o == nil || o.Lists == nil {
		var ret []MicrosoftGraphList
		return ret
	}
	return *o.Lists
}

// GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetListsOk() ([]MicrosoftGraphList, bool) {
	if o == nil || o.Lists == nil {
		var ret []MicrosoftGraphList
		return ret, false
	}
	return *o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasLists() bool {
	if o != nil && o.Lists != nil {
		return true
	}

	return false
}

// SetLists gets a reference to the given []MicrosoftGraphList and assigns it to the Lists field.
func (o *MicrosoftGraphSite) SetLists(v []MicrosoftGraphList) {
	o.Lists = &v
}

// GetSites returns the Sites field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetSites() []MicrosoftGraphSite {
	if o == nil || o.Sites == nil {
		var ret []MicrosoftGraphSite
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetSitesOk() ([]MicrosoftGraphSite, bool) {
	if o == nil || o.Sites == nil {
		var ret []MicrosoftGraphSite
		return ret, false
	}
	return *o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []MicrosoftGraphSite and assigns it to the Sites field.
func (o *MicrosoftGraphSite) SetSites(v []MicrosoftGraphSite) {
	o.Sites = &v
}

// GetOnenote returns the Onenote field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSite) GetOnenote() AnyOfmicrosoftGraphOnenote {
	if o == nil || o.Onenote == nil {
		var ret AnyOfmicrosoftGraphOnenote
		return ret
	}
	return *o.Onenote
}

// GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSite) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool) {
	if o == nil || o.Onenote == nil {
		var ret AnyOfmicrosoftGraphOnenote
		return ret, false
	}
	return *o.Onenote, true
}

// HasOnenote returns a boolean if a field has been set.
func (o *MicrosoftGraphSite) HasOnenote() bool {
	if o != nil && o.Onenote != nil {
		return true
	}

	return false
}

// SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.
func (o *MicrosoftGraphSite) SetOnenote(v AnyOfmicrosoftGraphOnenote) {
	o.Onenote = &v
}

// SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Onenote value is set to nil even if false is passed
func (o *MicrosoftGraphSite) SetOnenoteExplicitNull(b bool) {
	o.Onenote = nil
	o.isExplicitNullOnenote = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSite) MarshalJSON() ([]byte, error) {
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
	if o.Root == nil {
		if o.isExplicitNullRoot {
			toSerialize["root"] = o.Root
		}
	} else {
		toSerialize["root"] = o.Root
	}
	if o.SharepointIds == nil {
		if o.isExplicitNullSharepointIds {
			toSerialize["sharepointIds"] = o.SharepointIds
		}
	} else {
		toSerialize["sharepointIds"] = o.SharepointIds
	}
	if o.SiteCollection == nil {
		if o.isExplicitNullSiteCollection {
			toSerialize["siteCollection"] = o.SiteCollection
		}
	} else {
		toSerialize["siteCollection"] = o.SiteCollection
	}
	if o.Analytics == nil {
		if o.isExplicitNullAnalytics {
			toSerialize["analytics"] = o.Analytics
		}
	} else {
		toSerialize["analytics"] = o.Analytics
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
	if o.Drives != nil {
		toSerialize["drives"] = o.Drives
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Lists != nil {
		toSerialize["lists"] = o.Lists
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	if o.Onenote == nil {
		if o.isExplicitNullOnenote {
			toSerialize["onenote"] = o.Onenote
		}
	} else {
		toSerialize["onenote"] = o.Onenote
	}
	return json.Marshal(toSerialize)
}


