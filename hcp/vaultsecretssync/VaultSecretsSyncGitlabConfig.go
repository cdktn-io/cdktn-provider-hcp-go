// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretssync


type VaultSecretsSyncGitlabConfig struct {
	// ID of the group, if the scope is GROUP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_sync#group_id VaultSecretsSync#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// ID of the project, if the scope is PROJECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_sync#project_id VaultSecretsSync#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The scope to which sync applies. Defaults to GROUP. The valid options are GROUP and PROJECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_sync#scope VaultSecretsSync#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

