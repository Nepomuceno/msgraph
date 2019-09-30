# MicrosoftGraphPatternedRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to [**AnyOfmicrosoftGraphRecurrencePattern**](anyOf&lt;microsoft.graph.recurrencePattern&gt;.md) |  | [optional] 
**Range** | Pointer to [**AnyOfmicrosoftGraphRecurrenceRange**](anyOf&lt;microsoft.graph.recurrenceRange&gt;.md) |  | [optional] 

## Methods

### GetPattern

`func (o *MicrosoftGraphPatternedRecurrence) GetPattern() AnyOfmicrosoftGraphRecurrencePattern`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *MicrosoftGraphPatternedRecurrence) GetPatternOk() (AnyOfmicrosoftGraphRecurrencePattern, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPattern

`func (o *MicrosoftGraphPatternedRecurrence) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPattern

`func (o *MicrosoftGraphPatternedRecurrence) SetPattern(v AnyOfmicrosoftGraphRecurrencePattern)`

SetPattern gets a reference to the given AnyOfmicrosoftGraphRecurrencePattern and assigns it to the Pattern field.

### SetPatternExplicitNull

`func (o *MicrosoftGraphPatternedRecurrence) SetPatternExplicitNull(b bool)`

SetPatternExplicitNull (un)sets Pattern to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pattern value is set to nil even if false is passed
### GetRange

`func (o *MicrosoftGraphPatternedRecurrence) GetRange() AnyOfmicrosoftGraphRecurrenceRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MicrosoftGraphPatternedRecurrence) GetRangeOk() (AnyOfmicrosoftGraphRecurrenceRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRange

`func (o *MicrosoftGraphPatternedRecurrence) HasRange() bool`

HasRange returns a boolean if a field has been set.

### SetRange

`func (o *MicrosoftGraphPatternedRecurrence) SetRange(v AnyOfmicrosoftGraphRecurrenceRange)`

SetRange gets a reference to the given AnyOfmicrosoftGraphRecurrenceRange and assigns it to the Range field.

### SetRangeExplicitNull

`func (o *MicrosoftGraphPatternedRecurrence) SetRangeExplicitNull(b bool)`

SetRangeExplicitNull (un)sets Range to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Range value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


