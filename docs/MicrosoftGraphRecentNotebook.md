# MicrosoftGraphRecentNotebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**LastAccessedTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphRecentNotebookLinks**](anyOf&lt;microsoft.graph.recentNotebookLinks&gt;.md) |  | [optional] 
**SourceService** | Pointer to [**AnyOfmicrosoftGraphOnenoteSourceService**](anyOf&lt;microsoft.graph.onenoteSourceService&gt;.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphRecentNotebook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRecentNotebook) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphRecentNotebook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphRecentNotebook) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphRecentNotebook) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetLastAccessedTime

`func (o *MicrosoftGraphRecentNotebook) GetLastAccessedTime() time.Time`

GetLastAccessedTime returns the LastAccessedTime field if non-nil, zero value otherwise.

### GetLastAccessedTimeOk

`func (o *MicrosoftGraphRecentNotebook) GetLastAccessedTimeOk() (time.Time, bool)`

GetLastAccessedTimeOk returns a tuple with the LastAccessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastAccessedTime

`func (o *MicrosoftGraphRecentNotebook) HasLastAccessedTime() bool`

HasLastAccessedTime returns a boolean if a field has been set.

### SetLastAccessedTime

`func (o *MicrosoftGraphRecentNotebook) SetLastAccessedTime(v time.Time)`

SetLastAccessedTime gets a reference to the given time.Time and assigns it to the LastAccessedTime field.

### SetLastAccessedTimeExplicitNull

`func (o *MicrosoftGraphRecentNotebook) SetLastAccessedTimeExplicitNull(b bool)`

SetLastAccessedTimeExplicitNull (un)sets LastAccessedTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastAccessedTime value is set to nil even if false is passed
### GetLinks

`func (o *MicrosoftGraphRecentNotebook) GetLinks() AnyOfmicrosoftGraphRecentNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphRecentNotebook) GetLinksOk() (AnyOfmicrosoftGraphRecentNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *MicrosoftGraphRecentNotebook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *MicrosoftGraphRecentNotebook) SetLinks(v AnyOfmicrosoftGraphRecentNotebookLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphRecentNotebookLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *MicrosoftGraphRecentNotebook) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetSourceService

`func (o *MicrosoftGraphRecentNotebook) GetSourceService() AnyOfmicrosoftGraphOnenoteSourceService`

GetSourceService returns the SourceService field if non-nil, zero value otherwise.

### GetSourceServiceOk

`func (o *MicrosoftGraphRecentNotebook) GetSourceServiceOk() (AnyOfmicrosoftGraphOnenoteSourceService, bool)`

GetSourceServiceOk returns a tuple with the SourceService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceService

`func (o *MicrosoftGraphRecentNotebook) HasSourceService() bool`

HasSourceService returns a boolean if a field has been set.

### SetSourceService

`func (o *MicrosoftGraphRecentNotebook) SetSourceService(v AnyOfmicrosoftGraphOnenoteSourceService)`

SetSourceService gets a reference to the given AnyOfmicrosoftGraphOnenoteSourceService and assigns it to the SourceService field.

### SetSourceServiceExplicitNull

`func (o *MicrosoftGraphRecentNotebook) SetSourceServiceExplicitNull(b bool)`

SetSourceServiceExplicitNull (un)sets SourceService to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SourceService value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


