// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarsecretmanagervaultdedicated


type VaultRadarSecretManagerVaultDedicatedApprolePush struct {
	// Mount path of the AppRole auth method in Vault. Example 'approle'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#mount_path VaultRadarSecretManagerVaultDedicated#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Environment variable containing the AppRole role ID. Example: 'VAULT_APPROLE_ROLE_ID'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#role_id_env_var VaultRadarSecretManagerVaultDedicated#role_id_env_var}
	RoleIdEnvVar *string `field:"required" json:"roleIdEnvVar" yaml:"roleIdEnvVar"`
	// Environment variable containing the AppRole secret ID. Example: 'VAULT_APPROLE_SECRET_ID'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#secret_id_env_var VaultRadarSecretManagerVaultDedicated#secret_id_env_var}
	SecretIdEnvVar *string `field:"required" json:"secretIdEnvVar" yaml:"secretIdEnvVar"`
}

