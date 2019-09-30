# MicrosoftGraphDateTimeTimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTime** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 

## Methods

### GetDateTime

`func (o *MicrosoftGraphDateTimeTimeZone) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *MicrosoftGraphDateTimeTimeZone) GetDateTimeOk() (string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateTime

`func (o *MicrosoftGraphDateTimeTimeZone) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTime

`func (o *MicrosoftGraphDateTimeTimeZone) SetDateTime(v string)`

SetDateTime gets a reference to the given string and assigns it to the DateTime field.

### GetTimeZone

`func (o *MicrosoftGraphDateTimeTimeZone) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphDateTimeTimeZone) GetTimeZoneOk() (string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeZone

`func (o *MicrosoftGraphDateTimeTimeZone) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZone

`func (o *MicrosoftGraphDateTimeTimeZone) SetTimeZone(v string)`

SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.

### SetTimeZoneExplicitNull

`func (o *MicrosoftGraphDateTimeTimeZone) SetTimeZoneExplicitNull(b bool)`

SetTimeZoneExplicitNull (un)sets TimeZone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TimeZone value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


