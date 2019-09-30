# MacOsCustomConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadName** | Pointer to **string** | Name that is displayed to the user. | [optional] 
**PayloadFileName** | Pointer to **string** | Payload file name (*.mobileconfig | *.xml). | [optional] 
**Payload** | Pointer to **string** | Payload. (UTF8 encoded byte array) | [optional] 

## Methods

### GetPayloadName

`func (o *MacOsCustomConfiguration) GetPayloadName() string`

GetPayloadName returns the PayloadName field if non-nil, zero value otherwise.

### GetPayloadNameOk

`func (o *MacOsCustomConfiguration) GetPayloadNameOk() (string, bool)`

GetPayloadNameOk returns a tuple with the PayloadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayloadName

`func (o *MacOsCustomConfiguration) HasPayloadName() bool`

HasPayloadName returns a boolean if a field has been set.

### SetPayloadName

`func (o *MacOsCustomConfiguration) SetPayloadName(v string)`

SetPayloadName gets a reference to the given string and assigns it to the PayloadName field.

### GetPayloadFileName

`func (o *MacOsCustomConfiguration) GetPayloadFileName() string`

GetPayloadFileName returns the PayloadFileName field if non-nil, zero value otherwise.

### GetPayloadFileNameOk

`func (o *MacOsCustomConfiguration) GetPayloadFileNameOk() (string, bool)`

GetPayloadFileNameOk returns a tuple with the PayloadFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayloadFileName

`func (o *MacOsCustomConfiguration) HasPayloadFileName() bool`

HasPayloadFileName returns a boolean if a field has been set.

### SetPayloadFileName

`func (o *MacOsCustomConfiguration) SetPayloadFileName(v string)`

SetPayloadFileName gets a reference to the given string and assigns it to the PayloadFileName field.

### SetPayloadFileNameExplicitNull

`func (o *MacOsCustomConfiguration) SetPayloadFileNameExplicitNull(b bool)`

SetPayloadFileNameExplicitNull (un)sets PayloadFileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PayloadFileName value is set to nil even if false is passed
### GetPayload

`func (o *MacOsCustomConfiguration) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *MacOsCustomConfiguration) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *MacOsCustomConfiguration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *MacOsCustomConfiguration) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


