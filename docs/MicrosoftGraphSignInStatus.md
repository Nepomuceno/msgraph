# MicrosoftGraphSignInStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**AdditionalDetails** | Pointer to **string** |  | [optional] 

## Methods

### GetErrorCode

`func (o *MicrosoftGraphSignInStatus) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphSignInStatus) GetErrorCodeOk() (int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *MicrosoftGraphSignInStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *MicrosoftGraphSignInStatus) SetErrorCode(v int32)`

SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.

### SetErrorCodeExplicitNull

`func (o *MicrosoftGraphSignInStatus) SetErrorCodeExplicitNull(b bool)`

SetErrorCodeExplicitNull (un)sets ErrorCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ErrorCode value is set to nil even if false is passed
### GetFailureReason

`func (o *MicrosoftGraphSignInStatus) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MicrosoftGraphSignInStatus) GetFailureReasonOk() (string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailureReason

`func (o *MicrosoftGraphSignInStatus) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReason

`func (o *MicrosoftGraphSignInStatus) SetFailureReason(v string)`

SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.

### SetFailureReasonExplicitNull

`func (o *MicrosoftGraphSignInStatus) SetFailureReasonExplicitNull(b bool)`

SetFailureReasonExplicitNull (un)sets FailureReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FailureReason value is set to nil even if false is passed
### GetAdditionalDetails

`func (o *MicrosoftGraphSignInStatus) GetAdditionalDetails() string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *MicrosoftGraphSignInStatus) GetAdditionalDetailsOk() (string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalDetails

`func (o *MicrosoftGraphSignInStatus) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetails

`func (o *MicrosoftGraphSignInStatus) SetAdditionalDetails(v string)`

SetAdditionalDetails gets a reference to the given string and assigns it to the AdditionalDetails field.

### SetAdditionalDetailsExplicitNull

`func (o *MicrosoftGraphSignInStatus) SetAdditionalDetailsExplicitNull(b bool)`

SetAdditionalDetailsExplicitNull (un)sets AdditionalDetails to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AdditionalDetails value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


