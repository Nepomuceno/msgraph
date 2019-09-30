# TermsAndConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetCreatedDateTime

`func (o *TermsAndConditions) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *TermsAndConditions) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *TermsAndConditions) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *TermsAndConditions) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *TermsAndConditions) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *TermsAndConditions) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *TermsAndConditions) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *TermsAndConditions) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *TermsAndConditions) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TermsAndConditions) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *TermsAndConditions) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *TermsAndConditions) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *TermsAndConditions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TermsAndConditions) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *TermsAndConditions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *TermsAndConditions) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *TermsAndConditions) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetTitle

`func (o *TermsAndConditions) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TermsAndConditions) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *TermsAndConditions) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *TermsAndConditions) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *TermsAndConditions) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetBodyText

`func (o *TermsAndConditions) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *TermsAndConditions) GetBodyTextOk() (string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyText

`func (o *TermsAndConditions) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### SetBodyText

`func (o *TermsAndConditions) SetBodyText(v string)`

SetBodyText gets a reference to the given string and assigns it to the BodyText field.

### SetBodyTextExplicitNull

`func (o *TermsAndConditions) SetBodyTextExplicitNull(b bool)`

SetBodyTextExplicitNull (un)sets BodyText to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BodyText value is set to nil even if false is passed
### GetAcceptanceStatement

`func (o *TermsAndConditions) GetAcceptanceStatement() string`

GetAcceptanceStatement returns the AcceptanceStatement field if non-nil, zero value otherwise.

### GetAcceptanceStatementOk

`func (o *TermsAndConditions) GetAcceptanceStatementOk() (string, bool)`

GetAcceptanceStatementOk returns a tuple with the AcceptanceStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptanceStatement

`func (o *TermsAndConditions) HasAcceptanceStatement() bool`

HasAcceptanceStatement returns a boolean if a field has been set.

### SetAcceptanceStatement

`func (o *TermsAndConditions) SetAcceptanceStatement(v string)`

SetAcceptanceStatement gets a reference to the given string and assigns it to the AcceptanceStatement field.

### SetAcceptanceStatementExplicitNull

`func (o *TermsAndConditions) SetAcceptanceStatementExplicitNull(b bool)`

SetAcceptanceStatementExplicitNull (un)sets AcceptanceStatement to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AcceptanceStatement value is set to nil even if false is passed
### GetVersion

`func (o *TermsAndConditions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TermsAndConditions) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *TermsAndConditions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *TermsAndConditions) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *TermsAndConditions) GetAssignments() []MicrosoftGraphTermsAndConditionsAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *TermsAndConditions) GetAssignmentsOk() ([]MicrosoftGraphTermsAndConditionsAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *TermsAndConditions) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *TermsAndConditions) SetAssignments(v []MicrosoftGraphTermsAndConditionsAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTermsAndConditionsAssignment and assigns it to the Assignments field.

### GetAcceptanceStatuses

`func (o *TermsAndConditions) GetAcceptanceStatuses() []MicrosoftGraphTermsAndConditionsAcceptanceStatus`

GetAcceptanceStatuses returns the AcceptanceStatuses field if non-nil, zero value otherwise.

### GetAcceptanceStatusesOk

`func (o *TermsAndConditions) GetAcceptanceStatusesOk() ([]MicrosoftGraphTermsAndConditionsAcceptanceStatus, bool)`

GetAcceptanceStatusesOk returns a tuple with the AcceptanceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptanceStatuses

`func (o *TermsAndConditions) HasAcceptanceStatuses() bool`

HasAcceptanceStatuses returns a boolean if a field has been set.

### SetAcceptanceStatuses

`func (o *TermsAndConditions) SetAcceptanceStatuses(v []MicrosoftGraphTermsAndConditionsAcceptanceStatus)`

SetAcceptanceStatuses gets a reference to the given []MicrosoftGraphTermsAndConditionsAcceptanceStatus and assigns it to the AcceptanceStatuses field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


