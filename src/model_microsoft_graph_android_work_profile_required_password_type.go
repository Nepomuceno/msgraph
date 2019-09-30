/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphAndroidWorkProfileRequiredPasswordType the model 'MicrosoftGraphAndroidWorkProfileRequiredPasswordType'
type MicrosoftGraphAndroidWorkProfileRequiredPasswordType string

// List of microsoft.graph.androidWorkProfileRequiredPasswordType
const (
	DEVICE_DEFAULT MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "deviceDefault"
	LOW_SECURITY_BIOMETRIC MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "lowSecurityBiometric"
	REQUIRED MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "required"
	AT_LEAST_NUMERIC MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "atLeastNumeric"
	NUMERIC_COMPLEX MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "numericComplex"
	AT_LEAST_ALPHABETIC MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "atLeastAlphabetic"
	AT_LEAST_ALPHANUMERIC MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "atLeastAlphanumeric"
	ALPHANUMERIC_WITH_SYMBOLS MicrosoftGraphAndroidWorkProfileRequiredPasswordType = "alphanumericWithSymbols"
)

