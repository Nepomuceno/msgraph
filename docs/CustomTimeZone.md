# CustomTimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bias** | Pointer to **int32** |  | [optional] 
**StandardOffset** | Pointer to [**AnyOfmicrosoftGraphStandardTimeZoneOffset**](anyOf&lt;microsoft.graph.standardTimeZoneOffset&gt;.md) |  | [optional] 
**DaylightOffset** | Pointer to [**AnyOfmicrosoftGraphDaylightTimeZoneOffset**](anyOf&lt;microsoft.graph.daylightTimeZoneOffset&gt;.md) |  | [optional] 

## Methods

### GetBias

`func (o *CustomTimeZone) GetBias() int32`

GetBias returns the Bias field if non-nil, zero value otherwise.

### GetBiasOk

`func (o *CustomTimeZone) GetBiasOk() (int32, bool)`

GetBiasOk returns a tuple with the Bias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBias

`func (o *CustomTimeZone) HasBias() bool`

HasBias returns a boolean if a field has been set.

### SetBias

`func (o *CustomTimeZone) SetBias(v int32)`

SetBias gets a reference to the given int32 and assigns it to the Bias field.

### SetBiasExplicitNull

`func (o *CustomTimeZone) SetBiasExplicitNull(b bool)`

SetBiasExplicitNull (un)sets Bias to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Bias value is set to nil even if false is passed
### GetStandardOffset

`func (o *CustomTimeZone) GetStandardOffset() AnyOfmicrosoftGraphStandardTimeZoneOffset`

GetStandardOffset returns the StandardOffset field if non-nil, zero value otherwise.

### GetStandardOffsetOk

`func (o *CustomTimeZone) GetStandardOffsetOk() (AnyOfmicrosoftGraphStandardTimeZoneOffset, bool)`

GetStandardOffsetOk returns a tuple with the StandardOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStandardOffset

`func (o *CustomTimeZone) HasStandardOffset() bool`

HasStandardOffset returns a boolean if a field has been set.

### SetStandardOffset

`func (o *CustomTimeZone) SetStandardOffset(v AnyOfmicrosoftGraphStandardTimeZoneOffset)`

SetStandardOffset gets a reference to the given AnyOfmicrosoftGraphStandardTimeZoneOffset and assigns it to the StandardOffset field.

### SetStandardOffsetExplicitNull

`func (o *CustomTimeZone) SetStandardOffsetExplicitNull(b bool)`

SetStandardOffsetExplicitNull (un)sets StandardOffset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StandardOffset value is set to nil even if false is passed
### GetDaylightOffset

`func (o *CustomTimeZone) GetDaylightOffset() AnyOfmicrosoftGraphDaylightTimeZoneOffset`

GetDaylightOffset returns the DaylightOffset field if non-nil, zero value otherwise.

### GetDaylightOffsetOk

`func (o *CustomTimeZone) GetDaylightOffsetOk() (AnyOfmicrosoftGraphDaylightTimeZoneOffset, bool)`

GetDaylightOffsetOk returns a tuple with the DaylightOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaylightOffset

`func (o *CustomTimeZone) HasDaylightOffset() bool`

HasDaylightOffset returns a boolean if a field has been set.

### SetDaylightOffset

`func (o *CustomTimeZone) SetDaylightOffset(v AnyOfmicrosoftGraphDaylightTimeZoneOffset)`

SetDaylightOffset gets a reference to the given AnyOfmicrosoftGraphDaylightTimeZoneOffset and assigns it to the DaylightOffset field.

### SetDaylightOffsetExplicitNull

`func (o *CustomTimeZone) SetDaylightOffsetExplicitNull(b bool)`

SetDaylightOffsetExplicitNull (un)sets DaylightOffset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DaylightOffset value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


