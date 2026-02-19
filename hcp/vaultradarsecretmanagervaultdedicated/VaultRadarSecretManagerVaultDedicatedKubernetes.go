// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarsecretmanagervaultdedicated


type VaultRadarSecretManagerVaultDedicatedKubernetes struct {
	// Mount path where the Kubernetes auth method is enabled in Vault. Example 'kubernetes'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#mount_path VaultRadarSecretManagerVaultDedicated#mount_path}
	MountPath *string `field:"required" json:"mountPath" yaml:"mountPath"`
	// Kubernetes authentication role configured in Vault.  Example 'vault-radar-role'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#role_name VaultRadarSecretManagerVaultDedicated#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

