# MicrosoftGraphTermsAndConditionsAcceptanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserDisplayName** | Pointer to **string** | Display name of the user whose acceptance the entity represents. | [optional] 
**AcceptedVersion** | Pointer to **int32** | Most recent version number of the T&amp;C accepted by the user. | [optional] 
**AcceptedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime when the terms were last accepted by the user. | [optional] 
**TermsAndConditions** | Pointer to [**AnyOfmicrosoftGraphTermsAndConditions**](anyOf&lt;microsoft.graph.termsAndConditions&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetUserDisplayName

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetAcceptedVersion

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedVersion() int32`

GetAcceptedVersion returns the AcceptedVersion field if non-nil, zero value otherwise.

### GetAcceptedVersionOk

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedVersionOk() (int32, bool)`

GetAcceptedVersionOk returns a tuple with the AcceptedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptedVersion

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasAcceptedVersion() bool`

HasAcceptedVersion returns a boolean if a field has been set.

### SetAcceptedVersion

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetAcceptedVersion(v int32)`

SetAcceptedVersion gets a reference to the given int32 and assigns it to the AcceptedVersion field.

### GetAcceptedDateTime

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedDateTime() time.Time`

GetAcceptedDateTime returns the AcceptedDateTime field if non-nil, zero value otherwise.

### GetAcceptedDateTimeOk

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetAcceptedDateTimeOk() (time.Time, bool)`

GetAcceptedDateTimeOk returns a tuple with the AcceptedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptedDateTime

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasAcceptedDateTime() bool`

HasAcceptedDateTime returns a boolean if a field has been set.

### SetAcceptedDateTime

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(v time.Time)`

SetAcceptedDateTime gets a reference to the given time.Time and assigns it to the AcceptedDateTime field.

### GetTermsAndConditions

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetTermsAndConditions() AnyOfmicrosoftGraphTermsAndConditions`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) GetTermsAndConditionsOk() (AnyOfmicrosoftGraphTermsAndConditions, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTermsAndConditions

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### SetTermsAndConditions

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetTermsAndConditions(v AnyOfmicrosoftGraphTermsAndConditions)`

SetTermsAndConditions gets a reference to the given AnyOfmicrosoftGraphTermsAndConditions and assigns it to the TermsAndConditions field.

### SetTermsAndConditionsExplicitNull

`func (o *MicrosoftGraphTermsAndConditionsAcceptanceStatus) SetTermsAndConditionsExplicitNull(b bool)`

SetTermsAndConditionsExplicitNull (un)sets TermsAndConditions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TermsAndConditions value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


