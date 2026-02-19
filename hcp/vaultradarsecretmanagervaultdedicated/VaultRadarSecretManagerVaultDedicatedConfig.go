// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarsecretmanagervaultdedicated

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultRadarSecretManagerVaultDedicatedConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Specify the URL of the Vault instance without protocol. Example: 'acme-public-vault-abc.def.z1.hashicorp.cloud:8200'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#vault_url VaultRadarSecretManagerVaultDedicated#vault_url}
	VaultUrl *string `field:"required" json:"vaultUrl" yaml:"vaultUrl"`
	// Indicates if the auth method has read and write access to the secrets engine.
	//
	// Defaults to false. Set this to true if you want to copy secrets to this secret manager as part of remediation process. (see https://developer.hashicorp.com/hcp/docs/vault-radar/remediate-secrets/copy-secrets)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#access_read_write VaultRadarSecretManagerVaultDedicated#access_read_write}
	AccessReadWrite interface{} `field:"optional" json:"accessReadWrite" yaml:"accessReadWrite"`
	// Configuration for AppRole Push-based authentication. Only one authentication method may be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#approle_push VaultRadarSecretManagerVaultDedicated#approle_push}
	ApprolePush *VaultRadarSecretManagerVaultDedicatedApprolePush `field:"optional" json:"approlePush" yaml:"approlePush"`
	// Configuration for Kubernetes-based authentication. Only one authentication method may be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#kubernetes VaultRadarSecretManagerVaultDedicated#kubernetes}
	Kubernetes *VaultRadarSecretManagerVaultDedicatedKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// The ID of the HCP project where Vault Radar is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#project_id VaultRadarSecretManagerVaultDedicated#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Configuration for token-based authentication. Only one authentication method may be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated#token VaultRadarSecretManagerVaultDedicated#token}
	Token *VaultRadarSecretManagerVaultDedicatedToken `field:"optional" json:"token" yaml:"token"`
}

