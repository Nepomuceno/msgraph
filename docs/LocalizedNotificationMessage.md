# LocalizedNotificationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**Locale** | Pointer to **string** | The Locale for which this message is destined. | [optional] 
**Subject** | Pointer to **string** | The Message Template Subject. | [optional] 
**MessageTemplate** | Pointer to **string** | The Message Template content. | [optional] 
**IsDefault** | Pointer to **bool** | Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message. | [optional] 

## Methods

### GetLastModifiedDateTime

`func (o *LocalizedNotificationMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *LocalizedNotificationMessage) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *LocalizedNotificationMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *LocalizedNotificationMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetLocale

`func (o *LocalizedNotificationMessage) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LocalizedNotificationMessage) GetLocaleOk() (string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocale

`func (o *LocalizedNotificationMessage) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocale

`func (o *LocalizedNotificationMessage) SetLocale(v string)`

SetLocale gets a reference to the given string and assigns it to the Locale field.

### GetSubject

`func (o *LocalizedNotificationMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LocalizedNotificationMessage) GetSubjectOk() (string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *LocalizedNotificationMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *LocalizedNotificationMessage) SetSubject(v string)`

SetSubject gets a reference to the given string and assigns it to the Subject field.

### GetMessageTemplate

`func (o *LocalizedNotificationMessage) GetMessageTemplate() string`

GetMessageTemplate returns the MessageTemplate field if non-nil, zero value otherwise.

### GetMessageTemplateOk

`func (o *LocalizedNotificationMessage) GetMessageTemplateOk() (string, bool)`

GetMessageTemplateOk returns a tuple with the MessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageTemplate

`func (o *LocalizedNotificationMessage) HasMessageTemplate() bool`

HasMessageTemplate returns a boolean if a field has been set.

### SetMessageTemplate

`func (o *LocalizedNotificationMessage) SetMessageTemplate(v string)`

SetMessageTemplate gets a reference to the given string and assigns it to the MessageTemplate field.

### GetIsDefault

`func (o *LocalizedNotificationMessage) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *LocalizedNotificationMessage) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *LocalizedNotificationMessage) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *LocalizedNotificationMessage) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


