# MicrosoftGraphManagedMobileLobApp

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
**AppAvailability** | Pointer to [**AnyOfmicrosoftGraphManagedAppAvailability**](anyOf&lt;microsoft.graph.managedAppAvailability&gt;.md) | The Application&#39;s availability. | [optional] 
**Version** | Pointer to **string** | The Application&#39;s version. | [optional] 
**CommittedContentVersion** | Pointer to **string** | The internal committed content version. | [optional] 
**FileName** | Pointer to **string** | The name of the main Lob application file. | [optional] 
**Size** | Pointer to **int64** | The total size, including all uploaded files. | [optional] 
**ContentVersions** | Pointer to [**[]MicrosoftGraphMobileAppContent**](microsoft.graph.mobileAppContent.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedMobileLobApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedMobileLobApp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedMobileLobApp) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphManagedMobileLobApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphManagedMobileLobApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedMobileLobApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphManagedMobileLobApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphManagedMobileLobApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphManagedMobileLobApp) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphManagedMobileLobApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphManagedMobileLobApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphManagedMobileLobApp) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphManagedMobileLobApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphManagedMobileLobApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphManagedMobileLobApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphManagedMobileLobApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedMobileLobApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedMobileLobApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedMobileLobApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedMobileLobApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedMobileLobApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphManagedMobileLobApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphManagedMobileLobApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphManagedMobileLobApp) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedMobileLobApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphManagedMobileLobApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedMobileLobApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphManagedMobileLobApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphManagedMobileLobApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphManagedMobileLobApp) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphManagedMobileLobApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphManagedMobileLobApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphManagedMobileLobApp) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphManagedMobileLobApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphManagedMobileLobApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphManagedMobileLobApp) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphManagedMobileLobApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphManagedMobileLobApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphManagedMobileLobApp) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphManagedMobileLobApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphManagedMobileLobApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphManagedMobileLobApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphManagedMobileLobApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphManagedMobileLobApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphManagedMobileLobApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphManagedMobileLobApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphManagedMobileLobApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphManagedMobileLobApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetAppAvailability

`func (o *MicrosoftGraphManagedMobileLobApp) GetAppAvailability() AnyOfmicrosoftGraphManagedAppAvailability`

GetAppAvailability returns the AppAvailability field if non-nil, zero value otherwise.

### GetAppAvailabilityOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetAppAvailabilityOk() (AnyOfmicrosoftGraphManagedAppAvailability, bool)`

GetAppAvailabilityOk returns a tuple with the AppAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppAvailability

`func (o *MicrosoftGraphManagedMobileLobApp) HasAppAvailability() bool`

HasAppAvailability returns a boolean if a field has been set.

### SetAppAvailability

`func (o *MicrosoftGraphManagedMobileLobApp) SetAppAvailability(v AnyOfmicrosoftGraphManagedAppAvailability)`

SetAppAvailability gets a reference to the given AnyOfmicrosoftGraphManagedAppAvailability and assigns it to the AppAvailability field.

### GetVersion

`func (o *MicrosoftGraphManagedMobileLobApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphManagedMobileLobApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphManagedMobileLobApp) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetCommittedContentVersion

`func (o *MicrosoftGraphManagedMobileLobApp) GetCommittedContentVersion() string`

GetCommittedContentVersion returns the CommittedContentVersion field if non-nil, zero value otherwise.

### GetCommittedContentVersionOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetCommittedContentVersionOk() (string, bool)`

GetCommittedContentVersionOk returns a tuple with the CommittedContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommittedContentVersion

`func (o *MicrosoftGraphManagedMobileLobApp) HasCommittedContentVersion() bool`

HasCommittedContentVersion returns a boolean if a field has been set.

### SetCommittedContentVersion

`func (o *MicrosoftGraphManagedMobileLobApp) SetCommittedContentVersion(v string)`

SetCommittedContentVersion gets a reference to the given string and assigns it to the CommittedContentVersion field.

### SetCommittedContentVersionExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetCommittedContentVersionExplicitNull(b bool)`

SetCommittedContentVersionExplicitNull (un)sets CommittedContentVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommittedContentVersion value is set to nil even if false is passed
### GetFileName

`func (o *MicrosoftGraphManagedMobileLobApp) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetFileNameOk() (string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileName

`func (o *MicrosoftGraphManagedMobileLobApp) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileName

`func (o *MicrosoftGraphManagedMobileLobApp) SetFileName(v string)`

SetFileName gets a reference to the given string and assigns it to the FileName field.

### SetFileNameExplicitNull

`func (o *MicrosoftGraphManagedMobileLobApp) SetFileNameExplicitNull(b bool)`

SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileName value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphManagedMobileLobApp) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphManagedMobileLobApp) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphManagedMobileLobApp) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### GetContentVersions

`func (o *MicrosoftGraphManagedMobileLobApp) GetContentVersions() []MicrosoftGraphMobileAppContent`

GetContentVersions returns the ContentVersions field if non-nil, zero value otherwise.

### GetContentVersionsOk

`func (o *MicrosoftGraphManagedMobileLobApp) GetContentVersionsOk() ([]MicrosoftGraphMobileAppContent, bool)`

GetContentVersionsOk returns a tuple with the ContentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentVersions

`func (o *MicrosoftGraphManagedMobileLobApp) HasContentVersions() bool`

HasContentVersions returns a boolean if a field has been set.

### SetContentVersions

`func (o *MicrosoftGraphManagedMobileLobApp) SetContentVersions(v []MicrosoftGraphMobileAppContent)`

SetContentVersions gets a reference to the given []MicrosoftGraphMobileAppContent and assigns it to the ContentVersions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


