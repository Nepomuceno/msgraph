# MicrosoftGraphItemActivityStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StartDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**EndDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Access** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) |  | [optional] 
**Create** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) |  | [optional] 
**Delete** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) |  | [optional] 
**Edit** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) |  | [optional] 
**Move** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) |  | [optional] 
**IsTrending** | Pointer to **bool** |  | [optional] 
**IncompleteData** | Pointer to [**AnyOfmicrosoftGraphIncompleteData**](anyOf&lt;microsoft.graph.incompleteData&gt;.md) |  | [optional] 
**Activities** | Pointer to [**[]MicrosoftGraphItemActivity**](microsoft.graph.itemActivity.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphItemActivityStat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphItemActivityStat) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphItemActivityStat) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphItemActivityStat) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetStartDateTime

`func (o *MicrosoftGraphItemActivityStat) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphItemActivityStat) GetStartDateTimeOk() (time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDateTime

`func (o *MicrosoftGraphItemActivityStat) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTime

`func (o *MicrosoftGraphItemActivityStat) SetStartDateTime(v time.Time)`

SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.

### SetStartDateTimeExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetStartDateTimeExplicitNull(b bool)`

SetStartDateTimeExplicitNull (un)sets StartDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartDateTime value is set to nil even if false is passed
### GetEndDateTime

`func (o *MicrosoftGraphItemActivityStat) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphItemActivityStat) GetEndDateTimeOk() (time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndDateTime

`func (o *MicrosoftGraphItemActivityStat) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTime

`func (o *MicrosoftGraphItemActivityStat) SetEndDateTime(v time.Time)`

SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.

### SetEndDateTimeExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetEndDateTimeExplicitNull(b bool)`

SetEndDateTimeExplicitNull (un)sets EndDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EndDateTime value is set to nil even if false is passed
### GetAccess

`func (o *MicrosoftGraphItemActivityStat) GetAccess() AnyOfmicrosoftGraphItemActionStat`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MicrosoftGraphItemActivityStat) GetAccessOk() (AnyOfmicrosoftGraphItemActionStat, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccess

`func (o *MicrosoftGraphItemActivityStat) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccess

`func (o *MicrosoftGraphItemActivityStat) SetAccess(v AnyOfmicrosoftGraphItemActionStat)`

SetAccess gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Access field.

### SetAccessExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetAccessExplicitNull(b bool)`

SetAccessExplicitNull (un)sets Access to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Access value is set to nil even if false is passed
### GetCreate

`func (o *MicrosoftGraphItemActivityStat) GetCreate() AnyOfmicrosoftGraphItemActionStat`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MicrosoftGraphItemActivityStat) GetCreateOk() (AnyOfmicrosoftGraphItemActionStat, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreate

`func (o *MicrosoftGraphItemActivityStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### SetCreate

`func (o *MicrosoftGraphItemActivityStat) SetCreate(v AnyOfmicrosoftGraphItemActionStat)`

SetCreate gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Create field.

### SetCreateExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetCreateExplicitNull(b bool)`

SetCreateExplicitNull (un)sets Create to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Create value is set to nil even if false is passed
### GetDelete

`func (o *MicrosoftGraphItemActivityStat) GetDelete() AnyOfmicrosoftGraphItemActionStat`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *MicrosoftGraphItemActivityStat) GetDeleteOk() (AnyOfmicrosoftGraphItemActionStat, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDelete

`func (o *MicrosoftGraphItemActivityStat) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDelete

`func (o *MicrosoftGraphItemActivityStat) SetDelete(v AnyOfmicrosoftGraphItemActionStat)`

SetDelete gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Delete field.

### SetDeleteExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetDeleteExplicitNull(b bool)`

SetDeleteExplicitNull (un)sets Delete to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Delete value is set to nil even if false is passed
### GetEdit

`func (o *MicrosoftGraphItemActivityStat) GetEdit() AnyOfmicrosoftGraphItemActionStat`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *MicrosoftGraphItemActivityStat) GetEditOk() (AnyOfmicrosoftGraphItemActionStat, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEdit

`func (o *MicrosoftGraphItemActivityStat) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### SetEdit

`func (o *MicrosoftGraphItemActivityStat) SetEdit(v AnyOfmicrosoftGraphItemActionStat)`

SetEdit gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Edit field.

### SetEditExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetEditExplicitNull(b bool)`

SetEditExplicitNull (un)sets Edit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Edit value is set to nil even if false is passed
### GetMove

`func (o *MicrosoftGraphItemActivityStat) GetMove() AnyOfmicrosoftGraphItemActionStat`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *MicrosoftGraphItemActivityStat) GetMoveOk() (AnyOfmicrosoftGraphItemActionStat, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMove

`func (o *MicrosoftGraphItemActivityStat) HasMove() bool`

HasMove returns a boolean if a field has been set.

### SetMove

`func (o *MicrosoftGraphItemActivityStat) SetMove(v AnyOfmicrosoftGraphItemActionStat)`

SetMove gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Move field.

### SetMoveExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetMoveExplicitNull(b bool)`

SetMoveExplicitNull (un)sets Move to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Move value is set to nil even if false is passed
### GetIsTrending

`func (o *MicrosoftGraphItemActivityStat) GetIsTrending() bool`

GetIsTrending returns the IsTrending field if non-nil, zero value otherwise.

### GetIsTrendingOk

`func (o *MicrosoftGraphItemActivityStat) GetIsTrendingOk() (bool, bool)`

GetIsTrendingOk returns a tuple with the IsTrending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsTrending

`func (o *MicrosoftGraphItemActivityStat) HasIsTrending() bool`

HasIsTrending returns a boolean if a field has been set.

### SetIsTrending

`func (o *MicrosoftGraphItemActivityStat) SetIsTrending(v bool)`

SetIsTrending gets a reference to the given bool and assigns it to the IsTrending field.

### SetIsTrendingExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetIsTrendingExplicitNull(b bool)`

SetIsTrendingExplicitNull (un)sets IsTrending to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsTrending value is set to nil even if false is passed
### GetIncompleteData

`func (o *MicrosoftGraphItemActivityStat) GetIncompleteData() AnyOfmicrosoftGraphIncompleteData`

GetIncompleteData returns the IncompleteData field if non-nil, zero value otherwise.

### GetIncompleteDataOk

`func (o *MicrosoftGraphItemActivityStat) GetIncompleteDataOk() (AnyOfmicrosoftGraphIncompleteData, bool)`

GetIncompleteDataOk returns a tuple with the IncompleteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncompleteData

`func (o *MicrosoftGraphItemActivityStat) HasIncompleteData() bool`

HasIncompleteData returns a boolean if a field has been set.

### SetIncompleteData

`func (o *MicrosoftGraphItemActivityStat) SetIncompleteData(v AnyOfmicrosoftGraphIncompleteData)`

SetIncompleteData gets a reference to the given AnyOfmicrosoftGraphIncompleteData and assigns it to the IncompleteData field.

### SetIncompleteDataExplicitNull

`func (o *MicrosoftGraphItemActivityStat) SetIncompleteDataExplicitNull(b bool)`

SetIncompleteDataExplicitNull (un)sets IncompleteData to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IncompleteData value is set to nil even if false is passed
### GetActivities

`func (o *MicrosoftGraphItemActivityStat) GetActivities() []MicrosoftGraphItemActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *MicrosoftGraphItemActivityStat) GetActivitiesOk() ([]MicrosoftGraphItemActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivities

`func (o *MicrosoftGraphItemActivityStat) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivities

`func (o *MicrosoftGraphItemActivityStat) SetActivities(v []MicrosoftGraphItemActivity)`

SetActivities gets a reference to the given []MicrosoftGraphItemActivity and assigns it to the Activities field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


