// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration


type VaultSecretsIntegrationMongodbAtlasStaticCredentials struct {
	// Private key used alongside the public key to authenticate against the target project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#api_private_key VaultSecretsIntegration#api_private_key}
	ApiPrivateKey *string `field:"required" json:"apiPrivateKey" yaml:"apiPrivateKey"`
	// Public key used alongside the private key to authenticate against the target project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#api_public_key VaultSecretsIntegration#api_public_key}
	ApiPublicKey *string `field:"required" json:"apiPublicKey" yaml:"apiPublicKey"`
}

