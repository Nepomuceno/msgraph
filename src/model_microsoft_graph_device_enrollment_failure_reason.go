/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphDeviceEnrollmentFailureReason the model 'MicrosoftGraphDeviceEnrollmentFailureReason'
type MicrosoftGraphDeviceEnrollmentFailureReason string

// List of microsoft.graph.deviceEnrollmentFailureReason
const (
	UNKNOWN MicrosoftGraphDeviceEnrollmentFailureReason = "unknown"
	AUTHENTICATION MicrosoftGraphDeviceEnrollmentFailureReason = "authentication"
	AUTHORIZATION MicrosoftGraphDeviceEnrollmentFailureReason = "authorization"
	ACCOUNT_VALIDATION MicrosoftGraphDeviceEnrollmentFailureReason = "accountValidation"
	USER_VALIDATION MicrosoftGraphDeviceEnrollmentFailureReason = "userValidation"
	DEVICE_NOT_SUPPORTED MicrosoftGraphDeviceEnrollmentFailureReason = "deviceNotSupported"
	IN_MAINTENANCE MicrosoftGraphDeviceEnrollmentFailureReason = "inMaintenance"
	BAD_REQUEST MicrosoftGraphDeviceEnrollmentFailureReason = "badRequest"
	FEATURE_NOT_SUPPORTED MicrosoftGraphDeviceEnrollmentFailureReason = "featureNotSupported"
	ENROLLMENT_RESTRICTIONS_ENFORCED MicrosoftGraphDeviceEnrollmentFailureReason = "enrollmentRestrictionsEnforced"
	CLIENT_DISCONNECTED MicrosoftGraphDeviceEnrollmentFailureReason = "clientDisconnected"
	USER_ABANDONMENT MicrosoftGraphDeviceEnrollmentFailureReason = "userAbandonment"
)

