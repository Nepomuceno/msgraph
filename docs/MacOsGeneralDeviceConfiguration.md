# MacOsGeneralDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompliantAppsList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements. | [optional] 
**CompliantAppListType** | Pointer to [**AnyOfmicrosoftGraphAppListType**](anyOf&lt;microsoft.graph.appListType&gt;.md) | List that is in the CompliantAppsList. | [optional] 
**EmailInDomainSuffixes** | Pointer to **[]string** | An email address lacking a suffix that matches any of these strings will be considered out-of-domain. | [optional] 
**PasswordBlockSimple** | Pointer to **bool** | Block simple passwords. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | Number of character sets a password must contain. Valid values 0 to 4 | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of passwords. | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity required before a password is required. | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity required before the screen times out. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | Type of password that is required. | [optional] 
**PasswordRequired** | Pointer to **bool** | Whether or not to require a password. | [optional] 

## Methods

### GetCompliantAppsList

`func (o *MacOsGeneralDeviceConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *MacOsGeneralDeviceConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *MacOsGeneralDeviceConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *MacOsGeneralDeviceConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *MacOsGeneralDeviceConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *MacOsGeneralDeviceConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *MacOsGeneralDeviceConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *MacOsGeneralDeviceConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetEmailInDomainSuffixes

`func (o *MacOsGeneralDeviceConfiguration) GetEmailInDomainSuffixes() []string`

GetEmailInDomainSuffixes returns the EmailInDomainSuffixes field if non-nil, zero value otherwise.

### GetEmailInDomainSuffixesOk

`func (o *MacOsGeneralDeviceConfiguration) GetEmailInDomainSuffixesOk() ([]string, bool)`

GetEmailInDomainSuffixesOk returns a tuple with the EmailInDomainSuffixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailInDomainSuffixes

`func (o *MacOsGeneralDeviceConfiguration) HasEmailInDomainSuffixes() bool`

HasEmailInDomainSuffixes returns a boolean if a field has been set.

### SetEmailInDomainSuffixes

`func (o *MacOsGeneralDeviceConfiguration) SetEmailInDomainSuffixes(v []string)`

SetEmailInDomainSuffixes gets a reference to the given []string and assigns it to the EmailInDomainSuffixes field.

### GetPasswordBlockSimple

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeLock

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordRequired

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MacOsGeneralDeviceConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MacOsGeneralDeviceConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MacOsGeneralDeviceConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


