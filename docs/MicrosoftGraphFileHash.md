# MicrosoftGraphFileHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashType** | Pointer to [**AnyOfmicrosoftGraphFileHashType**](anyOf&lt;microsoft.graph.fileHashType&gt;.md) |  | [optional] 
**HashValue** | Pointer to **string** |  | [optional] 

## Methods

### GetHashType

`func (o *MicrosoftGraphFileHash) GetHashType() AnyOfmicrosoftGraphFileHashType`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *MicrosoftGraphFileHash) GetHashTypeOk() (AnyOfmicrosoftGraphFileHashType, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHashType

`func (o *MicrosoftGraphFileHash) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### SetHashType

`func (o *MicrosoftGraphFileHash) SetHashType(v AnyOfmicrosoftGraphFileHashType)`

SetHashType gets a reference to the given AnyOfmicrosoftGraphFileHashType and assigns it to the HashType field.

### SetHashTypeExplicitNull

`func (o *MicrosoftGraphFileHash) SetHashTypeExplicitNull(b bool)`

SetHashTypeExplicitNull (un)sets HashType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HashType value is set to nil even if false is passed
### GetHashValue

`func (o *MicrosoftGraphFileHash) GetHashValue() string`

GetHashValue returns the HashValue field if non-nil, zero value otherwise.

### GetHashValueOk

`func (o *MicrosoftGraphFileHash) GetHashValueOk() (string, bool)`

GetHashValueOk returns a tuple with the HashValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHashValue

`func (o *MicrosoftGraphFileHash) HasHashValue() bool`

HasHashValue returns a boolean if a field has been set.

### SetHashValue

`func (o *MicrosoftGraphFileHash) SetHashValue(v string)`

SetHashValue gets a reference to the given string and assigns it to the HashValue field.

### SetHashValueExplicitNull

`func (o *MicrosoftGraphFileHash) SetHashValueExplicitNull(b bool)`

SetHashValueExplicitNull (un)sets HashValue to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HashValue value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


