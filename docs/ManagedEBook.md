# ManagedEBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### GetDisplayName

`func (o *ManagedEBook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManagedEBook) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ManagedEBook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ManagedEBook) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *ManagedEBook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedEBook) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *ManagedEBook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *ManagedEBook) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *ManagedEBook) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *ManagedEBook) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ManagedEBook) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *ManagedEBook) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *ManagedEBook) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *ManagedEBook) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetPublishedDateTime

`func (o *ManagedEBook) GetPublishedDateTime() time.Time`

GetPublishedDateTime returns the PublishedDateTime field if non-nil, zero value otherwise.

### GetPublishedDateTimeOk

`func (o *ManagedEBook) GetPublishedDateTimeOk() (time.Time, bool)`

GetPublishedDateTimeOk returns a tuple with the PublishedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishedDateTime

`func (o *ManagedEBook) HasPublishedDateTime() bool`

HasPublishedDateTime returns a boolean if a field has been set.

### SetPublishedDateTime

`func (o *ManagedEBook) SetPublishedDateTime(v time.Time)`

SetPublishedDateTime gets a reference to the given time.Time and assigns it to the PublishedDateTime field.

### GetLargeCover

`func (o *ManagedEBook) GetLargeCover() AnyOfmicrosoftGraphMimeContent`

GetLargeCover returns the LargeCover field if non-nil, zero value otherwise.

### GetLargeCoverOk

`func (o *ManagedEBook) GetLargeCoverOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeCoverOk returns a tuple with the LargeCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeCover

`func (o *ManagedEBook) HasLargeCover() bool`

HasLargeCover returns a boolean if a field has been set.

### SetLargeCover

`func (o *ManagedEBook) SetLargeCover(v AnyOfmicrosoftGraphMimeContent)`

SetLargeCover gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeCover field.

### SetLargeCoverExplicitNull

`func (o *ManagedEBook) SetLargeCoverExplicitNull(b bool)`

SetLargeCoverExplicitNull (un)sets LargeCover to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeCover value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *ManagedEBook) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ManagedEBook) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *ManagedEBook) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *ManagedEBook) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *ManagedEBook) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *ManagedEBook) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *ManagedEBook) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *ManagedEBook) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetInformationUrl

`func (o *ManagedEBook) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *ManagedEBook) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *ManagedEBook) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *ManagedEBook) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *ManagedEBook) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetPrivacyInformationUrl

`func (o *ManagedEBook) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *ManagedEBook) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *ManagedEBook) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *ManagedEBook) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *ManagedEBook) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetAssignments

`func (o *ManagedEBook) GetAssignments() []MicrosoftGraphManagedEBookAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ManagedEBook) GetAssignmentsOk() ([]MicrosoftGraphManagedEBookAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *ManagedEBook) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *ManagedEBook) SetAssignments(v []MicrosoftGraphManagedEBookAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphManagedEBookAssignment and assigns it to the Assignments field.

### GetInstallSummary

`func (o *ManagedEBook) GetInstallSummary() AnyOfmicrosoftGraphEBookInstallSummary`

GetInstallSummary returns the InstallSummary field if non-nil, zero value otherwise.

### GetInstallSummaryOk

`func (o *ManagedEBook) GetInstallSummaryOk() (AnyOfmicrosoftGraphEBookInstallSummary, bool)`

GetInstallSummaryOk returns a tuple with the InstallSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallSummary

`func (o *ManagedEBook) HasInstallSummary() bool`

HasInstallSummary returns a boolean if a field has been set.

### SetInstallSummary

`func (o *ManagedEBook) SetInstallSummary(v AnyOfmicrosoftGraphEBookInstallSummary)`

SetInstallSummary gets a reference to the given AnyOfmicrosoftGraphEBookInstallSummary and assigns it to the InstallSummary field.

### SetInstallSummaryExplicitNull

`func (o *ManagedEBook) SetInstallSummaryExplicitNull(b bool)`

SetInstallSummaryExplicitNull (un)sets InstallSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstallSummary value is set to nil even if false is passed
### GetDeviceStates

`func (o *ManagedEBook) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *ManagedEBook) GetDeviceStatesOk() ([]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStates

`func (o *ManagedEBook) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.

### SetDeviceStates

`func (o *ManagedEBook) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates gets a reference to the given []MicrosoftGraphDeviceInstallState and assigns it to the DeviceStates field.

### GetUserStateSummary

`func (o *ManagedEBook) GetUserStateSummary() []MicrosoftGraphUserInstallStateSummary`

GetUserStateSummary returns the UserStateSummary field if non-nil, zero value otherwise.

### GetUserStateSummaryOk

`func (o *ManagedEBook) GetUserStateSummaryOk() ([]MicrosoftGraphUserInstallStateSummary, bool)`

GetUserStateSummaryOk returns a tuple with the UserStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStateSummary

`func (o *ManagedEBook) HasUserStateSummary() bool`

HasUserStateSummary returns a boolean if a field has been set.

### SetUserStateSummary

`func (o *ManagedEBook) SetUserStateSummary(v []MicrosoftGraphUserInstallStateSummary)`

SetUserStateSummary gets a reference to the given []MicrosoftGraphUserInstallStateSummary and assigns it to the UserStateSummary field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


