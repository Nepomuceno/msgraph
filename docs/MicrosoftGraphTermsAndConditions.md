# MicrosoftGraphTermsAndConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**DisplayName** | Pointer to **string** | Administrator-supplied name for the T&amp;C policy.  | [optional] 
**Description** | Pointer to **string** | Administrator-supplied description of the T&amp;C policy. | [optional] 
**Title** | Pointer to **string** | Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**BodyText** | Pointer to **string** | Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**AcceptanceStatement** | Pointer to **string** | Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&amp;C policy. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**Version** | Pointer to **int32** | Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&amp;C policy. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTermsAndConditionsAssignment**](microsoft.graph.termsAndConditionsAssignment.md) |  | [optional] 
**AcceptanceStatuses** | Pointer to [**[]MicrosoftGraphTermsAndConditionsAcceptanceStatus**](microsoft.graph.termsAndConditionsAcceptanceStatus.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTermsAndConditions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermsAndConditions) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTermsAndConditions) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTermsAndConditions) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphTermsAndConditions) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTermsAndConditions) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphTermsAndConditions) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTermsAndConditions) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTermsAndConditions) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTermsAndConditions) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTermsAndConditions) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTermsAndConditions) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphTermsAndConditions) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTermsAndConditions) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphTermsAndConditions) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphTermsAndConditions) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphTermsAndConditions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTermsAndConditions) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphTermsAndConditions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphTermsAndConditions) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphTermsAndConditions) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphTermsAndConditions) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphTermsAndConditions) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphTermsAndConditions) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphTermsAndConditions) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphTermsAndConditions) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetBodyText

`func (o *MicrosoftGraphTermsAndConditions) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *MicrosoftGraphTermsAndConditions) GetBodyTextOk() (string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyText

`func (o *MicrosoftGraphTermsAndConditions) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### SetBodyText

`func (o *MicrosoftGraphTermsAndConditions) SetBodyText(v string)`

SetBodyText gets a reference to the given string and assigns it to the BodyText field.

### SetBodyTextExplicitNull

`func (o *MicrosoftGraphTermsAndConditions) SetBodyTextExplicitNull(b bool)`

SetBodyTextExplicitNull (un)sets BodyText to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BodyText value is set to nil even if false is passed
### GetAcceptanceStatement

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatement() string`

GetAcceptanceStatement returns the AcceptanceStatement field if non-nil, zero value otherwise.

### GetAcceptanceStatementOk

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatementOk() (string, bool)`

GetAcceptanceStatementOk returns a tuple with the AcceptanceStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptanceStatement

`func (o *MicrosoftGraphTermsAndConditions) HasAcceptanceStatement() bool`

HasAcceptanceStatement returns a boolean if a field has been set.

### SetAcceptanceStatement

`func (o *MicrosoftGraphTermsAndConditions) SetAcceptanceStatement(v string)`

SetAcceptanceStatement gets a reference to the given string and assigns it to the AcceptanceStatement field.

### SetAcceptanceStatementExplicitNull

`func (o *MicrosoftGraphTermsAndConditions) SetAcceptanceStatementExplicitNull(b bool)`

SetAcceptanceStatementExplicitNull (un)sets AcceptanceStatement to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AcceptanceStatement value is set to nil even if false is passed
### GetVersion

`func (o *MicrosoftGraphTermsAndConditions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphTermsAndConditions) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphTermsAndConditions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphTermsAndConditions) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphTermsAndConditions) GetAssignments() []MicrosoftGraphTermsAndConditionsAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphTermsAndConditions) GetAssignmentsOk() ([]MicrosoftGraphTermsAndConditionsAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphTermsAndConditions) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphTermsAndConditions) SetAssignments(v []MicrosoftGraphTermsAndConditionsAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTermsAndConditionsAssignment and assigns it to the Assignments field.

### GetAcceptanceStatuses

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatuses() []MicrosoftGraphTermsAndConditionsAcceptanceStatus`

GetAcceptanceStatuses returns the AcceptanceStatuses field if non-nil, zero value otherwise.

### GetAcceptanceStatusesOk

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatusesOk() ([]MicrosoftGraphTermsAndConditionsAcceptanceStatus, bool)`

GetAcceptanceStatusesOk returns a tuple with the AcceptanceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptanceStatuses

`func (o *MicrosoftGraphTermsAndConditions) HasAcceptanceStatuses() bool`

HasAcceptanceStatuses returns a boolean if a field has been set.

### SetAcceptanceStatuses

`func (o *MicrosoftGraphTermsAndConditions) SetAcceptanceStatuses(v []MicrosoftGraphTermsAndConditionsAcceptanceStatus)`

SetAcceptanceStatuses gets a reference to the given []MicrosoftGraphTermsAndConditionsAcceptanceStatus and assigns it to the AcceptanceStatuses field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


