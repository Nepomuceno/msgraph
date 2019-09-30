/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
// MicrosoftGraphFirewallCertificateRevocationListCheckMethodType the model 'MicrosoftGraphFirewallCertificateRevocationListCheckMethodType'
type MicrosoftGraphFirewallCertificateRevocationListCheckMethodType string

// List of microsoft.graph.firewallCertificateRevocationListCheckMethodType
const (
	DEVICE_DEFAULT MicrosoftGraphFirewallCertificateRevocationListCheckMethodType = "deviceDefault"
	NONE MicrosoftGraphFirewallCertificateRevocationListCheckMethodType = "none"
	ATTEMPT MicrosoftGraphFirewallCertificateRevocationListCheckMethodType = "attempt"
	REQUIRE MicrosoftGraphFirewallCertificateRevocationListCheckMethodType = "require"
)

