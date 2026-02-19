// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationgcp


type VaultSecretsIntegrationGcpFederatedWorkloadIdentityA struct {
	// Audience configured on the GCP identity provider to federate access with HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_gcp#audience VaultSecretsIntegrationGcp#audience}
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// GCP service account email that HVS will impersonate to carry operations for the appropriate capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_gcp#service_account_email VaultSecretsIntegrationGcp#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

