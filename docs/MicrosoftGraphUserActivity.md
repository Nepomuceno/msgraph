# MicrosoftGraphUserActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VisualElements** | Pointer to [**MicrosoftGraphVisualInfo**](microsoft.graph.visualInfo.md) |  | [optional] 
**ActivitySourceHost** | Pointer to **string** |  | [optional] 
**ActivationUrl** | Pointer to **string** |  | [optional] 
**AppActivityId** | Pointer to **string** |  | [optional] 
**AppDisplayName** | Pointer to **string** |  | [optional] 
**ContentUrl** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**FallbackUrl** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UserTimezone** | Pointer to **string** |  | [optional] 
**ContentInfo** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphStatus**](anyOf&lt;microsoft.graph.status&gt;.md) |  | [optional] 
**HistoryItems** | Pointer to [**[]MicrosoftGraphActivityHistoryItem**](microsoft.graph.activityHistoryItem.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphUserActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserActivity) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphUserActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphUserActivity) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetVisualElements

`func (o *MicrosoftGraphUserActivity) GetVisualElements() MicrosoftGraphVisualInfo`

GetVisualElements returns the VisualElements field if non-nil, zero value otherwise.

### GetVisualElementsOk

`func (o *MicrosoftGraphUserActivity) GetVisualElementsOk() (MicrosoftGraphVisualInfo, bool)`

GetVisualElementsOk returns a tuple with the VisualElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisualElements

`func (o *MicrosoftGraphUserActivity) HasVisualElements() bool`

HasVisualElements returns a boolean if a field has been set.

### SetVisualElements

`func (o *MicrosoftGraphUserActivity) SetVisualElements(v MicrosoftGraphVisualInfo)`

SetVisualElements gets a reference to the given MicrosoftGraphVisualInfo and assigns it to the VisualElements field.

### GetActivitySourceHost

`func (o *MicrosoftGraphUserActivity) GetActivitySourceHost() string`

GetActivitySourceHost returns the ActivitySourceHost field if non-nil, zero value otherwise.

### GetActivitySourceHostOk

`func (o *MicrosoftGraphUserActivity) GetActivitySourceHostOk() (string, bool)`

GetActivitySourceHostOk returns a tuple with the ActivitySourceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivitySourceHost

`func (o *MicrosoftGraphUserActivity) HasActivitySourceHost() bool`

HasActivitySourceHost returns a boolean if a field has been set.

### SetActivitySourceHost

`func (o *MicrosoftGraphUserActivity) SetActivitySourceHost(v string)`

SetActivitySourceHost gets a reference to the given string and assigns it to the ActivitySourceHost field.

### GetActivationUrl

`func (o *MicrosoftGraphUserActivity) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *MicrosoftGraphUserActivity) GetActivationUrlOk() (string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationUrl

`func (o *MicrosoftGraphUserActivity) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.

### SetActivationUrl

`func (o *MicrosoftGraphUserActivity) SetActivationUrl(v string)`

SetActivationUrl gets a reference to the given string and assigns it to the ActivationUrl field.

### GetAppActivityId

`func (o *MicrosoftGraphUserActivity) GetAppActivityId() string`

GetAppActivityId returns the AppActivityId field if non-nil, zero value otherwise.

### GetAppActivityIdOk

`func (o *MicrosoftGraphUserActivity) GetAppActivityIdOk() (string, bool)`

GetAppActivityIdOk returns a tuple with the AppActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppActivityId

`func (o *MicrosoftGraphUserActivity) HasAppActivityId() bool`

HasAppActivityId returns a boolean if a field has been set.

### SetAppActivityId

`func (o *MicrosoftGraphUserActivity) SetAppActivityId(v string)`

SetAppActivityId gets a reference to the given string and assigns it to the AppActivityId field.

### GetAppDisplayName

`func (o *MicrosoftGraphUserActivity) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphUserActivity) GetAppDisplayNameOk() (string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDisplayName

`func (o *MicrosoftGraphUserActivity) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphUserActivity) SetAppDisplayName(v string)`

SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.

### SetAppDisplayNameExplicitNull

`func (o *MicrosoftGraphUserActivity) SetAppDisplayNameExplicitNull(b bool)`

SetAppDisplayNameExplicitNull (un)sets AppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppDisplayName value is set to nil even if false is passed
### GetContentUrl

`func (o *MicrosoftGraphUserActivity) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *MicrosoftGraphUserActivity) GetContentUrlOk() (string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentUrl

`func (o *MicrosoftGraphUserActivity) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrl

`func (o *MicrosoftGraphUserActivity) SetContentUrl(v string)`

SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.

### SetContentUrlExplicitNull

`func (o *MicrosoftGraphUserActivity) SetContentUrlExplicitNull(b bool)`

SetContentUrlExplicitNull (un)sets ContentUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentUrl value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphUserActivity) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphUserActivity) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphUserActivity) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphUserActivity) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphUserActivity) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *MicrosoftGraphUserActivity) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphUserActivity) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *MicrosoftGraphUserActivity) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphUserActivity) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### SetExpirationDateTimeExplicitNull

`func (o *MicrosoftGraphUserActivity) SetExpirationDateTimeExplicitNull(b bool)`

SetExpirationDateTimeExplicitNull (un)sets ExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExpirationDateTime value is set to nil even if false is passed
### GetFallbackUrl

`func (o *MicrosoftGraphUserActivity) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *MicrosoftGraphUserActivity) GetFallbackUrlOk() (string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFallbackUrl

`func (o *MicrosoftGraphUserActivity) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### SetFallbackUrl

`func (o *MicrosoftGraphUserActivity) SetFallbackUrl(v string)`

SetFallbackUrl gets a reference to the given string and assigns it to the FallbackUrl field.

### SetFallbackUrlExplicitNull

`func (o *MicrosoftGraphUserActivity) SetFallbackUrlExplicitNull(b bool)`

SetFallbackUrlExplicitNull (un)sets FallbackUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FallbackUrl value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphUserActivity) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphUserActivity) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphUserActivity) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphUserActivity) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphUserActivity) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetUserTimezone

`func (o *MicrosoftGraphUserActivity) GetUserTimezone() string`

GetUserTimezone returns the UserTimezone field if non-nil, zero value otherwise.

### GetUserTimezoneOk

`func (o *MicrosoftGraphUserActivity) GetUserTimezoneOk() (string, bool)`

GetUserTimezoneOk returns a tuple with the UserTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserTimezone

`func (o *MicrosoftGraphUserActivity) HasUserTimezone() bool`

HasUserTimezone returns a boolean if a field has been set.

### SetUserTimezone

`func (o *MicrosoftGraphUserActivity) SetUserTimezone(v string)`

SetUserTimezone gets a reference to the given string and assigns it to the UserTimezone field.

### SetUserTimezoneExplicitNull

`func (o *MicrosoftGraphUserActivity) SetUserTimezoneExplicitNull(b bool)`

SetUserTimezoneExplicitNull (un)sets UserTimezone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserTimezone value is set to nil even if false is passed
### GetContentInfo

`func (o *MicrosoftGraphUserActivity) GetContentInfo() AnyOfobject`

GetContentInfo returns the ContentInfo field if non-nil, zero value otherwise.

### GetContentInfoOk

`func (o *MicrosoftGraphUserActivity) GetContentInfoOk() (AnyOfobject, bool)`

GetContentInfoOk returns a tuple with the ContentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentInfo

`func (o *MicrosoftGraphUserActivity) HasContentInfo() bool`

HasContentInfo returns a boolean if a field has been set.

### SetContentInfo

`func (o *MicrosoftGraphUserActivity) SetContentInfo(v AnyOfobject)`

SetContentInfo gets a reference to the given AnyOfobject and assigns it to the ContentInfo field.

### SetContentInfoExplicitNull

`func (o *MicrosoftGraphUserActivity) SetContentInfoExplicitNull(b bool)`

SetContentInfoExplicitNull (un)sets ContentInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentInfo value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphUserActivity) GetStatus() AnyOfmicrosoftGraphStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphUserActivity) GetStatusOk() (AnyOfmicrosoftGraphStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphUserActivity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphUserActivity) SetStatus(v AnyOfmicrosoftGraphStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphUserActivity) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetHistoryItems

`func (o *MicrosoftGraphUserActivity) GetHistoryItems() []MicrosoftGraphActivityHistoryItem`

GetHistoryItems returns the HistoryItems field if non-nil, zero value otherwise.

### GetHistoryItemsOk

`func (o *MicrosoftGraphUserActivity) GetHistoryItemsOk() ([]MicrosoftGraphActivityHistoryItem, bool)`

GetHistoryItemsOk returns a tuple with the HistoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHistoryItems

`func (o *MicrosoftGraphUserActivity) HasHistoryItems() bool`

HasHistoryItems returns a boolean if a field has been set.

### SetHistoryItems

`func (o *MicrosoftGraphUserActivity) SetHistoryItems(v []MicrosoftGraphActivityHistoryItem)`

SetHistoryItems gets a reference to the given []MicrosoftGraphActivityHistoryItem and assigns it to the HistoryItems field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


