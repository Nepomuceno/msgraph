# MicrosoftGraphRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 

## Methods

### GetEmailAddress

`func (o *MicrosoftGraphRecipient) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphRecipient) GetEmailAddressOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddress

`func (o *MicrosoftGraphRecipient) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddress

`func (o *MicrosoftGraphRecipient) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddress field.

### SetEmailAddressExplicitNull

`func (o *MicrosoftGraphRecipient) SetEmailAddressExplicitNull(b bool)`

SetEmailAddressExplicitNull (un)sets EmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmailAddress value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


