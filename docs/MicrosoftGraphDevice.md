# MicrosoftGraphDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AccountEnabled** | Pointer to **bool** |  | [optional] 
**AlternativeSecurityIds** | Pointer to [**[]MicrosoftGraphAlternativeSecurityId**](microsoft.graph.alternativeSecurityId.md) |  | [optional] 
**ApproximateLastSignInDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ComplianceExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**DeviceMetadata** | Pointer to **string** |  | [optional] 
**DeviceVersion** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsCompliant** | Pointer to **bool** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**OnPremisesSyncEnabled** | Pointer to **bool** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**OperatingSystemVersion** | Pointer to **string** |  | [optional] 
**PhysicalIds** | Pointer to **[]string** |  | [optional] 
**ProfileType** | Pointer to **string** |  | [optional] 
**SystemLabels** | Pointer to **[]string** |  | [optional] 
**TrustType** | Pointer to **string** |  | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**RegisteredOwners** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**RegisteredUsers** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDevice) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDevice) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeletedDateTime

`func (o *MicrosoftGraphDevice) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphDevice) GetDeletedDateTimeOk() (time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDateTime

`func (o *MicrosoftGraphDevice) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphDevice) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.

### SetDeletedDateTimeExplicitNull

`func (o *MicrosoftGraphDevice) SetDeletedDateTimeExplicitNull(b bool)`

SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeletedDateTime value is set to nil even if false is passed
### GetAccountEnabled

`func (o *MicrosoftGraphDevice) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *MicrosoftGraphDevice) GetAccountEnabledOk() (bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountEnabled

`func (o *MicrosoftGraphDevice) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabled

`func (o *MicrosoftGraphDevice) SetAccountEnabled(v bool)`

SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.

### SetAccountEnabledExplicitNull

`func (o *MicrosoftGraphDevice) SetAccountEnabledExplicitNull(b bool)`

SetAccountEnabledExplicitNull (un)sets AccountEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountEnabled value is set to nil even if false is passed
### GetAlternativeSecurityIds

`func (o *MicrosoftGraphDevice) GetAlternativeSecurityIds() []MicrosoftGraphAlternativeSecurityId`

GetAlternativeSecurityIds returns the AlternativeSecurityIds field if non-nil, zero value otherwise.

### GetAlternativeSecurityIdsOk

`func (o *MicrosoftGraphDevice) GetAlternativeSecurityIdsOk() ([]MicrosoftGraphAlternativeSecurityId, bool)`

GetAlternativeSecurityIdsOk returns a tuple with the AlternativeSecurityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlternativeSecurityIds

`func (o *MicrosoftGraphDevice) HasAlternativeSecurityIds() bool`

HasAlternativeSecurityIds returns a boolean if a field has been set.

### SetAlternativeSecurityIds

`func (o *MicrosoftGraphDevice) SetAlternativeSecurityIds(v []MicrosoftGraphAlternativeSecurityId)`

SetAlternativeSecurityIds gets a reference to the given []MicrosoftGraphAlternativeSecurityId and assigns it to the AlternativeSecurityIds field.

### GetApproximateLastSignInDateTime

`func (o *MicrosoftGraphDevice) GetApproximateLastSignInDateTime() time.Time`

GetApproximateLastSignInDateTime returns the ApproximateLastSignInDateTime field if non-nil, zero value otherwise.

### GetApproximateLastSignInDateTimeOk

`func (o *MicrosoftGraphDevice) GetApproximateLastSignInDateTimeOk() (time.Time, bool)`

GetApproximateLastSignInDateTimeOk returns a tuple with the ApproximateLastSignInDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApproximateLastSignInDateTime

`func (o *MicrosoftGraphDevice) HasApproximateLastSignInDateTime() bool`

HasApproximateLastSignInDateTime returns a boolean if a field has been set.

### SetApproximateLastSignInDateTime

`func (o *MicrosoftGraphDevice) SetApproximateLastSignInDateTime(v time.Time)`

SetApproximateLastSignInDateTime gets a reference to the given time.Time and assigns it to the ApproximateLastSignInDateTime field.

### SetApproximateLastSignInDateTimeExplicitNull

`func (o *MicrosoftGraphDevice) SetApproximateLastSignInDateTimeExplicitNull(b bool)`

SetApproximateLastSignInDateTimeExplicitNull (un)sets ApproximateLastSignInDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApproximateLastSignInDateTime value is set to nil even if false is passed
### GetComplianceExpirationDateTime

`func (o *MicrosoftGraphDevice) GetComplianceExpirationDateTime() time.Time`

GetComplianceExpirationDateTime returns the ComplianceExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceExpirationDateTimeOk

`func (o *MicrosoftGraphDevice) GetComplianceExpirationDateTimeOk() (time.Time, bool)`

GetComplianceExpirationDateTimeOk returns a tuple with the ComplianceExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceExpirationDateTime

`func (o *MicrosoftGraphDevice) HasComplianceExpirationDateTime() bool`

HasComplianceExpirationDateTime returns a boolean if a field has been set.

### SetComplianceExpirationDateTime

`func (o *MicrosoftGraphDevice) SetComplianceExpirationDateTime(v time.Time)`

SetComplianceExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceExpirationDateTime field.

### SetComplianceExpirationDateTimeExplicitNull

`func (o *MicrosoftGraphDevice) SetComplianceExpirationDateTimeExplicitNull(b bool)`

SetComplianceExpirationDateTimeExplicitNull (un)sets ComplianceExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ComplianceExpirationDateTime value is set to nil even if false is passed
### GetDeviceId

`func (o *MicrosoftGraphDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDevice) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *MicrosoftGraphDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *MicrosoftGraphDevice) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *MicrosoftGraphDevice) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetDeviceMetadata

`func (o *MicrosoftGraphDevice) GetDeviceMetadata() string`

GetDeviceMetadata returns the DeviceMetadata field if non-nil, zero value otherwise.

### GetDeviceMetadataOk

`func (o *MicrosoftGraphDevice) GetDeviceMetadataOk() (string, bool)`

GetDeviceMetadataOk returns a tuple with the DeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceMetadata

`func (o *MicrosoftGraphDevice) HasDeviceMetadata() bool`

HasDeviceMetadata returns a boolean if a field has been set.

### SetDeviceMetadata

`func (o *MicrosoftGraphDevice) SetDeviceMetadata(v string)`

SetDeviceMetadata gets a reference to the given string and assigns it to the DeviceMetadata field.

### SetDeviceMetadataExplicitNull

`func (o *MicrosoftGraphDevice) SetDeviceMetadataExplicitNull(b bool)`

SetDeviceMetadataExplicitNull (un)sets DeviceMetadata to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceMetadata value is set to nil even if false is passed
### GetDeviceVersion

`func (o *MicrosoftGraphDevice) GetDeviceVersion() int32`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *MicrosoftGraphDevice) GetDeviceVersionOk() (int32, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceVersion

`func (o *MicrosoftGraphDevice) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### SetDeviceVersion

`func (o *MicrosoftGraphDevice) SetDeviceVersion(v int32)`

SetDeviceVersion gets a reference to the given int32 and assigns it to the DeviceVersion field.

### SetDeviceVersionExplicitNull

`func (o *MicrosoftGraphDevice) SetDeviceVersionExplicitNull(b bool)`

SetDeviceVersionExplicitNull (un)sets DeviceVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceVersion value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphDevice) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDevice) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDevice) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDevice) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDevice) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetIsCompliant

`func (o *MicrosoftGraphDevice) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *MicrosoftGraphDevice) GetIsCompliantOk() (bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsCompliant

`func (o *MicrosoftGraphDevice) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### SetIsCompliant

`func (o *MicrosoftGraphDevice) SetIsCompliant(v bool)`

SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.

### SetIsCompliantExplicitNull

`func (o *MicrosoftGraphDevice) SetIsCompliantExplicitNull(b bool)`

SetIsCompliantExplicitNull (un)sets IsCompliant to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsCompliant value is set to nil even if false is passed
### GetIsManaged

`func (o *MicrosoftGraphDevice) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *MicrosoftGraphDevice) GetIsManagedOk() (bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsManaged

`func (o *MicrosoftGraphDevice) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### SetIsManaged

`func (o *MicrosoftGraphDevice) SetIsManaged(v bool)`

SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.

### SetIsManagedExplicitNull

`func (o *MicrosoftGraphDevice) SetIsManagedExplicitNull(b bool)`

SetIsManagedExplicitNull (un)sets IsManaged to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsManaged value is set to nil even if false is passed
### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphDevice) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphDevice) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphDevice) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphDevice) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *MicrosoftGraphDevice) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphDevice) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphDevice) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphDevice) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphDevice) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *MicrosoftGraphDevice) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetOperatingSystem

`func (o *MicrosoftGraphDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphDevice) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *MicrosoftGraphDevice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphDevice) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *MicrosoftGraphDevice) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetOperatingSystemVersion

`func (o *MicrosoftGraphDevice) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *MicrosoftGraphDevice) GetOperatingSystemVersionOk() (string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystemVersion

`func (o *MicrosoftGraphDevice) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### SetOperatingSystemVersion

`func (o *MicrosoftGraphDevice) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion gets a reference to the given string and assigns it to the OperatingSystemVersion field.

### SetOperatingSystemVersionExplicitNull

`func (o *MicrosoftGraphDevice) SetOperatingSystemVersionExplicitNull(b bool)`

SetOperatingSystemVersionExplicitNull (un)sets OperatingSystemVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystemVersion value is set to nil even if false is passed
### GetPhysicalIds

`func (o *MicrosoftGraphDevice) GetPhysicalIds() []string`

GetPhysicalIds returns the PhysicalIds field if non-nil, zero value otherwise.

### GetPhysicalIdsOk

`func (o *MicrosoftGraphDevice) GetPhysicalIdsOk() ([]string, bool)`

GetPhysicalIdsOk returns a tuple with the PhysicalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhysicalIds

`func (o *MicrosoftGraphDevice) HasPhysicalIds() bool`

HasPhysicalIds returns a boolean if a field has been set.

### SetPhysicalIds

`func (o *MicrosoftGraphDevice) SetPhysicalIds(v []string)`

SetPhysicalIds gets a reference to the given []string and assigns it to the PhysicalIds field.

### GetProfileType

`func (o *MicrosoftGraphDevice) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *MicrosoftGraphDevice) GetProfileTypeOk() (string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileType

`func (o *MicrosoftGraphDevice) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### SetProfileType

`func (o *MicrosoftGraphDevice) SetProfileType(v string)`

SetProfileType gets a reference to the given string and assigns it to the ProfileType field.

### SetProfileTypeExplicitNull

`func (o *MicrosoftGraphDevice) SetProfileTypeExplicitNull(b bool)`

SetProfileTypeExplicitNull (un)sets ProfileType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProfileType value is set to nil even if false is passed
### GetSystemLabels

`func (o *MicrosoftGraphDevice) GetSystemLabels() []string`

GetSystemLabels returns the SystemLabels field if non-nil, zero value otherwise.

### GetSystemLabelsOk

`func (o *MicrosoftGraphDevice) GetSystemLabelsOk() ([]string, bool)`

GetSystemLabelsOk returns a tuple with the SystemLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystemLabels

`func (o *MicrosoftGraphDevice) HasSystemLabels() bool`

HasSystemLabels returns a boolean if a field has been set.

### SetSystemLabels

`func (o *MicrosoftGraphDevice) SetSystemLabels(v []string)`

SetSystemLabels gets a reference to the given []string and assigns it to the SystemLabels field.

### GetTrustType

`func (o *MicrosoftGraphDevice) GetTrustType() string`

GetTrustType returns the TrustType field if non-nil, zero value otherwise.

### GetTrustTypeOk

`func (o *MicrosoftGraphDevice) GetTrustTypeOk() (string, bool)`

GetTrustTypeOk returns a tuple with the TrustType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTrustType

`func (o *MicrosoftGraphDevice) HasTrustType() bool`

HasTrustType returns a boolean if a field has been set.

### SetTrustType

`func (o *MicrosoftGraphDevice) SetTrustType(v string)`

SetTrustType gets a reference to the given string and assigns it to the TrustType field.

### SetTrustTypeExplicitNull

`func (o *MicrosoftGraphDevice) SetTrustTypeExplicitNull(b bool)`

SetTrustTypeExplicitNull (un)sets TrustType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TrustType value is set to nil even if false is passed
### GetMemberOf

`func (o *MicrosoftGraphDevice) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MicrosoftGraphDevice) GetMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberOf

`func (o *MicrosoftGraphDevice) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOf

`func (o *MicrosoftGraphDevice) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.

### GetRegisteredOwners

`func (o *MicrosoftGraphDevice) GetRegisteredOwners() []MicrosoftGraphDirectoryObject`

GetRegisteredOwners returns the RegisteredOwners field if non-nil, zero value otherwise.

### GetRegisteredOwnersOk

`func (o *MicrosoftGraphDevice) GetRegisteredOwnersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredOwnersOk returns a tuple with the RegisteredOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegisteredOwners

`func (o *MicrosoftGraphDevice) HasRegisteredOwners() bool`

HasRegisteredOwners returns a boolean if a field has been set.

### SetRegisteredOwners

`func (o *MicrosoftGraphDevice) SetRegisteredOwners(v []MicrosoftGraphDirectoryObject)`

SetRegisteredOwners gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RegisteredOwners field.

### GetRegisteredUsers

`func (o *MicrosoftGraphDevice) GetRegisteredUsers() []MicrosoftGraphDirectoryObject`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *MicrosoftGraphDevice) GetRegisteredUsersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegisteredUsers

`func (o *MicrosoftGraphDevice) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.

### SetRegisteredUsers

`func (o *MicrosoftGraphDevice) SetRegisteredUsers(v []MicrosoftGraphDirectoryObject)`

SetRegisteredUsers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RegisteredUsers field.

### GetTransitiveMemberOf

`func (o *MicrosoftGraphDevice) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *MicrosoftGraphDevice) GetTransitiveMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMemberOf

`func (o *MicrosoftGraphDevice) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### SetTransitiveMemberOf

`func (o *MicrosoftGraphDevice) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.

### GetExtensions

`func (o *MicrosoftGraphDevice) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphDevice) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphDevice) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphDevice) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


