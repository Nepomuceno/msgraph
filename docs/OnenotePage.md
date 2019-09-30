# OnenotePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**CreatedByAppId** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphPageLinks**](anyOf&lt;microsoft.graph.pageLinks&gt;.md) |  | [optional] 
**ContentUrl** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**UserTags** | Pointer to **[]string** |  | [optional] 
**ParentSection** | Pointer to [**AnyOfmicrosoftGraphOnenoteSection**](anyOf&lt;microsoft.graph.onenoteSection&gt;.md) |  | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) |  | [optional] 

## Methods

### GetTitle

`func (o *OnenotePage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OnenotePage) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *OnenotePage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *OnenotePage) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *OnenotePage) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetCreatedByAppId

`func (o *OnenotePage) GetCreatedByAppId() string`

GetCreatedByAppId returns the CreatedByAppId field if non-nil, zero value otherwise.

### GetCreatedByAppIdOk

`func (o *OnenotePage) GetCreatedByAppIdOk() (string, bool)`

GetCreatedByAppIdOk returns a tuple with the CreatedByAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByAppId

`func (o *OnenotePage) HasCreatedByAppId() bool`

HasCreatedByAppId returns a boolean if a field has been set.

### SetCreatedByAppId

`func (o *OnenotePage) SetCreatedByAppId(v string)`

SetCreatedByAppId gets a reference to the given string and assigns it to the CreatedByAppId field.

### SetCreatedByAppIdExplicitNull

`func (o *OnenotePage) SetCreatedByAppIdExplicitNull(b bool)`

SetCreatedByAppIdExplicitNull (un)sets CreatedByAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByAppId value is set to nil even if false is passed
### GetLinks

`func (o *OnenotePage) GetLinks() AnyOfmicrosoftGraphPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OnenotePage) GetLinksOk() (AnyOfmicrosoftGraphPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *OnenotePage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *OnenotePage) SetLinks(v AnyOfmicrosoftGraphPageLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphPageLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *OnenotePage) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetContentUrl

`func (o *OnenotePage) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *OnenotePage) GetContentUrlOk() (string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentUrl

`func (o *OnenotePage) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrl

`func (o *OnenotePage) SetContentUrl(v string)`

SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.

### SetContentUrlExplicitNull

`func (o *OnenotePage) SetContentUrlExplicitNull(b bool)`

SetContentUrlExplicitNull (un)sets ContentUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentUrl value is set to nil even if false is passed
### GetContent

`func (o *OnenotePage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OnenotePage) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *OnenotePage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *OnenotePage) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *OnenotePage) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *OnenotePage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *OnenotePage) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *OnenotePage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *OnenotePage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *OnenotePage) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetLevel

`func (o *OnenotePage) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *OnenotePage) GetLevelOk() (int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLevel

`func (o *OnenotePage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevel

`func (o *OnenotePage) SetLevel(v int32)`

SetLevel gets a reference to the given int32 and assigns it to the Level field.

### SetLevelExplicitNull

`func (o *OnenotePage) SetLevelExplicitNull(b bool)`

SetLevelExplicitNull (un)sets Level to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Level value is set to nil even if false is passed
### GetOrder

`func (o *OnenotePage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OnenotePage) GetOrderOk() (int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrder

`func (o *OnenotePage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrder

`func (o *OnenotePage) SetOrder(v int32)`

SetOrder gets a reference to the given int32 and assigns it to the Order field.

### SetOrderExplicitNull

`func (o *OnenotePage) SetOrderExplicitNull(b bool)`

SetOrderExplicitNull (un)sets Order to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Order value is set to nil even if false is passed
### GetUserTags

`func (o *OnenotePage) GetUserTags() []string`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *OnenotePage) GetUserTagsOk() ([]string, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserTags

`func (o *OnenotePage) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### SetUserTags

`func (o *OnenotePage) SetUserTags(v []string)`

SetUserTags gets a reference to the given []string and assigns it to the UserTags field.

### GetParentSection

`func (o *OnenotePage) GetParentSection() AnyOfmicrosoftGraphOnenoteSection`

GetParentSection returns the ParentSection field if non-nil, zero value otherwise.

### GetParentSectionOk

`func (o *OnenotePage) GetParentSectionOk() (AnyOfmicrosoftGraphOnenoteSection, bool)`

GetParentSectionOk returns a tuple with the ParentSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentSection

`func (o *OnenotePage) HasParentSection() bool`

HasParentSection returns a boolean if a field has been set.

### SetParentSection

`func (o *OnenotePage) SetParentSection(v AnyOfmicrosoftGraphOnenoteSection)`

SetParentSection gets a reference to the given AnyOfmicrosoftGraphOnenoteSection and assigns it to the ParentSection field.

### SetParentSectionExplicitNull

`func (o *OnenotePage) SetParentSectionExplicitNull(b bool)`

SetParentSectionExplicitNull (un)sets ParentSection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentSection value is set to nil even if false is passed
### GetParentNotebook

`func (o *OnenotePage) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *OnenotePage) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentNotebook

`func (o *OnenotePage) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebook

`func (o *OnenotePage) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.

### SetParentNotebookExplicitNull

`func (o *OnenotePage) SetParentNotebookExplicitNull(b bool)`

SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentNotebook value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


