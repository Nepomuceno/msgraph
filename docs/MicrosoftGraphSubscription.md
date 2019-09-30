# MicrosoftGraphSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**ChangeType** | Pointer to **string** |  | [optional] 
**ClientState** | Pointer to **string** |  | [optional] 
**NotificationUrl** | Pointer to **string** |  | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSubscription) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSubscription) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetResource

`func (o *MicrosoftGraphSubscription) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphSubscription) GetResourceOk() (string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResource

`func (o *MicrosoftGraphSubscription) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResource

`func (o *MicrosoftGraphSubscription) SetResource(v string)`

SetResource gets a reference to the given string and assigns it to the Resource field.

### GetChangeType

`func (o *MicrosoftGraphSubscription) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *MicrosoftGraphSubscription) GetChangeTypeOk() (string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeType

`func (o *MicrosoftGraphSubscription) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### SetChangeType

`func (o *MicrosoftGraphSubscription) SetChangeType(v string)`

SetChangeType gets a reference to the given string and assigns it to the ChangeType field.

### GetClientState

`func (o *MicrosoftGraphSubscription) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *MicrosoftGraphSubscription) GetClientStateOk() (string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientState

`func (o *MicrosoftGraphSubscription) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### SetClientState

`func (o *MicrosoftGraphSubscription) SetClientState(v string)`

SetClientState gets a reference to the given string and assigns it to the ClientState field.

### SetClientStateExplicitNull

`func (o *MicrosoftGraphSubscription) SetClientStateExplicitNull(b bool)`

SetClientStateExplicitNull (un)sets ClientState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClientState value is set to nil even if false is passed
### GetNotificationUrl

`func (o *MicrosoftGraphSubscription) GetNotificationUrl() string`

GetNotificationUrl returns the NotificationUrl field if non-nil, zero value otherwise.

### GetNotificationUrlOk

`func (o *MicrosoftGraphSubscription) GetNotificationUrlOk() (string, bool)`

GetNotificationUrlOk returns a tuple with the NotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationUrl

`func (o *MicrosoftGraphSubscription) HasNotificationUrl() bool`

HasNotificationUrl returns a boolean if a field has been set.

### SetNotificationUrl

`func (o *MicrosoftGraphSubscription) SetNotificationUrl(v string)`

SetNotificationUrl gets a reference to the given string and assigns it to the NotificationUrl field.

### GetExpirationDateTime

`func (o *MicrosoftGraphSubscription) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphSubscription) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *MicrosoftGraphSubscription) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphSubscription) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### GetApplicationId

`func (o *MicrosoftGraphSubscription) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MicrosoftGraphSubscription) GetApplicationIdOk() (string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *MicrosoftGraphSubscription) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *MicrosoftGraphSubscription) SetApplicationId(v string)`

SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.

### SetApplicationIdExplicitNull

`func (o *MicrosoftGraphSubscription) SetApplicationIdExplicitNull(b bool)`

SetApplicationIdExplicitNull (un)sets ApplicationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationId value is set to nil even if false is passed
### GetCreatorId

`func (o *MicrosoftGraphSubscription) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *MicrosoftGraphSubscription) GetCreatorIdOk() (string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatorId

`func (o *MicrosoftGraphSubscription) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorId

`func (o *MicrosoftGraphSubscription) SetCreatorId(v string)`

SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.

### SetCreatorIdExplicitNull

`func (o *MicrosoftGraphSubscription) SetCreatorIdExplicitNull(b bool)`

SetCreatorIdExplicitNull (un)sets CreatorId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatorId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


