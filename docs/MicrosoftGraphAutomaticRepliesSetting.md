# MicrosoftGraphAutomaticRepliesSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AnyOfmicrosoftGraphAutomaticRepliesStatus**](anyOf&lt;microsoft.graph.automaticRepliesStatus&gt;.md) |  | [optional] 
**ExternalAudience** | Pointer to [**AnyOfmicrosoftGraphExternalAudienceScope**](anyOf&lt;microsoft.graph.externalAudienceScope&gt;.md) |  | [optional] 
**ScheduledStartDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**ScheduledEndDateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**InternalReplyMessage** | Pointer to **string** |  | [optional] 
**ExternalReplyMessage** | Pointer to **string** |  | [optional] 

## Methods

### GetStatus

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetStatus() AnyOfmicrosoftGraphAutomaticRepliesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetStatusOk() (AnyOfmicrosoftGraphAutomaticRepliesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetStatus(v AnyOfmicrosoftGraphAutomaticRepliesStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphAutomaticRepliesStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetExternalAudience

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalAudience() AnyOfmicrosoftGraphExternalAudienceScope`

GetExternalAudience returns the ExternalAudience field if non-nil, zero value otherwise.

### GetExternalAudienceOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalAudienceOk() (AnyOfmicrosoftGraphExternalAudienceScope, bool)`

GetExternalAudienceOk returns a tuple with the ExternalAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalAudience

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasExternalAudience() bool`

HasExternalAudience returns a boolean if a field has been set.

### SetExternalAudience

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalAudience(v AnyOfmicrosoftGraphExternalAudienceScope)`

SetExternalAudience gets a reference to the given AnyOfmicrosoftGraphExternalAudienceScope and assigns it to the ExternalAudience field.

### SetExternalAudienceExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalAudienceExplicitNull(b bool)`

SetExternalAudienceExplicitNull (un)sets ExternalAudience to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalAudience value is set to nil even if false is passed
### GetScheduledStartDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledStartDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledStartDateTime returns the ScheduledStartDateTime field if non-nil, zero value otherwise.

### GetScheduledStartDateTimeOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledStartDateTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledStartDateTimeOk returns a tuple with the ScheduledStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledStartDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasScheduledStartDateTime() bool`

HasScheduledStartDateTime returns a boolean if a field has been set.

### SetScheduledStartDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledStartDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledStartDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ScheduledStartDateTime field.

### SetScheduledStartDateTimeExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledStartDateTimeExplicitNull(b bool)`

SetScheduledStartDateTimeExplicitNull (un)sets ScheduledStartDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScheduledStartDateTime value is set to nil even if false is passed
### GetScheduledEndDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledEndDateTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetScheduledEndDateTime returns the ScheduledEndDateTime field if non-nil, zero value otherwise.

### GetScheduledEndDateTimeOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledEndDateTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetScheduledEndDateTimeOk returns a tuple with the ScheduledEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledEndDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasScheduledEndDateTime() bool`

HasScheduledEndDateTime returns a boolean if a field has been set.

### SetScheduledEndDateTime

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledEndDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetScheduledEndDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ScheduledEndDateTime field.

### SetScheduledEndDateTimeExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledEndDateTimeExplicitNull(b bool)`

SetScheduledEndDateTimeExplicitNull (un)sets ScheduledEndDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScheduledEndDateTime value is set to nil even if false is passed
### GetInternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetInternalReplyMessage() string`

GetInternalReplyMessage returns the InternalReplyMessage field if non-nil, zero value otherwise.

### GetInternalReplyMessageOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetInternalReplyMessageOk() (string, bool)`

GetInternalReplyMessageOk returns a tuple with the InternalReplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasInternalReplyMessage() bool`

HasInternalReplyMessage returns a boolean if a field has been set.

### SetInternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetInternalReplyMessage(v string)`

SetInternalReplyMessage gets a reference to the given string and assigns it to the InternalReplyMessage field.

### SetInternalReplyMessageExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetInternalReplyMessageExplicitNull(b bool)`

SetInternalReplyMessageExplicitNull (un)sets InternalReplyMessage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InternalReplyMessage value is set to nil even if false is passed
### GetExternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalReplyMessage() string`

GetExternalReplyMessage returns the ExternalReplyMessage field if non-nil, zero value otherwise.

### GetExternalReplyMessageOk

`func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalReplyMessageOk() (string, bool)`

GetExternalReplyMessageOk returns a tuple with the ExternalReplyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) HasExternalReplyMessage() bool`

HasExternalReplyMessage returns a boolean if a field has been set.

### SetExternalReplyMessage

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalReplyMessage(v string)`

SetExternalReplyMessage gets a reference to the given string and assigns it to the ExternalReplyMessage field.

### SetExternalReplyMessageExplicitNull

`func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalReplyMessageExplicitNull(b bool)`

SetExternalReplyMessageExplicitNull (un)sets ExternalReplyMessage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalReplyMessage value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


