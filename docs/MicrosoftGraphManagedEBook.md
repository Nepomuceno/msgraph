# MicrosoftGraphManagedEBook

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

## Methods

### GetId

`func (o *MicrosoftGraphManagedEBook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedEBook) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedEBook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedEBook) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphManagedEBook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedEBook) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphManagedEBook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedEBook) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphManagedEBook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedEBook) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphManagedEBook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphManagedEBook) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphManagedEBook) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphManagedEBook) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphManagedEBook) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphManagedEBook) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphManagedEBook) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphManagedEBook) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetPublishedDateTime

`func (o *MicrosoftGraphManagedEBook) GetPublishedDateTime() time.Time`

GetPublishedDateTime returns the PublishedDateTime field if non-nil, zero value otherwise.

### GetPublishedDateTimeOk

`func (o *MicrosoftGraphManagedEBook) GetPublishedDateTimeOk() (time.Time, bool)`

GetPublishedDateTimeOk returns a tuple with the PublishedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublishedDateTime

`func (o *MicrosoftGraphManagedEBook) HasPublishedDateTime() bool`

HasPublishedDateTime returns a boolean if a field has been set.

### SetPublishedDateTime

`func (o *MicrosoftGraphManagedEBook) SetPublishedDateTime(v time.Time)`

SetPublishedDateTime gets a reference to the given time.Time and assigns it to the PublishedDateTime field.

### GetLargeCover

`func (o *MicrosoftGraphManagedEBook) GetLargeCover() AnyOfmicrosoftGraphMimeContent`

GetLargeCover returns the LargeCover field if non-nil, zero value otherwise.

### GetLargeCoverOk

`func (o *MicrosoftGraphManagedEBook) GetLargeCoverOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeCoverOk returns a tuple with the LargeCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLargeCover

`func (o *MicrosoftGraphManagedEBook) HasLargeCover() bool`

HasLargeCover returns a boolean if a field has been set.

### SetLargeCover

`func (o *MicrosoftGraphManagedEBook) SetLargeCover(v AnyOfmicrosoftGraphMimeContent)`

SetLargeCover gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LargeCover field.

### SetLargeCoverExplicitNull

`func (o *MicrosoftGraphManagedEBook) SetLargeCoverExplicitNull(b bool)`

SetLargeCoverExplicitNull (un)sets LargeCover to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LargeCover value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphManagedEBook) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedEBook) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedEBook) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedEBook) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedEBook) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedEBook) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedEBook) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedEBook) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetInformationUrl

`func (o *MicrosoftGraphManagedEBook) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphManagedEBook) GetInformationUrlOk() (string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInformationUrl

`func (o *MicrosoftGraphManagedEBook) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrl

`func (o *MicrosoftGraphManagedEBook) SetInformationUrl(v string)`

SetInformationUrl gets a reference to the given string and assigns it to the InformationUrl field.

### SetInformationUrlExplicitNull

`func (o *MicrosoftGraphManagedEBook) SetInformationUrlExplicitNull(b bool)`

SetInformationUrlExplicitNull (un)sets InformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InformationUrl value is set to nil even if false is passed
### GetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedEBook) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphManagedEBook) GetPrivacyInformationUrlOk() (string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphManagedEBook) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedEBook) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl gets a reference to the given string and assigns it to the PrivacyInformationUrl field.

### SetPrivacyInformationUrlExplicitNull

`func (o *MicrosoftGraphManagedEBook) SetPrivacyInformationUrlExplicitNull(b bool)`

SetPrivacyInformationUrlExplicitNull (un)sets PrivacyInformationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyInformationUrl value is set to nil even if false is passed
### GetAssignments

`func (o *MicrosoftGraphManagedEBook) GetAssignments() []MicrosoftGraphManagedEBookAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphManagedEBook) GetAssignmentsOk() ([]MicrosoftGraphManagedEBookAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphManagedEBook) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphManagedEBook) SetAssignments(v []MicrosoftGraphManagedEBookAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphManagedEBookAssignment and assigns it to the Assignments field.

### GetInstallSummary

`func (o *MicrosoftGraphManagedEBook) GetInstallSummary() AnyOfmicrosoftGraphEBookInstallSummary`

GetInstallSummary returns the InstallSummary field if non-nil, zero value otherwise.

### GetInstallSummaryOk

`func (o *MicrosoftGraphManagedEBook) GetInstallSummaryOk() (AnyOfmicrosoftGraphEBookInstallSummary, bool)`

GetInstallSummaryOk returns a tuple with the InstallSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallSummary

`func (o *MicrosoftGraphManagedEBook) HasInstallSummary() bool`

HasInstallSummary returns a boolean if a field has been set.

### SetInstallSummary

`func (o *MicrosoftGraphManagedEBook) SetInstallSummary(v AnyOfmicrosoftGraphEBookInstallSummary)`

SetInstallSummary gets a reference to the given AnyOfmicrosoftGraphEBookInstallSummary and assigns it to the InstallSummary field.

### SetInstallSummaryExplicitNull

`func (o *MicrosoftGraphManagedEBook) SetInstallSummaryExplicitNull(b bool)`

SetInstallSummaryExplicitNull (un)sets InstallSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstallSummary value is set to nil even if false is passed
### GetDeviceStates

`func (o *MicrosoftGraphManagedEBook) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *MicrosoftGraphManagedEBook) GetDeviceStatesOk() ([]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStates

`func (o *MicrosoftGraphManagedEBook) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.

### SetDeviceStates

`func (o *MicrosoftGraphManagedEBook) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates gets a reference to the given []MicrosoftGraphDeviceInstallState and assigns it to the DeviceStates field.

### GetUserStateSummary

`func (o *MicrosoftGraphManagedEBook) GetUserStateSummary() []MicrosoftGraphUserInstallStateSummary`

GetUserStateSummary returns the UserStateSummary field if non-nil, zero value otherwise.

### GetUserStateSummaryOk

`func (o *MicrosoftGraphManagedEBook) GetUserStateSummaryOk() ([]MicrosoftGraphUserInstallStateSummary, bool)`

GetUserStateSummaryOk returns a tuple with the UserStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStateSummary

`func (o *MicrosoftGraphManagedEBook) HasUserStateSummary() bool`

HasUserStateSummary returns a boolean if a field has been set.

### SetUserStateSummary

`func (o *MicrosoftGraphManagedEBook) SetUserStateSummary(v []MicrosoftGraphUserInstallStateSummary)`

SetUserStateSummary gets a reference to the given []MicrosoftGraphUserInstallStateSummary and assigns it to the UserStateSummary field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


