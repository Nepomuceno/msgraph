/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphFreeBusyStatus the model 'MicrosoftGraphFreeBusyStatus'
type MicrosoftGraphFreeBusyStatus string

// List of microsoft.graph.freeBusyStatus
const (
	FREE MicrosoftGraphFreeBusyStatus = "free"
	TENTATIVE MicrosoftGraphFreeBusyStatus = "tentative"
	BUSY MicrosoftGraphFreeBusyStatus = "busy"
	OOF MicrosoftGraphFreeBusyStatus = "oof"
	WORKING_ELSEWHERE MicrosoftGraphFreeBusyStatus = "workingElsewhere"
	UNKNOWN MicrosoftGraphFreeBusyStatus = "unknown"
)

