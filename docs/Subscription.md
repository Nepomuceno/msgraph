# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** |  | [optional] 
**ChangeType** | Pointer to **string** |  | [optional] 
**ClientState** | Pointer to **string** |  | [optional] 
**NotificationUrl** | Pointer to **string** |  | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 

## Methods

### GetResource

`func (o *Subscription) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Subscription) GetResourceOk() (string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResource

`func (o *Subscription) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResource

`func (o *Subscription) SetResource(v string)`

SetResource gets a reference to the given string and assigns it to the Resource field.

### GetChangeType

`func (o *Subscription) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *Subscription) GetChangeTypeOk() (string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeType

`func (o *Subscription) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### SetChangeType

`func (o *Subscription) SetChangeType(v string)`

SetChangeType gets a reference to the given string and assigns it to the ChangeType field.

### GetClientState

`func (o *Subscription) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *Subscription) GetClientStateOk() (string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientState

`func (o *Subscription) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### SetClientState

`func (o *Subscription) SetClientState(v string)`

SetClientState gets a reference to the given string and assigns it to the ClientState field.

### SetClientStateExplicitNull

`func (o *Subscription) SetClientStateExplicitNull(b bool)`

SetClientStateExplicitNull (un)sets ClientState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClientState value is set to nil even if false is passed
### GetNotificationUrl

`func (o *Subscription) GetNotificationUrl() string`

GetNotificationUrl returns the NotificationUrl field if non-nil, zero value otherwise.

### GetNotificationUrlOk

`func (o *Subscription) GetNotificationUrlOk() (string, bool)`

GetNotificationUrlOk returns a tuple with the NotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationUrl

`func (o *Subscription) HasNotificationUrl() bool`

HasNotificationUrl returns a boolean if a field has been set.

### SetNotificationUrl

`func (o *Subscription) SetNotificationUrl(v string)`

SetNotificationUrl gets a reference to the given string and assigns it to the NotificationUrl field.

### GetExpirationDateTime

`func (o *Subscription) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *Subscription) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *Subscription) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *Subscription) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### GetApplicationId

`func (o *Subscription) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Subscription) GetApplicationIdOk() (string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Subscription) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Subscription) SetApplicationId(v string)`

SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.

### SetApplicationIdExplicitNull

`func (o *Subscription) SetApplicationIdExplicitNull(b bool)`

SetApplicationIdExplicitNull (un)sets ApplicationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationId value is set to nil even if false is passed
### GetCreatorId

`func (o *Subscription) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Subscription) GetCreatorIdOk() (string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatorId

`func (o *Subscription) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorId

`func (o *Subscription) SetCreatorId(v string)`

SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.

### SetCreatorIdExplicitNull

`func (o *Subscription) SetCreatorIdExplicitNull(b bool)`

SetCreatorIdExplicitNull (un)sets CreatorId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatorId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


