/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphTeamVisibilityType the model 'MicrosoftGraphTeamVisibilityType'
type MicrosoftGraphTeamVisibilityType string

// List of microsoft.graph.teamVisibilityType
const (
	PRIVATE MicrosoftGraphTeamVisibilityType = "private"
	PUBLIC MicrosoftGraphTeamVisibilityType = "public"
	HIDDEN_MEMBERSHIP MicrosoftGraphTeamVisibilityType = "hiddenMembership"
	UNKNOWN_FUTURE_VALUE MicrosoftGraphTeamVisibilityType = "unknownFutureValue"
)

