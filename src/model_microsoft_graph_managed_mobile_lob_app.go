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
// MicrosoftGraphManagedMobileLobApp struct for MicrosoftGraphManagedMobileLobApp
type MicrosoftGraphManagedMobileLobApp struct {
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

	// The Application's availability.
	AppAvailability *AnyOfmicrosoftGraphManagedAppAvailability `json:"appAvailability,omitempty"`

	// The Application's version.
	Version *string `json:"version,omitempty"`
	isExplicitNullVersion bool `json:"-"`
	// The internal committed content version.
	CommittedContentVersion *string `json:"committedContentVersion,omitempty"`
	isExplicitNullCommittedContentVersion bool `json:"-"`
	// The name of the main Lob application file.
	FileName *string `json:"fileName,omitempty"`
	isExplicitNullFileName bool `json:"-"`
	// The total size, including all uploaded files.
	Size *int64 `json:"size,omitempty"`

	ContentVersions *[]MicrosoftGraphMobileAppContent `json:"contentVersions,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphManagedMobileLobApp) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphManagedMobileLobApp) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphManagedMobileLobApp) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetPublisher returns the Publisher field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetPublisherOk() (string, bool) {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret, false
	}
	return *o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *MicrosoftGraphManagedMobileLobApp) SetPublisher(v string) {
	o.Publisher = &v
}

// SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Publisher value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetPublisherExplicitNull(b bool) {
	o.Publisher = nil
	o.isExplicitNullPublisher = b
}
// GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent {
	if o == nil || o.LargeIcon == nil {
		var ret AnyOfmicrosoftGraphMimeContent
		return ret
	}
	return *o.LargeIcon
}

// GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool) {
	if o == nil || o.LargeIcon == nil {
		var ret AnyOfmicrosoftGraphMimeContent
		return ret, false
	}
	return *o.LargeIcon, true
}

// HasLargeIcon returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasLargeIcon() bool {
	if o != nil && o.LargeIcon != nil {
		return true
	}

	return false
}

// SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.
func (o *MicrosoftGraphManagedMobileLobApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent) {
	o.LargeIcon = &v
}

// SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LargeIcon value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetLargeIconExplicitNull(b bool) {
	o.LargeIcon = nil
	o.isExplicitNullLargeIcon = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphManagedMobileLobApp) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphManagedMobileLobApp) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetIsFeatured() bool {
	if o == nil || o.IsFeatured == nil {
		var ret bool
		return ret
	}
	return *o.IsFeatured
}

// GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetIsFeaturedOk() (bool, bool) {
	if o == nil || o.IsFeatured == nil {
		var ret bool
		return ret, false
	}
	return *o.IsFeatured, true
}

// HasIsFeatured returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasIsFeatured() bool {
	if o != nil && o.IsFeatured != nil {
		return true
	}

	return false
}

// SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.
func (o *MicrosoftGraphManagedMobileLobApp) SetIsFeatured(v bool) {
	o.IsFeatured = &v
}

// GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetPrivacyInformationUrl() string {
	if o == nil || o.PrivacyInformationUrl == nil {
		var ret string
		return ret
	}
	return *o.PrivacyInformationUrl
}

// GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetPrivacyInformationUrlOk() (string, bool) {
	if o == nil || o.PrivacyInformationUrl == nil {
		var ret string
		return ret, false
	}
	return *o.PrivacyInformationUrl, true
}

// HasPrivacyInformationUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasPrivacyInformationUrl() bool {
	if o != nil && o.PrivacyInformationUrl != nil {
		return true
	}

	return false
}

// SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.
func (o *MicrosoftGraphManagedMobileLobApp) SetPrivacyInformationUrl(v string) {
	o.PrivacyInformationUrl = &v
}

// SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PrivacyInformationUrl value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetPrivacyInformationUrlExplicitNull(b bool) {
	o.PrivacyInformationUrl = nil
	o.isExplicitNullPrivacyInformationUrl = b
}
// GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetInformationUrl() string {
	if o == nil || o.InformationUrl == nil {
		var ret string
		return ret
	}
	return *o.InformationUrl
}

// GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetInformationUrlOk() (string, bool) {
	if o == nil || o.InformationUrl == nil {
		var ret string
		return ret, false
	}
	return *o.InformationUrl, true
}

// HasInformationUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasInformationUrl() bool {
	if o != nil && o.InformationUrl != nil {
		return true
	}

	return false
}

// SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.
func (o *MicrosoftGraphManagedMobileLobApp) SetInformationUrl(v string) {
	o.InformationUrl = &v
}

// SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InformationUrl value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetInformationUrlExplicitNull(b bool) {
	o.InformationUrl = nil
	o.isExplicitNullInformationUrl = b
}
// GetOwner returns the Owner field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetOwnerOk() (string, bool) {
	if o == nil || o.Owner == nil {
		var ret string
		return ret, false
	}
	return *o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *MicrosoftGraphManagedMobileLobApp) SetOwner(v string) {
	o.Owner = &v
}

// SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Owner value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetOwnerExplicitNull(b bool) {
	o.Owner = nil
	o.isExplicitNullOwner = b
}
// GetDeveloper returns the Developer field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetDeveloper() string {
	if o == nil || o.Developer == nil {
		var ret string
		return ret
	}
	return *o.Developer
}

// GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetDeveloperOk() (string, bool) {
	if o == nil || o.Developer == nil {
		var ret string
		return ret, false
	}
	return *o.Developer, true
}

// HasDeveloper returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasDeveloper() bool {
	if o != nil && o.Developer != nil {
		return true
	}

	return false
}

// SetDeveloper gets a reference to the given string and assigns it to the Developer field.
func (o *MicrosoftGraphManagedMobileLobApp) SetDeveloper(v string) {
	o.Developer = &v
}

// SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Developer value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetDeveloperExplicitNull(b bool) {
	o.Developer = nil
	o.isExplicitNullDeveloper = b
}
// GetNotes returns the Notes field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetNotesOk() (string, bool) {
	if o == nil || o.Notes == nil {
		var ret string
		return ret, false
	}
	return *o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *MicrosoftGraphManagedMobileLobApp) SetNotes(v string) {
	o.Notes = &v
}

// SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Notes value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetNotesExplicitNull(b bool) {
	o.Notes = nil
	o.isExplicitNullNotes = b
}
// GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState {
	if o == nil || o.PublishingState == nil {
		var ret AnyOfmicrosoftGraphMobileAppPublishingState
		return ret
	}
	return *o.PublishingState
}

// GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool) {
	if o == nil || o.PublishingState == nil {
		var ret AnyOfmicrosoftGraphMobileAppPublishingState
		return ret, false
	}
	return *o.PublishingState, true
}

// HasPublishingState returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasPublishingState() bool {
	if o != nil && o.PublishingState != nil {
		return true
	}

	return false
}

// SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.
func (o *MicrosoftGraphManagedMobileLobApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState) {
	o.PublishingState = &v
}

// GetCategories returns the Categories field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetCategories() []MicrosoftGraphMobileAppCategory {
	if o == nil || o.Categories == nil {
		var ret []MicrosoftGraphMobileAppCategory
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool) {
	if o == nil || o.Categories == nil {
		var ret []MicrosoftGraphMobileAppCategory
		return ret, false
	}
	return *o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.
func (o *MicrosoftGraphManagedMobileLobApp) SetCategories(v []MicrosoftGraphMobileAppCategory) {
	o.Categories = &v
}

// GetAssignments returns the Assignments field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetAssignments() []MicrosoftGraphMobileAppAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphMobileAppAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool) {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphMobileAppAssignment
		return ret, false
	}
	return *o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.
func (o *MicrosoftGraphManagedMobileLobApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment) {
	o.Assignments = &v
}

// GetAppAvailability returns the AppAvailability field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetAppAvailability() AnyOfmicrosoftGraphManagedAppAvailability {
	if o == nil || o.AppAvailability == nil {
		var ret AnyOfmicrosoftGraphManagedAppAvailability
		return ret
	}
	return *o.AppAvailability
}

// GetAppAvailabilityOk returns a tuple with the AppAvailability field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetAppAvailabilityOk() (AnyOfmicrosoftGraphManagedAppAvailability, bool) {
	if o == nil || o.AppAvailability == nil {
		var ret AnyOfmicrosoftGraphManagedAppAvailability
		return ret, false
	}
	return *o.AppAvailability, true
}

// HasAppAvailability returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasAppAvailability() bool {
	if o != nil && o.AppAvailability != nil {
		return true
	}

	return false
}

// SetAppAvailability gets a reference to the given AnyOfmicrosoftGraphManagedAppAvailability and assigns it to the AppAvailability field.
func (o *MicrosoftGraphManagedMobileLobApp) SetAppAvailability(v AnyOfmicrosoftGraphManagedAppAvailability) {
	o.AppAvailability = &v
}

// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetVersionOk() (string, bool) {
	if o == nil || o.Version == nil {
		var ret string
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MicrosoftGraphManagedMobileLobApp) SetVersion(v string) {
	o.Version = &v
}

// SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Version value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetVersionExplicitNull(b bool) {
	o.Version = nil
	o.isExplicitNullVersion = b
}
// GetCommittedContentVersion returns the CommittedContentVersion field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetCommittedContentVersion() string {
	if o == nil || o.CommittedContentVersion == nil {
		var ret string
		return ret
	}
	return *o.CommittedContentVersion
}

// GetCommittedContentVersionOk returns a tuple with the CommittedContentVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetCommittedContentVersionOk() (string, bool) {
	if o == nil || o.CommittedContentVersion == nil {
		var ret string
		return ret, false
	}
	return *o.CommittedContentVersion, true
}

// HasCommittedContentVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasCommittedContentVersion() bool {
	if o != nil && o.CommittedContentVersion != nil {
		return true
	}

	return false
}

// SetCommittedContentVersion gets a reference to the given string and assigns it to the CommittedContentVersion field.
func (o *MicrosoftGraphManagedMobileLobApp) SetCommittedContentVersion(v string) {
	o.CommittedContentVersion = &v
}

// SetCommittedContentVersionExplicitNull (un)sets CommittedContentVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CommittedContentVersion value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetCommittedContentVersionExplicitNull(b bool) {
	o.CommittedContentVersion = nil
	o.isExplicitNullCommittedContentVersion = b
}
// GetFileName returns the FileName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetFileNameOk() (string, bool) {
	if o == nil || o.FileName == nil {
		var ret string
		return ret, false
	}
	return *o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *MicrosoftGraphManagedMobileLobApp) SetFileName(v string) {
	o.FileName = &v
}

// SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FileName value is set to nil even if false is passed
func (o *MicrosoftGraphManagedMobileLobApp) SetFileNameExplicitNull(b bool) {
	o.FileName = nil
	o.isExplicitNullFileName = b
}
// GetSize returns the Size field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetSizeOk() (int64, bool) {
	if o == nil || o.Size == nil {
		var ret int64
		return ret, false
	}
	return *o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *MicrosoftGraphManagedMobileLobApp) SetSize(v int64) {
	o.Size = &v
}

// GetContentVersions returns the ContentVersions field if non-nil, zero value otherwise.
func (o *MicrosoftGraphManagedMobileLobApp) GetContentVersions() []MicrosoftGraphMobileAppContent {
	if o == nil || o.ContentVersions == nil {
		var ret []MicrosoftGraphMobileAppContent
		return ret
	}
	return *o.ContentVersions
}

// GetContentVersionsOk returns a tuple with the ContentVersions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphManagedMobileLobApp) GetContentVersionsOk() ([]MicrosoftGraphMobileAppContent, bool) {
	if o == nil || o.ContentVersions == nil {
		var ret []MicrosoftGraphMobileAppContent
		return ret, false
	}
	return *o.ContentVersions, true
}

// HasContentVersions returns a boolean if a field has been set.
func (o *MicrosoftGraphManagedMobileLobApp) HasContentVersions() bool {
	if o != nil && o.ContentVersions != nil {
		return true
	}

	return false
}

// SetContentVersions gets a reference to the given []MicrosoftGraphMobileAppContent and assigns it to the ContentVersions field.
func (o *MicrosoftGraphManagedMobileLobApp) SetContentVersions(v []MicrosoftGraphMobileAppContent) {
	o.ContentVersions = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphManagedMobileLobApp) MarshalJSON() ([]byte, error) {
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
	if o.AppAvailability != nil {
		toSerialize["appAvailability"] = o.AppAvailability
	}
	if o.Version == nil {
		if o.isExplicitNullVersion {
			toSerialize["version"] = o.Version
		}
	} else {
		toSerialize["version"] = o.Version
	}
	if o.CommittedContentVersion == nil {
		if o.isExplicitNullCommittedContentVersion {
			toSerialize["committedContentVersion"] = o.CommittedContentVersion
		}
	} else {
		toSerialize["committedContentVersion"] = o.CommittedContentVersion
	}
	if o.FileName == nil {
		if o.isExplicitNullFileName {
			toSerialize["fileName"] = o.FileName
		}
	} else {
		toSerialize["fileName"] = o.FileName
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.ContentVersions != nil {
		toSerialize["contentVersions"] = o.ContentVersions
	}
	return json.Marshal(toSerialize)
}

