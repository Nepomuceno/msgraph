# MicrosoftGraphMeetingTimeSuggestionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeetingTimeSuggestions** | Pointer to [**[]AnyOfmicrosoftGraphMeetingTimeSuggestion**](anyOf&lt;microsoft.graph.meetingTimeSuggestion&gt;.md) |  | [optional] 
**EmptySuggestionsReason** | Pointer to **string** |  | [optional] 

## Methods

### GetMeetingTimeSuggestions

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetMeetingTimeSuggestions() []AnyOfmicrosoftGraphMeetingTimeSuggestion`

GetMeetingTimeSuggestions returns the MeetingTimeSuggestions field if non-nil, zero value otherwise.

### GetMeetingTimeSuggestionsOk

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetMeetingTimeSuggestionsOk() ([]AnyOfmicrosoftGraphMeetingTimeSuggestion, bool)`

GetMeetingTimeSuggestionsOk returns a tuple with the MeetingTimeSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeetingTimeSuggestions

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) HasMeetingTimeSuggestions() bool`

HasMeetingTimeSuggestions returns a boolean if a field has been set.

### SetMeetingTimeSuggestions

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) SetMeetingTimeSuggestions(v []AnyOfmicrosoftGraphMeetingTimeSuggestion)`

SetMeetingTimeSuggestions gets a reference to the given []AnyOfmicrosoftGraphMeetingTimeSuggestion and assigns it to the MeetingTimeSuggestions field.

### GetEmptySuggestionsReason

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetEmptySuggestionsReason() string`

GetEmptySuggestionsReason returns the EmptySuggestionsReason field if non-nil, zero value otherwise.

### GetEmptySuggestionsReasonOk

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) GetEmptySuggestionsReasonOk() (string, bool)`

GetEmptySuggestionsReasonOk returns a tuple with the EmptySuggestionsReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmptySuggestionsReason

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) HasEmptySuggestionsReason() bool`

HasEmptySuggestionsReason returns a boolean if a field has been set.

### SetEmptySuggestionsReason

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) SetEmptySuggestionsReason(v string)`

SetEmptySuggestionsReason gets a reference to the given string and assigns it to the EmptySuggestionsReason field.

### SetEmptySuggestionsReasonExplicitNull

`func (o *MicrosoftGraphMeetingTimeSuggestionsResult) SetEmptySuggestionsReasonExplicitNull(b bool)`

SetEmptySuggestionsReasonExplicitNull (un)sets EmptySuggestionsReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmptySuggestionsReason value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


