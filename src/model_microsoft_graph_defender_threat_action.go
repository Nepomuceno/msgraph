/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphDefenderThreatAction the model 'MicrosoftGraphDefenderThreatAction'
type MicrosoftGraphDefenderThreatAction string

// List of microsoft.graph.defenderThreatAction
const (
	DEVICE_DEFAULT MicrosoftGraphDefenderThreatAction = "deviceDefault"
	CLEAN MicrosoftGraphDefenderThreatAction = "clean"
	QUARANTINE MicrosoftGraphDefenderThreatAction = "quarantine"
	REMOVE MicrosoftGraphDefenderThreatAction = "remove"
	ALLOW MicrosoftGraphDefenderThreatAction = "allow"
	USER_DEFINED MicrosoftGraphDefenderThreatAction = "userDefined"
	BLOCK MicrosoftGraphDefenderThreatAction = "block"
)

