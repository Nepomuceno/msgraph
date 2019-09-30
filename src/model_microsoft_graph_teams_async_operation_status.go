/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphTeamsAsyncOperationStatus the model 'MicrosoftGraphTeamsAsyncOperationStatus'
type MicrosoftGraphTeamsAsyncOperationStatus string

// List of microsoft.graph.teamsAsyncOperationStatus
const (
	INVALID MicrosoftGraphTeamsAsyncOperationStatus = "invalid"
	NOT_STARTED MicrosoftGraphTeamsAsyncOperationStatus = "notStarted"
	IN_PROGRESS MicrosoftGraphTeamsAsyncOperationStatus = "inProgress"
	SUCCEEDED MicrosoftGraphTeamsAsyncOperationStatus = "succeeded"
	FAILED MicrosoftGraphTeamsAsyncOperationStatus = "failed"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamsAsyncOperationStatus = "unknownFutureValue"
)


