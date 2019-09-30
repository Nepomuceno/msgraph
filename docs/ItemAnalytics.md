# ItemAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemActivityStats** | Pointer to [**[]MicrosoftGraphItemActivityStat**](microsoft.graph.itemActivityStat.md) |  | [optional] 
**AllTime** | Pointer to [**AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md) |  | [optional] 
**LastSevenDays** | Pointer to [**AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md) |  | [optional] 

## Methods

### GetItemActivityStats

`func (o *ItemAnalytics) GetItemActivityStats() []MicrosoftGraphItemActivityStat`

GetItemActivityStats returns the ItemActivityStats field if non-nil, zero value otherwise.

### GetItemActivityStatsOk

`func (o *ItemAnalytics) GetItemActivityStatsOk() ([]MicrosoftGraphItemActivityStat, bool)`

GetItemActivityStatsOk returns a tuple with the ItemActivityStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItemActivityStats

`func (o *ItemAnalytics) HasItemActivityStats() bool`

HasItemActivityStats returns a boolean if a field has been set.

### SetItemActivityStats

`func (o *ItemAnalytics) SetItemActivityStats(v []MicrosoftGraphItemActivityStat)`

SetItemActivityStats gets a reference to the given []MicrosoftGraphItemActivityStat and assigns it to the ItemActivityStats field.

### GetAllTime

`func (o *ItemAnalytics) GetAllTime() AnyOfmicrosoftGraphItemActivityStat`

GetAllTime returns the AllTime field if non-nil, zero value otherwise.

### GetAllTimeOk

`func (o *ItemAnalytics) GetAllTimeOk() (AnyOfmicrosoftGraphItemActivityStat, bool)`

GetAllTimeOk returns a tuple with the AllTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllTime

`func (o *ItemAnalytics) HasAllTime() bool`

HasAllTime returns a boolean if a field has been set.

### SetAllTime

`func (o *ItemAnalytics) SetAllTime(v AnyOfmicrosoftGraphItemActivityStat)`

SetAllTime gets a reference to the given AnyOfmicrosoftGraphItemActivityStat and assigns it to the AllTime field.

### SetAllTimeExplicitNull

`func (o *ItemAnalytics) SetAllTimeExplicitNull(b bool)`

SetAllTimeExplicitNull (un)sets AllTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AllTime value is set to nil even if false is passed
### GetLastSevenDays

`func (o *ItemAnalytics) GetLastSevenDays() AnyOfmicrosoftGraphItemActivityStat`

GetLastSevenDays returns the LastSevenDays field if non-nil, zero value otherwise.

### GetLastSevenDaysOk

`func (o *ItemAnalytics) GetLastSevenDaysOk() (AnyOfmicrosoftGraphItemActivityStat, bool)`

GetLastSevenDaysOk returns a tuple with the LastSevenDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSevenDays

`func (o *ItemAnalytics) HasLastSevenDays() bool`

HasLastSevenDays returns a boolean if a field has been set.

### SetLastSevenDays

`func (o *ItemAnalytics) SetLastSevenDays(v AnyOfmicrosoftGraphItemActivityStat)`

SetLastSevenDays gets a reference to the given AnyOfmicrosoftGraphItemActivityStat and assigns it to the LastSevenDays field.

### SetLastSevenDaysExplicitNull

`func (o *ItemAnalytics) SetLastSevenDaysExplicitNull(b bool)`

SetLastSevenDaysExplicitNull (un)sets LastSevenDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastSevenDays value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


