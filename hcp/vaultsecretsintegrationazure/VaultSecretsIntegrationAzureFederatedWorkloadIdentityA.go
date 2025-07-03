// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationazure


type VaultSecretsIntegrationAzureFederatedWorkloadIdentityA struct {
	// Audience configured on the Azure federated identity credentials to federate access with HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_azure#audience VaultSecretsIntegrationAzure#audience}
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Azure client ID corresponding to the Azure application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_azure#client_id VaultSecretsIntegrationAzure#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Azure tenant ID corresponding to the Azure application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_azure#tenant_id VaultSecretsIntegrationAzure#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

