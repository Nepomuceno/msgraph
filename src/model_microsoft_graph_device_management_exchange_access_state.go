/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphDeviceManagementExchangeAccessState the model 'MicrosoftGraphDeviceManagementExchangeAccessState'
type MicrosoftGraphDeviceManagementExchangeAccessState string

// List of microsoft.graph.deviceManagementExchangeAccessState
const (
	NONE MicrosoftGraphDeviceManagementExchangeAccessState = "none"
	UNKNOWN MicrosoftGraphDeviceManagementExchangeAccessState = "unknown"
	ALLOWED MicrosoftGraphDeviceManagementExchangeAccessState = "allowed"
	BLOCKED MicrosoftGraphDeviceManagementExchangeAccessState = "blocked"
	QUARANTINED MicrosoftGraphDeviceManagementExchangeAccessState = "quarantined"
)


