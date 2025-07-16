// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationAzureFederatedWorkloadIdentity struct {
	// Audience configured on the Azure federated identity credentials to federate access with HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration#audience VaultSecretsIntegration#audience}
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Azure client ID corresponding to the Azure application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration#client_id VaultSecretsIntegration#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Azure tenant ID corresponding to the Azure application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration#tenant_id VaultSecretsIntegration#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

