# ApplePushNotificationCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppleIdentifier** | Pointer to **string** | Apple Id of the account used to create the MDM push certificate. | [optional] 
**TopicIdentifier** | Pointer to **string** | Topic Id. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last modified date and time for Apple push notification certificate. | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The expiration date and time for Apple push notification certificate. | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 

## Methods

### GetAppleIdentifier

`func (o *ApplePushNotificationCertificate) GetAppleIdentifier() string`

GetAppleIdentifier returns the AppleIdentifier field if non-nil, zero value otherwise.

### GetAppleIdentifierOk

`func (o *ApplePushNotificationCertificate) GetAppleIdentifierOk() (string, bool)`

GetAppleIdentifierOk returns a tuple with the AppleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleIdentifier

`func (o *ApplePushNotificationCertificate) HasAppleIdentifier() bool`

HasAppleIdentifier returns a boolean if a field has been set.

### SetAppleIdentifier

`func (o *ApplePushNotificationCertificate) SetAppleIdentifier(v string)`

SetAppleIdentifier gets a reference to the given string and assigns it to the AppleIdentifier field.

### SetAppleIdentifierExplicitNull

`func (o *ApplePushNotificationCertificate) SetAppleIdentifierExplicitNull(b bool)`

SetAppleIdentifierExplicitNull (un)sets AppleIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppleIdentifier value is set to nil even if false is passed
### GetTopicIdentifier

`func (o *ApplePushNotificationCertificate) GetTopicIdentifier() string`

GetTopicIdentifier returns the TopicIdentifier field if non-nil, zero value otherwise.

### GetTopicIdentifierOk

`func (o *ApplePushNotificationCertificate) GetTopicIdentifierOk() (string, bool)`

GetTopicIdentifierOk returns a tuple with the TopicIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTopicIdentifier

`func (o *ApplePushNotificationCertificate) HasTopicIdentifier() bool`

HasTopicIdentifier returns a boolean if a field has been set.

### SetTopicIdentifier

`func (o *ApplePushNotificationCertificate) SetTopicIdentifier(v string)`

SetTopicIdentifier gets a reference to the given string and assigns it to the TopicIdentifier field.

### SetTopicIdentifierExplicitNull

`func (o *ApplePushNotificationCertificate) SetTopicIdentifierExplicitNull(b bool)`

SetTopicIdentifierExplicitNull (un)sets TopicIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TopicIdentifier value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *ApplePushNotificationCertificate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *ApplePushNotificationCertificate) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *ApplePushNotificationCertificate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *ApplePushNotificationCertificate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetExpirationDateTime

`func (o *ApplePushNotificationCertificate) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *ApplePushNotificationCertificate) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *ApplePushNotificationCertificate) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *ApplePushNotificationCertificate) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### GetCertificate

`func (o *ApplePushNotificationCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ApplePushNotificationCertificate) GetCertificateOk() (string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificate

`func (o *ApplePushNotificationCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificate

`func (o *ApplePushNotificationCertificate) SetCertificate(v string)`

SetCertificate gets a reference to the given string and assigns it to the Certificate field.

### SetCertificateExplicitNull

`func (o *ApplePushNotificationCertificate) SetCertificateExplicitNull(b bool)`

SetCertificateExplicitNull (un)sets Certificate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Certificate value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


