# MicrosoftGraphAutomaticRepliesMailTips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**MessageLanguage** | Pointer to [**AnyOfmicrosoftGraphLocaleInfo**](anyOf&lt;microsoft.graph.localeInfo&gt;.md) |  | [optional] 
**ScheduledStartTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**ScheduledEndTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 

## Methods

### GetMessage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### SetMessageExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessageExplicitNull(b bool)`

SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Message value is set to nil even if false is passed
### GetMessageLanguage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessageLanguage() AnyOfmicrosoftGraphLocaleInfo`

GetMessageLanguage returns the MessageLanguage field if non-nil, zero value otherwise.

### GetMessageLanguageOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetMessageLanguageOk() (AnyOfmicrosoftGraphLocaleInfo, bool)`

GetMessageLanguageOk returns a tuple with the MessageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageLanguage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasMessageLanguage() bool`

HasMessageLanguage returns a boolean if a field has been set.

### SetMessageLanguage

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessageLanguage(v AnyOfmicrosoftGraphLocaleInfo)`

SetMessageLanguage gets a reference to the given AnyOfmicrosoftGraphLocaleInfo and assigns it to the MessageLanguage field.

### SetMessageLanguageExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetMessageLanguageExplicitNull(b bool)`

SetMessageLanguageExplicitNull (un)sets MessageLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MessageLanguage value is set to nil even if false is passed
### GetScheduledStartTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledStartTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledStartTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledStartTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ScheduledStartTime field.

### SetScheduledStartTimeExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledStartTimeExplicitNull(b bool)`

SetScheduledStartTimeExplicitNull (un)sets ScheduledStartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScheduledStartTime value is set to nil even if false is passed
### GetScheduledEndTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledEndTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *MicrosoftGraphAutomaticRepliesMailTips) GetScheduledEndTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledEndTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTime

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ScheduledEndTime field.

### SetScheduledEndTimeExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesMailTips) SetScheduledEndTimeExplicitNull(b bool)`

SetScheduledEndTimeExplicitNull (un)sets ScheduledEndTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScheduledEndTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


