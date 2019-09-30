# MicrosoftGraphDeviceManagementTroubleshootingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**EventDateTime** | Pointer to [**time.Time**](time.Time.md) | Time when the event occurred . | [optional] 
**CorrelationId** | Pointer to **string** | Id used for tracing the failure in the service. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetEventDateTime

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTimeOk() (time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventDateTime

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTime

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetEventDateTime(v time.Time)`

SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.

### GetCorrelationId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


