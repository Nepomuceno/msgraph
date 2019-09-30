# DeviceManagementTroubleshootingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDateTime** | Pointer to [**time.Time**](time.Time.md) | Time when the event occurred . | [optional] 
**CorrelationId** | Pointer to **string** | Id used for tracing the failure in the service. | [optional] 

## Methods

### GetEventDateTime

`func (o *DeviceManagementTroubleshootingEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *DeviceManagementTroubleshootingEvent) GetEventDateTimeOk() (time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventDateTime

`func (o *DeviceManagementTroubleshootingEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTime

`func (o *DeviceManagementTroubleshootingEvent) SetEventDateTime(v time.Time)`

SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.

### GetCorrelationId

`func (o *DeviceManagementTroubleshootingEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DeviceManagementTroubleshootingEvent) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *DeviceManagementTroubleshootingEvent) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *DeviceManagementTroubleshootingEvent) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *DeviceManagementTroubleshootingEvent) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


