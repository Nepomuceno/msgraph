/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphAppListType the model 'MicrosoftGraphAppListType'
type MicrosoftGraphAppListType string

// List of microsoft.graph.appListType
const (
	NONE MicrosoftGraphAppListType = "none"
	APPS_IN_LIST_COMPLIANT MicrosoftGraphAppListType = "appsInListCompliant"
	APPS_NOT_IN_LIST_COMPLIANT MicrosoftGraphAppListType = "appsNotInListCompliant"
)

