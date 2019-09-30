# MicrosoftGraphManagedAppRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | Date and time of creation | [optional] 
**LastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | Date and time of last the app synced with management service. | [optional] 
**ApplicationVersion** | Pointer to **string** | App version | [optional] 
**ManagementSdkVersion** | Pointer to **string** | App management SDK version | [optional] 
**PlatformVersion** | Pointer to **string** | Operating System version | [optional] 
**DeviceType** | Pointer to **string** | Host device type | [optional] 
**DeviceTag** | Pointer to **string** | App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions. | [optional] 
**DeviceName** | Pointer to **string** | Host device name | [optional] 
**FlaggedReasons** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppFlaggedReason**](anyOf&lt;microsoft.graph.managedAppFlaggedReason&gt;.md) | Zero or more reasons an app registration is flagged. E.g. app running on rooted device | [optional] 
**UserId** | Pointer to **string** | The user Id to who this app registration belongs. | [optional] 
**AppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The app package Identifier | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 
**AppliedPolicies** | Pointer to [**[]MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md) |  | [optional] 
**IntendedPolicies** | Pointer to [**[]MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md) |  | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphManagedAppOperation**](microsoft.graph.managedAppOperation.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedAppRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedAppRegistration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedAppRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedAppRegistration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphManagedAppRegistration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedAppRegistration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedAppRegistration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedAppRegistration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastSyncDateTime

`func (o *MicrosoftGraphManagedAppRegistration) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphManagedAppRegistration) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *MicrosoftGraphManagedAppRegistration) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphManagedAppRegistration) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetApplicationVersion

`func (o *MicrosoftGraphManagedAppRegistration) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *MicrosoftGraphManagedAppRegistration) GetApplicationVersionOk() (string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationVersion

`func (o *MicrosoftGraphManagedAppRegistration) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersion

`func (o *MicrosoftGraphManagedAppRegistration) SetApplicationVersion(v string)`

SetApplicationVersion gets a reference to the given string and assigns it to the ApplicationVersion field.

### SetApplicationVersionExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetApplicationVersionExplicitNull(b bool)`

SetApplicationVersionExplicitNull (un)sets ApplicationVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationVersion value is set to nil even if false is passed
### GetManagementSdkVersion

`func (o *MicrosoftGraphManagedAppRegistration) GetManagementSdkVersion() string`

GetManagementSdkVersion returns the ManagementSdkVersion field if non-nil, zero value otherwise.

### GetManagementSdkVersionOk

`func (o *MicrosoftGraphManagedAppRegistration) GetManagementSdkVersionOk() (string, bool)`

GetManagementSdkVersionOk returns a tuple with the ManagementSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagementSdkVersion

`func (o *MicrosoftGraphManagedAppRegistration) HasManagementSdkVersion() bool`

HasManagementSdkVersion returns a boolean if a field has been set.

### SetManagementSdkVersion

`func (o *MicrosoftGraphManagedAppRegistration) SetManagementSdkVersion(v string)`

SetManagementSdkVersion gets a reference to the given string and assigns it to the ManagementSdkVersion field.

### SetManagementSdkVersionExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetManagementSdkVersionExplicitNull(b bool)`

SetManagementSdkVersionExplicitNull (un)sets ManagementSdkVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagementSdkVersion value is set to nil even if false is passed
### GetPlatformVersion

`func (o *MicrosoftGraphManagedAppRegistration) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *MicrosoftGraphManagedAppRegistration) GetPlatformVersionOk() (string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformVersion

`func (o *MicrosoftGraphManagedAppRegistration) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### SetPlatformVersion

`func (o *MicrosoftGraphManagedAppRegistration) SetPlatformVersion(v string)`

SetPlatformVersion gets a reference to the given string and assigns it to the PlatformVersion field.

### SetPlatformVersionExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetPlatformVersionExplicitNull(b bool)`

SetPlatformVersionExplicitNull (un)sets PlatformVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PlatformVersion value is set to nil even if false is passed
### GetDeviceType

`func (o *MicrosoftGraphManagedAppRegistration) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MicrosoftGraphManagedAppRegistration) GetDeviceTypeOk() (string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceType

`func (o *MicrosoftGraphManagedAppRegistration) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceType

`func (o *MicrosoftGraphManagedAppRegistration) SetDeviceType(v string)`

SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.

### SetDeviceTypeExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetDeviceTypeExplicitNull(b bool)`

SetDeviceTypeExplicitNull (un)sets DeviceType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceType value is set to nil even if false is passed
### GetDeviceTag

`func (o *MicrosoftGraphManagedAppRegistration) GetDeviceTag() string`

GetDeviceTag returns the DeviceTag field if non-nil, zero value otherwise.

### GetDeviceTagOk

`func (o *MicrosoftGraphManagedAppRegistration) GetDeviceTagOk() (string, bool)`

GetDeviceTagOk returns a tuple with the DeviceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceTag

`func (o *MicrosoftGraphManagedAppRegistration) HasDeviceTag() bool`

HasDeviceTag returns a boolean if a field has been set.

### SetDeviceTag

`func (o *MicrosoftGraphManagedAppRegistration) SetDeviceTag(v string)`

SetDeviceTag gets a reference to the given string and assigns it to the DeviceTag field.

### SetDeviceTagExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetDeviceTagExplicitNull(b bool)`

SetDeviceTagExplicitNull (un)sets DeviceTag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceTag value is set to nil even if false is passed
### GetDeviceName

`func (o *MicrosoftGraphManagedAppRegistration) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphManagedAppRegistration) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *MicrosoftGraphManagedAppRegistration) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *MicrosoftGraphManagedAppRegistration) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetFlaggedReasons

`func (o *MicrosoftGraphManagedAppRegistration) GetFlaggedReasons() []AnyOfmicrosoftGraphManagedAppFlaggedReason`

GetFlaggedReasons returns the FlaggedReasons field if non-nil, zero value otherwise.

### GetFlaggedReasonsOk

`func (o *MicrosoftGraphManagedAppRegistration) GetFlaggedReasonsOk() ([]AnyOfmicrosoftGraphManagedAppFlaggedReason, bool)`

GetFlaggedReasonsOk returns a tuple with the FlaggedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlaggedReasons

`func (o *MicrosoftGraphManagedAppRegistration) HasFlaggedReasons() bool`

HasFlaggedReasons returns a boolean if a field has been set.

### SetFlaggedReasons

`func (o *MicrosoftGraphManagedAppRegistration) SetFlaggedReasons(v []AnyOfmicrosoftGraphManagedAppFlaggedReason)`

SetFlaggedReasons gets a reference to the given []AnyOfmicrosoftGraphManagedAppFlaggedReason and assigns it to the FlaggedReasons field.

### GetUserId

`func (o *MicrosoftGraphManagedAppRegistration) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphManagedAppRegistration) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphManagedAppRegistration) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphManagedAppRegistration) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetAppIdentifier

`func (o *MicrosoftGraphManagedAppRegistration) GetAppIdentifier() AnyOfobject`

GetAppIdentifier returns the AppIdentifier field if non-nil, zero value otherwise.

### GetAppIdentifierOk

`func (o *MicrosoftGraphManagedAppRegistration) GetAppIdentifierOk() (AnyOfobject, bool)`

GetAppIdentifierOk returns a tuple with the AppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppIdentifier

`func (o *MicrosoftGraphManagedAppRegistration) HasAppIdentifier() bool`

HasAppIdentifier returns a boolean if a field has been set.

### SetAppIdentifier

`func (o *MicrosoftGraphManagedAppRegistration) SetAppIdentifier(v AnyOfobject)`

SetAppIdentifier gets a reference to the given AnyOfobject and assigns it to the AppIdentifier field.

### SetAppIdentifierExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetAppIdentifierExplicitNull(b bool)`

SetAppIdentifierExplicitNull (un)sets AppIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppIdentifier value is set to nil even if false is passed
### GetVersion

`func (o *MicrosoftGraphManagedAppRegistration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedAppRegistration) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphManagedAppRegistration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphManagedAppRegistration) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphManagedAppRegistration) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetAppliedPolicies

`func (o *MicrosoftGraphManagedAppRegistration) GetAppliedPolicies() []MicrosoftGraphManagedAppPolicy`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *MicrosoftGraphManagedAppRegistration) GetAppliedPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliedPolicies

`func (o *MicrosoftGraphManagedAppRegistration) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### SetAppliedPolicies

`func (o *MicrosoftGraphManagedAppRegistration) SetAppliedPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetAppliedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the AppliedPolicies field.

### GetIntendedPolicies

`func (o *MicrosoftGraphManagedAppRegistration) GetIntendedPolicies() []MicrosoftGraphManagedAppPolicy`

GetIntendedPolicies returns the IntendedPolicies field if non-nil, zero value otherwise.

### GetIntendedPoliciesOk

`func (o *MicrosoftGraphManagedAppRegistration) GetIntendedPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool)`

GetIntendedPoliciesOk returns a tuple with the IntendedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntendedPolicies

`func (o *MicrosoftGraphManagedAppRegistration) HasIntendedPolicies() bool`

HasIntendedPolicies returns a boolean if a field has been set.

### SetIntendedPolicies

`func (o *MicrosoftGraphManagedAppRegistration) SetIntendedPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetIntendedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the IntendedPolicies field.

### GetOperations

`func (o *MicrosoftGraphManagedAppRegistration) GetOperations() []MicrosoftGraphManagedAppOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphManagedAppRegistration) GetOperationsOk() ([]MicrosoftGraphManagedAppOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperations

`func (o *MicrosoftGraphManagedAppRegistration) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperations

`func (o *MicrosoftGraphManagedAppRegistration) SetOperations(v []MicrosoftGraphManagedAppOperation)`

SetOperations gets a reference to the given []MicrosoftGraphManagedAppOperation and assigns it to the Operations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


