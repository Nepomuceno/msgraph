# Attendee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AnyOfmicrosoftGraphResponseStatus**](anyOf&lt;microsoft.graph.responseStatus&gt;.md) |  | [optional] 

## Methods

### GetStatus

`func (o *Attendee) GetStatus() AnyOfmicrosoftGraphResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Attendee) GetStatusOk() (AnyOfmicrosoftGraphResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *Attendee) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *Attendee) SetStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphResponseStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *Attendee) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


