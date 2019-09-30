# InlineObject636

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddresses** | Pointer to **[]string** |  | [optional] 
**MailTipsOptions** | Pointer to [**AnyOfmicrosoftGraphMailTipsType**](anyOf&lt;microsoft.graph.mailTipsType&gt;.md) |  | [optional] 

## Methods

### GetEmailAddresses

`func (o *InlineObject636) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *InlineObject636) GetEmailAddressesOk() ([]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddresses

`func (o *InlineObject636) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### SetEmailAddresses

`func (o *InlineObject636) SetEmailAddresses(v []string)`

SetEmailAddresses gets a reference to the given []string and assigns it to the EmailAddresses field.

### GetMailTipsOptions

`func (o *InlineObject636) GetMailTipsOptions() AnyOfmicrosoftGraphMailTipsType`

GetMailTipsOptions returns the MailTipsOptions field if non-nil, zero value otherwise.

### GetMailTipsOptionsOk

`func (o *InlineObject636) GetMailTipsOptionsOk() (AnyOfmicrosoftGraphMailTipsType, bool)`

GetMailTipsOptionsOk returns a tuple with the MailTipsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailTipsOptions

`func (o *InlineObject636) HasMailTipsOptions() bool`

HasMailTipsOptions returns a boolean if a field has been set.

### SetMailTipsOptions

`func (o *InlineObject636) SetMailTipsOptions(v AnyOfmicrosoftGraphMailTipsType)`

SetMailTipsOptions gets a reference to the given AnyOfmicrosoftGraphMailTipsType and assigns it to the MailTipsOptions field.

### SetMailTipsOptionsExplicitNull

`func (o *InlineObject636) SetMailTipsOptionsExplicitNull(b bool)`

SetMailTipsOptionsExplicitNull (un)sets MailTipsOptions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailTipsOptions value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


