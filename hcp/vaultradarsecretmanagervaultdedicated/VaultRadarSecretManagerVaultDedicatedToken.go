// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultradarsecretmanagervaultdedicated


type VaultRadarSecretManagerVaultDedicatedToken struct {
	// Environment variable name containing the Vault token. Example: 'VAULT_TOKEN'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#token_env_var VaultRadarSecretManagerVaultDedicated#token_env_var}
	TokenEnvVar *string `field:"required" json:"tokenEnvVar" yaml:"tokenEnvVar"`
}

