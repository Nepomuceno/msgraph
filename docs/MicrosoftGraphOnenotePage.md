# MicrosoftGraphOnenotePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Self** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphOnenotePage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenotePage) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOnenotePage) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOnenotePage) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSelf

`func (o *MicrosoftGraphOnenotePage) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphOnenotePage) GetSelfOk() (string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelf

`func (o *MicrosoftGraphOnenotePage) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelf

`func (o *MicrosoftGraphOnenotePage) SetSelf(v string)`

SetSelf gets a reference to the given string and assigns it to the Self field.

### SetSelfExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetSelfExplicitNull(b bool)`

SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Self value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphOnenotePage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenotePage) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenotePage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenotePage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphOnenotePage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphOnenotePage) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphOnenotePage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphOnenotePage) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetCreatedByAppId

`func (o *MicrosoftGraphOnenotePage) GetCreatedByAppId() string`

GetCreatedByAppId returns the CreatedByAppId field if non-nil, zero value otherwise.

### GetCreatedByAppIdOk

`func (o *MicrosoftGraphOnenotePage) GetCreatedByAppIdOk() (string, bool)`

GetCreatedByAppIdOk returns a tuple with the CreatedByAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByAppId

`func (o *MicrosoftGraphOnenotePage) HasCreatedByAppId() bool`

HasCreatedByAppId returns a boolean if a field has been set.

### SetCreatedByAppId

`func (o *MicrosoftGraphOnenotePage) SetCreatedByAppId(v string)`

SetCreatedByAppId gets a reference to the given string and assigns it to the CreatedByAppId field.

### SetCreatedByAppIdExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetCreatedByAppIdExplicitNull(b bool)`

SetCreatedByAppIdExplicitNull (un)sets CreatedByAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByAppId value is set to nil even if false is passed
### GetLinks

`func (o *MicrosoftGraphOnenotePage) GetLinks() AnyOfmicrosoftGraphPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphOnenotePage) GetLinksOk() (AnyOfmicrosoftGraphPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *MicrosoftGraphOnenotePage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *MicrosoftGraphOnenotePage) SetLinks(v AnyOfmicrosoftGraphPageLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphPageLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetContentUrl

`func (o *MicrosoftGraphOnenotePage) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *MicrosoftGraphOnenotePage) GetContentUrlOk() (string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentUrl

`func (o *MicrosoftGraphOnenotePage) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrl

`func (o *MicrosoftGraphOnenotePage) SetContentUrl(v string)`

SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.

### SetContentUrlExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetContentUrlExplicitNull(b bool)`

SetContentUrlExplicitNull (un)sets ContentUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentUrl value is set to nil even if false is passed
### GetContent

`func (o *MicrosoftGraphOnenotePage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphOnenotePage) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *MicrosoftGraphOnenotePage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *MicrosoftGraphOnenotePage) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOnenotePage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOnenotePage) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOnenotePage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOnenotePage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetLevel

`func (o *MicrosoftGraphOnenotePage) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MicrosoftGraphOnenotePage) GetLevelOk() (int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLevel

`func (o *MicrosoftGraphOnenotePage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevel

`func (o *MicrosoftGraphOnenotePage) SetLevel(v int32)`

SetLevel gets a reference to the given int32 and assigns it to the Level field.

### SetLevelExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetLevelExplicitNull(b bool)`

SetLevelExplicitNull (un)sets Level to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Level value is set to nil even if false is passed
### GetOrder

`func (o *MicrosoftGraphOnenotePage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MicrosoftGraphOnenotePage) GetOrderOk() (int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrder

`func (o *MicrosoftGraphOnenotePage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrder

`func (o *MicrosoftGraphOnenotePage) SetOrder(v int32)`

SetOrder gets a reference to the given int32 and assigns it to the Order field.

### SetOrderExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetOrderExplicitNull(b bool)`

SetOrderExplicitNull (un)sets Order to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Order value is set to nil even if false is passed
### GetUserTags

`func (o *MicrosoftGraphOnenotePage) GetUserTags() []string`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *MicrosoftGraphOnenotePage) GetUserTagsOk() ([]string, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserTags

`func (o *MicrosoftGraphOnenotePage) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### SetUserTags

`func (o *MicrosoftGraphOnenotePage) SetUserTags(v []string)`

SetUserTags gets a reference to the given []string and assigns it to the UserTags field.

### GetParentSection

`func (o *MicrosoftGraphOnenotePage) GetParentSection() AnyOfmicrosoftGraphOnenoteSection`

GetParentSection returns the ParentSection field if non-nil, zero value otherwise.

### GetParentSectionOk

`func (o *MicrosoftGraphOnenotePage) GetParentSectionOk() (AnyOfmicrosoftGraphOnenoteSection, bool)`

GetParentSectionOk returns a tuple with the ParentSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentSection

`func (o *MicrosoftGraphOnenotePage) HasParentSection() bool`

HasParentSection returns a boolean if a field has been set.

### SetParentSection

`func (o *MicrosoftGraphOnenotePage) SetParentSection(v AnyOfmicrosoftGraphOnenoteSection)`

SetParentSection gets a reference to the given AnyOfmicrosoftGraphOnenoteSection and assigns it to the ParentSection field.

### SetParentSectionExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetParentSectionExplicitNull(b bool)`

SetParentSectionExplicitNull (un)sets ParentSection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentSection value is set to nil even if false is passed
### GetParentNotebook

`func (o *MicrosoftGraphOnenotePage) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *MicrosoftGraphOnenotePage) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentNotebook

`func (o *MicrosoftGraphOnenotePage) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebook

`func (o *MicrosoftGraphOnenotePage) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.

### SetParentNotebookExplicitNull

`func (o *MicrosoftGraphOnenotePage) SetParentNotebookExplicitNull(b bool)`

SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentNotebook value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


