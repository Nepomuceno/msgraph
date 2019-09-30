# MicrosoftGraphCustomTimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Bias** | Pointer to **int32** |  | [optional] 
**StandardOffset** | Pointer to [**AnyOfmicrosoftGraphStandardTimeZoneOffset**](anyOf&lt;microsoft.graph.standardTimeZoneOffset&gt;.md) |  | [optional] 
**DaylightOffset** | Pointer to [**AnyOfmicrosoftGraphDaylightTimeZoneOffset**](anyOf&lt;microsoft.graph.daylightTimeZoneOffset&gt;.md) |  | [optional] 

## Methods

### GetName

`func (o *MicrosoftGraphCustomTimeZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphCustomTimeZone) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphCustomTimeZone) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphCustomTimeZone) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphCustomTimeZone) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetBias

`func (o *MicrosoftGraphCustomTimeZone) GetBias() int32`

GetBias returns the Bias field if non-nil, zero value otherwise.

### GetBiasOk

`func (o *MicrosoftGraphCustomTimeZone) GetBiasOk() (int32, bool)`

GetBiasOk returns a tuple with the Bias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBias

`func (o *MicrosoftGraphCustomTimeZone) HasBias() bool`

HasBias returns a boolean if a field has been set.

### SetBias

`func (o *MicrosoftGraphCustomTimeZone) SetBias(v int32)`

SetBias gets a reference to the given int32 and assigns it to the Bias field.

### SetBiasExplicitNull

`func (o *MicrosoftGraphCustomTimeZone) SetBiasExplicitNull(b bool)`

SetBiasExplicitNull (un)sets Bias to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Bias value is set to nil even if false is passed
### GetStandardOffset

`func (o *MicrosoftGraphCustomTimeZone) GetStandardOffset() AnyOfmicrosoftGraphStandardTimeZoneOffset`

GetStandardOffset returns the StandardOffset field if non-nil, zero value otherwise.

### GetStandardOffsetOk

`func (o *MicrosoftGraphCustomTimeZone) GetStandardOffsetOk() (AnyOfmicrosoftGraphStandardTimeZoneOffset, bool)`

GetStandardOffsetOk returns a tuple with the StandardOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStandardOffset

`func (o *MicrosoftGraphCustomTimeZone) HasStandardOffset() bool`

HasStandardOffset returns a boolean if a field has been set.

### SetStandardOffset

`func (o *MicrosoftGraphCustomTimeZone) SetStandardOffset(v AnyOfmicrosoftGraphStandardTimeZoneOffset)`

SetStandardOffset gets a reference to the given AnyOfmicrosoftGraphStandardTimeZoneOffset and assigns it to the StandardOffset field.

### SetStandardOffsetExplicitNull

`func (o *MicrosoftGraphCustomTimeZone) SetStandardOffsetExplicitNull(b bool)`

SetStandardOffsetExplicitNull (un)sets StandardOffset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StandardOffset value is set to nil even if false is passed
### GetDaylightOffset

`func (o *MicrosoftGraphCustomTimeZone) GetDaylightOffset() AnyOfmicrosoftGraphDaylightTimeZoneOffset`

GetDaylightOffset returns the DaylightOffset field if non-nil, zero value otherwise.

### GetDaylightOffsetOk

`func (o *MicrosoftGraphCustomTimeZone) GetDaylightOffsetOk() (AnyOfmicrosoftGraphDaylightTimeZoneOffset, bool)`

GetDaylightOffsetOk returns a tuple with the DaylightOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaylightOffset

`func (o *MicrosoftGraphCustomTimeZone) HasDaylightOffset() bool`

HasDaylightOffset returns a boolean if a field has been set.

### SetDaylightOffset

`func (o *MicrosoftGraphCustomTimeZone) SetDaylightOffset(v AnyOfmicrosoftGraphDaylightTimeZoneOffset)`

SetDaylightOffset gets a reference to the given AnyOfmicrosoftGraphDaylightTimeZoneOffset and assigns it to the DaylightOffset field.

### SetDaylightOffsetExplicitNull

`func (o *MicrosoftGraphCustomTimeZone) SetDaylightOffsetExplicitNull(b bool)`

SetDaylightOffsetExplicitNull (un)sets DaylightOffset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DaylightOffset value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


