/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphEdgeCookiePolicy the model 'MicrosoftGraphEdgeCookiePolicy'
type MicrosoftGraphEdgeCookiePolicy string

// List of microsoft.graph.edgeCookiePolicy
const (
	USER_DEFINED MicrosoftGraphEdgeCookiePolicy = "userDefined"
	ALLOW MicrosoftGraphEdgeCookiePolicy = "allow"
	BLOCK_THIRD_PARTY MicrosoftGraphEdgeCookiePolicy = "blockThirdParty"
	BLOCK_ALL MicrosoftGraphEdgeCookiePolicy = "blockAll"
)

