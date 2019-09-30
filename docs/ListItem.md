# ListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to [**AnyOfmicrosoftGraphContentTypeInfo**](anyOf&lt;microsoft.graph.contentTypeInfo&gt;.md) |  | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**Analytics** | Pointer to [**AnyOfmicrosoftGraphItemAnalytics**](anyOf&lt;microsoft.graph.itemAnalytics&gt;.md) |  | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Fields** | Pointer to [**AnyOfmicrosoftGraphFieldValueSet**](anyOf&lt;microsoft.graph.fieldValueSet&gt;.md) |  | [optional] 
**Versions** | Pointer to [**[]MicrosoftGraphListItemVersion**](microsoft.graph.listItemVersion.md) |  | [optional] 

## Methods

### GetContentType

`func (o *ListItem) GetContentType() AnyOfmicrosoftGraphContentTypeInfo`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ListItem) GetContentTypeOk() (AnyOfmicrosoftGraphContentTypeInfo, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *ListItem) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *ListItem) SetContentType(v AnyOfmicrosoftGraphContentTypeInfo)`

SetContentType gets a reference to the given AnyOfmicrosoftGraphContentTypeInfo and assigns it to the ContentType field.

### SetContentTypeExplicitNull

`func (o *ListItem) SetContentTypeExplicitNull(b bool)`

SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentType value is set to nil even if false is passed
### GetSharepointIds

`func (o *ListItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *ListItem) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *ListItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *ListItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *ListItem) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetAnalytics

`func (o *ListItem) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *ListItem) GetAnalyticsOk() (AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnalytics

`func (o *ListItem) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalytics

`func (o *ListItem) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics gets a reference to the given AnyOfmicrosoftGraphItemAnalytics and assigns it to the Analytics field.

### SetAnalyticsExplicitNull

`func (o *ListItem) SetAnalyticsExplicitNull(b bool)`

SetAnalyticsExplicitNull (un)sets Analytics to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Analytics value is set to nil even if false is passed
### GetDriveItem

`func (o *ListItem) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *ListItem) GetDriveItemOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveItem

`func (o *ListItem) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItem

`func (o *ListItem) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.

### SetDriveItemExplicitNull

`func (o *ListItem) SetDriveItemExplicitNull(b bool)`

SetDriveItemExplicitNull (un)sets DriveItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveItem value is set to nil even if false is passed
### GetFields

`func (o *ListItem) GetFields() AnyOfmicrosoftGraphFieldValueSet`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ListItem) GetFieldsOk() (AnyOfmicrosoftGraphFieldValueSet, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFields

`func (o *ListItem) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFields

`func (o *ListItem) SetFields(v AnyOfmicrosoftGraphFieldValueSet)`

SetFields gets a reference to the given AnyOfmicrosoftGraphFieldValueSet and assigns it to the Fields field.

### SetFieldsExplicitNull

`func (o *ListItem) SetFieldsExplicitNull(b bool)`

SetFieldsExplicitNull (un)sets Fields to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fields value is set to nil even if false is passed
### GetVersions

`func (o *ListItem) GetVersions() []MicrosoftGraphListItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListItem) GetVersionsOk() ([]MicrosoftGraphListItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersions

`func (o *ListItem) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersions

`func (o *ListItem) SetVersions(v []MicrosoftGraphListItemVersion)`

SetVersions gets a reference to the given []MicrosoftGraphListItemVersion and assigns it to the Versions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


