/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphManagedAppDataEncryptionType the model 'MicrosoftGraphManagedAppDataEncryptionType'
type MicrosoftGraphManagedAppDataEncryptionType string

// List of microsoft.graph.managedAppDataEncryptionType
const (
	USE_DEVICE_SETTINGS MicrosoftGraphManagedAppDataEncryptionType = "useDeviceSettings"
	AFTER_DEVICE_RESTART MicrosoftGraphManagedAppDataEncryptionType = "afterDeviceRestart"
	WHEN_DEVICE_LOCKED_EXCEPT_OPEN_FILES MicrosoftGraphManagedAppDataEncryptionType = "whenDeviceLockedExceptOpenFiles"
	WHEN_DEVICE_LOCKED MicrosoftGraphManagedAppDataEncryptionType = "whenDeviceLocked"
)

