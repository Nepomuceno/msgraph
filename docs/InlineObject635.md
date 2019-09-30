# InlineObject635

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendees** | Pointer to [**[]AnyOfmicrosoftGraphAttendeeBase**](anyOf&lt;microsoft.graph.attendeeBase&gt;.md) |  | [optional] 
**LocationConstraint** | Pointer to [**AnyOfmicrosoftGraphLocationConstraint**](anyOf&lt;microsoft.graph.locationConstraint&gt;.md) |  | [optional] 
**TimeConstraint** | Pointer to [**AnyOfmicrosoftGraphTimeConstraint**](anyOf&lt;microsoft.graph.timeConstraint&gt;.md) |  | [optional] 
**MeetingDuration** | Pointer to **string** |  | [optional] 
**MaxCandidates** | Pointer to **int32** |  | [optional] 
**IsOrganizerOptional** | Pointer to **bool** |  | [optional] [default to false]
**ReturnSuggestionReasons** | Pointer to **bool** |  | [optional] [default to false]
**MinimumAttendeePercentage** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 

## Methods

### GetAttendees

`func (o *InlineObject635) GetAttendees() []AnyOfmicrosoftGraphAttendeeBase`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *InlineObject635) GetAttendeesOk() ([]AnyOfmicrosoftGraphAttendeeBase, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttendees

`func (o *InlineObject635) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### SetAttendees

`func (o *InlineObject635) SetAttendees(v []AnyOfmicrosoftGraphAttendeeBase)`

SetAttendees gets a reference to the given []AnyOfmicrosoftGraphAttendeeBase and assigns it to the Attendees field.

### GetLocationConstraint

`func (o *InlineObject635) GetLocationConstraint() AnyOfmicrosoftGraphLocationConstraint`

GetLocationConstraint returns the LocationConstraint field if non-nil, zero value otherwise.

### GetLocationConstraintOk

`func (o *InlineObject635) GetLocationConstraintOk() (AnyOfmicrosoftGraphLocationConstraint, bool)`

GetLocationConstraintOk returns a tuple with the LocationConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationConstraint

`func (o *InlineObject635) HasLocationConstraint() bool`

HasLocationConstraint returns a boolean if a field has been set.

### SetLocationConstraint

`func (o *InlineObject635) SetLocationConstraint(v AnyOfmicrosoftGraphLocationConstraint)`

SetLocationConstraint gets a reference to the given AnyOfmicrosoftGraphLocationConstraint and assigns it to the LocationConstraint field.

### SetLocationConstraintExplicitNull

`func (o *InlineObject635) SetLocationConstraintExplicitNull(b bool)`

SetLocationConstraintExplicitNull (un)sets LocationConstraint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LocationConstraint value is set to nil even if false is passed
### GetTimeConstraint

`func (o *InlineObject635) GetTimeConstraint() AnyOfmicrosoftGraphTimeConstraint`

GetTimeConstraint returns the TimeConstraint field if non-nil, zero value otherwise.

### GetTimeConstraintOk

`func (o *InlineObject635) GetTimeConstraintOk() (AnyOfmicrosoftGraphTimeConstraint, bool)`

GetTimeConstraintOk returns a tuple with the TimeConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeConstraint

`func (o *InlineObject635) HasTimeConstraint() bool`

HasTimeConstraint returns a boolean if a field has been set.

### SetTimeConstraint

`func (o *InlineObject635) SetTimeConstraint(v AnyOfmicrosoftGraphTimeConstraint)`

SetTimeConstraint gets a reference to the given AnyOfmicrosoftGraphTimeConstraint and assigns it to the TimeConstraint field.

### SetTimeConstraintExplicitNull

`func (o *InlineObject635) SetTimeConstraintExplicitNull(b bool)`

SetTimeConstraintExplicitNull (un)sets TimeConstraint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TimeConstraint value is set to nil even if false is passed
### GetMeetingDuration

`func (o *InlineObject635) GetMeetingDuration() string`

GetMeetingDuration returns the MeetingDuration field if non-nil, zero value otherwise.

### GetMeetingDurationOk

`func (o *InlineObject635) GetMeetingDurationOk() (string, bool)`

GetMeetingDurationOk returns a tuple with the MeetingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeetingDuration

`func (o *InlineObject635) HasMeetingDuration() bool`

HasMeetingDuration returns a boolean if a field has been set.

### SetMeetingDuration

`func (o *InlineObject635) SetMeetingDuration(v string)`

SetMeetingDuration gets a reference to the given string and assigns it to the MeetingDuration field.

### SetMeetingDurationExplicitNull

`func (o *InlineObject635) SetMeetingDurationExplicitNull(b bool)`

SetMeetingDurationExplicitNull (un)sets MeetingDuration to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MeetingDuration value is set to nil even if false is passed
### GetMaxCandidates

`func (o *InlineObject635) GetMaxCandidates() int32`

GetMaxCandidates returns the MaxCandidates field if non-nil, zero value otherwise.

### GetMaxCandidatesOk

`func (o *InlineObject635) GetMaxCandidatesOk() (int32, bool)`

GetMaxCandidatesOk returns a tuple with the MaxCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxCandidates

`func (o *InlineObject635) HasMaxCandidates() bool`

HasMaxCandidates returns a boolean if a field has been set.

### SetMaxCandidates

`func (o *InlineObject635) SetMaxCandidates(v int32)`

SetMaxCandidates gets a reference to the given int32 and assigns it to the MaxCandidates field.

### SetMaxCandidatesExplicitNull

`func (o *InlineObject635) SetMaxCandidatesExplicitNull(b bool)`

SetMaxCandidatesExplicitNull (un)sets MaxCandidates to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaxCandidates value is set to nil even if false is passed
### GetIsOrganizerOptional

`func (o *InlineObject635) GetIsOrganizerOptional() bool`

GetIsOrganizerOptional returns the IsOrganizerOptional field if non-nil, zero value otherwise.

### GetIsOrganizerOptionalOk

`func (o *InlineObject635) GetIsOrganizerOptionalOk() (bool, bool)`

GetIsOrganizerOptionalOk returns a tuple with the IsOrganizerOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsOrganizerOptional

`func (o *InlineObject635) HasIsOrganizerOptional() bool`

HasIsOrganizerOptional returns a boolean if a field has been set.

### SetIsOrganizerOptional

`func (o *InlineObject635) SetIsOrganizerOptional(v bool)`

SetIsOrganizerOptional gets a reference to the given bool and assigns it to the IsOrganizerOptional field.

### SetIsOrganizerOptionalExplicitNull

`func (o *InlineObject635) SetIsOrganizerOptionalExplicitNull(b bool)`

SetIsOrganizerOptionalExplicitNull (un)sets IsOrganizerOptional to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsOrganizerOptional value is set to nil even if false is passed
### GetReturnSuggestionReasons

`func (o *InlineObject635) GetReturnSuggestionReasons() bool`

GetReturnSuggestionReasons returns the ReturnSuggestionReasons field if non-nil, zero value otherwise.

### GetReturnSuggestionReasonsOk

`func (o *InlineObject635) GetReturnSuggestionReasonsOk() (bool, bool)`

GetReturnSuggestionReasonsOk returns a tuple with the ReturnSuggestionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReturnSuggestionReasons

`func (o *InlineObject635) HasReturnSuggestionReasons() bool`

HasReturnSuggestionReasons returns a boolean if a field has been set.

### SetReturnSuggestionReasons

`func (o *InlineObject635) SetReturnSuggestionReasons(v bool)`

SetReturnSuggestionReasons gets a reference to the given bool and assigns it to the ReturnSuggestionReasons field.

### SetReturnSuggestionReasonsExplicitNull

`func (o *InlineObject635) SetReturnSuggestionReasonsExplicitNull(b bool)`

SetReturnSuggestionReasonsExplicitNull (un)sets ReturnSuggestionReasons to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReturnSuggestionReasons value is set to nil even if false is passed
### GetMinimumAttendeePercentage

`func (o *InlineObject635) GetMinimumAttendeePercentage() AnyOfnumberstringstring`

GetMinimumAttendeePercentage returns the MinimumAttendeePercentage field if non-nil, zero value otherwise.

### GetMinimumAttendeePercentageOk

`func (o *InlineObject635) GetMinimumAttendeePercentageOk() (AnyOfnumberstringstring, bool)`

GetMinimumAttendeePercentageOk returns a tuple with the MinimumAttendeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumAttendeePercentage

`func (o *InlineObject635) HasMinimumAttendeePercentage() bool`

HasMinimumAttendeePercentage returns a boolean if a field has been set.

### SetMinimumAttendeePercentage

`func (o *InlineObject635) SetMinimumAttendeePercentage(v AnyOfnumberstringstring)`

SetMinimumAttendeePercentage gets a reference to the given AnyOfnumberstringstring and assigns it to the MinimumAttendeePercentage field.

### SetMinimumAttendeePercentageExplicitNull

`func (o *InlineObject635) SetMinimumAttendeePercentageExplicitNull(b bool)`

SetMinimumAttendeePercentageExplicitNull (un)sets MinimumAttendeePercentage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumAttendeePercentage value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

