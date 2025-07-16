// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationmongodbatlas


type VaultSecretsIntegrationMongodbatlasStaticCredentialDetails struct {
	// Private key used alongside the public key to authenticate against the target project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_mongodbatlas#api_private_key VaultSecretsIntegrationMongodbatlas#api_private_key}
	ApiPrivateKey *string `field:"required" json:"apiPrivateKey" yaml:"apiPrivateKey"`
	// Public key used alongside the private key to authenticate against the target project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_mongodbatlas#api_public_key VaultSecretsIntegrationMongodbatlas#api_public_key}
	ApiPublicKey *string `field:"required" json:"apiPublicKey" yaml:"apiPublicKey"`
}

