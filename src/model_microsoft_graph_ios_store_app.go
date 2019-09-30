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
// MicrosoftGraphIosStoreApp struct for MicrosoftGraphIosStoreApp
type MicrosoftGraphIosStoreApp struct {
	Id *string `json:"id,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	// The description of the app.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// The publisher of the app.
	Publisher *string `json:"publisher,omitempty"`
	isExplicitNullPublisher bool `json:"-"`
	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *AnyOfmicrosoftGraphMimeContent `json:"largeIcon,omitempty"`
	isExplicitNullLargeIcon bool `json:"-"`
	// The date and time the app was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl *string `json:"privacyInformationUrl,omitempty"`
	isExplicitNullPrivacyInformationUrl bool `json:"-"`
	// The more information Url.
	InformationUrl *string `json:"informationUrl,omitempty"`
	isExplicitNullInformationUrl bool `json:"-"`
	// The owner of the app.
	Owner *string `json:"owner,omitempty"`
	isExplicitNullOwner bool `json:"-"`
	// The developer of the app.
	Developer *string `json:"developer,omitempty"`
	isExplicitNullDeveloper bool `json:"-"`
	// Notes for the app.
	Notes *string `json:"notes,omitempty"`
	isExplicitNullNotes bool `json:"-"`
	// The publishing state for the app. The app cannot be assigned unless the app is published.
	PublishingState *AnyOfmicrosoftGraphMobileAppPublishingState `json:"publishingState,omitempty"`

	Categories *[]MicrosoftGraphMobileAppCategory `json:"categories,omitempty"`

	Assignments *[]MicrosoftGraphMobileAppAssignment `json:"assignments,omitempty"`

	// The Identity Name.
	BundleId *string `json:"bundleId,omitempty"`
	isExplicitNullBundleId bool `json:"-"`
	// The Apple App Store URL
	AppStoreUrl *string `json:"appStoreUrl,omitempty"`
	isExplicitNullAppStoreUrl bool `json:"-"`
	ApplicableDeviceType *MicrosoftGraphIosDeviceType `json:"applicableDeviceType,omitempty"`

	// The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *AnyOfmicrosoftGraphIosMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	isExplicitNullMinimumSupportedOperatingSystem bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphIosStoreApp) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphIosStoreApp) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphIosStoreApp) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetPublisher returns the Publisher field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetPublisherOk() (string, bool) {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret, false
	}
	return *o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *MicrosoftGraphIosStoreApp) SetPublisher(v string) {
	o.Publisher = &v
}

// SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Publisher value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetPublisherExplicitNull(b bool) {
	o.Publisher = nil
	o.isExplicitNullPublisher = b
}
// GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent {
	if o == nil || o.LargeIcon == nil {
		var ret AnyOfmicrosoftGraphMimeContent
		return ret
	}
	return *o.LargeIcon
}

// GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool) {
	if o == nil || o.LargeIcon == nil {
		var ret AnyOfmicrosoftGraphMimeContent
		return ret, false
	}
	return *o.LargeIcon, true
}

// HasLargeIcon returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasLargeIcon() bool {
	if o != nil && o.LargeIcon != nil {
		return true
	}

	return false
}

// SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.
func (o *MicrosoftGraphIosStoreApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent) {
	o.LargeIcon = &v
}

// SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LargeIcon value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetLargeIconExplicitNull(b bool) {
	o.LargeIcon = nil
	o.isExplicitNullLargeIcon = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphIosStoreApp) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphIosStoreApp) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetIsFeatured() bool {
	if o == nil || o.IsFeatured == nil {
		var ret bool
		return ret
	}
	return *o.IsFeatured
}

// GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetIsFeaturedOk() (bool, bool) {
	if o == nil || o.IsFeatured == nil {
		var ret bool
		return ret, false
	}
	return *o.IsFeatured, true
}

// HasIsFeatured returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasIsFeatured() bool {
	if o != nil && o.IsFeatured != nil {
		return true
	}

	return false
}

// SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.
func (o *MicrosoftGraphIosStoreApp) SetIsFeatured(v bool) {
	o.IsFeatured = &v
}

// GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetPrivacyInformationUrl() string {
	if o == nil || o.PrivacyInformationUrl == nil {
		var ret string
		return ret
	}
	return *o.PrivacyInformationUrl
}

// GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetPrivacyInformationUrlOk() (string, bool) {
	if o == nil || o.PrivacyInformationUrl == nil {
		var ret string
		return ret, false
	}
	return *o.PrivacyInformationUrl, true
}

// HasPrivacyInformationUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasPrivacyInformationUrl() bool {
	if o != nil && o.PrivacyInformationUrl != nil {
		return true
	}

	return false
}

// SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.
func (o *MicrosoftGraphIosStoreApp) SetPrivacyInformationUrl(v string) {
	o.PrivacyInformationUrl = &v
}

// SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PrivacyInformationUrl value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetPrivacyInformationUrlExplicitNull(b bool) {
	o.PrivacyInformationUrl = nil
	o.isExplicitNullPrivacyInformationUrl = b
}
// GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetInformationUrl() string {
	if o == nil || o.InformationUrl == nil {
		var ret string
		return ret
	}
	return *o.InformationUrl
}

// GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetInformationUrlOk() (string, bool) {
	if o == nil || o.InformationUrl == nil {
		var ret string
		return ret, false
	}
	return *o.InformationUrl, true
}

// HasInformationUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasInformationUrl() bool {
	if o != nil && o.InformationUrl != nil {
		return true
	}

	return false
}

// SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.
func (o *MicrosoftGraphIosStoreApp) SetInformationUrl(v string) {
	o.InformationUrl = &v
}

// SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InformationUrl value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetInformationUrlExplicitNull(b bool) {
	o.InformationUrl = nil
	o.isExplicitNullInformationUrl = b
}
// GetOwner returns the Owner field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetOwnerOk() (string, bool) {
	if o == nil || o.Owner == nil {
		var ret string
		return ret, false
	}
	return *o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *MicrosoftGraphIosStoreApp) SetOwner(v string) {
	o.Owner = &v
}

// SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Owner value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetOwnerExplicitNull(b bool) {
	o.Owner = nil
	o.isExplicitNullOwner = b
}
// GetDeveloper returns the Developer field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetDeveloper() string {
	if o == nil || o.Developer == nil {
		var ret string
		return ret
	}
	return *o.Developer
}

// GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetDeveloperOk() (string, bool) {
	if o == nil || o.Developer == nil {
		var ret string
		return ret, false
	}
	return *o.Developer, true
}

// HasDeveloper returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasDeveloper() bool {
	if o != nil && o.Developer != nil {
		return true
	}

	return false
}

// SetDeveloper gets a reference to the given string and assigns it to the Developer field.
func (o *MicrosoftGraphIosStoreApp) SetDeveloper(v string) {
	o.Developer = &v
}

// SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Developer value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetDeveloperExplicitNull(b bool) {
	o.Developer = nil
	o.isExplicitNullDeveloper = b
}
// GetNotes returns the Notes field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetNotesOk() (string, bool) {
	if o == nil || o.Notes == nil {
		var ret string
		return ret, false
	}
	return *o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *MicrosoftGraphIosStoreApp) SetNotes(v string) {
	o.Notes = &v
}

// SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Notes value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetNotesExplicitNull(b bool) {
	o.Notes = nil
	o.isExplicitNullNotes = b
}
// GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState {
	if o == nil || o.PublishingState == nil {
		var ret AnyOfmicrosoftGraphMobileAppPublishingState
		return ret
	}
	return *o.PublishingState
}

// GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool) {
	if o == nil || o.PublishingState == nil {
		var ret AnyOfmicrosoftGraphMobileAppPublishingState
		return ret, false
	}
	return *o.PublishingState, true
}

// HasPublishingState returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasPublishingState() bool {
	if o != nil && o.PublishingState != nil {
		return true
	}

	return false
}

// SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.
func (o *MicrosoftGraphIosStoreApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState) {
	o.PublishingState = &v
}

// GetCategories returns the Categories field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetCategories() []MicrosoftGraphMobileAppCategory {
	if o == nil || o.Categories == nil {
		var ret []MicrosoftGraphMobileAppCategory
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool) {
	if o == nil || o.Categories == nil {
		var ret []MicrosoftGraphMobileAppCategory
		return ret, false
	}
	return *o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.
func (o *MicrosoftGraphIosStoreApp) SetCategories(v []MicrosoftGraphMobileAppCategory) {
	o.Categories = &v
}

// GetAssignments returns the Assignments field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetAssignments() []MicrosoftGraphMobileAppAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphMobileAppAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool) {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphMobileAppAssignment
		return ret, false
	}
	return *o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.
func (o *MicrosoftGraphIosStoreApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment) {
	o.Assignments = &v
}

// GetBundleId returns the BundleId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetBundleIdOk() (string, bool) {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret, false
	}
	return *o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *MicrosoftGraphIosStoreApp) SetBundleId(v string) {
	o.BundleId = &v
}

// SetBundleIdExplicitNull (un)sets BundleId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BundleId value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetBundleIdExplicitNull(b bool) {
	o.BundleId = nil
	o.isExplicitNullBundleId = b
}
// GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetAppStoreUrl() string {
	if o == nil || o.AppStoreUrl == nil {
		var ret string
		return ret
	}
	return *o.AppStoreUrl
}

// GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetAppStoreUrlOk() (string, bool) {
	if o == nil || o.AppStoreUrl == nil {
		var ret string
		return ret, false
	}
	return *o.AppStoreUrl, true
}

// HasAppStoreUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasAppStoreUrl() bool {
	if o != nil && o.AppStoreUrl != nil {
		return true
	}

	return false
}

// SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.
func (o *MicrosoftGraphIosStoreApp) SetAppStoreUrl(v string) {
	o.AppStoreUrl = &v
}

// SetAppStoreUrlExplicitNull (un)sets AppStoreUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppStoreUrl value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetAppStoreUrlExplicitNull(b bool) {
	o.AppStoreUrl = nil
	o.isExplicitNullAppStoreUrl = b
}
// GetApplicableDeviceType returns the ApplicableDeviceType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetApplicableDeviceType() MicrosoftGraphIosDeviceType {
	if o == nil || o.ApplicableDeviceType == nil {
		var ret MicrosoftGraphIosDeviceType
		return ret
	}
	return *o.ApplicableDeviceType
}

// GetApplicableDeviceTypeOk returns a tuple with the ApplicableDeviceType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetApplicableDeviceTypeOk() (MicrosoftGraphIosDeviceType, bool) {
	if o == nil || o.ApplicableDeviceType == nil {
		var ret MicrosoftGraphIosDeviceType
		return ret, false
	}
	return *o.ApplicableDeviceType, true
}

// HasApplicableDeviceType returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasApplicableDeviceType() bool {
	if o != nil && o.ApplicableDeviceType != nil {
		return true
	}

	return false
}

// SetApplicableDeviceType gets a reference to the given MicrosoftGraphIosDeviceType and assigns it to the ApplicableDeviceType field.
func (o *MicrosoftGraphIosStoreApp) SetApplicableDeviceType(v MicrosoftGraphIosDeviceType) {
	o.ApplicableDeviceType = &v
}

// GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosStoreApp) GetMinimumSupportedOperatingSystem() AnyOfmicrosoftGraphIosMinimumOperatingSystem {
	if o == nil || o.MinimumSupportedOperatingSystem == nil {
		var ret AnyOfmicrosoftGraphIosMinimumOperatingSystem
		return ret
	}
	return *o.MinimumSupportedOperatingSystem
}

// GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosStoreApp) GetMinimumSupportedOperatingSystemOk() (AnyOfmicrosoftGraphIosMinimumOperatingSystem, bool) {
	if o == nil || o.MinimumSupportedOperatingSystem == nil {
		var ret AnyOfmicrosoftGraphIosMinimumOperatingSystem
		return ret, false
	}
	return *o.MinimumSupportedOperatingSystem, true
}

// HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.
func (o *MicrosoftGraphIosStoreApp) HasMinimumSupportedOperatingSystem() bool {
	if o != nil && o.MinimumSupportedOperatingSystem != nil {
		return true
	}

	return false
}

// SetMinimumSupportedOperatingSystem gets a reference to the given AnyOfmicrosoftGraphIosMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.
func (o *MicrosoftGraphIosStoreApp) SetMinimumSupportedOperatingSystem(v AnyOfmicrosoftGraphIosMinimumOperatingSystem) {
	o.MinimumSupportedOperatingSystem = &v
}

// SetMinimumSupportedOperatingSystemExplicitNull (un)sets MinimumSupportedOperatingSystem to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The MinimumSupportedOperatingSystem value is set to nil even if false is passed
func (o *MicrosoftGraphIosStoreApp) SetMinimumSupportedOperatingSystemExplicitNull(b bool) {
	o.MinimumSupportedOperatingSystem = nil
	o.isExplicitNullMinimumSupportedOperatingSystem = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphIosStoreApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.Publisher == nil {
		if o.isExplicitNullPublisher {
			toSerialize["publisher"] = o.Publisher
		}
	} else {
		toSerialize["publisher"] = o.Publisher
	}
	if o.LargeIcon == nil {
		if o.isExplicitNullLargeIcon {
			toSerialize["largeIcon"] = o.LargeIcon
		}
	} else {
		toSerialize["largeIcon"] = o.LargeIcon
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.IsFeatured != nil {
		toSerialize["isFeatured"] = o.IsFeatured
	}
	if o.PrivacyInformationUrl == nil {
		if o.isExplicitNullPrivacyInformationUrl {
			toSerialize["privacyInformationUrl"] = o.PrivacyInformationUrl
		}
	} else {
		toSerialize["privacyInformationUrl"] = o.PrivacyInformationUrl
	}
	if o.InformationUrl == nil {
		if o.isExplicitNullInformationUrl {
			toSerialize["informationUrl"] = o.InformationUrl
		}
	} else {
		toSerialize["informationUrl"] = o.InformationUrl
	}
	if o.Owner == nil {
		if o.isExplicitNullOwner {
			toSerialize["owner"] = o.Owner
		}
	} else {
		toSerialize["owner"] = o.Owner
	}
	if o.Developer == nil {
		if o.isExplicitNullDeveloper {
			toSerialize["developer"] = o.Developer
		}
	} else {
		toSerialize["developer"] = o.Developer
	}
	if o.Notes == nil {
		if o.isExplicitNullNotes {
			toSerialize["notes"] = o.Notes
		}
	} else {
		toSerialize["notes"] = o.Notes
	}
	if o.PublishingState != nil {
		toSerialize["publishingState"] = o.PublishingState
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.BundleId == nil {
		if o.isExplicitNullBundleId {
			toSerialize["bundleId"] = o.BundleId
		}
	} else {
		toSerialize["bundleId"] = o.BundleId
	}
	if o.AppStoreUrl == nil {
		if o.isExplicitNullAppStoreUrl {
			toSerialize["appStoreUrl"] = o.AppStoreUrl
		}
	} else {
		toSerialize["appStoreUrl"] = o.AppStoreUrl
	}
	if o.ApplicableDeviceType != nil {
		toSerialize["applicableDeviceType"] = o.ApplicableDeviceType
	}
	if o.MinimumSupportedOperatingSystem == nil {
		if o.isExplicitNullMinimumSupportedOperatingSystem {
			toSerialize["minimumSupportedOperatingSystem"] = o.MinimumSupportedOperatingSystem
		}
	} else {
		toSerialize["minimumSupportedOperatingSystem"] = o.MinimumSupportedOperatingSystem
	}
	return json.Marshal(toSerialize)
}


