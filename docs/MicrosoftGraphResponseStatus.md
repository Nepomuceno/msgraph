# MicrosoftGraphResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**AnyOfmicrosoftGraphResponseType**](anyOf&lt;microsoft.graph.responseType&gt;.md) |  | [optional] 
**Time** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetResponse

`func (o *MicrosoftGraphResponseStatus) GetResponse() AnyOfmicrosoftGraphResponseType`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MicrosoftGraphResponseStatus) GetResponseOk() (AnyOfmicrosoftGraphResponseType, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponse

`func (o *MicrosoftGraphResponseStatus) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponse

`func (o *MicrosoftGraphResponseStatus) SetResponse(v AnyOfmicrosoftGraphResponseType)`

SetResponse gets a reference to the given AnyOfmicrosoftGraphResponseType and assigns it to the Response field.

### SetResponseExplicitNull

`func (o *MicrosoftGraphResponseStatus) SetResponseExplicitNull(b bool)`

SetResponseExplicitNull (un)sets Response to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Response value is set to nil even if false is passed
### GetTime

`func (o *MicrosoftGraphResponseStatus) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MicrosoftGraphResponseStatus) GetTimeOk() (time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *MicrosoftGraphResponseStatus) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *MicrosoftGraphResponseStatus) SetTime(v time.Time)`

SetTime gets a reference to the given time.Time and assigns it to the Time field.

### SetTimeExplicitNull

`func (o *MicrosoftGraphResponseStatus) SetTimeExplicitNull(b bool)`

SetTimeExplicitNull (un)sets Time to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Time value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


