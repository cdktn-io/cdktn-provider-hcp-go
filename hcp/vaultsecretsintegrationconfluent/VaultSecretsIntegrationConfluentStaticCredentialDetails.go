// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationconfluent


type VaultSecretsIntegrationConfluentStaticCredentialDetails struct {
	// Public key used alongside the private key to authenticate for cloud apis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration_confluent#cloud_api_key_id VaultSecretsIntegrationConfluent#cloud_api_key_id}
	CloudApiKeyId *string `field:"required" json:"cloudApiKeyId" yaml:"cloudApiKeyId"`
	// Private key used alongside the public key to authenticate for cloud apis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration_confluent#cloud_api_secret VaultSecretsIntegrationConfluent#cloud_api_secret}
	CloudApiSecret *string `field:"required" json:"cloudApiSecret" yaml:"cloudApiSecret"`
}

