# MicrosoftGraphIosVppEBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Name of the eBook. | [optional] 
**Description** | Pointer to **string** | Description. | [optional] 
**Publisher** | Pointer to **string** | Publisher. | [optional] 
**PublishedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the eBook was published. | [optional] 
**LargeCover** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | Book cover. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the eBook file was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the eBook was last modified. | [optional] 
**InformationUrl** | Pointer to **string** | The more information Url. | [optional] 
**PrivacyInformationUrl** | Pointer to **string** | The privacy statement Url. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphManagedEBookAssignment**](microsoft.graph.managedEBookAssignment.md) |  | [optional] 
**InstallSummary** | Pointer to [**AnyOfmicrosoftGraphEBookInstallSummary**](anyOf&lt;microsoft.graph.eBookInstallSummary&gt;.md) |  | [optional] 
**DeviceStates** | Pointer to [**[]MicrosoftGraphDeviceInstallState**](microsoft.graph.deviceInstallState.md) |  | [optional] 
**UserStateSummary** | Pointer to [**[]MicrosoftGraphUserInstallStateSummary**](microsoft.graph.userInstallStateSummary.md) |  | [optional] 
**VppTokenId** | Pointer to **string** | The Vpp token ID. | [optional] 
**AppleId** | Pointer to **string** | The Apple ID associated with Vpp token. | [optional] 
**VppOrganizationName** | Pointer to **string** | The Vpp token&#39;s organization name. | [optional] 
**Genres** | Pointer to **[]string** | Genres. | [optional] 
**Language** | Pointer to **string** | Language. | [optional] 
**Seller** | Pointer to **string** | Seller. | [optional] 
**TotalLicenseCount** | Pointer to **int32** | Total license count. | [optional] 
**UsedLicenseCount** | Pointer to **int32** | Used license count. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphIosVppEBook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosVppEBook) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosVppEBook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosVppEBook) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphIosVppEBook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosVppEBook) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosVppEBook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosVppEBook) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphIosVppEBook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosVppEBook) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosVppEBook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosVppEBook) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphIosVppEBook) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphIosVppEBook) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphIosVppEBook) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphIosVppEBook) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetPublishedDateTime

`func (o *MicrosoftGraphIosVppEBook) GetPublishedDateTime() time.Time`

GetPublishedDateTime returns the PublishedDateTime field if non-nil, zero value otherwise.

### GetPublishedDateTimeOk

`func (o *MicrosoftGraphIosVppEBook) GetPublishedDateTimeOk() (time.Time, bool)`

GetPublishedDateTimeOk returns a tuple with the PublishedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishedDateTime

`func (o *MicrosoftGraphIosVppEBook) HasPublishedDateTime() bool`

HasPublishedDateTime returns a boolean if a field has been set.

### SetPublishedDateTime

`func (o *MicrosoftGraphIosVppEBook) SetPublishedDateTime(v time.Time)`

SetPublishedDateTime gets a reference to the given time.Time and assigns it to the PublishedDateTime field.

### GetLargeCover

`func (o *MicrosoftGraphIosVppEBook) GetLargeCover() AnyOfmicrosoftGraphMimeContent`

GetLargeCover returns the LargeCover field if non-nil, zero value otherwise.

### GetLargeCoverOk

`func (o *MicrosoftGraphIosVppEBook) GetLargeCoverOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeCoverOk returns a tuple with the LargeCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeCover

`func (o *MicrosoftGraphIosVppEBook) HasLargeCover() bool`

HasLargeCover returns a boolean if a field has been set.

### SetLargeCover

`func (o *MicrosoftGraphIosVppEBook) SetLargeCover(v AnyOfmicrosoftGraphMimeContent)`

SetLargeCover gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeCover field.

### SetLargeCoverExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetLargeCoverExplicitNull(b bool)`

SetLargeCoverExplicitNull (un)sets LargeCover to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeCover value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphIosVppEBook) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosVppEBook) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosVppEBook) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosVppEBook) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosVppEBook) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosVppEBook) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosVppEBook) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosVppEBook) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetInformationUrl

`func (o *MicrosoftGraphIosVppEBook) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphIosVppEBook) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphIosVppEBook) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphIosVppEBook) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetPrivacyInformationUrl

`func (o *MicrosoftGraphIosVppEBook) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphIosVppEBook) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphIosVppEBook) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphIosVppEBook) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetAssignments

`func (o *MicrosoftGraphIosVppEBook) GetAssignments() []MicrosoftGraphManagedEBookAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosVppEBook) GetAssignmentsOk() ([]MicrosoftGraphManagedEBookAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosVppEBook) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosVppEBook) SetAssignments(v []MicrosoftGraphManagedEBookAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphManagedEBookAssignment and assigns it to the Assignments field.

### GetInstallSummary

`func (o *MicrosoftGraphIosVppEBook) GetInstallSummary() AnyOfmicrosoftGraphEBookInstallSummary`

GetInstallSummary returns the InstallSummary field if non-nil, zero value otherwise.

### GetInstallSummaryOk

`func (o *MicrosoftGraphIosVppEBook) GetInstallSummaryOk() (AnyOfmicrosoftGraphEBookInstallSummary, bool)`

GetInstallSummaryOk returns a tuple with the InstallSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallSummary

`func (o *MicrosoftGraphIosVppEBook) HasInstallSummary() bool`

HasInstallSummary returns a boolean if a field has been set.

### SetInstallSummary

`func (o *MicrosoftGraphIosVppEBook) SetInstallSummary(v AnyOfmicrosoftGraphEBookInstallSummary)`

SetInstallSummary gets a reference to the given AnyOfmicrosoftGraphEBookInstallSummary and assigns it to the InstallSummary field.

### SetInstallSummaryExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetInstallSummaryExplicitNull(b bool)`

SetInstallSummaryExplicitNull (un)sets InstallSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstallSummary value is set to nil even if false is passed
### GetDeviceStates

`func (o *MicrosoftGraphIosVppEBook) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *MicrosoftGraphIosVppEBook) GetDeviceStatesOk() ([]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStates

`func (o *MicrosoftGraphIosVppEBook) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.

### SetDeviceStates

`func (o *MicrosoftGraphIosVppEBook) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates gets a reference to the given []MicrosoftGraphDeviceInstallState and assigns it to the DeviceStates field.

### GetUserStateSummary

`func (o *MicrosoftGraphIosVppEBook) GetUserStateSummary() []MicrosoftGraphUserInstallStateSummary`

GetUserStateSummary returns the UserStateSummary field if non-nil, zero value otherwise.

### GetUserStateSummaryOk

`func (o *MicrosoftGraphIosVppEBook) GetUserStateSummaryOk() ([]MicrosoftGraphUserInstallStateSummary, bool)`

GetUserStateSummaryOk returns a tuple with the UserStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStateSummary

`func (o *MicrosoftGraphIosVppEBook) HasUserStateSummary() bool`

HasUserStateSummary returns a boolean if a field has been set.

### SetUserStateSummary

`func (o *MicrosoftGraphIosVppEBook) SetUserStateSummary(v []MicrosoftGraphUserInstallStateSummary)`

SetUserStateSummary gets a reference to the given []MicrosoftGraphUserInstallStateSummary and assigns it to the UserStateSummary field.

### GetVppTokenId

`func (o *MicrosoftGraphIosVppEBook) GetVppTokenId() string`

GetVppTokenId returns the VppTokenId field if non-nil, zero value otherwise.

### GetVppTokenIdOk

`func (o *MicrosoftGraphIosVppEBook) GetVppTokenIdOk() (string, bool)`

GetVppTokenIdOk returns a tuple with the VppTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenId

`func (o *MicrosoftGraphIosVppEBook) HasVppTokenId() bool`

HasVppTokenId returns a boolean if a field has been set.

### SetVppTokenId

`func (o *MicrosoftGraphIosVppEBook) SetVppTokenId(v string)`

SetVppTokenId gets a reference to the given string and assigns it to the VppTokenId field.

### GetAppleId

`func (o *MicrosoftGraphIosVppEBook) GetAppleId() string`

GetAppleId returns the AppleId field if non-nil, zero value otherwise.

### GetAppleIdOk

`func (o *MicrosoftGraphIosVppEBook) GetAppleIdOk() (string, bool)`

GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleId

`func (o *MicrosoftGraphIosVppEBook) HasAppleId() bool`

HasAppleId returns a boolean if a field has been set.

### SetAppleId

`func (o *MicrosoftGraphIosVppEBook) SetAppleId(v string)`

SetAppleId gets a reference to the given string and assigns it to the AppleId field.

### SetAppleIdExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetAppleIdExplicitNull(b bool)`

SetAppleIdExplicitNull (un)sets AppleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppleId value is set to nil even if false is passed
### GetVppOrganizationName

`func (o *MicrosoftGraphIosVppEBook) GetVppOrganizationName() string`

GetVppOrganizationName returns the VppOrganizationName field if non-nil, zero value otherwise.

### GetVppOrganizationNameOk

`func (o *MicrosoftGraphIosVppEBook) GetVppOrganizationNameOk() (string, bool)`

GetVppOrganizationNameOk returns a tuple with the VppOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppOrganizationName

`func (o *MicrosoftGraphIosVppEBook) HasVppOrganizationName() bool`

HasVppOrganizationName returns a boolean if a field has been set.

### SetVppOrganizationName

`func (o *MicrosoftGraphIosVppEBook) SetVppOrganizationName(v string)`

SetVppOrganizationName gets a reference to the given string and assigns it to the VppOrganizationName field.

### SetVppOrganizationNameExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetVppOrganizationNameExplicitNull(b bool)`

SetVppOrganizationNameExplicitNull (un)sets VppOrganizationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VppOrganizationName value is set to nil even if false is passed
### GetGenres

`func (o *MicrosoftGraphIosVppEBook) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *MicrosoftGraphIosVppEBook) GetGenresOk() ([]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGenres

`func (o *MicrosoftGraphIosVppEBook) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenres

`func (o *MicrosoftGraphIosVppEBook) SetGenres(v []string)`

SetGenres gets a reference to the given []string and assigns it to the Genres field.

### GetLanguage

`func (o *MicrosoftGraphIosVppEBook) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphIosVppEBook) GetLanguageOk() (string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLanguage

`func (o *MicrosoftGraphIosVppEBook) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguage

`func (o *MicrosoftGraphIosVppEBook) SetLanguage(v string)`

SetLanguage gets a reference to the given string and assigns it to the Language field.

### SetLanguageExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetLanguageExplicitNull(b bool)`

SetLanguageExplicitNull (un)sets Language to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Language value is set to nil even if false is passed
### GetSeller

`func (o *MicrosoftGraphIosVppEBook) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *MicrosoftGraphIosVppEBook) GetSellerOk() (string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeller

`func (o *MicrosoftGraphIosVppEBook) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### SetSeller

`func (o *MicrosoftGraphIosVppEBook) SetSeller(v string)`

SetSeller gets a reference to the given string and assigns it to the Seller field.

### SetSellerExplicitNull

`func (o *MicrosoftGraphIosVppEBook) SetSellerExplicitNull(b bool)`

SetSellerExplicitNull (un)sets Seller to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Seller value is set to nil even if false is passed
### GetTotalLicenseCount

`func (o *MicrosoftGraphIosVppEBook) GetTotalLicenseCount() int32`

GetTotalLicenseCount returns the TotalLicenseCount field if non-nil, zero value otherwise.

### GetTotalLicenseCountOk

`func (o *MicrosoftGraphIosVppEBook) GetTotalLicenseCountOk() (int32, bool)`

GetTotalLicenseCountOk returns a tuple with the TotalLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalLicenseCount

`func (o *MicrosoftGraphIosVppEBook) HasTotalLicenseCount() bool`

HasTotalLicenseCount returns a boolean if a field has been set.

### SetTotalLicenseCount

`func (o *MicrosoftGraphIosVppEBook) SetTotalLicenseCount(v int32)`

SetTotalLicenseCount gets a reference to the given int32 and assigns it to the TotalLicenseCount field.

### GetUsedLicenseCount

`func (o *MicrosoftGraphIosVppEBook) GetUsedLicenseCount() int32`

GetUsedLicenseCount returns the UsedLicenseCount field if non-nil, zero value otherwise.

### GetUsedLicenseCountOk

`func (o *MicrosoftGraphIosVppEBook) GetUsedLicenseCountOk() (int32, bool)`

GetUsedLicenseCountOk returns a tuple with the UsedLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsedLicenseCount

`func (o *MicrosoftGraphIosVppEBook) HasUsedLicenseCount() bool`

HasUsedLicenseCount returns a boolean if a field has been set.

### SetUsedLicenseCount

`func (o *MicrosoftGraphIosVppEBook) SetUsedLicenseCount(v int32)`

SetUsedLicenseCount gets a reference to the given int32 and assigns it to the UsedLicenseCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


