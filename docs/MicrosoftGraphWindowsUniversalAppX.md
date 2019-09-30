# MicrosoftGraphWindowsUniversalAppX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | The admin provided or imported title of the app. | [optional] 
**Description** | Pointer to **string** | The description of the app. | [optional] 
**Publisher** | Pointer to **string** | The publisher of the app. | [optional] 
**LargeIcon** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | The large icon, to be displayed in the app details and used for upload of the icon. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time the app was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time the app was last modified. | [optional] 
**IsFeatured** | Pointer to **bool** | The value indicating whether the app is marked as featured by the admin. | [optional] 
**PrivacyInformationUrl** | Pointer to **string** | The privacy statement Url. | [optional] 
**InformationUrl** | Pointer to **string** | The more information Url. | [optional] 
**Owner** | Pointer to **string** | The owner of the app. | [optional] 
**Developer** | Pointer to **string** | The developer of the app. | [optional] 
**Notes** | Pointer to **string** | Notes for the app. | [optional] 
**PublishingState** | Pointer to [**AnyOfmicrosoftGraphMobileAppPublishingState**](anyOf&lt;microsoft.graph.mobileAppPublishingState&gt;.md) | The publishing state for the app. The app cannot be assigned unless the app is published. | [optional] 
**Categories** | Pointer to [**[]MicrosoftGraphMobileAppCategory**](microsoft.graph.mobileAppCategory.md) |  | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphMobileAppAssignment**](microsoft.graph.mobileAppAssignment.md) |  | [optional] 
**CommittedContentVersion** | Pointer to **string** | The internal committed content version. | [optional] 
**FileName** | Pointer to **string** | The name of the main Lob application file. | [optional] 
**Size** | Pointer to **int64** | The total size, including all uploaded files. | [optional] 
**ContentVersions** | Pointer to [**[]MicrosoftGraphMobileAppContent**](microsoft.graph.mobileAppContent.md) |  | [optional] 
**ApplicableArchitectures** | Pointer to [**AnyOfmicrosoftGraphWindowsArchitecture**](anyOf&lt;microsoft.graph.windowsArchitecture&gt;.md) | The Windows architecture(s) for which this app can run on. | [optional] 
**ApplicableDeviceTypes** | Pointer to [**AnyOfmicrosoftGraphWindowsDeviceType**](anyOf&lt;microsoft.graph.windowsDeviceType&gt;.md) | The Windows device type(s) for which this app can run on. | [optional] 
**IdentityName** | Pointer to **string** | The Identity Name. | [optional] 
**IdentityPublisherHash** | Pointer to **string** | The Identity Publisher Hash. | [optional] 
**IdentityResourceIdentifier** | Pointer to **string** | The Identity Resource Identifier. | [optional] 
**IsBundle** | Pointer to **bool** | Whether or not the app is a bundle. | [optional] 
**MinimumSupportedOperatingSystem** | Pointer to [**MicrosoftGraphWindowsMinimumOperatingSystem**](microsoft.graph.windowsMinimumOperatingSystem.md) |  | [optional] 
**IdentityVersion** | Pointer to **string** | The identity version. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsUniversalAppX) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsUniversalAppX) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsUniversalAppX) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphWindowsUniversalAppX) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsUniversalAppX) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsUniversalAppX) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphWindowsUniversalAppX) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsUniversalAppX) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsUniversalAppX) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphWindowsUniversalAppX) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphWindowsUniversalAppX) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphWindowsUniversalAppX) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphWindowsUniversalAppX) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphWindowsUniversalAppX) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphWindowsUniversalAppX) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsUniversalAppX) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsUniversalAppX) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsUniversalAppX) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsUniversalAppX) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindowsUniversalAppX) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsUniversalAppX) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphWindowsUniversalAppX) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphWindowsUniversalAppX) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphWindowsUniversalAppX) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphWindowsUniversalAppX) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphWindowsUniversalAppX) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphWindowsUniversalAppX) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphWindowsUniversalAppX) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphWindowsUniversalAppX) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphWindowsUniversalAppX) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphWindowsUniversalAppX) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphWindowsUniversalAppX) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphWindowsUniversalAppX) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphWindowsUniversalAppX) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphWindowsUniversalAppX) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphWindowsUniversalAppX) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphWindowsUniversalAppX) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphWindowsUniversalAppX) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphWindowsUniversalAppX) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphWindowsUniversalAppX) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphWindowsUniversalAppX) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphWindowsUniversalAppX) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphWindowsUniversalAppX) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphWindowsUniversalAppX) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindowsUniversalAppX) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindowsUniversalAppX) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetCommittedContentVersion

`func (o *MicrosoftGraphWindowsUniversalAppX) GetCommittedContentVersion() string`

GetCommittedContentVersion returns the CommittedContentVersion field if non-nil, zero value otherwise.

### GetCommittedContentVersionOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetCommittedContentVersionOk() (string, bool)`

GetCommittedContentVersionOk returns a tuple with the CommittedContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommittedContentVersion

`func (o *MicrosoftGraphWindowsUniversalAppX) HasCommittedContentVersion() bool`

HasCommittedContentVersion returns a boolean if a field has been set.

### SetCommittedContentVersion

`func (o *MicrosoftGraphWindowsUniversalAppX) SetCommittedContentVersion(v string)`

SetCommittedContentVersion gets a reference to the given string and assigns it to the CommittedContentVersion field.

### SetCommittedContentVersionExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetCommittedContentVersionExplicitNull(b bool)`

SetCommittedContentVersionExplicitNull (un)sets CommittedContentVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommittedContentVersion value is set to nil even if false is passed
### GetFileName

`func (o *MicrosoftGraphWindowsUniversalAppX) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetFileNameOk() (string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileName

`func (o *MicrosoftGraphWindowsUniversalAppX) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileName

`func (o *MicrosoftGraphWindowsUniversalAppX) SetFileName(v string)`

SetFileName gets a reference to the given string and assigns it to the FileName field.

### SetFileNameExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetFileNameExplicitNull(b bool)`

SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileName value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphWindowsUniversalAppX) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphWindowsUniversalAppX) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphWindowsUniversalAppX) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### GetContentVersions

`func (o *MicrosoftGraphWindowsUniversalAppX) GetContentVersions() []MicrosoftGraphMobileAppContent`

GetContentVersions returns the ContentVersions field if non-nil, zero value otherwise.

### GetContentVersionsOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetContentVersionsOk() ([]MicrosoftGraphMobileAppContent, bool)`

GetContentVersionsOk returns a tuple with the ContentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentVersions

`func (o *MicrosoftGraphWindowsUniversalAppX) HasContentVersions() bool`

HasContentVersions returns a boolean if a field has been set.

### SetContentVersions

`func (o *MicrosoftGraphWindowsUniversalAppX) SetContentVersions(v []MicrosoftGraphMobileAppContent)`

SetContentVersions gets a reference to the given []MicrosoftGraphMobileAppContent and assigns it to the ContentVersions field.

### GetApplicableArchitectures

`func (o *MicrosoftGraphWindowsUniversalAppX) GetApplicableArchitectures() AnyOfmicrosoftGraphWindowsArchitecture`

GetApplicableArchitectures returns the ApplicableArchitectures field if non-nil, zero value otherwise.

### GetApplicableArchitecturesOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetApplicableArchitecturesOk() (AnyOfmicrosoftGraphWindowsArchitecture, bool)`

GetApplicableArchitecturesOk returns a tuple with the ApplicableArchitectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableArchitectures

`func (o *MicrosoftGraphWindowsUniversalAppX) HasApplicableArchitectures() bool`

HasApplicableArchitectures returns a boolean if a field has been set.

### SetApplicableArchitectures

`func (o *MicrosoftGraphWindowsUniversalAppX) SetApplicableArchitectures(v AnyOfmicrosoftGraphWindowsArchitecture)`

SetApplicableArchitectures gets a reference to the given AnyOfmicrosoftGraphWindowsArchitecture and assigns it to the ApplicableArchitectures field.

### GetApplicableDeviceTypes

`func (o *MicrosoftGraphWindowsUniversalAppX) GetApplicableDeviceTypes() AnyOfmicrosoftGraphWindowsDeviceType`

GetApplicableDeviceTypes returns the ApplicableDeviceTypes field if non-nil, zero value otherwise.

### GetApplicableDeviceTypesOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetApplicableDeviceTypesOk() (AnyOfmicrosoftGraphWindowsDeviceType, bool)`

GetApplicableDeviceTypesOk returns a tuple with the ApplicableDeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableDeviceTypes

`func (o *MicrosoftGraphWindowsUniversalAppX) HasApplicableDeviceTypes() bool`

HasApplicableDeviceTypes returns a boolean if a field has been set.

### SetApplicableDeviceTypes

`func (o *MicrosoftGraphWindowsUniversalAppX) SetApplicableDeviceTypes(v AnyOfmicrosoftGraphWindowsDeviceType)`

SetApplicableDeviceTypes gets a reference to the given AnyOfmicrosoftGraphWindowsDeviceType and assigns it to the ApplicableDeviceTypes field.

### GetIdentityName

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityName() string`

GetIdentityName returns the IdentityName field if non-nil, zero value otherwise.

### GetIdentityNameOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityNameOk() (string, bool)`

GetIdentityNameOk returns a tuple with the IdentityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityName

`func (o *MicrosoftGraphWindowsUniversalAppX) HasIdentityName() bool`

HasIdentityName returns a boolean if a field has been set.

### SetIdentityName

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityName(v string)`

SetIdentityName gets a reference to the given string and assigns it to the IdentityName field.

### SetIdentityNameExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityNameExplicitNull(b bool)`

SetIdentityNameExplicitNull (un)sets IdentityName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdentityName value is set to nil even if false is passed
### GetIdentityPublisherHash

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityPublisherHash() string`

GetIdentityPublisherHash returns the IdentityPublisherHash field if non-nil, zero value otherwise.

### GetIdentityPublisherHashOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityPublisherHashOk() (string, bool)`

GetIdentityPublisherHashOk returns a tuple with the IdentityPublisherHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityPublisherHash

`func (o *MicrosoftGraphWindowsUniversalAppX) HasIdentityPublisherHash() bool`

HasIdentityPublisherHash returns a boolean if a field has been set.

### SetIdentityPublisherHash

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityPublisherHash(v string)`

SetIdentityPublisherHash gets a reference to the given string and assigns it to the IdentityPublisherHash field.

### GetIdentityResourceIdentifier

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityResourceIdentifier() string`

GetIdentityResourceIdentifier returns the IdentityResourceIdentifier field if non-nil, zero value otherwise.

### GetIdentityResourceIdentifierOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityResourceIdentifierOk() (string, bool)`

GetIdentityResourceIdentifierOk returns a tuple with the IdentityResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityResourceIdentifier

`func (o *MicrosoftGraphWindowsUniversalAppX) HasIdentityResourceIdentifier() bool`

HasIdentityResourceIdentifier returns a boolean if a field has been set.

### SetIdentityResourceIdentifier

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityResourceIdentifier(v string)`

SetIdentityResourceIdentifier gets a reference to the given string and assigns it to the IdentityResourceIdentifier field.

### SetIdentityResourceIdentifierExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityResourceIdentifierExplicitNull(b bool)`

SetIdentityResourceIdentifierExplicitNull (un)sets IdentityResourceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdentityResourceIdentifier value is set to nil even if false is passed
### GetIsBundle

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIsBundle() bool`

GetIsBundle returns the IsBundle field if non-nil, zero value otherwise.

### GetIsBundleOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIsBundleOk() (bool, bool)`

GetIsBundleOk returns a tuple with the IsBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsBundle

`func (o *MicrosoftGraphWindowsUniversalAppX) HasIsBundle() bool`

HasIsBundle returns a boolean if a field has been set.

### SetIsBundle

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIsBundle(v bool)`

SetIsBundle gets a reference to the given bool and assigns it to the IsBundle field.

### GetMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphWindowsUniversalAppX) GetMinimumSupportedOperatingSystem() MicrosoftGraphWindowsMinimumOperatingSystem`

GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumSupportedOperatingSystemOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetMinimumSupportedOperatingSystemOk() (MicrosoftGraphWindowsMinimumOperatingSystem, bool)`

GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphWindowsUniversalAppX) HasMinimumSupportedOperatingSystem() bool`

HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.

### SetMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphWindowsUniversalAppX) SetMinimumSupportedOperatingSystem(v MicrosoftGraphWindowsMinimumOperatingSystem)`

SetMinimumSupportedOperatingSystem gets a reference to the given MicrosoftGraphWindowsMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.

### GetIdentityVersion

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityVersion() string`

GetIdentityVersion returns the IdentityVersion field if non-nil, zero value otherwise.

### GetIdentityVersionOk

`func (o *MicrosoftGraphWindowsUniversalAppX) GetIdentityVersionOk() (string, bool)`

GetIdentityVersionOk returns a tuple with the IdentityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityVersion

`func (o *MicrosoftGraphWindowsUniversalAppX) HasIdentityVersion() bool`

HasIdentityVersion returns a boolean if a field has been set.

### SetIdentityVersion

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityVersion(v string)`

SetIdentityVersion gets a reference to the given string and assigns it to the IdentityVersion field.

### SetIdentityVersionExplicitNull

`func (o *MicrosoftGraphWindowsUniversalAppX) SetIdentityVersionExplicitNull(b bool)`

SetIdentityVersionExplicitNull (un)sets IdentityVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdentityVersion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


