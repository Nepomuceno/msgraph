# MicrosoftGraphAttendee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphAttendeeType**](anyOf&lt;microsoft.graph.attendeeType&gt;.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphResponseStatus**](anyOf&lt;microsoft.graph.responseStatus&gt;.md) |  | [optional] 

## Methods

### GetEmailAddress

`func (o *MicrosoftGraphAttendee) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphAttendee) GetEmailAddressOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddress

`func (o *MicrosoftGraphAttendee) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddress

`func (o *MicrosoftGraphAttendee) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddress field.

### SetEmailAddressExplicitNull

`func (o *MicrosoftGraphAttendee) SetEmailAddressExplicitNull(b bool)`

SetEmailAddressExplicitNull (un)sets EmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmailAddress value is set to nil even if false is passed
### GetType

`func (o *MicrosoftGraphAttendee) GetType() AnyOfmicrosoftGraphAttendeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphAttendee) GetTypeOk() (AnyOfmicrosoftGraphAttendeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *MicrosoftGraphAttendee) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *MicrosoftGraphAttendee) SetType(v AnyOfmicrosoftGraphAttendeeType)`

SetType gets a reference to the given AnyOfmicrosoftGraphAttendeeType and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *MicrosoftGraphAttendee) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphAttendee) GetStatus() AnyOfmicrosoftGraphResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAttendee) GetStatusOk() (AnyOfmicrosoftGraphResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphAttendee) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphAttendee) SetStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphResponseStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphAttendee) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


