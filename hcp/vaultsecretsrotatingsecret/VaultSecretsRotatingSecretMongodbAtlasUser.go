// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsrotatingsecret


type VaultSecretsRotatingSecretMongodbAtlasUser struct {
	// MongoDB Atlas database or cluster name to rotate the username and password for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#database_name VaultSecretsRotatingSecret#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// MongoDB Atlas project ID to rotate the username and password for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#project_id VaultSecretsRotatingSecret#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// MongoDB Atlas roles to assign to the rotating user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#roles VaultSecretsRotatingSecret#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

