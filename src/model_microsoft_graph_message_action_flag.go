/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphMessageActionFlag the model 'MicrosoftGraphMessageActionFlag'
type MicrosoftGraphMessageActionFlag string

// List of microsoft.graph.messageActionFlag
const (
	ANY MicrosoftGraphMessageActionFlag = "any"
	CALL MicrosoftGraphMessageActionFlag = "call"
	DO_NOT_FORWARD MicrosoftGraphMessageActionFlag = "doNotForward"
	FOLLOW_UP MicrosoftGraphMessageActionFlag = "followUp"
	FYI MicrosoftGraphMessageActionFlag = "fyi"
	FORWARD MicrosoftGraphMessageActionFlag = "forward"
	NO_RESPONSE_NECESSARY MicrosoftGraphMessageActionFlag = "noResponseNecessary"
	READ MicrosoftGraphMessageActionFlag = "read"
	REPLY MicrosoftGraphMessageActionFlag = "reply"
	REPLY_TO_ALL MicrosoftGraphMessageActionFlag = "replyToAll"
	REVIEW MicrosoftGraphMessageActionFlag = "review"
)

