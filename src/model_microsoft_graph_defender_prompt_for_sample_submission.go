/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphDefenderPromptForSampleSubmission the model 'MicrosoftGraphDefenderPromptForSampleSubmission'
type MicrosoftGraphDefenderPromptForSampleSubmission string

// List of microsoft.graph.defenderPromptForSampleSubmission
const (
	USER_DEFINED MicrosoftGraphDefenderPromptForSampleSubmission = "userDefined"
	ALWAYS_PROMPT MicrosoftGraphDefenderPromptForSampleSubmission = "alwaysPrompt"
	PROMPT_BEFORE_SENDING_PERSONAL_DATA MicrosoftGraphDefenderPromptForSampleSubmission = "promptBeforeSendingPersonalData"
	NEVER_SEND_DATA MicrosoftGraphDefenderPromptForSampleSubmission = "neverSendData"
	SEND_ALL_DATA_WITHOUT_PROMPTING MicrosoftGraphDefenderPromptForSampleSubmission = "sendAllDataWithoutPrompting"
)


