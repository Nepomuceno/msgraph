# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedLicenses** | Pointer to [**[]AnyOfmicrosoftGraphAssignedLicense**](anyOf&lt;microsoft.graph.assignedLicense&gt;.md) |  | [optional] 
**Classification** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**HasMembersWithLicenseErrors** | Pointer to **bool** |  | [optional] 
**GroupTypes** | Pointer to **[]string** |  | [optional] 
**LicenseProcessingState** | Pointer to [**AnyOfmicrosoftGraphLicenseProcessingState**](anyOf&lt;microsoft.graph.licenseProcessingState&gt;.md) |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 
**MailEnabled** | Pointer to **bool** |  | [optional] 
**MailNickname** | Pointer to **string** |  | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**OnPremisesProvisioningErrors** | Pointer to [**[]AnyOfmicrosoftGraphOnPremisesProvisioningError**](anyOf&lt;microsoft.graph.onPremisesProvisioningError&gt;.md) |  | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **string** |  | [optional] 
**OnPremisesSyncEnabled** | Pointer to **bool** |  | [optional] 
**PreferredDataLocation** | Pointer to **string** |  | [optional] 
**ProxyAddresses** | Pointer to **[]string** |  | [optional] 
**RenewedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**SecurityEnabled** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**AllowExternalSenders** | Pointer to **bool** |  | [optional] 
**AutoSubscribeNewMembers** | Pointer to **bool** |  | [optional] 
**IsSubscribedByMail** | Pointer to **bool** |  | [optional] 
**UnseenCount** | Pointer to **int32** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**MembersWithLicenseErrors** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**TransitiveMembers** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**CreatedOnBehalfOf** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) |  | [optional] 
**Owners** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**Settings** | Pointer to [**[]MicrosoftGraphGroupSetting**](microsoft.graph.groupSetting.md) |  | [optional] 
**Conversations** | Pointer to [**[]MicrosoftGraphConversation**](microsoft.graph.conversation.md) |  | [optional] 
**Photos** | Pointer to [**[]MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md) |  | [optional] 
**AcceptedSenders** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**RejectedSenders** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**Threads** | Pointer to [**[]MicrosoftGraphConversationThread**](microsoft.graph.conversationThread.md) |  | [optional] 
**Calendar** | Pointer to [**AnyOfmicrosoftGraphCalendar**](anyOf&lt;microsoft.graph.calendar&gt;.md) |  | [optional] 
**CalendarView** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 
**Events** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphProfilePhoto**](anyOf&lt;microsoft.graph.profilePhoto&gt;.md) |  | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) |  | [optional] 
**Drives** | Pointer to [**[]MicrosoftGraphDrive**](microsoft.graph.drive.md) |  | [optional] 
**Sites** | Pointer to [**[]MicrosoftGraphSite**](microsoft.graph.site.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 
**GroupLifecyclePolicies** | Pointer to [**[]MicrosoftGraphGroupLifecyclePolicy**](microsoft.graph.groupLifecyclePolicy.md) |  | [optional] 
**Planner** | Pointer to [**AnyOfmicrosoftGraphPlannerGroup**](anyOf&lt;microsoft.graph.plannerGroup&gt;.md) |  | [optional] 
**Onenote** | Pointer to [**AnyOfmicrosoftGraphOnenote**](anyOf&lt;microsoft.graph.onenote&gt;.md) |  | [optional] 
**Team** | Pointer to [**AnyOfmicrosoftGraphTeam**](anyOf&lt;microsoft.graph.team&gt;.md) |  | [optional] 

## Methods

### GetAssignedLicenses

`func (o *Group) GetAssignedLicenses() []AnyOfmicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *Group) GetAssignedLicensesOk() ([]AnyOfmicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedLicenses

`func (o *Group) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### SetAssignedLicenses

`func (o *Group) SetAssignedLicenses(v []AnyOfmicrosoftGraphAssignedLicense)`

SetAssignedLicenses gets a reference to the given []AnyOfmicrosoftGraphAssignedLicense and assigns it to the AssignedLicenses field.

### GetClassification

`func (o *Group) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Group) GetClassificationOk() (string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassification

`func (o *Group) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassification

`func (o *Group) SetClassification(v string)`

SetClassification gets a reference to the given string and assigns it to the Classification field.

### SetClassificationExplicitNull

`func (o *Group) SetClassificationExplicitNull(b bool)`

SetClassificationExplicitNull (un)sets Classification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Classification value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *Group) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Group) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *Group) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *Group) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *Group) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *Group) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *Group) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Group) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *Group) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *Group) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *Group) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetHasMembersWithLicenseErrors

`func (o *Group) GetHasMembersWithLicenseErrors() bool`

GetHasMembersWithLicenseErrors returns the HasMembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetHasMembersWithLicenseErrorsOk

`func (o *Group) GetHasMembersWithLicenseErrorsOk() (bool, bool)`

GetHasMembersWithLicenseErrorsOk returns a tuple with the HasMembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasMembersWithLicenseErrors

`func (o *Group) HasHasMembersWithLicenseErrors() bool`

HasHasMembersWithLicenseErrors returns a boolean if a field has been set.

### SetHasMembersWithLicenseErrors

`func (o *Group) SetHasMembersWithLicenseErrors(v bool)`

SetHasMembersWithLicenseErrors gets a reference to the given bool and assigns it to the HasMembersWithLicenseErrors field.

### SetHasMembersWithLicenseErrorsExplicitNull

`func (o *Group) SetHasMembersWithLicenseErrorsExplicitNull(b bool)`

SetHasMembersWithLicenseErrorsExplicitNull (un)sets HasMembersWithLicenseErrors to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasMembersWithLicenseErrors value is set to nil even if false is passed
### GetGroupTypes

`func (o *Group) GetGroupTypes() []string`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *Group) GetGroupTypesOk() ([]string, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupTypes

`func (o *Group) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### SetGroupTypes

`func (o *Group) SetGroupTypes(v []string)`

SetGroupTypes gets a reference to the given []string and assigns it to the GroupTypes field.

### GetLicenseProcessingState

`func (o *Group) GetLicenseProcessingState() AnyOfmicrosoftGraphLicenseProcessingState`

GetLicenseProcessingState returns the LicenseProcessingState field if non-nil, zero value otherwise.

### GetLicenseProcessingStateOk

`func (o *Group) GetLicenseProcessingStateOk() (AnyOfmicrosoftGraphLicenseProcessingState, bool)`

GetLicenseProcessingStateOk returns a tuple with the LicenseProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseProcessingState

`func (o *Group) HasLicenseProcessingState() bool`

HasLicenseProcessingState returns a boolean if a field has been set.

### SetLicenseProcessingState

`func (o *Group) SetLicenseProcessingState(v AnyOfmicrosoftGraphLicenseProcessingState)`

SetLicenseProcessingState gets a reference to the given AnyOfmicrosoftGraphLicenseProcessingState and assigns it to the LicenseProcessingState field.

### SetLicenseProcessingStateExplicitNull

`func (o *Group) SetLicenseProcessingStateExplicitNull(b bool)`

SetLicenseProcessingStateExplicitNull (un)sets LicenseProcessingState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LicenseProcessingState value is set to nil even if false is passed
### GetMail

`func (o *Group) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *Group) GetMailOk() (string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMail

`func (o *Group) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMail

`func (o *Group) SetMail(v string)`

SetMail gets a reference to the given string and assigns it to the Mail field.

### SetMailExplicitNull

`func (o *Group) SetMailExplicitNull(b bool)`

SetMailExplicitNull (un)sets Mail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mail value is set to nil even if false is passed
### GetMailEnabled

`func (o *Group) GetMailEnabled() bool`

GetMailEnabled returns the MailEnabled field if non-nil, zero value otherwise.

### GetMailEnabledOk

`func (o *Group) GetMailEnabledOk() (bool, bool)`

GetMailEnabledOk returns a tuple with the MailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailEnabled

`func (o *Group) HasMailEnabled() bool`

HasMailEnabled returns a boolean if a field has been set.

### SetMailEnabled

`func (o *Group) SetMailEnabled(v bool)`

SetMailEnabled gets a reference to the given bool and assigns it to the MailEnabled field.

### SetMailEnabledExplicitNull

`func (o *Group) SetMailEnabledExplicitNull(b bool)`

SetMailEnabledExplicitNull (un)sets MailEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailEnabled value is set to nil even if false is passed
### GetMailNickname

`func (o *Group) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *Group) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *Group) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *Group) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### SetMailNicknameExplicitNull

`func (o *Group) SetMailNicknameExplicitNull(b bool)`

SetMailNicknameExplicitNull (un)sets MailNickname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailNickname value is set to nil even if false is passed
### GetOnPremisesLastSyncDateTime

`func (o *Group) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Group) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *Group) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Group) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *Group) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesProvisioningErrors

`func (o *Group) GetOnPremisesProvisioningErrors() []AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *Group) GetOnPremisesProvisioningErrorsOk() ([]AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesProvisioningErrors

`func (o *Group) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### SetOnPremisesProvisioningErrors

`func (o *Group) SetOnPremisesProvisioningErrors(v []AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors gets a reference to the given []AnyOfmicrosoftGraphOnPremisesProvisioningError and assigns it to the OnPremisesProvisioningErrors field.

### GetOnPremisesSecurityIdentifier

`func (o *Group) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *Group) GetOnPremisesSecurityIdentifierOk() (string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSecurityIdentifier

`func (o *Group) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifier

`func (o *Group) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier gets a reference to the given string and assigns it to the OnPremisesSecurityIdentifier field.

### SetOnPremisesSecurityIdentifierExplicitNull

`func (o *Group) SetOnPremisesSecurityIdentifierExplicitNull(b bool)`

SetOnPremisesSecurityIdentifierExplicitNull (un)sets OnPremisesSecurityIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSecurityIdentifier value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *Group) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Group) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *Group) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *Group) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *Group) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetPreferredDataLocation

`func (o *Group) GetPreferredDataLocation() string`

GetPreferredDataLocation returns the PreferredDataLocation field if non-nil, zero value otherwise.

### GetPreferredDataLocationOk

`func (o *Group) GetPreferredDataLocationOk() (string, bool)`

GetPreferredDataLocationOk returns a tuple with the PreferredDataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredDataLocation

`func (o *Group) HasPreferredDataLocation() bool`

HasPreferredDataLocation returns a boolean if a field has been set.

### SetPreferredDataLocation

`func (o *Group) SetPreferredDataLocation(v string)`

SetPreferredDataLocation gets a reference to the given string and assigns it to the PreferredDataLocation field.

### SetPreferredDataLocationExplicitNull

`func (o *Group) SetPreferredDataLocationExplicitNull(b bool)`

SetPreferredDataLocationExplicitNull (un)sets PreferredDataLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredDataLocation value is set to nil even if false is passed
### GetProxyAddresses

`func (o *Group) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *Group) GetProxyAddressesOk() ([]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProxyAddresses

`func (o *Group) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### SetProxyAddresses

`func (o *Group) SetProxyAddresses(v []string)`

SetProxyAddresses gets a reference to the given []string and assigns it to the ProxyAddresses field.

### GetRenewedDateTime

`func (o *Group) GetRenewedDateTime() time.Time`

GetRenewedDateTime returns the RenewedDateTime field if non-nil, zero value otherwise.

### GetRenewedDateTimeOk

`func (o *Group) GetRenewedDateTimeOk() (time.Time, bool)`

GetRenewedDateTimeOk returns a tuple with the RenewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRenewedDateTime

`func (o *Group) HasRenewedDateTime() bool`

HasRenewedDateTime returns a boolean if a field has been set.

### SetRenewedDateTime

`func (o *Group) SetRenewedDateTime(v time.Time)`

SetRenewedDateTime gets a reference to the given time.Time and assigns it to the RenewedDateTime field.

### SetRenewedDateTimeExplicitNull

`func (o *Group) SetRenewedDateTimeExplicitNull(b bool)`

SetRenewedDateTimeExplicitNull (un)sets RenewedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RenewedDateTime value is set to nil even if false is passed
### GetSecurityEnabled

`func (o *Group) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *Group) GetSecurityEnabledOk() (bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityEnabled

`func (o *Group) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### SetSecurityEnabled

`func (o *Group) SetSecurityEnabled(v bool)`

SetSecurityEnabled gets a reference to the given bool and assigns it to the SecurityEnabled field.

### SetSecurityEnabledExplicitNull

`func (o *Group) SetSecurityEnabledExplicitNull(b bool)`

SetSecurityEnabledExplicitNull (un)sets SecurityEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SecurityEnabled value is set to nil even if false is passed
### GetVisibility

`func (o *Group) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Group) GetVisibilityOk() (string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisibility

`func (o *Group) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibility

`func (o *Group) SetVisibility(v string)`

SetVisibility gets a reference to the given string and assigns it to the Visibility field.

### SetVisibilityExplicitNull

`func (o *Group) SetVisibilityExplicitNull(b bool)`

SetVisibilityExplicitNull (un)sets Visibility to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Visibility value is set to nil even if false is passed
### GetAllowExternalSenders

`func (o *Group) GetAllowExternalSenders() bool`

GetAllowExternalSenders returns the AllowExternalSenders field if non-nil, zero value otherwise.

### GetAllowExternalSendersOk

`func (o *Group) GetAllowExternalSendersOk() (bool, bool)`

GetAllowExternalSendersOk returns a tuple with the AllowExternalSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowExternalSenders

`func (o *Group) HasAllowExternalSenders() bool`

HasAllowExternalSenders returns a boolean if a field has been set.

### SetAllowExternalSenders

`func (o *Group) SetAllowExternalSenders(v bool)`

SetAllowExternalSenders gets a reference to the given bool and assigns it to the AllowExternalSenders field.

### SetAllowExternalSendersExplicitNull

`func (o *Group) SetAllowExternalSendersExplicitNull(b bool)`

SetAllowExternalSendersExplicitNull (un)sets AllowExternalSenders to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AllowExternalSenders value is set to nil even if false is passed
### GetAutoSubscribeNewMembers

`func (o *Group) GetAutoSubscribeNewMembers() bool`

GetAutoSubscribeNewMembers returns the AutoSubscribeNewMembers field if non-nil, zero value otherwise.

### GetAutoSubscribeNewMembersOk

`func (o *Group) GetAutoSubscribeNewMembersOk() (bool, bool)`

GetAutoSubscribeNewMembersOk returns a tuple with the AutoSubscribeNewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutoSubscribeNewMembers

`func (o *Group) HasAutoSubscribeNewMembers() bool`

HasAutoSubscribeNewMembers returns a boolean if a field has been set.

### SetAutoSubscribeNewMembers

`func (o *Group) SetAutoSubscribeNewMembers(v bool)`

SetAutoSubscribeNewMembers gets a reference to the given bool and assigns it to the AutoSubscribeNewMembers field.

### SetAutoSubscribeNewMembersExplicitNull

`func (o *Group) SetAutoSubscribeNewMembersExplicitNull(b bool)`

SetAutoSubscribeNewMembersExplicitNull (un)sets AutoSubscribeNewMembers to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AutoSubscribeNewMembers value is set to nil even if false is passed
### GetIsSubscribedByMail

`func (o *Group) GetIsSubscribedByMail() bool`

GetIsSubscribedByMail returns the IsSubscribedByMail field if non-nil, zero value otherwise.

### GetIsSubscribedByMailOk

`func (o *Group) GetIsSubscribedByMailOk() (bool, bool)`

GetIsSubscribedByMailOk returns a tuple with the IsSubscribedByMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSubscribedByMail

`func (o *Group) HasIsSubscribedByMail() bool`

HasIsSubscribedByMail returns a boolean if a field has been set.

### SetIsSubscribedByMail

`func (o *Group) SetIsSubscribedByMail(v bool)`

SetIsSubscribedByMail gets a reference to the given bool and assigns it to the IsSubscribedByMail field.

### SetIsSubscribedByMailExplicitNull

`func (o *Group) SetIsSubscribedByMailExplicitNull(b bool)`

SetIsSubscribedByMailExplicitNull (un)sets IsSubscribedByMail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsSubscribedByMail value is set to nil even if false is passed
### GetUnseenCount

`func (o *Group) GetUnseenCount() int32`

GetUnseenCount returns the UnseenCount field if non-nil, zero value otherwise.

### GetUnseenCountOk

`func (o *Group) GetUnseenCountOk() (int32, bool)`

GetUnseenCountOk returns a tuple with the UnseenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnseenCount

`func (o *Group) HasUnseenCount() bool`

HasUnseenCount returns a boolean if a field has been set.

### SetUnseenCount

`func (o *Group) SetUnseenCount(v int32)`

SetUnseenCount gets a reference to the given int32 and assigns it to the UnseenCount field.

### SetUnseenCountExplicitNull

`func (o *Group) SetUnseenCountExplicitNull(b bool)`

SetUnseenCountExplicitNull (un)sets UnseenCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnseenCount value is set to nil even if false is passed
### GetIsArchived

`func (o *Group) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Group) GetIsArchivedOk() (bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsArchived

`func (o *Group) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchived

`func (o *Group) SetIsArchived(v bool)`

SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.

### SetIsArchivedExplicitNull

`func (o *Group) SetIsArchivedExplicitNull(b bool)`

SetIsArchivedExplicitNull (un)sets IsArchived to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsArchived value is set to nil even if false is passed
### GetMembers

`func (o *Group) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Group) GetMembersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *Group) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Members field.

### GetMemberOf

`func (o *Group) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *Group) GetMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberOf

`func (o *Group) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOf

`func (o *Group) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.

### GetMembersWithLicenseErrors

`func (o *Group) GetMembersWithLicenseErrors() []MicrosoftGraphDirectoryObject`

GetMembersWithLicenseErrors returns the MembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetMembersWithLicenseErrorsOk

`func (o *Group) GetMembersWithLicenseErrorsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMembersWithLicenseErrorsOk returns a tuple with the MembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembersWithLicenseErrors

`func (o *Group) HasMembersWithLicenseErrors() bool`

HasMembersWithLicenseErrors returns a boolean if a field has been set.

### SetMembersWithLicenseErrors

`func (o *Group) SetMembersWithLicenseErrors(v []MicrosoftGraphDirectoryObject)`

SetMembersWithLicenseErrors gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MembersWithLicenseErrors field.

### GetTransitiveMembers

`func (o *Group) GetTransitiveMembers() []MicrosoftGraphDirectoryObject`

GetTransitiveMembers returns the TransitiveMembers field if non-nil, zero value otherwise.

### GetTransitiveMembersOk

`func (o *Group) GetTransitiveMembersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMembersOk returns a tuple with the TransitiveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMembers

`func (o *Group) HasTransitiveMembers() bool`

HasTransitiveMembers returns a boolean if a field has been set.

### SetTransitiveMembers

`func (o *Group) SetTransitiveMembers(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMembers field.

### GetTransitiveMemberOf

`func (o *Group) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *Group) GetTransitiveMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMemberOf

`func (o *Group) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### SetTransitiveMemberOf

`func (o *Group) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.

### GetCreatedOnBehalfOf

`func (o *Group) GetCreatedOnBehalfOf() AnyOfmicrosoftGraphDirectoryObject`

GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field if non-nil, zero value otherwise.

### GetCreatedOnBehalfOfOk

`func (o *Group) GetCreatedOnBehalfOfOk() (AnyOfmicrosoftGraphDirectoryObject, bool)`

GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedOnBehalfOf

`func (o *Group) HasCreatedOnBehalfOf() bool`

HasCreatedOnBehalfOf returns a boolean if a field has been set.

### SetCreatedOnBehalfOf

`func (o *Group) SetCreatedOnBehalfOf(v AnyOfmicrosoftGraphDirectoryObject)`

SetCreatedOnBehalfOf gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the CreatedOnBehalfOf field.

### SetCreatedOnBehalfOfExplicitNull

`func (o *Group) SetCreatedOnBehalfOfExplicitNull(b bool)`

SetCreatedOnBehalfOfExplicitNull (un)sets CreatedOnBehalfOf to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedOnBehalfOf value is set to nil even if false is passed
### GetOwners

`func (o *Group) GetOwners() []MicrosoftGraphDirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Group) GetOwnersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwners

`func (o *Group) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### SetOwners

`func (o *Group) SetOwners(v []MicrosoftGraphDirectoryObject)`

SetOwners gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Owners field.

### GetSettings

`func (o *Group) GetSettings() []MicrosoftGraphGroupSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Group) GetSettingsOk() ([]MicrosoftGraphGroupSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *Group) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *Group) SetSettings(v []MicrosoftGraphGroupSetting)`

SetSettings gets a reference to the given []MicrosoftGraphGroupSetting and assigns it to the Settings field.

### GetConversations

`func (o *Group) GetConversations() []MicrosoftGraphConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *Group) GetConversationsOk() ([]MicrosoftGraphConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversations

`func (o *Group) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### SetConversations

`func (o *Group) SetConversations(v []MicrosoftGraphConversation)`

SetConversations gets a reference to the given []MicrosoftGraphConversation and assigns it to the Conversations field.

### GetPhotos

`func (o *Group) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *Group) GetPhotosOk() ([]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhotos

`func (o *Group) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### SetPhotos

`func (o *Group) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos gets a reference to the given []MicrosoftGraphProfilePhoto and assigns it to the Photos field.

### GetAcceptedSenders

`func (o *Group) GetAcceptedSenders() []MicrosoftGraphDirectoryObject`

GetAcceptedSenders returns the AcceptedSenders field if non-nil, zero value otherwise.

### GetAcceptedSendersOk

`func (o *Group) GetAcceptedSendersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetAcceptedSendersOk returns a tuple with the AcceptedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptedSenders

`func (o *Group) HasAcceptedSenders() bool`

HasAcceptedSenders returns a boolean if a field has been set.

### SetAcceptedSenders

`func (o *Group) SetAcceptedSenders(v []MicrosoftGraphDirectoryObject)`

SetAcceptedSenders gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the AcceptedSenders field.

### GetRejectedSenders

`func (o *Group) GetRejectedSenders() []MicrosoftGraphDirectoryObject`

GetRejectedSenders returns the RejectedSenders field if non-nil, zero value otherwise.

### GetRejectedSendersOk

`func (o *Group) GetRejectedSendersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRejectedSendersOk returns a tuple with the RejectedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRejectedSenders

`func (o *Group) HasRejectedSenders() bool`

HasRejectedSenders returns a boolean if a field has been set.

### SetRejectedSenders

`func (o *Group) SetRejectedSenders(v []MicrosoftGraphDirectoryObject)`

SetRejectedSenders gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RejectedSenders field.

### GetThreads

`func (o *Group) GetThreads() []MicrosoftGraphConversationThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *Group) GetThreadsOk() ([]MicrosoftGraphConversationThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThreads

`func (o *Group) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### SetThreads

`func (o *Group) SetThreads(v []MicrosoftGraphConversationThread)`

SetThreads gets a reference to the given []MicrosoftGraphConversationThread and assigns it to the Threads field.

### GetCalendar

`func (o *Group) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Group) GetCalendarOk() (AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendar

`func (o *Group) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendar

`func (o *Group) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar gets a reference to the given AnyOfmicrosoftGraphCalendar and assigns it to the Calendar field.

### SetCalendarExplicitNull

`func (o *Group) SetCalendarExplicitNull(b bool)`

SetCalendarExplicitNull (un)sets Calendar to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calendar value is set to nil even if false is passed
### GetCalendarView

`func (o *Group) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *Group) GetCalendarViewOk() ([]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarView

`func (o *Group) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### SetCalendarView

`func (o *Group) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView gets a reference to the given []MicrosoftGraphEvent and assigns it to the CalendarView field.

### GetEvents

`func (o *Group) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Group) GetEventsOk() ([]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *Group) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *Group) SetEvents(v []MicrosoftGraphEvent)`

SetEvents gets a reference to the given []MicrosoftGraphEvent and assigns it to the Events field.

### GetPhoto

`func (o *Group) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Group) GetPhotoOk() (AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *Group) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *Group) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphProfilePhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *Group) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetDrive

`func (o *Group) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *Group) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *Group) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *Group) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *Group) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetDrives

`func (o *Group) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *Group) GetDrivesOk() ([]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrives

`func (o *Group) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrives

`func (o *Group) SetDrives(v []MicrosoftGraphDrive)`

SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.

### GetSites

`func (o *Group) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Group) GetSitesOk() ([]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSites

`func (o *Group) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSites

`func (o *Group) SetSites(v []MicrosoftGraphSite)`

SetSites gets a reference to the given []MicrosoftGraphSite and assigns it to the Sites field.

### GetExtensions

`func (o *Group) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Group) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Group) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Group) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetGroupLifecyclePolicies

`func (o *Group) GetGroupLifecyclePolicies() []MicrosoftGraphGroupLifecyclePolicy`

GetGroupLifecyclePolicies returns the GroupLifecyclePolicies field if non-nil, zero value otherwise.

### GetGroupLifecyclePoliciesOk

`func (o *Group) GetGroupLifecyclePoliciesOk() ([]MicrosoftGraphGroupLifecyclePolicy, bool)`

GetGroupLifecyclePoliciesOk returns a tuple with the GroupLifecyclePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupLifecyclePolicies

`func (o *Group) HasGroupLifecyclePolicies() bool`

HasGroupLifecyclePolicies returns a boolean if a field has been set.

### SetGroupLifecyclePolicies

`func (o *Group) SetGroupLifecyclePolicies(v []MicrosoftGraphGroupLifecyclePolicy)`

SetGroupLifecyclePolicies gets a reference to the given []MicrosoftGraphGroupLifecyclePolicy and assigns it to the GroupLifecyclePolicies field.

### GetPlanner

`func (o *Group) GetPlanner() AnyOfmicrosoftGraphPlannerGroup`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *Group) GetPlannerOk() (AnyOfmicrosoftGraphPlannerGroup, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanner

`func (o *Group) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlanner

`func (o *Group) SetPlanner(v AnyOfmicrosoftGraphPlannerGroup)`

SetPlanner gets a reference to the given AnyOfmicrosoftGraphPlannerGroup and assigns it to the Planner field.

### SetPlannerExplicitNull

`func (o *Group) SetPlannerExplicitNull(b bool)`

SetPlannerExplicitNull (un)sets Planner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Planner value is set to nil even if false is passed
### GetOnenote

`func (o *Group) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *Group) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnenote

`func (o *Group) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenote

`func (o *Group) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.

### SetOnenoteExplicitNull

`func (o *Group) SetOnenoteExplicitNull(b bool)`

SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Onenote value is set to nil even if false is passed
### GetTeam

`func (o *Group) GetTeam() AnyOfmicrosoftGraphTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Group) GetTeamOk() (AnyOfmicrosoftGraphTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam

`func (o *Group) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeam

`func (o *Group) SetTeam(v AnyOfmicrosoftGraphTeam)`

SetTeam gets a reference to the given AnyOfmicrosoftGraphTeam and assigns it to the Team field.

### SetTeamExplicitNull

`func (o *Group) SetTeamExplicitNull(b bool)`

SetTeamExplicitNull (un)sets Team to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Team value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


