/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphProcessIntegrityLevel the model 'MicrosoftGraphProcessIntegrityLevel'
type MicrosoftGraphProcessIntegrityLevel string

// List of microsoft.graph.processIntegrityLevel
const (
	UNKNOWN MicrosoftGraphProcessIntegrityLevel = "unknown"
	UNTRUSTED MicrosoftGraphProcessIntegrityLevel = "untrusted"
	LOW MicrosoftGraphProcessIntegrityLevel = "low"
	MEDIUM MicrosoftGraphProcessIntegrityLevel = "medium"
	HIGH MicrosoftGraphProcessIntegrityLevel = "high"
	SYSTEM MicrosoftGraphProcessIntegrityLevel = "system"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphProcessIntegrityLevel = "unknownFutureValue"
)


