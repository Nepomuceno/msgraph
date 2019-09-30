# TermsAndConditionsAcceptanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserDisplayName** | Pointer to **string** | Display name of the user whose acceptance the entity represents. | [optional] 
**AcceptedVersion** | Pointer to **int32** | Most recent version number of the T&amp;C accepted by the user. | [optional] 
**AcceptedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime when the terms were last accepted by the user. | [optional] 
**TermsAndConditions** | Pointer to [**AnyOfmicrosoftGraphTermsAndConditions**](anyOf&lt;microsoft.graph.termsAndConditions&gt;.md) |  | [optional] 

## Methods

### GetUserDisplayName

`func (o *TermsAndConditionsAcceptanceStatus) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *TermsAndConditionsAcceptanceStatus) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *TermsAndConditionsAcceptanceStatus) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *TermsAndConditionsAcceptanceStatus) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *TermsAndConditionsAcceptanceStatus) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetAcceptedVersion

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedVersion() int32`

GetAcceptedVersion returns the AcceptedVersion field if non-nil, zero value otherwise.

### GetAcceptedVersionOk

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedVersionOk() (int32, bool)`

GetAcceptedVersionOk returns a tuple with the AcceptedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptedVersion

`func (o *TermsAndConditionsAcceptanceStatus) HasAcceptedVersion() bool`

HasAcceptedVersion returns a boolean if a field has been set.

### SetAcceptedVersion

`func (o *TermsAndConditionsAcceptanceStatus) SetAcceptedVersion(v int32)`

SetAcceptedVersion gets a reference to the given int32 and assigns it to the AcceptedVersion field.

### GetAcceptedDateTime

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTime() time.Time`

GetAcceptedDateTime returns the AcceptedDateTime field if non-nil, zero value otherwise.

### GetAcceptedDateTimeOk

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTimeOk() (time.Time, bool)`

GetAcceptedDateTimeOk returns a tuple with the AcceptedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptedDateTime

`func (o *TermsAndConditionsAcceptanceStatus) HasAcceptedDateTime() bool`

HasAcceptedDateTime returns a boolean if a field has been set.

### SetAcceptedDateTime

`func (o *TermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(v time.Time)`

SetAcceptedDateTime gets a reference to the given time.Time and assigns it to the AcceptedDateTime field.

### GetTermsAndConditions

`func (o *TermsAndConditionsAcceptanceStatus) GetTermsAndConditions() AnyOfmicrosoftGraphTermsAndConditions`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *TermsAndConditionsAcceptanceStatus) GetTermsAndConditionsOk() (AnyOfmicrosoftGraphTermsAndConditions, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTermsAndConditions

`func (o *TermsAndConditionsAcceptanceStatus) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### SetTermsAndConditions

`func (o *TermsAndConditionsAcceptanceStatus) SetTermsAndConditions(v AnyOfmicrosoftGraphTermsAndConditions)`

SetTermsAndConditions gets a reference to the given AnyOfmicrosoftGraphTermsAndConditions and assigns it to the TermsAndConditions field.

### SetTermsAndConditionsExplicitNull

`func (o *TermsAndConditionsAcceptanceStatus) SetTermsAndConditionsExplicitNull(b bool)`

SetTermsAndConditionsExplicitNull (un)sets TermsAndConditions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TermsAndConditions value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


