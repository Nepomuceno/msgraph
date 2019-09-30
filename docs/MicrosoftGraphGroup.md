# MicrosoftGraphGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphGroup) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphGroup) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeletedDateTime

`func (o *MicrosoftGraphGroup) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphGroup) GetDeletedDateTimeOk() (time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDateTime

`func (o *MicrosoftGraphGroup) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphGroup) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.

### SetDeletedDateTimeExplicitNull

`func (o *MicrosoftGraphGroup) SetDeletedDateTimeExplicitNull(b bool)`

SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeletedDateTime value is set to nil even if false is passed
### GetAssignedLicenses

`func (o *MicrosoftGraphGroup) GetAssignedLicenses() []AnyOfmicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *MicrosoftGraphGroup) GetAssignedLicensesOk() ([]AnyOfmicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedLicenses

`func (o *MicrosoftGraphGroup) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### SetAssignedLicenses

`func (o *MicrosoftGraphGroup) SetAssignedLicenses(v []AnyOfmicrosoftGraphAssignedLicense)`

SetAssignedLicenses gets a reference to the given []AnyOfmicrosoftGraphAssignedLicense and assigns it to the AssignedLicenses field.

### GetClassification

`func (o *MicrosoftGraphGroup) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *MicrosoftGraphGroup) GetClassificationOk() (string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassification

`func (o *MicrosoftGraphGroup) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassification

`func (o *MicrosoftGraphGroup) SetClassification(v string)`

SetClassification gets a reference to the given string and assigns it to the Classification field.

### SetClassificationExplicitNull

`func (o *MicrosoftGraphGroup) SetClassificationExplicitNull(b bool)`

SetClassificationExplicitNull (un)sets Classification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Classification value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphGroup) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphGroup) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphGroup) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphGroup) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphGroup) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphGroup) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphGroup) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphGroup) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphGroup) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphGroup) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphGroup) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetHasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) GetHasMembersWithLicenseErrors() bool`

GetHasMembersWithLicenseErrors returns the HasMembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetHasMembersWithLicenseErrorsOk

`func (o *MicrosoftGraphGroup) GetHasMembersWithLicenseErrorsOk() (bool, bool)`

GetHasMembersWithLicenseErrorsOk returns a tuple with the HasMembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) HasHasMembersWithLicenseErrors() bool`

HasHasMembersWithLicenseErrors returns a boolean if a field has been set.

### SetHasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) SetHasMembersWithLicenseErrors(v bool)`

SetHasMembersWithLicenseErrors gets a reference to the given bool and assigns it to the HasMembersWithLicenseErrors field.

### SetHasMembersWithLicenseErrorsExplicitNull

`func (o *MicrosoftGraphGroup) SetHasMembersWithLicenseErrorsExplicitNull(b bool)`

SetHasMembersWithLicenseErrorsExplicitNull (un)sets HasMembersWithLicenseErrors to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasMembersWithLicenseErrors value is set to nil even if false is passed
### GetGroupTypes

`func (o *MicrosoftGraphGroup) GetGroupTypes() []string`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *MicrosoftGraphGroup) GetGroupTypesOk() ([]string, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupTypes

`func (o *MicrosoftGraphGroup) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### SetGroupTypes

`func (o *MicrosoftGraphGroup) SetGroupTypes(v []string)`

SetGroupTypes gets a reference to the given []string and assigns it to the GroupTypes field.

### GetLicenseProcessingState

`func (o *MicrosoftGraphGroup) GetLicenseProcessingState() AnyOfmicrosoftGraphLicenseProcessingState`

GetLicenseProcessingState returns the LicenseProcessingState field if non-nil, zero value otherwise.

### GetLicenseProcessingStateOk

`func (o *MicrosoftGraphGroup) GetLicenseProcessingStateOk() (AnyOfmicrosoftGraphLicenseProcessingState, bool)`

GetLicenseProcessingStateOk returns a tuple with the LicenseProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseProcessingState

`func (o *MicrosoftGraphGroup) HasLicenseProcessingState() bool`

HasLicenseProcessingState returns a boolean if a field has been set.

### SetLicenseProcessingState

`func (o *MicrosoftGraphGroup) SetLicenseProcessingState(v AnyOfmicrosoftGraphLicenseProcessingState)`

SetLicenseProcessingState gets a reference to the given AnyOfmicrosoftGraphLicenseProcessingState and assigns it to the LicenseProcessingState field.

### SetLicenseProcessingStateExplicitNull

`func (o *MicrosoftGraphGroup) SetLicenseProcessingStateExplicitNull(b bool)`

SetLicenseProcessingStateExplicitNull (un)sets LicenseProcessingState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LicenseProcessingState value is set to nil even if false is passed
### GetMail

`func (o *MicrosoftGraphGroup) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *MicrosoftGraphGroup) GetMailOk() (string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMail

`func (o *MicrosoftGraphGroup) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMail

`func (o *MicrosoftGraphGroup) SetMail(v string)`

SetMail gets a reference to the given string and assigns it to the Mail field.

### SetMailExplicitNull

`func (o *MicrosoftGraphGroup) SetMailExplicitNull(b bool)`

SetMailExplicitNull (un)sets Mail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mail value is set to nil even if false is passed
### GetMailEnabled

`func (o *MicrosoftGraphGroup) GetMailEnabled() bool`

GetMailEnabled returns the MailEnabled field if non-nil, zero value otherwise.

### GetMailEnabledOk

`func (o *MicrosoftGraphGroup) GetMailEnabledOk() (bool, bool)`

GetMailEnabledOk returns a tuple with the MailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailEnabled

`func (o *MicrosoftGraphGroup) HasMailEnabled() bool`

HasMailEnabled returns a boolean if a field has been set.

### SetMailEnabled

`func (o *MicrosoftGraphGroup) SetMailEnabled(v bool)`

SetMailEnabled gets a reference to the given bool and assigns it to the MailEnabled field.

### SetMailEnabledExplicitNull

`func (o *MicrosoftGraphGroup) SetMailEnabledExplicitNull(b bool)`

SetMailEnabledExplicitNull (un)sets MailEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailEnabled value is set to nil even if false is passed
### GetMailNickname

`func (o *MicrosoftGraphGroup) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *MicrosoftGraphGroup) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *MicrosoftGraphGroup) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *MicrosoftGraphGroup) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### SetMailNicknameExplicitNull

`func (o *MicrosoftGraphGroup) SetMailNicknameExplicitNull(b bool)`

SetMailNicknameExplicitNull (un)sets MailNickname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailNickname value is set to nil even if false is passed
### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphGroup) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphGroup) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphGroup) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphGroup) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *MicrosoftGraphGroup) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesProvisioningErrors

`func (o *MicrosoftGraphGroup) GetOnPremisesProvisioningErrors() []AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *MicrosoftGraphGroup) GetOnPremisesProvisioningErrorsOk() ([]AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesProvisioningErrors

`func (o *MicrosoftGraphGroup) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### SetOnPremisesProvisioningErrors

`func (o *MicrosoftGraphGroup) SetOnPremisesProvisioningErrors(v []AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors gets a reference to the given []AnyOfmicrosoftGraphOnPremisesProvisioningError and assigns it to the OnPremisesProvisioningErrors field.

### GetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphGroup) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *MicrosoftGraphGroup) GetOnPremisesSecurityIdentifierOk() (string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphGroup) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphGroup) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier gets a reference to the given string and assigns it to the OnPremisesSecurityIdentifier field.

### SetOnPremisesSecurityIdentifierExplicitNull

`func (o *MicrosoftGraphGroup) SetOnPremisesSecurityIdentifierExplicitNull(b bool)`

SetOnPremisesSecurityIdentifierExplicitNull (un)sets OnPremisesSecurityIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSecurityIdentifier value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphGroup) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphGroup) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphGroup) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphGroup) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *MicrosoftGraphGroup) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetPreferredDataLocation

`func (o *MicrosoftGraphGroup) GetPreferredDataLocation() string`

GetPreferredDataLocation returns the PreferredDataLocation field if non-nil, zero value otherwise.

### GetPreferredDataLocationOk

`func (o *MicrosoftGraphGroup) GetPreferredDataLocationOk() (string, bool)`

GetPreferredDataLocationOk returns a tuple with the PreferredDataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredDataLocation

`func (o *MicrosoftGraphGroup) HasPreferredDataLocation() bool`

HasPreferredDataLocation returns a boolean if a field has been set.

### SetPreferredDataLocation

`func (o *MicrosoftGraphGroup) SetPreferredDataLocation(v string)`

SetPreferredDataLocation gets a reference to the given string and assigns it to the PreferredDataLocation field.

### SetPreferredDataLocationExplicitNull

`func (o *MicrosoftGraphGroup) SetPreferredDataLocationExplicitNull(b bool)`

SetPreferredDataLocationExplicitNull (un)sets PreferredDataLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredDataLocation value is set to nil even if false is passed
### GetProxyAddresses

`func (o *MicrosoftGraphGroup) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *MicrosoftGraphGroup) GetProxyAddressesOk() ([]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProxyAddresses

`func (o *MicrosoftGraphGroup) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### SetProxyAddresses

`func (o *MicrosoftGraphGroup) SetProxyAddresses(v []string)`

SetProxyAddresses gets a reference to the given []string and assigns it to the ProxyAddresses field.

### GetRenewedDateTime

`func (o *MicrosoftGraphGroup) GetRenewedDateTime() time.Time`

GetRenewedDateTime returns the RenewedDateTime field if non-nil, zero value otherwise.

### GetRenewedDateTimeOk

`func (o *MicrosoftGraphGroup) GetRenewedDateTimeOk() (time.Time, bool)`

GetRenewedDateTimeOk returns a tuple with the RenewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRenewedDateTime

`func (o *MicrosoftGraphGroup) HasRenewedDateTime() bool`

HasRenewedDateTime returns a boolean if a field has been set.

### SetRenewedDateTime

`func (o *MicrosoftGraphGroup) SetRenewedDateTime(v time.Time)`

SetRenewedDateTime gets a reference to the given time.Time and assigns it to the RenewedDateTime field.

### SetRenewedDateTimeExplicitNull

`func (o *MicrosoftGraphGroup) SetRenewedDateTimeExplicitNull(b bool)`

SetRenewedDateTimeExplicitNull (un)sets RenewedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RenewedDateTime value is set to nil even if false is passed
### GetSecurityEnabled

`func (o *MicrosoftGraphGroup) GetSecurityEnabled() bool`

GetSecurityEnabled returns the SecurityEnabled field if non-nil, zero value otherwise.

### GetSecurityEnabledOk

`func (o *MicrosoftGraphGroup) GetSecurityEnabledOk() (bool, bool)`

GetSecurityEnabledOk returns a tuple with the SecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityEnabled

`func (o *MicrosoftGraphGroup) HasSecurityEnabled() bool`

HasSecurityEnabled returns a boolean if a field has been set.

### SetSecurityEnabled

`func (o *MicrosoftGraphGroup) SetSecurityEnabled(v bool)`

SetSecurityEnabled gets a reference to the given bool and assigns it to the SecurityEnabled field.

### SetSecurityEnabledExplicitNull

`func (o *MicrosoftGraphGroup) SetSecurityEnabledExplicitNull(b bool)`

SetSecurityEnabledExplicitNull (un)sets SecurityEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SecurityEnabled value is set to nil even if false is passed
### GetVisibility

`func (o *MicrosoftGraphGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MicrosoftGraphGroup) GetVisibilityOk() (string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisibility

`func (o *MicrosoftGraphGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibility

`func (o *MicrosoftGraphGroup) SetVisibility(v string)`

SetVisibility gets a reference to the given string and assigns it to the Visibility field.

### SetVisibilityExplicitNull

`func (o *MicrosoftGraphGroup) SetVisibilityExplicitNull(b bool)`

SetVisibilityExplicitNull (un)sets Visibility to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Visibility value is set to nil even if false is passed
### GetAllowExternalSenders

`func (o *MicrosoftGraphGroup) GetAllowExternalSenders() bool`

GetAllowExternalSenders returns the AllowExternalSenders field if non-nil, zero value otherwise.

### GetAllowExternalSendersOk

`func (o *MicrosoftGraphGroup) GetAllowExternalSendersOk() (bool, bool)`

GetAllowExternalSendersOk returns a tuple with the AllowExternalSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowExternalSenders

`func (o *MicrosoftGraphGroup) HasAllowExternalSenders() bool`

HasAllowExternalSenders returns a boolean if a field has been set.

### SetAllowExternalSenders

`func (o *MicrosoftGraphGroup) SetAllowExternalSenders(v bool)`

SetAllowExternalSenders gets a reference to the given bool and assigns it to the AllowExternalSenders field.

### SetAllowExternalSendersExplicitNull

`func (o *MicrosoftGraphGroup) SetAllowExternalSendersExplicitNull(b bool)`

SetAllowExternalSendersExplicitNull (un)sets AllowExternalSenders to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AllowExternalSenders value is set to nil even if false is passed
### GetAutoSubscribeNewMembers

`func (o *MicrosoftGraphGroup) GetAutoSubscribeNewMembers() bool`

GetAutoSubscribeNewMembers returns the AutoSubscribeNewMembers field if non-nil, zero value otherwise.

### GetAutoSubscribeNewMembersOk

`func (o *MicrosoftGraphGroup) GetAutoSubscribeNewMembersOk() (bool, bool)`

GetAutoSubscribeNewMembersOk returns a tuple with the AutoSubscribeNewMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutoSubscribeNewMembers

`func (o *MicrosoftGraphGroup) HasAutoSubscribeNewMembers() bool`

HasAutoSubscribeNewMembers returns a boolean if a field has been set.

### SetAutoSubscribeNewMembers

`func (o *MicrosoftGraphGroup) SetAutoSubscribeNewMembers(v bool)`

SetAutoSubscribeNewMembers gets a reference to the given bool and assigns it to the AutoSubscribeNewMembers field.

### SetAutoSubscribeNewMembersExplicitNull

`func (o *MicrosoftGraphGroup) SetAutoSubscribeNewMembersExplicitNull(b bool)`

SetAutoSubscribeNewMembersExplicitNull (un)sets AutoSubscribeNewMembers to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AutoSubscribeNewMembers value is set to nil even if false is passed
### GetIsSubscribedByMail

`func (o *MicrosoftGraphGroup) GetIsSubscribedByMail() bool`

GetIsSubscribedByMail returns the IsSubscribedByMail field if non-nil, zero value otherwise.

### GetIsSubscribedByMailOk

`func (o *MicrosoftGraphGroup) GetIsSubscribedByMailOk() (bool, bool)`

GetIsSubscribedByMailOk returns a tuple with the IsSubscribedByMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSubscribedByMail

`func (o *MicrosoftGraphGroup) HasIsSubscribedByMail() bool`

HasIsSubscribedByMail returns a boolean if a field has been set.

### SetIsSubscribedByMail

`func (o *MicrosoftGraphGroup) SetIsSubscribedByMail(v bool)`

SetIsSubscribedByMail gets a reference to the given bool and assigns it to the IsSubscribedByMail field.

### SetIsSubscribedByMailExplicitNull

`func (o *MicrosoftGraphGroup) SetIsSubscribedByMailExplicitNull(b bool)`

SetIsSubscribedByMailExplicitNull (un)sets IsSubscribedByMail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsSubscribedByMail value is set to nil even if false is passed
### GetUnseenCount

`func (o *MicrosoftGraphGroup) GetUnseenCount() int32`

GetUnseenCount returns the UnseenCount field if non-nil, zero value otherwise.

### GetUnseenCountOk

`func (o *MicrosoftGraphGroup) GetUnseenCountOk() (int32, bool)`

GetUnseenCountOk returns a tuple with the UnseenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnseenCount

`func (o *MicrosoftGraphGroup) HasUnseenCount() bool`

HasUnseenCount returns a boolean if a field has been set.

### SetUnseenCount

`func (o *MicrosoftGraphGroup) SetUnseenCount(v int32)`

SetUnseenCount gets a reference to the given int32 and assigns it to the UnseenCount field.

### SetUnseenCountExplicitNull

`func (o *MicrosoftGraphGroup) SetUnseenCountExplicitNull(b bool)`

SetUnseenCountExplicitNull (un)sets UnseenCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnseenCount value is set to nil even if false is passed
### GetIsArchived

`func (o *MicrosoftGraphGroup) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *MicrosoftGraphGroup) GetIsArchivedOk() (bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsArchived

`func (o *MicrosoftGraphGroup) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchived

`func (o *MicrosoftGraphGroup) SetIsArchived(v bool)`

SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.

### SetIsArchivedExplicitNull

`func (o *MicrosoftGraphGroup) SetIsArchivedExplicitNull(b bool)`

SetIsArchivedExplicitNull (un)sets IsArchived to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsArchived value is set to nil even if false is passed
### GetMembers

`func (o *MicrosoftGraphGroup) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphGroup) GetMembersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *MicrosoftGraphGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *MicrosoftGraphGroup) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Members field.

### GetMemberOf

`func (o *MicrosoftGraphGroup) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MicrosoftGraphGroup) GetMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberOf

`func (o *MicrosoftGraphGroup) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOf

`func (o *MicrosoftGraphGroup) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.

### GetMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) GetMembersWithLicenseErrors() []MicrosoftGraphDirectoryObject`

GetMembersWithLicenseErrors returns the MembersWithLicenseErrors field if non-nil, zero value otherwise.

### GetMembersWithLicenseErrorsOk

`func (o *MicrosoftGraphGroup) GetMembersWithLicenseErrorsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMembersWithLicenseErrorsOk returns a tuple with the MembersWithLicenseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) HasMembersWithLicenseErrors() bool`

HasMembersWithLicenseErrors returns a boolean if a field has been set.

### SetMembersWithLicenseErrors

`func (o *MicrosoftGraphGroup) SetMembersWithLicenseErrors(v []MicrosoftGraphDirectoryObject)`

SetMembersWithLicenseErrors gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MembersWithLicenseErrors field.

### GetTransitiveMembers

`func (o *MicrosoftGraphGroup) GetTransitiveMembers() []MicrosoftGraphDirectoryObject`

GetTransitiveMembers returns the TransitiveMembers field if non-nil, zero value otherwise.

### GetTransitiveMembersOk

`func (o *MicrosoftGraphGroup) GetTransitiveMembersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMembersOk returns a tuple with the TransitiveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMembers

`func (o *MicrosoftGraphGroup) HasTransitiveMembers() bool`

HasTransitiveMembers returns a boolean if a field has been set.

### SetTransitiveMembers

`func (o *MicrosoftGraphGroup) SetTransitiveMembers(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMembers field.

### GetTransitiveMemberOf

`func (o *MicrosoftGraphGroup) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *MicrosoftGraphGroup) GetTransitiveMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMemberOf

`func (o *MicrosoftGraphGroup) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### SetTransitiveMemberOf

`func (o *MicrosoftGraphGroup) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.

### GetCreatedOnBehalfOf

`func (o *MicrosoftGraphGroup) GetCreatedOnBehalfOf() AnyOfmicrosoftGraphDirectoryObject`

GetCreatedOnBehalfOf returns the CreatedOnBehalfOf field if non-nil, zero value otherwise.

### GetCreatedOnBehalfOfOk

`func (o *MicrosoftGraphGroup) GetCreatedOnBehalfOfOk() (AnyOfmicrosoftGraphDirectoryObject, bool)`

GetCreatedOnBehalfOfOk returns a tuple with the CreatedOnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedOnBehalfOf

`func (o *MicrosoftGraphGroup) HasCreatedOnBehalfOf() bool`

HasCreatedOnBehalfOf returns a boolean if a field has been set.

### SetCreatedOnBehalfOf

`func (o *MicrosoftGraphGroup) SetCreatedOnBehalfOf(v AnyOfmicrosoftGraphDirectoryObject)`

SetCreatedOnBehalfOf gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the CreatedOnBehalfOf field.

### SetCreatedOnBehalfOfExplicitNull

`func (o *MicrosoftGraphGroup) SetCreatedOnBehalfOfExplicitNull(b bool)`

SetCreatedOnBehalfOfExplicitNull (un)sets CreatedOnBehalfOf to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedOnBehalfOf value is set to nil even if false is passed
### GetOwners

`func (o *MicrosoftGraphGroup) GetOwners() []MicrosoftGraphDirectoryObject`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *MicrosoftGraphGroup) GetOwnersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwners

`func (o *MicrosoftGraphGroup) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### SetOwners

`func (o *MicrosoftGraphGroup) SetOwners(v []MicrosoftGraphDirectoryObject)`

SetOwners gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Owners field.

### GetSettings

`func (o *MicrosoftGraphGroup) GetSettings() []MicrosoftGraphGroupSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphGroup) GetSettingsOk() ([]MicrosoftGraphGroupSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *MicrosoftGraphGroup) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *MicrosoftGraphGroup) SetSettings(v []MicrosoftGraphGroupSetting)`

SetSettings gets a reference to the given []MicrosoftGraphGroupSetting and assigns it to the Settings field.

### GetConversations

`func (o *MicrosoftGraphGroup) GetConversations() []MicrosoftGraphConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *MicrosoftGraphGroup) GetConversationsOk() ([]MicrosoftGraphConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversations

`func (o *MicrosoftGraphGroup) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### SetConversations

`func (o *MicrosoftGraphGroup) SetConversations(v []MicrosoftGraphConversation)`

SetConversations gets a reference to the given []MicrosoftGraphConversation and assigns it to the Conversations field.

### GetPhotos

`func (o *MicrosoftGraphGroup) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *MicrosoftGraphGroup) GetPhotosOk() ([]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhotos

`func (o *MicrosoftGraphGroup) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### SetPhotos

`func (o *MicrosoftGraphGroup) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos gets a reference to the given []MicrosoftGraphProfilePhoto and assigns it to the Photos field.

### GetAcceptedSenders

`func (o *MicrosoftGraphGroup) GetAcceptedSenders() []MicrosoftGraphDirectoryObject`

GetAcceptedSenders returns the AcceptedSenders field if non-nil, zero value otherwise.

### GetAcceptedSendersOk

`func (o *MicrosoftGraphGroup) GetAcceptedSendersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetAcceptedSendersOk returns a tuple with the AcceptedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcceptedSenders

`func (o *MicrosoftGraphGroup) HasAcceptedSenders() bool`

HasAcceptedSenders returns a boolean if a field has been set.

### SetAcceptedSenders

`func (o *MicrosoftGraphGroup) SetAcceptedSenders(v []MicrosoftGraphDirectoryObject)`

SetAcceptedSenders gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the AcceptedSenders field.

### GetRejectedSenders

`func (o *MicrosoftGraphGroup) GetRejectedSenders() []MicrosoftGraphDirectoryObject`

GetRejectedSenders returns the RejectedSenders field if non-nil, zero value otherwise.

### GetRejectedSendersOk

`func (o *MicrosoftGraphGroup) GetRejectedSendersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRejectedSendersOk returns a tuple with the RejectedSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRejectedSenders

`func (o *MicrosoftGraphGroup) HasRejectedSenders() bool`

HasRejectedSenders returns a boolean if a field has been set.

### SetRejectedSenders

`func (o *MicrosoftGraphGroup) SetRejectedSenders(v []MicrosoftGraphDirectoryObject)`

SetRejectedSenders gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RejectedSenders field.

### GetThreads

`func (o *MicrosoftGraphGroup) GetThreads() []MicrosoftGraphConversationThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *MicrosoftGraphGroup) GetThreadsOk() ([]MicrosoftGraphConversationThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThreads

`func (o *MicrosoftGraphGroup) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### SetThreads

`func (o *MicrosoftGraphGroup) SetThreads(v []MicrosoftGraphConversationThread)`

SetThreads gets a reference to the given []MicrosoftGraphConversationThread and assigns it to the Threads field.

### GetCalendar

`func (o *MicrosoftGraphGroup) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MicrosoftGraphGroup) GetCalendarOk() (AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendar

`func (o *MicrosoftGraphGroup) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendar

`func (o *MicrosoftGraphGroup) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar gets a reference to the given AnyOfmicrosoftGraphCalendar and assigns it to the Calendar field.

### SetCalendarExplicitNull

`func (o *MicrosoftGraphGroup) SetCalendarExplicitNull(b bool)`

SetCalendarExplicitNull (un)sets Calendar to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calendar value is set to nil even if false is passed
### GetCalendarView

`func (o *MicrosoftGraphGroup) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *MicrosoftGraphGroup) GetCalendarViewOk() ([]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarView

`func (o *MicrosoftGraphGroup) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### SetCalendarView

`func (o *MicrosoftGraphGroup) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView gets a reference to the given []MicrosoftGraphEvent and assigns it to the CalendarView field.

### GetEvents

`func (o *MicrosoftGraphGroup) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MicrosoftGraphGroup) GetEventsOk() ([]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *MicrosoftGraphGroup) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *MicrosoftGraphGroup) SetEvents(v []MicrosoftGraphEvent)`

SetEvents gets a reference to the given []MicrosoftGraphEvent and assigns it to the Events field.

### GetPhoto

`func (o *MicrosoftGraphGroup) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphGroup) GetPhotoOk() (AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *MicrosoftGraphGroup) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *MicrosoftGraphGroup) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphProfilePhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *MicrosoftGraphGroup) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetDrive

`func (o *MicrosoftGraphGroup) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *MicrosoftGraphGroup) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *MicrosoftGraphGroup) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *MicrosoftGraphGroup) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *MicrosoftGraphGroup) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetDrives

`func (o *MicrosoftGraphGroup) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *MicrosoftGraphGroup) GetDrivesOk() ([]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrives

`func (o *MicrosoftGraphGroup) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrives

`func (o *MicrosoftGraphGroup) SetDrives(v []MicrosoftGraphDrive)`

SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.

### GetSites

`func (o *MicrosoftGraphGroup) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *MicrosoftGraphGroup) GetSitesOk() ([]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSites

`func (o *MicrosoftGraphGroup) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSites

`func (o *MicrosoftGraphGroup) SetSites(v []MicrosoftGraphSite)`

SetSites gets a reference to the given []MicrosoftGraphSite and assigns it to the Sites field.

### GetExtensions

`func (o *MicrosoftGraphGroup) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphGroup) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphGroup) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphGroup) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetGroupLifecyclePolicies

`func (o *MicrosoftGraphGroup) GetGroupLifecyclePolicies() []MicrosoftGraphGroupLifecyclePolicy`

GetGroupLifecyclePolicies returns the GroupLifecyclePolicies field if non-nil, zero value otherwise.

### GetGroupLifecyclePoliciesOk

`func (o *MicrosoftGraphGroup) GetGroupLifecyclePoliciesOk() ([]MicrosoftGraphGroupLifecyclePolicy, bool)`

GetGroupLifecyclePoliciesOk returns a tuple with the GroupLifecyclePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupLifecyclePolicies

`func (o *MicrosoftGraphGroup) HasGroupLifecyclePolicies() bool`

HasGroupLifecyclePolicies returns a boolean if a field has been set.

### SetGroupLifecyclePolicies

`func (o *MicrosoftGraphGroup) SetGroupLifecyclePolicies(v []MicrosoftGraphGroupLifecyclePolicy)`

SetGroupLifecyclePolicies gets a reference to the given []MicrosoftGraphGroupLifecyclePolicy and assigns it to the GroupLifecyclePolicies field.

### GetPlanner

`func (o *MicrosoftGraphGroup) GetPlanner() AnyOfmicrosoftGraphPlannerGroup`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *MicrosoftGraphGroup) GetPlannerOk() (AnyOfmicrosoftGraphPlannerGroup, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanner

`func (o *MicrosoftGraphGroup) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlanner

`func (o *MicrosoftGraphGroup) SetPlanner(v AnyOfmicrosoftGraphPlannerGroup)`

SetPlanner gets a reference to the given AnyOfmicrosoftGraphPlannerGroup and assigns it to the Planner field.

### SetPlannerExplicitNull

`func (o *MicrosoftGraphGroup) SetPlannerExplicitNull(b bool)`

SetPlannerExplicitNull (un)sets Planner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Planner value is set to nil even if false is passed
### GetOnenote

`func (o *MicrosoftGraphGroup) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *MicrosoftGraphGroup) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnenote

`func (o *MicrosoftGraphGroup) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenote

`func (o *MicrosoftGraphGroup) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.

### SetOnenoteExplicitNull

`func (o *MicrosoftGraphGroup) SetOnenoteExplicitNull(b bool)`

SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Onenote value is set to nil even if false is passed
### GetTeam

`func (o *MicrosoftGraphGroup) GetTeam() AnyOfmicrosoftGraphTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MicrosoftGraphGroup) GetTeamOk() (AnyOfmicrosoftGraphTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeam

`func (o *MicrosoftGraphGroup) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeam

`func (o *MicrosoftGraphGroup) SetTeam(v AnyOfmicrosoftGraphTeam)`

SetTeam gets a reference to the given AnyOfmicrosoftGraphTeam and assigns it to the Team field.

### SetTeamExplicitNull

`func (o *MicrosoftGraphGroup) SetTeamExplicitNull(b bool)`

SetTeamExplicitNull (un)sets Team to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Team value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


