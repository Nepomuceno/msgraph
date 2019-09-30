# MicrosoftGraphActivityHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphStatus**](anyOf&lt;microsoft.graph.status&gt;.md) |  | [optional] 
**ActiveDurationSeconds** | Pointer to **int32** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastActiveDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**StartedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UserTimezone** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to [**MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphActivityHistoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphActivityHistoryItem) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphActivityHistoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphActivityHistoryItem) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetStatus

`func (o *MicrosoftGraphActivityHistoryItem) GetStatus() AnyOfmicrosoftGraphStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphActivityHistoryItem) GetStatusOk() (AnyOfmicrosoftGraphStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphActivityHistoryItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphActivityHistoryItem) SetStatus(v AnyOfmicrosoftGraphStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetActiveDurationSeconds

`func (o *MicrosoftGraphActivityHistoryItem) GetActiveDurationSeconds() int32`

GetActiveDurationSeconds returns the ActiveDurationSeconds field if non-nil, zero value otherwise.

### GetActiveDurationSecondsOk

`func (o *MicrosoftGraphActivityHistoryItem) GetActiveDurationSecondsOk() (int32, bool)`

GetActiveDurationSecondsOk returns a tuple with the ActiveDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveDurationSeconds

`func (o *MicrosoftGraphActivityHistoryItem) HasActiveDurationSeconds() bool`

HasActiveDurationSeconds returns a boolean if a field has been set.

### SetActiveDurationSeconds

`func (o *MicrosoftGraphActivityHistoryItem) SetActiveDurationSeconds(v int32)`

SetActiveDurationSeconds gets a reference to the given int32 and assigns it to the ActiveDurationSeconds field.

### SetActiveDurationSecondsExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetActiveDurationSecondsExplicitNull(b bool)`

SetActiveDurationSecondsExplicitNull (un)sets ActiveDurationSeconds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActiveDurationSeconds value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastActiveDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetLastActiveDateTime() time.Time`

GetLastActiveDateTime returns the LastActiveDateTime field if non-nil, zero value otherwise.

### GetLastActiveDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetLastActiveDateTimeOk() (time.Time, bool)`

GetLastActiveDateTimeOk returns a tuple with the LastActiveDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActiveDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasLastActiveDateTime() bool`

HasLastActiveDateTime returns a boolean if a field has been set.

### SetLastActiveDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetLastActiveDateTime(v time.Time)`

SetLastActiveDateTime gets a reference to the given time.Time and assigns it to the LastActiveDateTime field.

### SetLastActiveDateTimeExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetLastActiveDateTimeExplicitNull(b bool)`

SetLastActiveDateTimeExplicitNull (un)sets LastActiveDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastActiveDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### SetExpirationDateTimeExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetExpirationDateTimeExplicitNull(b bool)`

SetExpirationDateTimeExplicitNull (un)sets ExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExpirationDateTime value is set to nil even if false is passed
### GetStartedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) GetStartedDateTime() time.Time`

GetStartedDateTime returns the StartedDateTime field if non-nil, zero value otherwise.

### GetStartedDateTimeOk

`func (o *MicrosoftGraphActivityHistoryItem) GetStartedDateTimeOk() (time.Time, bool)`

GetStartedDateTimeOk returns a tuple with the StartedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) HasStartedDateTime() bool`

HasStartedDateTime returns a boolean if a field has been set.

### SetStartedDateTime

`func (o *MicrosoftGraphActivityHistoryItem) SetStartedDateTime(v time.Time)`

SetStartedDateTime gets a reference to the given time.Time and assigns it to the StartedDateTime field.

### GetUserTimezone

`func (o *MicrosoftGraphActivityHistoryItem) GetUserTimezone() string`

GetUserTimezone returns the UserTimezone field if non-nil, zero value otherwise.

### GetUserTimezoneOk

`func (o *MicrosoftGraphActivityHistoryItem) GetUserTimezoneOk() (string, bool)`

GetUserTimezoneOk returns a tuple with the UserTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserTimezone

`func (o *MicrosoftGraphActivityHistoryItem) HasUserTimezone() bool`

HasUserTimezone returns a boolean if a field has been set.

### SetUserTimezone

`func (o *MicrosoftGraphActivityHistoryItem) SetUserTimezone(v string)`

SetUserTimezone gets a reference to the given string and assigns it to the UserTimezone field.

### SetUserTimezoneExplicitNull

`func (o *MicrosoftGraphActivityHistoryItem) SetUserTimezoneExplicitNull(b bool)`

SetUserTimezoneExplicitNull (un)sets UserTimezone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserTimezone value is set to nil even if false is passed
### GetActivity

`func (o *MicrosoftGraphActivityHistoryItem) GetActivity() MicrosoftGraphUserActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MicrosoftGraphActivityHistoryItem) GetActivityOk() (MicrosoftGraphUserActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivity

`func (o *MicrosoftGraphActivityHistoryItem) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### SetActivity

`func (o *MicrosoftGraphActivityHistoryItem) SetActivity(v MicrosoftGraphUserActivity)`

SetActivity gets a reference to the given MicrosoftGraphUserActivity and assigns it to the Activity field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


