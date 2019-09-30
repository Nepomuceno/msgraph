# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAccountEnabled

`func (o *Device) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *Device) GetAccountEnabledOk() (bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountEnabled

`func (o *Device) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabled

`func (o *Device) SetAccountEnabled(v bool)`

SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.

### SetAccountEnabledExplicitNull

`func (o *Device) SetAccountEnabledExplicitNull(b bool)`

SetAccountEnabledExplicitNull (un)sets AccountEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountEnabled value is set to nil even if false is passed
### GetAlternativeSecurityIds

`func (o *Device) GetAlternativeSecurityIds() []MicrosoftGraphAlternativeSecurityId`

GetAlternativeSecurityIds returns the AlternativeSecurityIds field if non-nil, zero value otherwise.

### GetAlternativeSecurityIdsOk

`func (o *Device) GetAlternativeSecurityIdsOk() ([]MicrosoftGraphAlternativeSecurityId, bool)`

GetAlternativeSecurityIdsOk returns a tuple with the AlternativeSecurityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlternativeSecurityIds

`func (o *Device) HasAlternativeSecurityIds() bool`

HasAlternativeSecurityIds returns a boolean if a field has been set.

### SetAlternativeSecurityIds

`func (o *Device) SetAlternativeSecurityIds(v []MicrosoftGraphAlternativeSecurityId)`

SetAlternativeSecurityIds gets a reference to the given []MicrosoftGraphAlternativeSecurityId and assigns it to the AlternativeSecurityIds field.

### GetApproximateLastSignInDateTime

`func (o *Device) GetApproximateLastSignInDateTime() time.Time`

GetApproximateLastSignInDateTime returns the ApproximateLastSignInDateTime field if non-nil, zero value otherwise.

### GetApproximateLastSignInDateTimeOk

`func (o *Device) GetApproximateLastSignInDateTimeOk() (time.Time, bool)`

GetApproximateLastSignInDateTimeOk returns a tuple with the ApproximateLastSignInDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApproximateLastSignInDateTime

`func (o *Device) HasApproximateLastSignInDateTime() bool`

HasApproximateLastSignInDateTime returns a boolean if a field has been set.

### SetApproximateLastSignInDateTime

`func (o *Device) SetApproximateLastSignInDateTime(v time.Time)`

SetApproximateLastSignInDateTime gets a reference to the given time.Time and assigns it to the ApproximateLastSignInDateTime field.

### SetApproximateLastSignInDateTimeExplicitNull

`func (o *Device) SetApproximateLastSignInDateTimeExplicitNull(b bool)`

SetApproximateLastSignInDateTimeExplicitNull (un)sets ApproximateLastSignInDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApproximateLastSignInDateTime value is set to nil even if false is passed
### GetComplianceExpirationDateTime

`func (o *Device) GetComplianceExpirationDateTime() time.Time`

GetComplianceExpirationDateTime returns the ComplianceExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceExpirationDateTimeOk

`func (o *Device) GetComplianceExpirationDateTimeOk() (time.Time, bool)`

GetComplianceExpirationDateTimeOk returns a tuple with the ComplianceExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceExpirationDateTime

`func (o *Device) HasComplianceExpirationDateTime() bool`

HasComplianceExpirationDateTime returns a boolean if a field has been set.

### SetComplianceExpirationDateTime

`func (o *Device) SetComplianceExpirationDateTime(v time.Time)`

SetComplianceExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceExpirationDateTime field.

### SetComplianceExpirationDateTimeExplicitNull

`func (o *Device) SetComplianceExpirationDateTimeExplicitNull(b bool)`

SetComplianceExpirationDateTimeExplicitNull (un)sets ComplianceExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ComplianceExpirationDateTime value is set to nil even if false is passed
### GetDeviceId

`func (o *Device) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Device) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *Device) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *Device) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *Device) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetDeviceMetadata

`func (o *Device) GetDeviceMetadata() string`

GetDeviceMetadata returns the DeviceMetadata field if non-nil, zero value otherwise.

### GetDeviceMetadataOk

`func (o *Device) GetDeviceMetadataOk() (string, bool)`

GetDeviceMetadataOk returns a tuple with the DeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceMetadata

`func (o *Device) HasDeviceMetadata() bool`

HasDeviceMetadata returns a boolean if a field has been set.

### SetDeviceMetadata

`func (o *Device) SetDeviceMetadata(v string)`

SetDeviceMetadata gets a reference to the given string and assigns it to the DeviceMetadata field.

### SetDeviceMetadataExplicitNull

`func (o *Device) SetDeviceMetadataExplicitNull(b bool)`

SetDeviceMetadataExplicitNull (un)sets DeviceMetadata to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceMetadata value is set to nil even if false is passed
### GetDeviceVersion

`func (o *Device) GetDeviceVersion() int32`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *Device) GetDeviceVersionOk() (int32, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceVersion

`func (o *Device) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### SetDeviceVersion

`func (o *Device) SetDeviceVersion(v int32)`

SetDeviceVersion gets a reference to the given int32 and assigns it to the DeviceVersion field.

### SetDeviceVersionExplicitNull

`func (o *Device) SetDeviceVersionExplicitNull(b bool)`

SetDeviceVersionExplicitNull (un)sets DeviceVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceVersion value is set to nil even if false is passed
### GetDisplayName

`func (o *Device) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Device) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *Device) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *Device) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *Device) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetIsCompliant

`func (o *Device) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *Device) GetIsCompliantOk() (bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsCompliant

`func (o *Device) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### SetIsCompliant

`func (o *Device) SetIsCompliant(v bool)`

SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.

### SetIsCompliantExplicitNull

`func (o *Device) SetIsCompliantExplicitNull(b bool)`

SetIsCompliantExplicitNull (un)sets IsCompliant to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsCompliant value is set to nil even if false is passed
### GetIsManaged

`func (o *Device) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *Device) GetIsManagedOk() (bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsManaged

`func (o *Device) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### SetIsManaged

`func (o *Device) SetIsManaged(v bool)`

SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.

### SetIsManagedExplicitNull

`func (o *Device) SetIsManagedExplicitNull(b bool)`

SetIsManagedExplicitNull (un)sets IsManaged to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsManaged value is set to nil even if false is passed
### GetOnPremisesLastSyncDateTime

`func (o *Device) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Device) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *Device) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Device) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *Device) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *Device) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Device) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *Device) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *Device) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *Device) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetOperatingSystem

`func (o *Device) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Device) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *Device) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *Device) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *Device) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetOperatingSystemVersion

`func (o *Device) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *Device) GetOperatingSystemVersionOk() (string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystemVersion

`func (o *Device) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### SetOperatingSystemVersion

`func (o *Device) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion gets a reference to the given string and assigns it to the OperatingSystemVersion field.

### SetOperatingSystemVersionExplicitNull

`func (o *Device) SetOperatingSystemVersionExplicitNull(b bool)`

SetOperatingSystemVersionExplicitNull (un)sets OperatingSystemVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystemVersion value is set to nil even if false is passed
### GetPhysicalIds

`func (o *Device) GetPhysicalIds() []string`

GetPhysicalIds returns the PhysicalIds field if non-nil, zero value otherwise.

### GetPhysicalIdsOk

`func (o *Device) GetPhysicalIdsOk() ([]string, bool)`

GetPhysicalIdsOk returns a tuple with the PhysicalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhysicalIds

`func (o *Device) HasPhysicalIds() bool`

HasPhysicalIds returns a boolean if a field has been set.

### SetPhysicalIds

`func (o *Device) SetPhysicalIds(v []string)`

SetPhysicalIds gets a reference to the given []string and assigns it to the PhysicalIds field.

### GetProfileType

`func (o *Device) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *Device) GetProfileTypeOk() (string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileType

`func (o *Device) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### SetProfileType

`func (o *Device) SetProfileType(v string)`

SetProfileType gets a reference to the given string and assigns it to the ProfileType field.

### SetProfileTypeExplicitNull

`func (o *Device) SetProfileTypeExplicitNull(b bool)`

SetProfileTypeExplicitNull (un)sets ProfileType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProfileType value is set to nil even if false is passed
### GetSystemLabels

`func (o *Device) GetSystemLabels() []string`

GetSystemLabels returns the SystemLabels field if non-nil, zero value otherwise.

### GetSystemLabelsOk

`func (o *Device) GetSystemLabelsOk() ([]string, bool)`

GetSystemLabelsOk returns a tuple with the SystemLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystemLabels

`func (o *Device) HasSystemLabels() bool`

HasSystemLabels returns a boolean if a field has been set.

### SetSystemLabels

`func (o *Device) SetSystemLabels(v []string)`

SetSystemLabels gets a reference to the given []string and assigns it to the SystemLabels field.

### GetTrustType

`func (o *Device) GetTrustType() string`

GetTrustType returns the TrustType field if non-nil, zero value otherwise.

### GetTrustTypeOk

`func (o *Device) GetTrustTypeOk() (string, bool)`

GetTrustTypeOk returns a tuple with the TrustType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTrustType

`func (o *Device) HasTrustType() bool`

HasTrustType returns a boolean if a field has been set.

### SetTrustType

`func (o *Device) SetTrustType(v string)`

SetTrustType gets a reference to the given string and assigns it to the TrustType field.

### SetTrustTypeExplicitNull

`func (o *Device) SetTrustTypeExplicitNull(b bool)`

SetTrustTypeExplicitNull (un)sets TrustType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TrustType value is set to nil even if false is passed
### GetMemberOf

`func (o *Device) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *Device) GetMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberOf

`func (o *Device) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOf

`func (o *Device) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.

### GetRegisteredOwners

`func (o *Device) GetRegisteredOwners() []MicrosoftGraphDirectoryObject`

GetRegisteredOwners returns the RegisteredOwners field if non-nil, zero value otherwise.

### GetRegisteredOwnersOk

`func (o *Device) GetRegisteredOwnersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredOwnersOk returns a tuple with the RegisteredOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegisteredOwners

`func (o *Device) HasRegisteredOwners() bool`

HasRegisteredOwners returns a boolean if a field has been set.

### SetRegisteredOwners

`func (o *Device) SetRegisteredOwners(v []MicrosoftGraphDirectoryObject)`

SetRegisteredOwners gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RegisteredOwners field.

### GetRegisteredUsers

`func (o *Device) GetRegisteredUsers() []MicrosoftGraphDirectoryObject`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *Device) GetRegisteredUsersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegisteredUsers

`func (o *Device) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.

### SetRegisteredUsers

`func (o *Device) SetRegisteredUsers(v []MicrosoftGraphDirectoryObject)`

SetRegisteredUsers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RegisteredUsers field.

### GetTransitiveMemberOf

`func (o *Device) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *Device) GetTransitiveMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMemberOf

`func (o *Device) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### SetTransitiveMemberOf

`func (o *Device) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.

### GetExtensions

`func (o *Device) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Device) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Device) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Device) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


