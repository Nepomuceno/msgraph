# MicrosoftGraphManagedAndroidStoreApp

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
**PackageId** | Pointer to **string** | The app&#39;s package ID. | [optional] 
**AppStoreUrl** | Pointer to **string** | The Android AppStoreUrl. | [optional] 
**MinimumSupportedOperatingSystem** | Pointer to [**MicrosoftGraphAndroidMinimumOperatingSystem**](microsoft.graph.androidMinimumOperatingSystem.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetAppAvailability

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetAppAvailability() AnyOfmicrosoftGraphManagedAppAvailability`

GetAppAvailability returns the AppAvailability field if non-nil, zero value otherwise.

### GetAppAvailabilityOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetAppAvailabilityOk() (AnyOfmicrosoftGraphManagedAppAvailability, bool)`

GetAppAvailabilityOk returns a tuple with the AppAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppAvailability

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasAppAvailability() bool`

HasAppAvailability returns a boolean if a field has been set.

### SetAppAvailability

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetAppAvailability(v AnyOfmicrosoftGraphManagedAppAvailability)`

SetAppAvailability gets a reference to the given AnyOfmicrosoftGraphManagedAppAvailability and assigns it to the AppAvailability field.

### GetVersion

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetPackageId

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetPackageIdOk() (string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackageId

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageId

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPackageId(v string)`

SetPackageId gets a reference to the given string and assigns it to the PackageId field.

### SetPackageIdExplicitNull

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetPackageIdExplicitNull(b bool)`

SetPackageIdExplicitNull (un)sets PackageId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PackageId value is set to nil even if false is passed
### GetAppStoreUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetAppStoreUrlOk() (string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrl

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetAppStoreUrl(v string)`

SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.

### GetMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetMinimumSupportedOperatingSystem() MicrosoftGraphAndroidMinimumOperatingSystem`

GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumSupportedOperatingSystemOk

`func (o *MicrosoftGraphManagedAndroidStoreApp) GetMinimumSupportedOperatingSystemOk() (MicrosoftGraphAndroidMinimumOperatingSystem, bool)`

GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphManagedAndroidStoreApp) HasMinimumSupportedOperatingSystem() bool`

HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.

### SetMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphManagedAndroidStoreApp) SetMinimumSupportedOperatingSystem(v MicrosoftGraphAndroidMinimumOperatingSystem)`

SetMinimumSupportedOperatingSystem gets a reference to the given MicrosoftGraphAndroidMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


