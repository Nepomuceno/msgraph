# MicrosoftGraphItemActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Access** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**ActivityDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Actor** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphItemActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphItemActivity) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphItemActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphItemActivity) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetAccess

`func (o *MicrosoftGraphItemActivity) GetAccess() AnyOfobject`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MicrosoftGraphItemActivity) GetAccessOk() (AnyOfobject, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccess

`func (o *MicrosoftGraphItemActivity) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccess

`func (o *MicrosoftGraphItemActivity) SetAccess(v AnyOfobject)`

SetAccess gets a reference to the given AnyOfobject and assigns it to the Access field.

### SetAccessExplicitNull

`func (o *MicrosoftGraphItemActivity) SetAccessExplicitNull(b bool)`

SetAccessExplicitNull (un)sets Access to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Access value is set to nil even if false is passed
### GetActivityDateTime

`func (o *MicrosoftGraphItemActivity) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *MicrosoftGraphItemActivity) GetActivityDateTimeOk() (time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDateTime

`func (o *MicrosoftGraphItemActivity) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### SetActivityDateTime

`func (o *MicrosoftGraphItemActivity) SetActivityDateTime(v time.Time)`

SetActivityDateTime gets a reference to the given time.Time and assigns it to the ActivityDateTime field.

### SetActivityDateTimeExplicitNull

`func (o *MicrosoftGraphItemActivity) SetActivityDateTimeExplicitNull(b bool)`

SetActivityDateTimeExplicitNull (un)sets ActivityDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActivityDateTime value is set to nil even if false is passed
### GetActor

`func (o *MicrosoftGraphItemActivity) GetActor() AnyOfmicrosoftGraphIdentitySet`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *MicrosoftGraphItemActivity) GetActorOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActor

`func (o *MicrosoftGraphItemActivity) HasActor() bool`

HasActor returns a boolean if a field has been set.

### SetActor

`func (o *MicrosoftGraphItemActivity) SetActor(v AnyOfmicrosoftGraphIdentitySet)`

SetActor gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Actor field.

### SetActorExplicitNull

`func (o *MicrosoftGraphItemActivity) SetActorExplicitNull(b bool)`

SetActorExplicitNull (un)sets Actor to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Actor value is set to nil even if false is passed
### GetDriveItem

`func (o *MicrosoftGraphItemActivity) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *MicrosoftGraphItemActivity) GetDriveItemOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveItem

`func (o *MicrosoftGraphItemActivity) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItem

`func (o *MicrosoftGraphItemActivity) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.

### SetDriveItemExplicitNull

`func (o *MicrosoftGraphItemActivity) SetDriveItemExplicitNull(b bool)`

SetDriveItemExplicitNull (un)sets DriveItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveItem value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


