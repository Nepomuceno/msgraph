# MicrosoftGraphWindowsFirewallNetworkProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallEnabled** | Pointer to [**AnyOfmicrosoftGraphStateManagementSetting**](anyOf&lt;microsoft.graph.stateManagementSetting&gt;.md) | Configures the host device to allow or block the firewall and advanced security enforcement for the network profile. | [optional] 
**StealthModeBlocked** | Pointer to **bool** | Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority. | [optional] 
**IncomingTrafficBlocked** | Pointer to **bool** | Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority. | [optional] 
**UnicastResponsesToMulticastBroadcastsBlocked** | Pointer to **bool** | Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority. | [optional] 
**InboundNotificationsBlocked** | Pointer to **bool** | Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority. | [optional] 
**AuthorizedApplicationRulesFromGroupPolicyMerged** | Pointer to **bool** | Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority. | [optional] 
**GlobalPortRulesFromGroupPolicyMerged** | Pointer to **bool** | Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority. | [optional] 
**ConnectionSecurityRulesFromGroupPolicyMerged** | Pointer to **bool** | Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority. | [optional] 
**OutboundConnectionsBlocked** | Pointer to **bool** | Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. | [optional] 
**InboundConnectionsBlocked** | Pointer to **bool** | Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority. | [optional] 
**SecuredPacketExemptionAllowed** | Pointer to **bool** | Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority. | [optional] 
**PolicyRulesFromGroupPolicyMerged** | Pointer to **bool** | Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority. | [optional] 

## Methods

### GetFirewallEnabled

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetFirewallEnabled() AnyOfmicrosoftGraphStateManagementSetting`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetFirewallEnabledOk() (AnyOfmicrosoftGraphStateManagementSetting, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallEnabled

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### SetFirewallEnabled

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetFirewallEnabled(v AnyOfmicrosoftGraphStateManagementSetting)`

SetFirewallEnabled gets a reference to the given AnyOfmicrosoftGraphStateManagementSetting and assigns it to the FirewallEnabled field.

### GetStealthModeBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetStealthModeBlocked() bool`

GetStealthModeBlocked returns the StealthModeBlocked field if non-nil, zero value otherwise.

### GetStealthModeBlockedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetStealthModeBlockedOk() (bool, bool)`

GetStealthModeBlockedOk returns a tuple with the StealthModeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStealthModeBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasStealthModeBlocked() bool`

HasStealthModeBlocked returns a boolean if a field has been set.

### SetStealthModeBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetStealthModeBlocked(v bool)`

SetStealthModeBlocked gets a reference to the given bool and assigns it to the StealthModeBlocked field.

### GetIncomingTrafficBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetIncomingTrafficBlocked() bool`

GetIncomingTrafficBlocked returns the IncomingTrafficBlocked field if non-nil, zero value otherwise.

### GetIncomingTrafficBlockedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetIncomingTrafficBlockedOk() (bool, bool)`

GetIncomingTrafficBlockedOk returns a tuple with the IncomingTrafficBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncomingTrafficBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasIncomingTrafficBlocked() bool`

HasIncomingTrafficBlocked returns a boolean if a field has been set.

### SetIncomingTrafficBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetIncomingTrafficBlocked(v bool)`

SetIncomingTrafficBlocked gets a reference to the given bool and assigns it to the IncomingTrafficBlocked field.

### GetUnicastResponsesToMulticastBroadcastsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetUnicastResponsesToMulticastBroadcastsBlocked() bool`

GetUnicastResponsesToMulticastBroadcastsBlocked returns the UnicastResponsesToMulticastBroadcastsBlocked field if non-nil, zero value otherwise.

### GetUnicastResponsesToMulticastBroadcastsBlockedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetUnicastResponsesToMulticastBroadcastsBlockedOk() (bool, bool)`

GetUnicastResponsesToMulticastBroadcastsBlockedOk returns a tuple with the UnicastResponsesToMulticastBroadcastsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnicastResponsesToMulticastBroadcastsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasUnicastResponsesToMulticastBroadcastsBlocked() bool`

HasUnicastResponsesToMulticastBroadcastsBlocked returns a boolean if a field has been set.

### SetUnicastResponsesToMulticastBroadcastsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetUnicastResponsesToMulticastBroadcastsBlocked(v bool)`

SetUnicastResponsesToMulticastBroadcastsBlocked gets a reference to the given bool and assigns it to the UnicastResponsesToMulticastBroadcastsBlocked field.

### GetInboundNotificationsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetInboundNotificationsBlocked() bool`

GetInboundNotificationsBlocked returns the InboundNotificationsBlocked field if non-nil, zero value otherwise.

### GetInboundNotificationsBlockedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetInboundNotificationsBlockedOk() (bool, bool)`

GetInboundNotificationsBlockedOk returns a tuple with the InboundNotificationsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInboundNotificationsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasInboundNotificationsBlocked() bool`

HasInboundNotificationsBlocked returns a boolean if a field has been set.

### SetInboundNotificationsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetInboundNotificationsBlocked(v bool)`

SetInboundNotificationsBlocked gets a reference to the given bool and assigns it to the InboundNotificationsBlocked field.

### GetAuthorizedApplicationRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetAuthorizedApplicationRulesFromGroupPolicyMerged() bool`

GetAuthorizedApplicationRulesFromGroupPolicyMerged returns the AuthorizedApplicationRulesFromGroupPolicyMerged field if non-nil, zero value otherwise.

### GetAuthorizedApplicationRulesFromGroupPolicyMergedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetAuthorizedApplicationRulesFromGroupPolicyMergedOk() (bool, bool)`

GetAuthorizedApplicationRulesFromGroupPolicyMergedOk returns a tuple with the AuthorizedApplicationRulesFromGroupPolicyMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthorizedApplicationRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasAuthorizedApplicationRulesFromGroupPolicyMerged() bool`

HasAuthorizedApplicationRulesFromGroupPolicyMerged returns a boolean if a field has been set.

### SetAuthorizedApplicationRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetAuthorizedApplicationRulesFromGroupPolicyMerged(v bool)`

SetAuthorizedApplicationRulesFromGroupPolicyMerged gets a reference to the given bool and assigns it to the AuthorizedApplicationRulesFromGroupPolicyMerged field.

### GetGlobalPortRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetGlobalPortRulesFromGroupPolicyMerged() bool`

GetGlobalPortRulesFromGroupPolicyMerged returns the GlobalPortRulesFromGroupPolicyMerged field if non-nil, zero value otherwise.

### GetGlobalPortRulesFromGroupPolicyMergedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetGlobalPortRulesFromGroupPolicyMergedOk() (bool, bool)`

GetGlobalPortRulesFromGroupPolicyMergedOk returns a tuple with the GlobalPortRulesFromGroupPolicyMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGlobalPortRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasGlobalPortRulesFromGroupPolicyMerged() bool`

HasGlobalPortRulesFromGroupPolicyMerged returns a boolean if a field has been set.

### SetGlobalPortRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetGlobalPortRulesFromGroupPolicyMerged(v bool)`

SetGlobalPortRulesFromGroupPolicyMerged gets a reference to the given bool and assigns it to the GlobalPortRulesFromGroupPolicyMerged field.

### GetConnectionSecurityRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetConnectionSecurityRulesFromGroupPolicyMerged() bool`

GetConnectionSecurityRulesFromGroupPolicyMerged returns the ConnectionSecurityRulesFromGroupPolicyMerged field if non-nil, zero value otherwise.

### GetConnectionSecurityRulesFromGroupPolicyMergedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetConnectionSecurityRulesFromGroupPolicyMergedOk() (bool, bool)`

GetConnectionSecurityRulesFromGroupPolicyMergedOk returns a tuple with the ConnectionSecurityRulesFromGroupPolicyMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionSecurityRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasConnectionSecurityRulesFromGroupPolicyMerged() bool`

HasConnectionSecurityRulesFromGroupPolicyMerged returns a boolean if a field has been set.

### SetConnectionSecurityRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetConnectionSecurityRulesFromGroupPolicyMerged(v bool)`

SetConnectionSecurityRulesFromGroupPolicyMerged gets a reference to the given bool and assigns it to the ConnectionSecurityRulesFromGroupPolicyMerged field.

### GetOutboundConnectionsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetOutboundConnectionsBlocked() bool`

GetOutboundConnectionsBlocked returns the OutboundConnectionsBlocked field if non-nil, zero value otherwise.

### GetOutboundConnectionsBlockedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetOutboundConnectionsBlockedOk() (bool, bool)`

GetOutboundConnectionsBlockedOk returns a tuple with the OutboundConnectionsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutboundConnectionsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasOutboundConnectionsBlocked() bool`

HasOutboundConnectionsBlocked returns a boolean if a field has been set.

### SetOutboundConnectionsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetOutboundConnectionsBlocked(v bool)`

SetOutboundConnectionsBlocked gets a reference to the given bool and assigns it to the OutboundConnectionsBlocked field.

### GetInboundConnectionsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetInboundConnectionsBlocked() bool`

GetInboundConnectionsBlocked returns the InboundConnectionsBlocked field if non-nil, zero value otherwise.

### GetInboundConnectionsBlockedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetInboundConnectionsBlockedOk() (bool, bool)`

GetInboundConnectionsBlockedOk returns a tuple with the InboundConnectionsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInboundConnectionsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasInboundConnectionsBlocked() bool`

HasInboundConnectionsBlocked returns a boolean if a field has been set.

### SetInboundConnectionsBlocked

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetInboundConnectionsBlocked(v bool)`

SetInboundConnectionsBlocked gets a reference to the given bool and assigns it to the InboundConnectionsBlocked field.

### GetSecuredPacketExemptionAllowed

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetSecuredPacketExemptionAllowed() bool`

GetSecuredPacketExemptionAllowed returns the SecuredPacketExemptionAllowed field if non-nil, zero value otherwise.

### GetSecuredPacketExemptionAllowedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetSecuredPacketExemptionAllowedOk() (bool, bool)`

GetSecuredPacketExemptionAllowedOk returns a tuple with the SecuredPacketExemptionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecuredPacketExemptionAllowed

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasSecuredPacketExemptionAllowed() bool`

HasSecuredPacketExemptionAllowed returns a boolean if a field has been set.

### SetSecuredPacketExemptionAllowed

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetSecuredPacketExemptionAllowed(v bool)`

SetSecuredPacketExemptionAllowed gets a reference to the given bool and assigns it to the SecuredPacketExemptionAllowed field.

### GetPolicyRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetPolicyRulesFromGroupPolicyMerged() bool`

GetPolicyRulesFromGroupPolicyMerged returns the PolicyRulesFromGroupPolicyMerged field if non-nil, zero value otherwise.

### GetPolicyRulesFromGroupPolicyMergedOk

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) GetPolicyRulesFromGroupPolicyMergedOk() (bool, bool)`

GetPolicyRulesFromGroupPolicyMergedOk returns a tuple with the PolicyRulesFromGroupPolicyMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicyRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) HasPolicyRulesFromGroupPolicyMerged() bool`

HasPolicyRulesFromGroupPolicyMerged returns a boolean if a field has been set.

### SetPolicyRulesFromGroupPolicyMerged

`func (o *MicrosoftGraphWindowsFirewallNetworkProfile) SetPolicyRulesFromGroupPolicyMerged(v bool)`

SetPolicyRulesFromGroupPolicyMerged gets a reference to the given bool and assigns it to the PolicyRulesFromGroupPolicyMerged field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


