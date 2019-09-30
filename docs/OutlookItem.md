# OutlookItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 

## Methods

### GetCreatedDateTime

`func (o *OutlookItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OutlookItem) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *OutlookItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *OutlookItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *OutlookItem) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *OutlookItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *OutlookItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *OutlookItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *OutlookItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *OutlookItem) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetChangeKey

`func (o *OutlookItem) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *OutlookItem) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *OutlookItem) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *OutlookItem) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *OutlookItem) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCategories

`func (o *OutlookItem) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *OutlookItem) GetCategoriesOk() ([]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *OutlookItem) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *OutlookItem) SetCategories(v []string)`

SetCategories gets a reference to the given []string and assigns it to the Categories field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


