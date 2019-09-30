# MicrosoftGraphWindowsMobileMsi

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
**CommandLine** | Pointer to **string** | The command line. | [optional] 
**ProductCode** | Pointer to **string** | The product code. | [optional] 
**ProductVersion** | Pointer to **string** | The product version of Windows Mobile MSI Line of Business (LoB) app. | [optional] 
**IgnoreVersionDetection** | Pointer to **bool** | A boolean to control whether the app&#39;s version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsMobileMsi) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsMobileMsi) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsMobileMsi) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphWindowsMobileMsi) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsMobileMsi) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsMobileMsi) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphWindowsMobileMsi) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsMobileMsi) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsMobileMsi) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphWindowsMobileMsi) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphWindowsMobileMsi) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphWindowsMobileMsi) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetLargeIcon

`func (o *MicrosoftGraphWindowsMobileMsi) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetLargeIconOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeIcon

`func (o *MicrosoftGraphWindowsMobileMsi) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIcon

`func (o *MicrosoftGraphWindowsMobileMsi) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeIcon field.

### SetLargeIconExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetLargeIconExplicitNull(b bool)`

SetLargeIconExplicitNull (un)sets LargeIcon to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeIcon value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsMobileMsi) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsMobileMsi) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsMobileMsi) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsMobileMsi) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindowsMobileMsi) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsMobileMsi) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetIsFeatured

`func (o *MicrosoftGraphWindowsMobileMsi) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetIsFeaturedOk() (bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFeatured

`func (o *MicrosoftGraphWindowsMobileMsi) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### SetIsFeatured

`func (o *MicrosoftGraphWindowsMobileMsi) SetIsFeatured(v bool)`

SetIsFeatured gets a reference to the given bool and assigns it to the IsFeatured field.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphWindowsMobileMsi) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphWindowsMobileMsi) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphWindowsMobileMsi) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetInformationUrl

`func (o *MicrosoftGraphWindowsMobileMsi) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphWindowsMobileMsi) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphWindowsMobileMsi) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphWindowsMobileMsi) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphWindowsMobileMsi) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphWindowsMobileMsi) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDeveloper

`func (o *MicrosoftGraphWindowsMobileMsi) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetDeveloperOk() (string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeveloper

`func (o *MicrosoftGraphWindowsMobileMsi) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloper

`func (o *MicrosoftGraphWindowsMobileMsi) SetDeveloper(v string)`

SetDeveloper gets a reference to the given string and assigns it to the Developer field.

### SetDeveloperExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetDeveloperExplicitNull(b bool)`

SetDeveloperExplicitNull (un)sets Developer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Developer value is set to nil even if false is passed
### GetNotes

`func (o *MicrosoftGraphWindowsMobileMsi) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetNotesOk() (string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotes

`func (o *MicrosoftGraphWindowsMobileMsi) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotes

`func (o *MicrosoftGraphWindowsMobileMsi) SetNotes(v string)`

SetNotes gets a reference to the given string and assigns it to the Notes field.

### SetNotesExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetNotesExplicitNull(b bool)`

SetNotesExplicitNull (un)sets Notes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Notes value is set to nil even if false is passed
### GetPublishingState

`func (o *MicrosoftGraphWindowsMobileMsi) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetPublishingStateOk() (AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishingState

`func (o *MicrosoftGraphWindowsMobileMsi) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingState

`func (o *MicrosoftGraphWindowsMobileMsi) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState gets a reference to the given AnyOfmicrosoftGraphMobileAppPublishingState and assigns it to the PublishingState field.

### GetCategories

`func (o *MicrosoftGraphWindowsMobileMsi) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphWindowsMobileMsi) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphWindowsMobileMsi) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the Categories field.

### GetAssignments

`func (o *MicrosoftGraphWindowsMobileMsi) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetAssignmentsOk() ([]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindowsMobileMsi) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindowsMobileMsi) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphMobileAppAssignment and assigns it to the Assignments field.

### GetCommittedContentVersion

`func (o *MicrosoftGraphWindowsMobileMsi) GetCommittedContentVersion() string`

GetCommittedContentVersion returns the CommittedContentVersion field if non-nil, zero value otherwise.

### GetCommittedContentVersionOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetCommittedContentVersionOk() (string, bool)`

GetCommittedContentVersionOk returns a tuple with the CommittedContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommittedContentVersion

`func (o *MicrosoftGraphWindowsMobileMsi) HasCommittedContentVersion() bool`

HasCommittedContentVersion returns a boolean if a field has been set.

### SetCommittedContentVersion

`func (o *MicrosoftGraphWindowsMobileMsi) SetCommittedContentVersion(v string)`

SetCommittedContentVersion gets a reference to the given string and assigns it to the CommittedContentVersion field.

### SetCommittedContentVersionExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetCommittedContentVersionExplicitNull(b bool)`

SetCommittedContentVersionExplicitNull (un)sets CommittedContentVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommittedContentVersion value is set to nil even if false is passed
### GetFileName

`func (o *MicrosoftGraphWindowsMobileMsi) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetFileNameOk() (string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileName

`func (o *MicrosoftGraphWindowsMobileMsi) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileName

`func (o *MicrosoftGraphWindowsMobileMsi) SetFileName(v string)`

SetFileName gets a reference to the given string and assigns it to the FileName field.

### SetFileNameExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetFileNameExplicitNull(b bool)`

SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileName value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphWindowsMobileMsi) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphWindowsMobileMsi) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphWindowsMobileMsi) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### GetContentVersions

`func (o *MicrosoftGraphWindowsMobileMsi) GetContentVersions() []MicrosoftGraphMobileAppContent`

GetContentVersions returns the ContentVersions field if non-nil, zero value otherwise.

### GetContentVersionsOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetContentVersionsOk() ([]MicrosoftGraphMobileAppContent, bool)`

GetContentVersionsOk returns a tuple with the ContentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentVersions

`func (o *MicrosoftGraphWindowsMobileMsi) HasContentVersions() bool`

HasContentVersions returns a boolean if a field has been set.

### SetContentVersions

`func (o *MicrosoftGraphWindowsMobileMsi) SetContentVersions(v []MicrosoftGraphMobileAppContent)`

SetContentVersions gets a reference to the given []MicrosoftGraphMobileAppContent and assigns it to the ContentVersions field.

### GetCommandLine

`func (o *MicrosoftGraphWindowsMobileMsi) GetCommandLine() string`

GetCommandLine returns the CommandLine field if non-nil, zero value otherwise.

### GetCommandLineOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetCommandLineOk() (string, bool)`

GetCommandLineOk returns a tuple with the CommandLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommandLine

`func (o *MicrosoftGraphWindowsMobileMsi) HasCommandLine() bool`

HasCommandLine returns a boolean if a field has been set.

### SetCommandLine

`func (o *MicrosoftGraphWindowsMobileMsi) SetCommandLine(v string)`

SetCommandLine gets a reference to the given string and assigns it to the CommandLine field.

### SetCommandLineExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetCommandLineExplicitNull(b bool)`

SetCommandLineExplicitNull (un)sets CommandLine to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommandLine value is set to nil even if false is passed
### GetProductCode

`func (o *MicrosoftGraphWindowsMobileMsi) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetProductCodeOk() (string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductCode

`func (o *MicrosoftGraphWindowsMobileMsi) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCode

`func (o *MicrosoftGraphWindowsMobileMsi) SetProductCode(v string)`

SetProductCode gets a reference to the given string and assigns it to the ProductCode field.

### SetProductCodeExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetProductCodeExplicitNull(b bool)`

SetProductCodeExplicitNull (un)sets ProductCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductCode value is set to nil even if false is passed
### GetProductVersion

`func (o *MicrosoftGraphWindowsMobileMsi) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetProductVersionOk() (string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductVersion

`func (o *MicrosoftGraphWindowsMobileMsi) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### SetProductVersion

`func (o *MicrosoftGraphWindowsMobileMsi) SetProductVersion(v string)`

SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.

### SetProductVersionExplicitNull

`func (o *MicrosoftGraphWindowsMobileMsi) SetProductVersionExplicitNull(b bool)`

SetProductVersionExplicitNull (un)sets ProductVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductVersion value is set to nil even if false is passed
### GetIgnoreVersionDetection

`func (o *MicrosoftGraphWindowsMobileMsi) GetIgnoreVersionDetection() bool`

GetIgnoreVersionDetection returns the IgnoreVersionDetection field if non-nil, zero value otherwise.

### GetIgnoreVersionDetectionOk

`func (o *MicrosoftGraphWindowsMobileMsi) GetIgnoreVersionDetectionOk() (bool, bool)`

GetIgnoreVersionDetectionOk returns a tuple with the IgnoreVersionDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIgnoreVersionDetection

`func (o *MicrosoftGraphWindowsMobileMsi) HasIgnoreVersionDetection() bool`

HasIgnoreVersionDetection returns a boolean if a field has been set.

### SetIgnoreVersionDetection

`func (o *MicrosoftGraphWindowsMobileMsi) SetIgnoreVersionDetection(v bool)`

SetIgnoreVersionDetection gets a reference to the given bool and assigns it to the IgnoreVersionDetection field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


