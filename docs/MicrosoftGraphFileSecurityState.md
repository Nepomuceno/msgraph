# MicrosoftGraphFileSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHash** | Pointer to [**AnyOfmicrosoftGraphFileHash**](anyOf&lt;microsoft.graph.fileHash&gt;.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RiskScore** | Pointer to **string** |  | [optional] 

## Methods

### GetFileHash

`func (o *MicrosoftGraphFileSecurityState) GetFileHash() AnyOfmicrosoftGraphFileHash`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *MicrosoftGraphFileSecurityState) GetFileHashOk() (AnyOfmicrosoftGraphFileHash, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileHash

`func (o *MicrosoftGraphFileSecurityState) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHash

`func (o *MicrosoftGraphFileSecurityState) SetFileHash(v AnyOfmicrosoftGraphFileHash)`

SetFileHash gets a reference to the given AnyOfmicrosoftGraphFileHash and assigns it to the FileHash field.

### SetFileHashExplicitNull

`func (o *MicrosoftGraphFileSecurityState) SetFileHashExplicitNull(b bool)`

SetFileHashExplicitNull (un)sets FileHash to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileHash value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphFileSecurityState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphFileSecurityState) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphFileSecurityState) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphFileSecurityState) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphFileSecurityState) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetPath

`func (o *MicrosoftGraphFileSecurityState) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MicrosoftGraphFileSecurityState) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *MicrosoftGraphFileSecurityState) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *MicrosoftGraphFileSecurityState) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### SetPathExplicitNull

`func (o *MicrosoftGraphFileSecurityState) SetPathExplicitNull(b bool)`

SetPathExplicitNull (un)sets Path to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Path value is set to nil even if false is passed
### GetRiskScore

`func (o *MicrosoftGraphFileSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphFileSecurityState) GetRiskScoreOk() (string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskScore

`func (o *MicrosoftGraphFileSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScore

`func (o *MicrosoftGraphFileSecurityState) SetRiskScore(v string)`

SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.

### SetRiskScoreExplicitNull

`func (o *MicrosoftGraphFileSecurityState) SetRiskScoreExplicitNull(b bool)`

SetRiskScoreExplicitNull (un)sets RiskScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskScore value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


