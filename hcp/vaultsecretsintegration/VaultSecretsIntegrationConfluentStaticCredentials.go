// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationConfluentStaticCredentials struct {
	// Public key used alongside the private key to authenticate for cloud apis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#cloud_api_key_id VaultSecretsIntegration#cloud_api_key_id}
	CloudApiKeyId *string `field:"required" json:"cloudApiKeyId" yaml:"cloudApiKeyId"`
	// Private key used alongside the public key to authenticate for cloud apis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#cloud_api_secret VaultSecretsIntegration#cloud_api_secret}
	CloudApiSecret *string `field:"required" json:"cloudApiSecret" yaml:"cloudApiSecret"`
}

