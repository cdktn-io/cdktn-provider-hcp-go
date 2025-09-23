// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationAzureClientSecret struct {
	// Azure client ID corresponding to the Azure application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration#client_id VaultSecretsIntegration#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Secret value corresponding to the Azure client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration#client_secret VaultSecretsIntegration#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Azure tenant ID corresponding to the Azure application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration#tenant_id VaultSecretsIntegration#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

