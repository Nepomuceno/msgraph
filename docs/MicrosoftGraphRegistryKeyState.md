# MicrosoftGraphRegistryKeyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hive** | Pointer to [**AnyOfmicrosoftGraphRegistryHive**](anyOf&lt;microsoft.graph.registryHive&gt;.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**OldKey** | Pointer to **string** |  | [optional] 
**OldValueData** | Pointer to **string** |  | [optional] 
**OldValueName** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to [**AnyOfmicrosoftGraphRegistryOperation**](anyOf&lt;microsoft.graph.registryOperation&gt;.md) |  | [optional] 
**ProcessId** | Pointer to **int32** |  | [optional] 
**ValueData** | Pointer to **string** |  | [optional] 
**ValueName** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to [**AnyOfmicrosoftGraphRegistryValueType**](anyOf&lt;microsoft.graph.registryValueType&gt;.md) |  | [optional] 

## Methods

### GetHive

`func (o *MicrosoftGraphRegistryKeyState) GetHive() AnyOfmicrosoftGraphRegistryHive`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *MicrosoftGraphRegistryKeyState) GetHiveOk() (AnyOfmicrosoftGraphRegistryHive, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHive

`func (o *MicrosoftGraphRegistryKeyState) HasHive() bool`

HasHive returns a boolean if a field has been set.

### SetHive

`func (o *MicrosoftGraphRegistryKeyState) SetHive(v AnyOfmicrosoftGraphRegistryHive)`

SetHive gets a reference to the given AnyOfmicrosoftGraphRegistryHive and assigns it to the Hive field.

### SetHiveExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetHiveExplicitNull(b bool)`

SetHiveExplicitNull (un)sets Hive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hive value is set to nil even if false is passed
### GetKey

`func (o *MicrosoftGraphRegistryKeyState) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MicrosoftGraphRegistryKeyState) GetKeyOk() (string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKey

`func (o *MicrosoftGraphRegistryKeyState) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKey

`func (o *MicrosoftGraphRegistryKeyState) SetKey(v string)`

SetKey gets a reference to the given string and assigns it to the Key field.

### SetKeyExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetKeyExplicitNull(b bool)`

SetKeyExplicitNull (un)sets Key to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Key value is set to nil even if false is passed
### GetOldKey

`func (o *MicrosoftGraphRegistryKeyState) GetOldKey() string`

GetOldKey returns the OldKey field if non-nil, zero value otherwise.

### GetOldKeyOk

`func (o *MicrosoftGraphRegistryKeyState) GetOldKeyOk() (string, bool)`

GetOldKeyOk returns a tuple with the OldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldKey

`func (o *MicrosoftGraphRegistryKeyState) HasOldKey() bool`

HasOldKey returns a boolean if a field has been set.

### SetOldKey

`func (o *MicrosoftGraphRegistryKeyState) SetOldKey(v string)`

SetOldKey gets a reference to the given string and assigns it to the OldKey field.

### SetOldKeyExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetOldKeyExplicitNull(b bool)`

SetOldKeyExplicitNull (un)sets OldKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OldKey value is set to nil even if false is passed
### GetOldValueData

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueData() string`

GetOldValueData returns the OldValueData field if non-nil, zero value otherwise.

### GetOldValueDataOk

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueDataOk() (string, bool)`

GetOldValueDataOk returns a tuple with the OldValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldValueData

`func (o *MicrosoftGraphRegistryKeyState) HasOldValueData() bool`

HasOldValueData returns a boolean if a field has been set.

### SetOldValueData

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueData(v string)`

SetOldValueData gets a reference to the given string and assigns it to the OldValueData field.

### SetOldValueDataExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueDataExplicitNull(b bool)`

SetOldValueDataExplicitNull (un)sets OldValueData to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OldValueData value is set to nil even if false is passed
### GetOldValueName

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueName() string`

GetOldValueName returns the OldValueName field if non-nil, zero value otherwise.

### GetOldValueNameOk

`func (o *MicrosoftGraphRegistryKeyState) GetOldValueNameOk() (string, bool)`

GetOldValueNameOk returns a tuple with the OldValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldValueName

`func (o *MicrosoftGraphRegistryKeyState) HasOldValueName() bool`

HasOldValueName returns a boolean if a field has been set.

### SetOldValueName

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueName(v string)`

SetOldValueName gets a reference to the given string and assigns it to the OldValueName field.

### SetOldValueNameExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetOldValueNameExplicitNull(b bool)`

SetOldValueNameExplicitNull (un)sets OldValueName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OldValueName value is set to nil even if false is passed
### GetOperation

`func (o *MicrosoftGraphRegistryKeyState) GetOperation() AnyOfmicrosoftGraphRegistryOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MicrosoftGraphRegistryKeyState) GetOperationOk() (AnyOfmicrosoftGraphRegistryOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperation

`func (o *MicrosoftGraphRegistryKeyState) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperation

`func (o *MicrosoftGraphRegistryKeyState) SetOperation(v AnyOfmicrosoftGraphRegistryOperation)`

SetOperation gets a reference to the given AnyOfmicrosoftGraphRegistryOperation and assigns it to the Operation field.

### SetOperationExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetOperationExplicitNull(b bool)`

SetOperationExplicitNull (un)sets Operation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Operation value is set to nil even if false is passed
### GetProcessId

`func (o *MicrosoftGraphRegistryKeyState) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *MicrosoftGraphRegistryKeyState) GetProcessIdOk() (int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessId

`func (o *MicrosoftGraphRegistryKeyState) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### SetProcessId

`func (o *MicrosoftGraphRegistryKeyState) SetProcessId(v int32)`

SetProcessId gets a reference to the given int32 and assigns it to the ProcessId field.

### SetProcessIdExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetProcessIdExplicitNull(b bool)`

SetProcessIdExplicitNull (un)sets ProcessId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProcessId value is set to nil even if false is passed
### GetValueData

`func (o *MicrosoftGraphRegistryKeyState) GetValueData() string`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *MicrosoftGraphRegistryKeyState) GetValueDataOk() (string, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueData

`func (o *MicrosoftGraphRegistryKeyState) HasValueData() bool`

HasValueData returns a boolean if a field has been set.

### SetValueData

`func (o *MicrosoftGraphRegistryKeyState) SetValueData(v string)`

SetValueData gets a reference to the given string and assigns it to the ValueData field.

### SetValueDataExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetValueDataExplicitNull(b bool)`

SetValueDataExplicitNull (un)sets ValueData to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueData value is set to nil even if false is passed
### GetValueName

`func (o *MicrosoftGraphRegistryKeyState) GetValueName() string`

GetValueName returns the ValueName field if non-nil, zero value otherwise.

### GetValueNameOk

`func (o *MicrosoftGraphRegistryKeyState) GetValueNameOk() (string, bool)`

GetValueNameOk returns a tuple with the ValueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueName

`func (o *MicrosoftGraphRegistryKeyState) HasValueName() bool`

HasValueName returns a boolean if a field has been set.

### SetValueName

`func (o *MicrosoftGraphRegistryKeyState) SetValueName(v string)`

SetValueName gets a reference to the given string and assigns it to the ValueName field.

### SetValueNameExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetValueNameExplicitNull(b bool)`

SetValueNameExplicitNull (un)sets ValueName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueName value is set to nil even if false is passed
### GetValueType

`func (o *MicrosoftGraphRegistryKeyState) GetValueType() AnyOfmicrosoftGraphRegistryValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *MicrosoftGraphRegistryKeyState) GetValueTypeOk() (AnyOfmicrosoftGraphRegistryValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueType

`func (o *MicrosoftGraphRegistryKeyState) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueType

`func (o *MicrosoftGraphRegistryKeyState) SetValueType(v AnyOfmicrosoftGraphRegistryValueType)`

SetValueType gets a reference to the given AnyOfmicrosoftGraphRegistryValueType and assigns it to the ValueType field.

### SetValueTypeExplicitNull

`func (o *MicrosoftGraphRegistryKeyState) SetValueTypeExplicitNull(b bool)`

SetValueTypeExplicitNull (un)sets ValueType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueType value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


