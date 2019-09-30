# MicrosoftGraphDeviceManagementExchangeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | Last sync time for the Exchange Connector | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus**](anyOf&lt;microsoft.graph.deviceManagementExchangeConnectorStatus&gt;.md) | Exchange Connector Status | [optional] 
**PrimarySmtpAddress** | Pointer to **string** | Email address used to configure the Service To Service Exchange Connector. | [optional] 
**ServerName** | Pointer to **string** | The name of the Exchange server. | [optional] 
**ConnectorServerName** | Pointer to **string** | The name of the server hosting the Exchange Connector. | [optional] 
**ExchangeConnectorType** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType**](anyOf&lt;microsoft.graph.deviceManagementExchangeConnectorType&gt;.md) | The type of Exchange Connector Configured. | [optional] 
**Version** | Pointer to **string** | The version of the ExchangeConnectorAgent | [optional] 
**ExchangeAlias** | Pointer to **string** | An alias assigned to the Exchange server | [optional] 
**ExchangeOrganization** | Pointer to **string** | Exchange Organization to the Exchange server | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastSyncDateTime

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetStatus

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetStatus() AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetStatusOk() (AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetStatus(v AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus and assigns it to the Status field.

### GetPrimarySmtpAddress

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetPrimarySmtpAddress() string`

GetPrimarySmtpAddress returns the PrimarySmtpAddress field if non-nil, zero value otherwise.

### GetPrimarySmtpAddressOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetPrimarySmtpAddressOk() (string, bool)`

GetPrimarySmtpAddressOk returns a tuple with the PrimarySmtpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimarySmtpAddress

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasPrimarySmtpAddress() bool`

HasPrimarySmtpAddress returns a boolean if a field has been set.

### SetPrimarySmtpAddress

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetPrimarySmtpAddress(v string)`

SetPrimarySmtpAddress gets a reference to the given string and assigns it to the PrimarySmtpAddress field.

### SetPrimarySmtpAddressExplicitNull

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetPrimarySmtpAddressExplicitNull(b bool)`

SetPrimarySmtpAddressExplicitNull (un)sets PrimarySmtpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrimarySmtpAddress value is set to nil even if false is passed
### GetServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetServerNameOk() (string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetServerName(v string)`

SetServerName gets a reference to the given string and assigns it to the ServerName field.

### SetServerNameExplicitNull

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetServerNameExplicitNull(b bool)`

SetServerNameExplicitNull (un)sets ServerName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ServerName value is set to nil even if false is passed
### GetConnectorServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetConnectorServerName() string`

GetConnectorServerName returns the ConnectorServerName field if non-nil, zero value otherwise.

### GetConnectorServerNameOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetConnectorServerNameOk() (string, bool)`

GetConnectorServerNameOk returns a tuple with the ConnectorServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectorServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasConnectorServerName() bool`

HasConnectorServerName returns a boolean if a field has been set.

### SetConnectorServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetConnectorServerName(v string)`

SetConnectorServerName gets a reference to the given string and assigns it to the ConnectorServerName field.

### SetConnectorServerNameExplicitNull

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetConnectorServerNameExplicitNull(b bool)`

SetConnectorServerNameExplicitNull (un)sets ConnectorServerName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConnectorServerName value is set to nil even if false is passed
### GetExchangeConnectorType

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeConnectorType() AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType`

GetExchangeConnectorType returns the ExchangeConnectorType field if non-nil, zero value otherwise.

### GetExchangeConnectorTypeOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeConnectorTypeOk() (AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType, bool)`

GetExchangeConnectorTypeOk returns a tuple with the ExchangeConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeConnectorType

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasExchangeConnectorType() bool`

HasExchangeConnectorType returns a boolean if a field has been set.

### SetExchangeConnectorType

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeConnectorType(v AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType)`

SetExchangeConnectorType gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType and assigns it to the ExchangeConnectorType field.

### GetVersion

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetExchangeAlias

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeAlias() string`

GetExchangeAlias returns the ExchangeAlias field if non-nil, zero value otherwise.

### GetExchangeAliasOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeAliasOk() (string, bool)`

GetExchangeAliasOk returns a tuple with the ExchangeAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeAlias

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasExchangeAlias() bool`

HasExchangeAlias returns a boolean if a field has been set.

### SetExchangeAlias

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeAlias(v string)`

SetExchangeAlias gets a reference to the given string and assigns it to the ExchangeAlias field.

### SetExchangeAliasExplicitNull

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeAliasExplicitNull(b bool)`

SetExchangeAliasExplicitNull (un)sets ExchangeAlias to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExchangeAlias value is set to nil even if false is passed
### GetExchangeOrganization

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeOrganization() string`

GetExchangeOrganization returns the ExchangeOrganization field if non-nil, zero value otherwise.

### GetExchangeOrganizationOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeOrganizationOk() (string, bool)`

GetExchangeOrganizationOk returns a tuple with the ExchangeOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeOrganization

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasExchangeOrganization() bool`

HasExchangeOrganization returns a boolean if a field has been set.

### SetExchangeOrganization

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeOrganization(v string)`

SetExchangeOrganization gets a reference to the given string and assigns it to the ExchangeOrganization field.

### SetExchangeOrganizationExplicitNull

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeOrganizationExplicitNull(b bool)`

SetExchangeOrganizationExplicitNull (un)sets ExchangeOrganization to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExchangeOrganization value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


