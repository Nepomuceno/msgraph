# MicrosoftGraphProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** |  | [optional] 
**CommandLine** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**FileHash** | Pointer to [**AnyOfmicrosoftGraphFileHash**](anyOf&lt;microsoft.graph.fileHash&gt;.md) |  | [optional] 
**IntegrityLevel** | Pointer to [**AnyOfmicrosoftGraphProcessIntegrityLevel**](anyOf&lt;microsoft.graph.processIntegrityLevel&gt;.md) |  | [optional] 
**IsElevated** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentProcessCreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ParentProcessId** | Pointer to **int32** |  | [optional] 
**ParentProcessName** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ProcessId** | Pointer to **int32** |  | [optional] 

## Methods

### GetAccountName

`func (o *MicrosoftGraphProcess) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *MicrosoftGraphProcess) GetAccountNameOk() (string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountName

`func (o *MicrosoftGraphProcess) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountName

`func (o *MicrosoftGraphProcess) SetAccountName(v string)`

SetAccountName gets a reference to the given string and assigns it to the AccountName field.

### SetAccountNameExplicitNull

`func (o *MicrosoftGraphProcess) SetAccountNameExplicitNull(b bool)`

SetAccountNameExplicitNull (un)sets AccountName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountName value is set to nil even if false is passed
### GetCommandLine

`func (o *MicrosoftGraphProcess) GetCommandLine() string`

GetCommandLine returns the CommandLine field if non-nil, zero value otherwise.

### GetCommandLineOk

`func (o *MicrosoftGraphProcess) GetCommandLineOk() (string, bool)`

GetCommandLineOk returns a tuple with the CommandLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommandLine

`func (o *MicrosoftGraphProcess) HasCommandLine() bool`

HasCommandLine returns a boolean if a field has been set.

### SetCommandLine

`func (o *MicrosoftGraphProcess) SetCommandLine(v string)`

SetCommandLine gets a reference to the given string and assigns it to the CommandLine field.

### SetCommandLineExplicitNull

`func (o *MicrosoftGraphProcess) SetCommandLineExplicitNull(b bool)`

SetCommandLineExplicitNull (un)sets CommandLine to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommandLine value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphProcess) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphProcess) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphProcess) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphProcess) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphProcess) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetFileHash

`func (o *MicrosoftGraphProcess) GetFileHash() AnyOfmicrosoftGraphFileHash`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *MicrosoftGraphProcess) GetFileHashOk() (AnyOfmicrosoftGraphFileHash, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileHash

`func (o *MicrosoftGraphProcess) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHash

`func (o *MicrosoftGraphProcess) SetFileHash(v AnyOfmicrosoftGraphFileHash)`

SetFileHash gets a reference to the given AnyOfmicrosoftGraphFileHash and assigns it to the FileHash field.

### SetFileHashExplicitNull

`func (o *MicrosoftGraphProcess) SetFileHashExplicitNull(b bool)`

SetFileHashExplicitNull (un)sets FileHash to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileHash value is set to nil even if false is passed
### GetIntegrityLevel

`func (o *MicrosoftGraphProcess) GetIntegrityLevel() AnyOfmicrosoftGraphProcessIntegrityLevel`

GetIntegrityLevel returns the IntegrityLevel field if non-nil, zero value otherwise.

### GetIntegrityLevelOk

`func (o *MicrosoftGraphProcess) GetIntegrityLevelOk() (AnyOfmicrosoftGraphProcessIntegrityLevel, bool)`

GetIntegrityLevelOk returns a tuple with the IntegrityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrityLevel

`func (o *MicrosoftGraphProcess) HasIntegrityLevel() bool`

HasIntegrityLevel returns a boolean if a field has been set.

### SetIntegrityLevel

`func (o *MicrosoftGraphProcess) SetIntegrityLevel(v AnyOfmicrosoftGraphProcessIntegrityLevel)`

SetIntegrityLevel gets a reference to the given AnyOfmicrosoftGraphProcessIntegrityLevel and assigns it to the IntegrityLevel field.

### SetIntegrityLevelExplicitNull

`func (o *MicrosoftGraphProcess) SetIntegrityLevelExplicitNull(b bool)`

SetIntegrityLevelExplicitNull (un)sets IntegrityLevel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IntegrityLevel value is set to nil even if false is passed
### GetIsElevated

`func (o *MicrosoftGraphProcess) GetIsElevated() bool`

GetIsElevated returns the IsElevated field if non-nil, zero value otherwise.

### GetIsElevatedOk

`func (o *MicrosoftGraphProcess) GetIsElevatedOk() (bool, bool)`

GetIsElevatedOk returns a tuple with the IsElevated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsElevated

`func (o *MicrosoftGraphProcess) HasIsElevated() bool`

HasIsElevated returns a boolean if a field has been set.

### SetIsElevated

`func (o *MicrosoftGraphProcess) SetIsElevated(v bool)`

SetIsElevated gets a reference to the given bool and assigns it to the IsElevated field.

### SetIsElevatedExplicitNull

`func (o *MicrosoftGraphProcess) SetIsElevatedExplicitNull(b bool)`

SetIsElevatedExplicitNull (un)sets IsElevated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsElevated value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphProcess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphProcess) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphProcess) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphProcess) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphProcess) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentProcessCreatedDateTime

`func (o *MicrosoftGraphProcess) GetParentProcessCreatedDateTime() time.Time`

GetParentProcessCreatedDateTime returns the ParentProcessCreatedDateTime field if non-nil, zero value otherwise.

### GetParentProcessCreatedDateTimeOk

`func (o *MicrosoftGraphProcess) GetParentProcessCreatedDateTimeOk() (time.Time, bool)`

GetParentProcessCreatedDateTimeOk returns a tuple with the ParentProcessCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentProcessCreatedDateTime

`func (o *MicrosoftGraphProcess) HasParentProcessCreatedDateTime() bool`

HasParentProcessCreatedDateTime returns a boolean if a field has been set.

### SetParentProcessCreatedDateTime

`func (o *MicrosoftGraphProcess) SetParentProcessCreatedDateTime(v time.Time)`

SetParentProcessCreatedDateTime gets a reference to the given time.Time and assigns it to the ParentProcessCreatedDateTime field.

### SetParentProcessCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphProcess) SetParentProcessCreatedDateTimeExplicitNull(b bool)`

SetParentProcessCreatedDateTimeExplicitNull (un)sets ParentProcessCreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentProcessCreatedDateTime value is set to nil even if false is passed
### GetParentProcessId

`func (o *MicrosoftGraphProcess) GetParentProcessId() int32`

GetParentProcessId returns the ParentProcessId field if non-nil, zero value otherwise.

### GetParentProcessIdOk

`func (o *MicrosoftGraphProcess) GetParentProcessIdOk() (int32, bool)`

GetParentProcessIdOk returns a tuple with the ParentProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentProcessId

`func (o *MicrosoftGraphProcess) HasParentProcessId() bool`

HasParentProcessId returns a boolean if a field has been set.

### SetParentProcessId

`func (o *MicrosoftGraphProcess) SetParentProcessId(v int32)`

SetParentProcessId gets a reference to the given int32 and assigns it to the ParentProcessId field.

### SetParentProcessIdExplicitNull

`func (o *MicrosoftGraphProcess) SetParentProcessIdExplicitNull(b bool)`

SetParentProcessIdExplicitNull (un)sets ParentProcessId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentProcessId value is set to nil even if false is passed
### GetParentProcessName

`func (o *MicrosoftGraphProcess) GetParentProcessName() string`

GetParentProcessName returns the ParentProcessName field if non-nil, zero value otherwise.

### GetParentProcessNameOk

`func (o *MicrosoftGraphProcess) GetParentProcessNameOk() (string, bool)`

GetParentProcessNameOk returns a tuple with the ParentProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentProcessName

`func (o *MicrosoftGraphProcess) HasParentProcessName() bool`

HasParentProcessName returns a boolean if a field has been set.

### SetParentProcessName

`func (o *MicrosoftGraphProcess) SetParentProcessName(v string)`

SetParentProcessName gets a reference to the given string and assigns it to the ParentProcessName field.

### SetParentProcessNameExplicitNull

`func (o *MicrosoftGraphProcess) SetParentProcessNameExplicitNull(b bool)`

SetParentProcessNameExplicitNull (un)sets ParentProcessName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentProcessName value is set to nil even if false is passed
### GetPath

`func (o *MicrosoftGraphProcess) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MicrosoftGraphProcess) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *MicrosoftGraphProcess) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *MicrosoftGraphProcess) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### SetPathExplicitNull

`func (o *MicrosoftGraphProcess) SetPathExplicitNull(b bool)`

SetPathExplicitNull (un)sets Path to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Path value is set to nil even if false is passed
### GetProcessId

`func (o *MicrosoftGraphProcess) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *MicrosoftGraphProcess) GetProcessIdOk() (int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessId

`func (o *MicrosoftGraphProcess) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### SetProcessId

`func (o *MicrosoftGraphProcess) SetProcessId(v int32)`

SetProcessId gets a reference to the given int32 and assigns it to the ProcessId field.

### SetProcessIdExplicitNull

`func (o *MicrosoftGraphProcess) SetProcessIdExplicitNull(b bool)`

SetProcessIdExplicitNull (un)sets ProcessId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProcessId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


