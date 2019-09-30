# MicrosoftGraphAndroidStoreApp

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
**PackageId** | Pointer to **string** | The package identifier. | [optional] 
**AppStoreUrl** | Pointer to **string** | The Android app store URL. | [optional] 
**MinimumSupportedOperatingSystem** | Pointer to [**AnyOfmicrosoftGraphAndroidMinimumOperatingSystem**](anyOf&lt;microsoft.graph.androidMinimumOperatingSystem&gt;.md) | The value for the minimum applicable operating system. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAndroidStoreApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidStoreApp) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAndroidStoreApp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAndroidStoreApp) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphAndroidStoreApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAndroidStoreApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphAndroidStoreApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphAndroidStoreApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphAndroidStoreApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAndroidStoreApp) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphAndroidStoreApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphAndroidStoreApp) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphAndroidStoreApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphAndroidStoreApp) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphAndroidStoreApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphAndroidStoreApp) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphAndroidStoreApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphAndroidStoreApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphAndroidStoreApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphAndroidStoreApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidStoreApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidStoreApp) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidStoreApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidStoreApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidStoreApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAndroidStoreApp) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAndroidStoreApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidStoreApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphAndroidStoreApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphAndroidStoreApp) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphAndroidStoreApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphAndroidStoreApp) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphAndroidStoreApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphAndroidStoreApp) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphAndroidStoreApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphAndroidStoreApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphAndroidStoreApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphAndroidStoreApp) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphAndroidStoreApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphAndroidStoreApp) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphAndroidStoreApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphAndroidStoreApp) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphAndroidStoreApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphAndroidStoreApp) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphAndroidStoreApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphAndroidStoreApp) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphAndroidStoreApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphAndroidStoreApp) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphAndroidStoreApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphAndroidStoreApp) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphAndroidStoreApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphAndroidStoreApp) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphAndroidStoreApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphAndroidStoreApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphAndroidStoreApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphAndroidStoreApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphAndroidStoreApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphAndroidStoreApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphAndroidStoreApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphAndroidStoreApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphAndroidStoreApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphAndroidStoreApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphAndroidStoreApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphAndroidStoreApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetPackageId

`func (o *MicrosoftGraphAndroidStoreApp) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *MicrosoftGraphAndroidStoreApp) GetPackageIdOk() (string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackageId

`func (o *MicrosoftGraphAndroidStoreApp) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageId

`func (o *MicrosoftGraphAndroidStoreApp) SetPackageId(v string)`

SetPackageId gets a reference to the given string and assigns it to the PackageId field.

### SetPackageIdExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetPackageIdExplicitNull(b bool)`

SetPackageIdExplicitNull (un)sets PackageId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PackageId value is set to nil even if false is passed
### GetAppStoreUrl

`func (o *MicrosoftGraphAndroidStoreApp) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *MicrosoftGraphAndroidStoreApp) GetAppStoreUrlOk() (string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreUrl

`func (o *MicrosoftGraphAndroidStoreApp) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrl

`func (o *MicrosoftGraphAndroidStoreApp) SetAppStoreUrl(v string)`

SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.

### SetAppStoreUrlExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetAppStoreUrlExplicitNull(b bool)`

SetAppStoreUrlExplicitNull (un)sets AppStoreUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppStoreUrl value is set to nil even if false is passed
### GetMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphAndroidStoreApp) GetMinimumSupportedOperatingSystem() AnyOfmicrosoftGraphAndroidMinimumOperatingSystem`

GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumSupportedOperatingSystemOk

`func (o *MicrosoftGraphAndroidStoreApp) GetMinimumSupportedOperatingSystemOk() (AnyOfmicrosoftGraphAndroidMinimumOperatingSystem, bool)`

GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphAndroidStoreApp) HasMinimumSupportedOperatingSystem() bool`

HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.

### SetMinimumSupportedOperatingSystem

`func (o *MicrosoftGraphAndroidStoreApp) SetMinimumSupportedOperatingSystem(v AnyOfmicrosoftGraphAndroidMinimumOperatingSystem)`

SetMinimumSupportedOperatingSystem gets a reference to the given AnyOfmicrosoftGraphAndroidMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.

### SetMinimumSupportedOperatingSystemExplicitNull

`func (o *MicrosoftGraphAndroidStoreApp) SetMinimumSupportedOperatingSystemExplicitNull(b bool)`

SetMinimumSupportedOperatingSystemExplicitNull (un)sets MinimumSupportedOperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumSupportedOperatingSystem value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


