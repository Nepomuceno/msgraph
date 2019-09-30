# MicrosoftGraphMicrosoftStoreForBusinessApp

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
**UsedLicenseCount** | Pointer to **int32** | The number of Microsoft Store for Business licenses in use. | [optional] 
**TotalLicenseCount** | Pointer to **int32** | The total number of Microsoft Store for Business licenses. | [optional] 
**ProductKey** | Pointer to **string** | The app product key | [optional] 
**LicenseType** | Pointer to [**AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType**](anyOf&lt;microsoft.graph.microsoftStoreForBusinessLicenseType&gt;.md) | The app license type | [optional] 
**PackageIdentityName** | Pointer to **string** | The app package identifier | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetUsedLicenseCount

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetUsedLicenseCount() int32`

GetUsedLicenseCount returns the UsedLicenseCount field if non-nil, zero value otherwise.

### GetUsedLicenseCountOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetUsedLicenseCountOk() (int32, bool)`

GetUsedLicenseCountOk returns a tuple with the UsedLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsedLicenseCount

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasUsedLicenseCount() bool`

HasUsedLicenseCount returns a boolean if a field has been set.

### SetUsedLicenseCount

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetUsedLicenseCount(v int32)`

SetUsedLicenseCount gets a reference to the given int32 and assigns it to the UsedLicenseCount field.

### GetTotalLicenseCount

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetTotalLicenseCount() int32`

GetTotalLicenseCount returns the TotalLicenseCount field if non-nil, zero value otherwise.

### GetTotalLicenseCountOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetTotalLicenseCountOk() (int32, bool)`

GetTotalLicenseCountOk returns a tuple with the TotalLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalLicenseCount

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasTotalLicenseCount() bool`

HasTotalLicenseCount returns a boolean if a field has been set.

### SetTotalLicenseCount

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetTotalLicenseCount(v int32)`

SetTotalLicenseCount gets a reference to the given int32 and assigns it to the TotalLicenseCount field.

### GetProductKey

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetProductKeyOk() (string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductKey

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKey

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetProductKey(v string)`

SetProductKey gets a reference to the given string and assigns it to the ProductKey field.

### SetProductKeyExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetProductKeyExplicitNull(b bool)`

SetProductKeyExplicitNull (un)sets ProductKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductKey value is set to nil even if false is passed
### GetLicenseType

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetLicenseType() AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetLicenseTypeOk() (AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseType

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseType

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetLicenseType(v AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType)`

SetLicenseType gets a reference to the given AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType and assigns it to the LicenseType field.

### GetPackageIdentityName

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPackageIdentityName() string`

GetPackageIdentityName returns the PackageIdentityName field if non-nil, zero value otherwise.

### GetPackageIdentityNameOk

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) GetPackageIdentityNameOk() (string, bool)`

GetPackageIdentityNameOk returns a tuple with the PackageIdentityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackageIdentityName

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) HasPackageIdentityName() bool`

HasPackageIdentityName returns a boolean if a field has been set.

### SetPackageIdentityName

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPackageIdentityName(v string)`

SetPackageIdentityName gets a reference to the given string and assigns it to the PackageIdentityName field.

### SetPackageIdentityNameExplicitNull

`func (o *MicrosoftGraphMicrosoftStoreForBusinessApp) SetPackageIdentityNameExplicitNull(b bool)`

SetPackageIdentityNameExplicitNull (un)sets PackageIdentityName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PackageIdentityName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


