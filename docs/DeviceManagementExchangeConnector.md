# DeviceManagementExchangeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetLastSyncDateTime

`func (o *DeviceManagementExchangeConnector) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *DeviceManagementExchangeConnector) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *DeviceManagementExchangeConnector) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *DeviceManagementExchangeConnector) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetStatus

`func (o *DeviceManagementExchangeConnector) GetStatus() AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceManagementExchangeConnector) GetStatusOk() (AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *DeviceManagementExchangeConnector) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *DeviceManagementExchangeConnector) SetStatus(v AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus and assigns it to the Status field.

### GetPrimarySmtpAddress

`func (o *DeviceManagementExchangeConnector) GetPrimarySmtpAddress() string`

GetPrimarySmtpAddress returns the PrimarySmtpAddress field if non-nil, zero value otherwise.

### GetPrimarySmtpAddressOk

`func (o *DeviceManagementExchangeConnector) GetPrimarySmtpAddressOk() (string, bool)`

GetPrimarySmtpAddressOk returns a tuple with the PrimarySmtpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimarySmtpAddress

`func (o *DeviceManagementExchangeConnector) HasPrimarySmtpAddress() bool`

HasPrimarySmtpAddress returns a boolean if a field has been set.

### SetPrimarySmtpAddress

`func (o *DeviceManagementExchangeConnector) SetPrimarySmtpAddress(v string)`

SetPrimarySmtpAddress gets a reference to the given string and assigns it to the PrimarySmtpAddress field.

### SetPrimarySmtpAddressExplicitNull

`func (o *DeviceManagementExchangeConnector) SetPrimarySmtpAddressExplicitNull(b bool)`

SetPrimarySmtpAddressExplicitNull (un)sets PrimarySmtpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrimarySmtpAddress value is set to nil even if false is passed
### GetServerName

`func (o *DeviceManagementExchangeConnector) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DeviceManagementExchangeConnector) GetServerNameOk() (string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServerName

`func (o *DeviceManagementExchangeConnector) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerName

`func (o *DeviceManagementExchangeConnector) SetServerName(v string)`

SetServerName gets a reference to the given string and assigns it to the ServerName field.

### SetServerNameExplicitNull

`func (o *DeviceManagementExchangeConnector) SetServerNameExplicitNull(b bool)`

SetServerNameExplicitNull (un)sets ServerName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ServerName value is set to nil even if false is passed
### GetConnectorServerName

`func (o *DeviceManagementExchangeConnector) GetConnectorServerName() string`

GetConnectorServerName returns the ConnectorServerName field if non-nil, zero value otherwise.

### GetConnectorServerNameOk

`func (o *DeviceManagementExchangeConnector) GetConnectorServerNameOk() (string, bool)`

GetConnectorServerNameOk returns a tuple with the ConnectorServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectorServerName

`func (o *DeviceManagementExchangeConnector) HasConnectorServerName() bool`

HasConnectorServerName returns a boolean if a field has been set.

### SetConnectorServerName

`func (o *DeviceManagementExchangeConnector) SetConnectorServerName(v string)`

SetConnectorServerName gets a reference to the given string and assigns it to the ConnectorServerName field.

### SetConnectorServerNameExplicitNull

`func (o *DeviceManagementExchangeConnector) SetConnectorServerNameExplicitNull(b bool)`

SetConnectorServerNameExplicitNull (un)sets ConnectorServerName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConnectorServerName value is set to nil even if false is passed
### GetExchangeConnectorType

`func (o *DeviceManagementExchangeConnector) GetExchangeConnectorType() AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType`

GetExchangeConnectorType returns the ExchangeConnectorType field if non-nil, zero value otherwise.

### GetExchangeConnectorTypeOk

`func (o *DeviceManagementExchangeConnector) GetExchangeConnectorTypeOk() (AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType, bool)`

GetExchangeConnectorTypeOk returns a tuple with the ExchangeConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeConnectorType

`func (o *DeviceManagementExchangeConnector) HasExchangeConnectorType() bool`

HasExchangeConnectorType returns a boolean if a field has been set.

### SetExchangeConnectorType

`func (o *DeviceManagementExchangeConnector) SetExchangeConnectorType(v AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType)`

SetExchangeConnectorType gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType and assigns it to the ExchangeConnectorType field.

### GetVersion

`func (o *DeviceManagementExchangeConnector) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceManagementExchangeConnector) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *DeviceManagementExchangeConnector) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *DeviceManagementExchangeConnector) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *DeviceManagementExchangeConnector) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetExchangeAlias

`func (o *DeviceManagementExchangeConnector) GetExchangeAlias() string`

GetExchangeAlias returns the ExchangeAlias field if non-nil, zero value otherwise.

### GetExchangeAliasOk

`func (o *DeviceManagementExchangeConnector) GetExchangeAliasOk() (string, bool)`

GetExchangeAliasOk returns a tuple with the ExchangeAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeAlias

`func (o *DeviceManagementExchangeConnector) HasExchangeAlias() bool`

HasExchangeAlias returns a boolean if a field has been set.

### SetExchangeAlias

`func (o *DeviceManagementExchangeConnector) SetExchangeAlias(v string)`

SetExchangeAlias gets a reference to the given string and assigns it to the ExchangeAlias field.

### SetExchangeAliasExplicitNull

`func (o *DeviceManagementExchangeConnector) SetExchangeAliasExplicitNull(b bool)`

SetExchangeAliasExplicitNull (un)sets ExchangeAlias to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExchangeAlias value is set to nil even if false is passed
### GetExchangeOrganization

`func (o *DeviceManagementExchangeConnector) GetExchangeOrganization() string`

GetExchangeOrganization returns the ExchangeOrganization field if non-nil, zero value otherwise.

### GetExchangeOrganizationOk

`func (o *DeviceManagementExchangeConnector) GetExchangeOrganizationOk() (string, bool)`

GetExchangeOrganizationOk returns a tuple with the ExchangeOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeOrganization

`func (o *DeviceManagementExchangeConnector) HasExchangeOrganization() bool`

HasExchangeOrganization returns a boolean if a field has been set.

### SetExchangeOrganization

`func (o *DeviceManagementExchangeConnector) SetExchangeOrganization(v string)`

SetExchangeOrganization gets a reference to the given string and assigns it to the ExchangeOrganization field.

### SetExchangeOrganizationExplicitNull

`func (o *DeviceManagementExchangeConnector) SetExchangeOrganizationExplicitNull(b bool)`

SetExchangeOrganizationExplicitNull (un)sets ExchangeOrganization to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExchangeOrganization value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


