// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationGcpFederatedWorkloadIdentity struct {
	// Audience configured on the GCP identity provider to federate access with HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration#audience VaultSecretsIntegration#audience}
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// GCP service account email that HVS will impersonate to carry operations for the appropriate capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration#service_account_email VaultSecretsIntegration#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

