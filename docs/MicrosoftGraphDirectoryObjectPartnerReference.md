# MicrosoftGraphDirectoryObjectPartnerReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalPartnerTenantId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeletedDateTime

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetDeletedDateTimeOk() (time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDateTime

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.

### SetDeletedDateTimeExplicitNull

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetDeletedDateTimeExplicitNull(b bool)`

SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeletedDateTime value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetExternalPartnerTenantId

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetExternalPartnerTenantId() string`

GetExternalPartnerTenantId returns the ExternalPartnerTenantId field if non-nil, zero value otherwise.

### GetExternalPartnerTenantIdOk

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetExternalPartnerTenantIdOk() (string, bool)`

GetExternalPartnerTenantIdOk returns a tuple with the ExternalPartnerTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalPartnerTenantId

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) HasExternalPartnerTenantId() bool`

HasExternalPartnerTenantId returns a boolean if a field has been set.

### SetExternalPartnerTenantId

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetExternalPartnerTenantId(v string)`

SetExternalPartnerTenantId gets a reference to the given string and assigns it to the ExternalPartnerTenantId field.

### SetExternalPartnerTenantIdExplicitNull

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetExternalPartnerTenantIdExplicitNull(b bool)`

SetExternalPartnerTenantIdExplicitNull (un)sets ExternalPartnerTenantId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalPartnerTenantId value is set to nil even if false is passed
### GetObjectType

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) GetObjectTypeOk() (string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasObjectType

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectType

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetObjectType(v string)`

SetObjectType gets a reference to the given string and assigns it to the ObjectType field.

### SetObjectTypeExplicitNull

`func (o *MicrosoftGraphDirectoryObjectPartnerReference) SetObjectTypeExplicitNull(b bool)`

SetObjectTypeExplicitNull (un)sets ObjectType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ObjectType value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


