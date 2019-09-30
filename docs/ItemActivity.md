# ItemActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**ActivityDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Actor** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 

## Methods

### GetAccess

`func (o *ItemActivity) GetAccess() AnyOfobject`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ItemActivity) GetAccessOk() (AnyOfobject, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccess

`func (o *ItemActivity) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccess

`func (o *ItemActivity) SetAccess(v AnyOfobject)`

SetAccess gets a reference to the given AnyOfobject and assigns it to the Access field.

### SetAccessExplicitNull

`func (o *ItemActivity) SetAccessExplicitNull(b bool)`

SetAccessExplicitNull (un)sets Access to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Access value is set to nil even if false is passed
### GetActivityDateTime

`func (o *ItemActivity) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *ItemActivity) GetActivityDateTimeOk() (time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDateTime

`func (o *ItemActivity) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### SetActivityDateTime

`func (o *ItemActivity) SetActivityDateTime(v time.Time)`

SetActivityDateTime gets a reference to the given time.Time and assigns it to the ActivityDateTime field.

### SetActivityDateTimeExplicitNull

`func (o *ItemActivity) SetActivityDateTimeExplicitNull(b bool)`

SetActivityDateTimeExplicitNull (un)sets ActivityDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActivityDateTime value is set to nil even if false is passed
### GetActor

`func (o *ItemActivity) GetActor() AnyOfmicrosoftGraphIdentitySet`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ItemActivity) GetActorOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActor

`func (o *ItemActivity) HasActor() bool`

HasActor returns a boolean if a field has been set.

### SetActor

`func (o *ItemActivity) SetActor(v AnyOfmicrosoftGraphIdentitySet)`

SetActor gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Actor field.

### SetActorExplicitNull

`func (o *ItemActivity) SetActorExplicitNull(b bool)`

SetActorExplicitNull (un)sets Actor to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Actor value is set to nil even if false is passed
### GetDriveItem

`func (o *ItemActivity) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *ItemActivity) GetDriveItemOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveItem

`func (o *ItemActivity) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItem

`func (o *ItemActivity) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.

### SetDriveItemExplicitNull

`func (o *ItemActivity) SetDriveItemExplicitNull(b bool)`

SetDriveItemExplicitNull (un)sets DriveItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveItem value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


