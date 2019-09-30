# UserActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetVisualElements

`func (o *UserActivity) GetVisualElements() MicrosoftGraphVisualInfo`

GetVisualElements returns the VisualElements field if non-nil, zero value otherwise.

### GetVisualElementsOk

`func (o *UserActivity) GetVisualElementsOk() (MicrosoftGraphVisualInfo, bool)`

GetVisualElementsOk returns a tuple with the VisualElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisualElements

`func (o *UserActivity) HasVisualElements() bool`

HasVisualElements returns a boolean if a field has been set.

### SetVisualElements

`func (o *UserActivity) SetVisualElements(v MicrosoftGraphVisualInfo)`

SetVisualElements gets a reference to the given MicrosoftGraphVisualInfo and assigns it to the VisualElements field.

### GetActivitySourceHost

`func (o *UserActivity) GetActivitySourceHost() string`

GetActivitySourceHost returns the ActivitySourceHost field if non-nil, zero value otherwise.

### GetActivitySourceHostOk

`func (o *UserActivity) GetActivitySourceHostOk() (string, bool)`

GetActivitySourceHostOk returns a tuple with the ActivitySourceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivitySourceHost

`func (o *UserActivity) HasActivitySourceHost() bool`

HasActivitySourceHost returns a boolean if a field has been set.

### SetActivitySourceHost

`func (o *UserActivity) SetActivitySourceHost(v string)`

SetActivitySourceHost gets a reference to the given string and assigns it to the ActivitySourceHost field.

### GetActivationUrl

`func (o *UserActivity) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *UserActivity) GetActivationUrlOk() (string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationUrl

`func (o *UserActivity) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.

### SetActivationUrl

`func (o *UserActivity) SetActivationUrl(v string)`

SetActivationUrl gets a reference to the given string and assigns it to the ActivationUrl field.

### GetAppActivityId

`func (o *UserActivity) GetAppActivityId() string`

GetAppActivityId returns the AppActivityId field if non-nil, zero value otherwise.

### GetAppActivityIdOk

`func (o *UserActivity) GetAppActivityIdOk() (string, bool)`

GetAppActivityIdOk returns a tuple with the AppActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppActivityId

`func (o *UserActivity) HasAppActivityId() bool`

HasAppActivityId returns a boolean if a field has been set.

### SetAppActivityId

`func (o *UserActivity) SetAppActivityId(v string)`

SetAppActivityId gets a reference to the given string and assigns it to the AppActivityId field.

### GetAppDisplayName

`func (o *UserActivity) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *UserActivity) GetAppDisplayNameOk() (string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDisplayName

`func (o *UserActivity) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayName

`func (o *UserActivity) SetAppDisplayName(v string)`

SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.

### SetAppDisplayNameExplicitNull

`func (o *UserActivity) SetAppDisplayNameExplicitNull(b bool)`

SetAppDisplayNameExplicitNull (un)sets AppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppDisplayName value is set to nil even if false is passed
### GetContentUrl

`func (o *UserActivity) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *UserActivity) GetContentUrlOk() (string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentUrl

`func (o *UserActivity) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrl

`func (o *UserActivity) SetContentUrl(v string)`

SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.

### SetContentUrlExplicitNull

`func (o *UserActivity) SetContentUrlExplicitNull(b bool)`

SetContentUrlExplicitNull (un)sets ContentUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentUrl value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *UserActivity) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *UserActivity) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *UserActivity) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *UserActivity) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *UserActivity) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *UserActivity) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *UserActivity) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *UserActivity) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *UserActivity) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### SetExpirationDateTimeExplicitNull

`func (o *UserActivity) SetExpirationDateTimeExplicitNull(b bool)`

SetExpirationDateTimeExplicitNull (un)sets ExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExpirationDateTime value is set to nil even if false is passed
### GetFallbackUrl

`func (o *UserActivity) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *UserActivity) GetFallbackUrlOk() (string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFallbackUrl

`func (o *UserActivity) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### SetFallbackUrl

`func (o *UserActivity) SetFallbackUrl(v string)`

SetFallbackUrl gets a reference to the given string and assigns it to the FallbackUrl field.

### SetFallbackUrlExplicitNull

`func (o *UserActivity) SetFallbackUrlExplicitNull(b bool)`

SetFallbackUrlExplicitNull (un)sets FallbackUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FallbackUrl value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *UserActivity) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *UserActivity) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *UserActivity) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *UserActivity) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *UserActivity) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetUserTimezone

`func (o *UserActivity) GetUserTimezone() string`

GetUserTimezone returns the UserTimezone field if non-nil, zero value otherwise.

### GetUserTimezoneOk

`func (o *UserActivity) GetUserTimezoneOk() (string, bool)`

GetUserTimezoneOk returns a tuple with the UserTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserTimezone

`func (o *UserActivity) HasUserTimezone() bool`

HasUserTimezone returns a boolean if a field has been set.

### SetUserTimezone

`func (o *UserActivity) SetUserTimezone(v string)`

SetUserTimezone gets a reference to the given string and assigns it to the UserTimezone field.

### SetUserTimezoneExplicitNull

`func (o *UserActivity) SetUserTimezoneExplicitNull(b bool)`

SetUserTimezoneExplicitNull (un)sets UserTimezone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserTimezone value is set to nil even if false is passed
### GetContentInfo

`func (o *UserActivity) GetContentInfo() AnyOfobject`

GetContentInfo returns the ContentInfo field if non-nil, zero value otherwise.

### GetContentInfoOk

`func (o *UserActivity) GetContentInfoOk() (AnyOfobject, bool)`

GetContentInfoOk returns a tuple with the ContentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentInfo

`func (o *UserActivity) HasContentInfo() bool`

HasContentInfo returns a boolean if a field has been set.

### SetContentInfo

`func (o *UserActivity) SetContentInfo(v AnyOfobject)`

SetContentInfo gets a reference to the given AnyOfobject and assigns it to the ContentInfo field.

### SetContentInfoExplicitNull

`func (o *UserActivity) SetContentInfoExplicitNull(b bool)`

SetContentInfoExplicitNull (un)sets ContentInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentInfo value is set to nil even if false is passed
### GetStatus

`func (o *UserActivity) GetStatus() AnyOfmicrosoftGraphStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserActivity) GetStatusOk() (AnyOfmicrosoftGraphStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *UserActivity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *UserActivity) SetStatus(v AnyOfmicrosoftGraphStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *UserActivity) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetHistoryItems

`func (o *UserActivity) GetHistoryItems() []MicrosoftGraphActivityHistoryItem`

GetHistoryItems returns the HistoryItems field if non-nil, zero value otherwise.

### GetHistoryItemsOk

`func (o *UserActivity) GetHistoryItemsOk() ([]MicrosoftGraphActivityHistoryItem, bool)`

GetHistoryItemsOk returns a tuple with the HistoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHistoryItems

`func (o *UserActivity) HasHistoryItems() bool`

HasHistoryItems returns a boolean if a field has been set.

### SetHistoryItems

`func (o *UserActivity) SetHistoryItems(v []MicrosoftGraphActivityHistoryItem)`

SetHistoryItems gets a reference to the given []MicrosoftGraphActivityHistoryItem and assigns it to the HistoryItems field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


