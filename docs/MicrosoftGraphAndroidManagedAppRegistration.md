# MicrosoftGraphAndroidManagedAppRegistration

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

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastSyncDateTime

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetApplicationVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetApplicationVersionOk() (string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetApplicationVersion(v string)`

SetApplicationVersion gets a reference to the given string and assigns it to the ApplicationVersion field.

### SetApplicationVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetApplicationVersionExplicitNull(b bool)`

SetApplicationVersionExplicitNull (un)sets ApplicationVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationVersion value is set to nil even if false is passed
### GetManagementSdkVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetManagementSdkVersion() string`

GetManagementSdkVersion returns the ManagementSdkVersion field if non-nil, zero value otherwise.

### GetManagementSdkVersionOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetManagementSdkVersionOk() (string, bool)`

GetManagementSdkVersionOk returns a tuple with the ManagementSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagementSdkVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasManagementSdkVersion() bool`

HasManagementSdkVersion returns a boolean if a field has been set.

### SetManagementSdkVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetManagementSdkVersion(v string)`

SetManagementSdkVersion gets a reference to the given string and assigns it to the ManagementSdkVersion field.

### SetManagementSdkVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetManagementSdkVersionExplicitNull(b bool)`

SetManagementSdkVersionExplicitNull (un)sets ManagementSdkVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagementSdkVersion value is set to nil even if false is passed
### GetPlatformVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetPlatformVersionOk() (string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### SetPlatformVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetPlatformVersion(v string)`

SetPlatformVersion gets a reference to the given string and assigns it to the PlatformVersion field.

### SetPlatformVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetPlatformVersionExplicitNull(b bool)`

SetPlatformVersionExplicitNull (un)sets PlatformVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PlatformVersion value is set to nil even if false is passed
### GetDeviceType

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetDeviceTypeOk() (string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceType

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceType

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetDeviceType(v string)`

SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.

### SetDeviceTypeExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetDeviceTypeExplicitNull(b bool)`

SetDeviceTypeExplicitNull (un)sets DeviceType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceType value is set to nil even if false is passed
### GetDeviceTag

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetDeviceTag() string`

GetDeviceTag returns the DeviceTag field if non-nil, zero value otherwise.

### GetDeviceTagOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetDeviceTagOk() (string, bool)`

GetDeviceTagOk returns a tuple with the DeviceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceTag

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasDeviceTag() bool`

HasDeviceTag returns a boolean if a field has been set.

### SetDeviceTag

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetDeviceTag(v string)`

SetDeviceTag gets a reference to the given string and assigns it to the DeviceTag field.

### SetDeviceTagExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetDeviceTagExplicitNull(b bool)`

SetDeviceTagExplicitNull (un)sets DeviceTag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceTag value is set to nil even if false is passed
### GetDeviceName

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetFlaggedReasons

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetFlaggedReasons() []AnyOfmicrosoftGraphManagedAppFlaggedReason`

GetFlaggedReasons returns the FlaggedReasons field if non-nil, zero value otherwise.

### GetFlaggedReasonsOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetFlaggedReasonsOk() ([]AnyOfmicrosoftGraphManagedAppFlaggedReason, bool)`

GetFlaggedReasonsOk returns a tuple with the FlaggedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlaggedReasons

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasFlaggedReasons() bool`

HasFlaggedReasons returns a boolean if a field has been set.

### SetFlaggedReasons

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetFlaggedReasons(v []AnyOfmicrosoftGraphManagedAppFlaggedReason)`

SetFlaggedReasons gets a reference to the given []AnyOfmicrosoftGraphManagedAppFlaggedReason and assigns it to the FlaggedReasons field.

### GetUserId

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetAppIdentifier

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetAppIdentifier() AnyOfobject`

GetAppIdentifier returns the AppIdentifier field if non-nil, zero value otherwise.

### GetAppIdentifierOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetAppIdentifierOk() (AnyOfobject, bool)`

GetAppIdentifierOk returns a tuple with the AppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppIdentifier

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasAppIdentifier() bool`

HasAppIdentifier returns a boolean if a field has been set.

### SetAppIdentifier

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetAppIdentifier(v AnyOfobject)`

SetAppIdentifier gets a reference to the given AnyOfobject and assigns it to the AppIdentifier field.

### SetAppIdentifierExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetAppIdentifierExplicitNull(b bool)`

SetAppIdentifierExplicitNull (un)sets AppIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppIdentifier value is set to nil even if false is passed
### GetVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetAppliedPolicies

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetAppliedPolicies() []MicrosoftGraphManagedAppPolicy`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetAppliedPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliedPolicies

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### SetAppliedPolicies

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetAppliedPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetAppliedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the AppliedPolicies field.

### GetIntendedPolicies

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetIntendedPolicies() []MicrosoftGraphManagedAppPolicy`

GetIntendedPolicies returns the IntendedPolicies field if non-nil, zero value otherwise.

### GetIntendedPoliciesOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetIntendedPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool)`

GetIntendedPoliciesOk returns a tuple with the IntendedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntendedPolicies

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasIntendedPolicies() bool`

HasIntendedPolicies returns a boolean if a field has been set.

### SetIntendedPolicies

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetIntendedPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetIntendedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the IntendedPolicies field.

### GetOperations

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetOperations() []MicrosoftGraphManagedAppOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphAndroidManagedAppRegistration) GetOperationsOk() ([]MicrosoftGraphManagedAppOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperations

`func (o *MicrosoftGraphAndroidManagedAppRegistration) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperations

`func (o *MicrosoftGraphAndroidManagedAppRegistration) SetOperations(v []MicrosoftGraphManagedAppOperation)`

SetOperations gets a reference to the given []MicrosoftGraphManagedAppOperation and assigns it to the Operations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


