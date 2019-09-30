# MicrosoftGraphDomainState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**LastActionDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetStatus

`func (o *MicrosoftGraphDomainState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDomainState) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphDomainState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphDomainState) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphDomainState) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetOperation

`func (o *MicrosoftGraphDomainState) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MicrosoftGraphDomainState) GetOperationOk() (string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperation

`func (o *MicrosoftGraphDomainState) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperation

`func (o *MicrosoftGraphDomainState) SetOperation(v string)`

SetOperation gets a reference to the given string and assigns it to the Operation field.

### SetOperationExplicitNull

`func (o *MicrosoftGraphDomainState) SetOperationExplicitNull(b bool)`

SetOperationExplicitNull (un)sets Operation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Operation value is set to nil even if false is passed
### GetLastActionDateTime

`func (o *MicrosoftGraphDomainState) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *MicrosoftGraphDomainState) GetLastActionDateTimeOk() (time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActionDateTime

`func (o *MicrosoftGraphDomainState) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTime

`func (o *MicrosoftGraphDomainState) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.

### SetLastActionDateTimeExplicitNull

`func (o *MicrosoftGraphDomainState) SetLastActionDateTimeExplicitNull(b bool)`

SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastActionDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


