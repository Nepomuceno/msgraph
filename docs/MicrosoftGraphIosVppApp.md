# MicrosoftGraphIosVppApp

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
**UsedLicenseCount** | Pointer to **int32** | The number of VPP licenses in use. | [optional] 
**TotalLicenseCount** | Pointer to **int32** | The total number of VPP licenses. | [optional] 
**ReleaseDateTime** | Pointer to [**time.Time**](time.Time.md) | The VPP application release date and time. | [optional] 
**AppStoreUrl** | Pointer to **string** | The store URL. | [optional] 
**LicensingType** | Pointer to [**AnyOfmicrosoftGraphVppLicensingType**](anyOf&lt;microsoft.graph.vppLicensingType&gt;.md) | The supported License Type. | [optional] 
**ApplicableDeviceType** | Pointer to [**AnyOfmicrosoftGraphIosDeviceType**](anyOf&lt;microsoft.graph.iosDeviceType&gt;.md) | The applicable iOS Device Type. | [optional] 
**VppTokenOrganizationName** | Pointer to **string** | The organization associated with the Apple Volume Purchase Program Token | [optional] 
**VppTokenAccountType** | Pointer to [**AnyOfmicrosoftGraphVppTokenAccountType**](anyOf&lt;microsoft.graph.vppTokenAccountType&gt;.md) | The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: &#x60;business&#x60;, &#x60;education&#x60;. | [optional] 
**VppTokenAppleId** | Pointer to **string** | The Apple Id associated with the given Apple Volume Purchase Program Token. | [optional] 
**BundleId** | Pointer to **string** | The Identity Name. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphIosVppApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosVppApp) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosVppApp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosVppApp) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphIosVppApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosVppApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosVppApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosVppApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphIosVppApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosVppApp) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosVppApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosVppApp) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphIosVppApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphIosVppApp) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphIosVppApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphIosVppApp) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphIosVppApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphIosVppApp) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphIosVppApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphIosVppApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphIosVppApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosVppApp) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosVppApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosVppApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosVppApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosVppApp) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosVppApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosVppApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphIosVppApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphIosVppApp) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphIosVppApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphIosVppApp) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphIosVppApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphIosVppApp) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphIosVppApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphIosVppApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphIosVppApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphIosVppApp) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphIosVppApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphIosVppApp) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphIosVppApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphIosVppApp) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphIosVppApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphIosVppApp) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphIosVppApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphIosVppApp) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphIosVppApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphIosVppApp) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphIosVppApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphIosVppApp) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphIosVppApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphIosVppApp) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphIosVppApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphIosVppApp) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphIosVppApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphIosVppApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphIosVppApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphIosVppApp) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphIosVppApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphIosVppApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphIosVppApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosVppApp) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosVppApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosVppApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetUsedLicenseCount

`func (o *MicrosoftGraphIosVppApp) GetUsedLicenseCount() int32`

GetUsedLicenseCount returns the UsedLicenseCount field if non-nil, zero value otherwise.

### GetUsedLicenseCountOk

`func (o *MicrosoftGraphIosVppApp) GetUsedLicenseCountOk() (int32, bool)`

GetUsedLicenseCountOk returns a tuple with the UsedLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsedLicenseCount

`func (o *MicrosoftGraphIosVppApp) HasUsedLicenseCount() bool`

HasUsedLicenseCount returns a boolean if a field has been set.

### SetUsedLicenseCount

`func (o *MicrosoftGraphIosVppApp) SetUsedLicenseCount(v int32)`

SetUsedLicenseCount gets a reference to the given int32 and assigns it to the UsedLicenseCount field.

### GetTotalLicenseCount

`func (o *MicrosoftGraphIosVppApp) GetTotalLicenseCount() int32`

GetTotalLicenseCount returns the TotalLicenseCount field if non-nil, zero value otherwise.

### GetTotalLicenseCountOk

`func (o *MicrosoftGraphIosVppApp) GetTotalLicenseCountOk() (int32, bool)`

GetTotalLicenseCountOk returns a tuple with the TotalLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalLicenseCount

`func (o *MicrosoftGraphIosVppApp) HasTotalLicenseCount() bool`

HasTotalLicenseCount returns a boolean if a field has been set.

### SetTotalLicenseCount

`func (o *MicrosoftGraphIosVppApp) SetTotalLicenseCount(v int32)`

SetTotalLicenseCount gets a reference to the given int32 and assigns it to the TotalLicenseCount field.

### GetReleaseDateTime

`func (o *MicrosoftGraphIosVppApp) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *MicrosoftGraphIosVppApp) GetReleaseDateTimeOk() (time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReleaseDateTime

`func (o *MicrosoftGraphIosVppApp) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTime

`func (o *MicrosoftGraphIosVppApp) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime gets a reference to the given time.Time and assigns it to the ReleaseDateTime field.

### SetReleaseDateTimeExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetReleaseDateTimeExplicitNull(b bool)`

SetReleaseDateTimeExplicitNull (un)sets ReleaseDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReleaseDateTime value is set to nil even if false is passed
### GetAppStoreUrl

`func (o *MicrosoftGraphIosVppApp) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *MicrosoftGraphIosVppApp) GetAppStoreUrlOk() (string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreUrl

`func (o *MicrosoftGraphIosVppApp) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrl

`func (o *MicrosoftGraphIosVppApp) SetAppStoreUrl(v string)`

SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.

### SetAppStoreUrlExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetAppStoreUrlExplicitNull(b bool)`

SetAppStoreUrlExplicitNull (un)sets AppStoreUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppStoreUrl value is set to nil even if false is passed
### GetLicensingType

`func (o *MicrosoftGraphIosVppApp) GetLicensingType() AnyOfmicrosoftGraphVppLicensingType`

GetLicensingType returns the LicensingType field if non-nil, zero value otherwise.

### GetLicensingTypeOk

`func (o *MicrosoftGraphIosVppApp) GetLicensingTypeOk() (AnyOfmicrosoftGraphVppLicensingType, bool)`

GetLicensingTypeOk returns a tuple with the LicensingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicensingType

`func (o *MicrosoftGraphIosVppApp) HasLicensingType() bool`

HasLicensingType returns a boolean if a field has been set.

### SetLicensingType

`func (o *MicrosoftGraphIosVppApp) SetLicensingType(v AnyOfmicrosoftGraphVppLicensingType)`

SetLicensingType gets a reference to the given AnyOfmicrosoftGraphVppLicensingType and assigns it to the LicensingType field.

### SetLicensingTypeExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetLicensingTypeExplicitNull(b bool)`

SetLicensingTypeExplicitNull (un)sets LicensingType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LicensingType value is set to nil even if false is passed
### GetApplicableDeviceType

`func (o *MicrosoftGraphIosVppApp) GetApplicableDeviceType() AnyOfmicrosoftGraphIosDeviceType`

GetApplicableDeviceType returns the ApplicableDeviceType field if non-nil, zero value otherwise.

### GetApplicableDeviceTypeOk

`func (o *MicrosoftGraphIosVppApp) GetApplicableDeviceTypeOk() (AnyOfmicrosoftGraphIosDeviceType, bool)`

GetApplicableDeviceTypeOk returns a tuple with the ApplicableDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableDeviceType

`func (o *MicrosoftGraphIosVppApp) HasApplicableDeviceType() bool`

HasApplicableDeviceType returns a boolean if a field has been set.

### SetApplicableDeviceType

`func (o *MicrosoftGraphIosVppApp) SetApplicableDeviceType(v AnyOfmicrosoftGraphIosDeviceType)`

SetApplicableDeviceType gets a reference to the given AnyOfmicrosoftGraphIosDeviceType and assigns it to the ApplicableDeviceType field.

### SetApplicableDeviceTypeExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetApplicableDeviceTypeExplicitNull(b bool)`

SetApplicableDeviceTypeExplicitNull (un)sets ApplicableDeviceType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicableDeviceType value is set to nil even if false is passed
### GetVppTokenOrganizationName

`func (o *MicrosoftGraphIosVppApp) GetVppTokenOrganizationName() string`

GetVppTokenOrganizationName returns the VppTokenOrganizationName field if non-nil, zero value otherwise.

### GetVppTokenOrganizationNameOk

`func (o *MicrosoftGraphIosVppApp) GetVppTokenOrganizationNameOk() (string, bool)`

GetVppTokenOrganizationNameOk returns a tuple with the VppTokenOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenOrganizationName

`func (o *MicrosoftGraphIosVppApp) HasVppTokenOrganizationName() bool`

HasVppTokenOrganizationName returns a boolean if a field has been set.

### SetVppTokenOrganizationName

`func (o *MicrosoftGraphIosVppApp) SetVppTokenOrganizationName(v string)`

SetVppTokenOrganizationName gets a reference to the given string and assigns it to the VppTokenOrganizationName field.

### SetVppTokenOrganizationNameExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetVppTokenOrganizationNameExplicitNull(b bool)`

SetVppTokenOrganizationNameExplicitNull (un)sets VppTokenOrganizationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VppTokenOrganizationName value is set to nil even if false is passed
### GetVppTokenAccountType

`func (o *MicrosoftGraphIosVppApp) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType`

GetVppTokenAccountType returns the VppTokenAccountType field if non-nil, zero value otherwise.

### GetVppTokenAccountTypeOk

`func (o *MicrosoftGraphIosVppApp) GetVppTokenAccountTypeOk() (AnyOfmicrosoftGraphVppTokenAccountType, bool)`

GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenAccountType

`func (o *MicrosoftGraphIosVppApp) HasVppTokenAccountType() bool`

HasVppTokenAccountType returns a boolean if a field has been set.

### SetVppTokenAccountType

`func (o *MicrosoftGraphIosVppApp) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType)`

SetVppTokenAccountType gets a reference to the given AnyOfmicrosoftGraphVppTokenAccountType and assigns it to the VppTokenAccountType field.

### GetVppTokenAppleId

`func (o *MicrosoftGraphIosVppApp) GetVppTokenAppleId() string`

GetVppTokenAppleId returns the VppTokenAppleId field if non-nil, zero value otherwise.

### GetVppTokenAppleIdOk

`func (o *MicrosoftGraphIosVppApp) GetVppTokenAppleIdOk() (string, bool)`

GetVppTokenAppleIdOk returns a tuple with the VppTokenAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenAppleId

`func (o *MicrosoftGraphIosVppApp) HasVppTokenAppleId() bool`

HasVppTokenAppleId returns a boolean if a field has been set.

### SetVppTokenAppleId

`func (o *MicrosoftGraphIosVppApp) SetVppTokenAppleId(v string)`

SetVppTokenAppleId gets a reference to the given string and assigns it to the VppTokenAppleId field.

### SetVppTokenAppleIdExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetVppTokenAppleIdExplicitNull(b bool)`

SetVppTokenAppleIdExplicitNull (un)sets VppTokenAppleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VppTokenAppleId value is set to nil even if false is passed
### GetBundleId

`func (o *MicrosoftGraphIosVppApp) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *MicrosoftGraphIosVppApp) GetBundleIdOk() (string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleId

`func (o *MicrosoftGraphIosVppApp) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### SetBundleId

`func (o *MicrosoftGraphIosVppApp) SetBundleId(v string)`

SetBundleId gets a reference to the given string and assigns it to the BundleId field.

### SetBundleIdExplicitNull

`func (o *MicrosoftGraphIosVppApp) SetBundleIdExplicitNull(b bool)`

SetBundleIdExplicitNull (un)sets BundleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BundleId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


