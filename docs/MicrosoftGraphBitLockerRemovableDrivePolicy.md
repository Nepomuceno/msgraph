# MicrosoftGraphBitLockerRemovableDrivePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionMethod** | Pointer to [**AnyOfmicrosoftGraphBitLockerEncryptionMethod**](anyOf&lt;microsoft.graph.bitLockerEncryptionMethod&gt;.md) | Select the encryption method for removable  drives. | [optional] 
**RequireEncryptionForWriteAccess** | Pointer to **bool** | Indicates whether to block write access to devices configured in another organization.  If requireEncryptionForWriteAccess is false, this value does not affect. | [optional] 
**BlockCrossOrganizationWriteAccess** | Pointer to **bool** | This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a computer. | [optional] 

## Methods

### GetEncryptionMethod

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetEncryptionMethod() AnyOfmicrosoftGraphBitLockerEncryptionMethod`

GetEncryptionMethod returns the EncryptionMethod field if non-nil, zero value otherwise.

### GetEncryptionMethodOk

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetEncryptionMethodOk() (AnyOfmicrosoftGraphBitLockerEncryptionMethod, bool)`

GetEncryptionMethodOk returns a tuple with the EncryptionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncryptionMethod

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) HasEncryptionMethod() bool`

HasEncryptionMethod returns a boolean if a field has been set.

### SetEncryptionMethod

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetEncryptionMethod(v AnyOfmicrosoftGraphBitLockerEncryptionMethod)`

SetEncryptionMethod gets a reference to the given AnyOfmicrosoftGraphBitLockerEncryptionMethod and assigns it to the EncryptionMethod field.

### SetEncryptionMethodExplicitNull

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetEncryptionMethodExplicitNull(b bool)`

SetEncryptionMethodExplicitNull (un)sets EncryptionMethod to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EncryptionMethod value is set to nil even if false is passed
### GetRequireEncryptionForWriteAccess

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetRequireEncryptionForWriteAccess() bool`

GetRequireEncryptionForWriteAccess returns the RequireEncryptionForWriteAccess field if non-nil, zero value otherwise.

### GetRequireEncryptionForWriteAccessOk

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetRequireEncryptionForWriteAccessOk() (bool, bool)`

GetRequireEncryptionForWriteAccessOk returns a tuple with the RequireEncryptionForWriteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequireEncryptionForWriteAccess

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) HasRequireEncryptionForWriteAccess() bool`

HasRequireEncryptionForWriteAccess returns a boolean if a field has been set.

### SetRequireEncryptionForWriteAccess

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetRequireEncryptionForWriteAccess(v bool)`

SetRequireEncryptionForWriteAccess gets a reference to the given bool and assigns it to the RequireEncryptionForWriteAccess field.

### GetBlockCrossOrganizationWriteAccess

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetBlockCrossOrganizationWriteAccess() bool`

GetBlockCrossOrganizationWriteAccess returns the BlockCrossOrganizationWriteAccess field if non-nil, zero value otherwise.

### GetBlockCrossOrganizationWriteAccessOk

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetBlockCrossOrganizationWriteAccessOk() (bool, bool)`

GetBlockCrossOrganizationWriteAccessOk returns a tuple with the BlockCrossOrganizationWriteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockCrossOrganizationWriteAccess

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) HasBlockCrossOrganizationWriteAccess() bool`

HasBlockCrossOrganizationWriteAccess returns a boolean if a field has been set.

### SetBlockCrossOrganizationWriteAccess

`func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetBlockCrossOrganizationWriteAccess(v bool)`

SetBlockCrossOrganizationWriteAccess gets a reference to the given bool and assigns it to the BlockCrossOrganizationWriteAccess field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


