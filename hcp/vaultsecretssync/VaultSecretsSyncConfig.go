// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretssync

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultSecretsSyncConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Vault Secrets integration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_sync#integration_name VaultSecretsSync#integration_name}
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The Vault Secrets Sync name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_sync#name VaultSecretsSync#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration parameters used to determine the sync destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_sync#gitlab_config VaultSecretsSync#gitlab_config}
	GitlabConfig *VaultSecretsSyncGitlabConfig `field:"optional" json:"gitlabConfig" yaml:"gitlabConfig"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_sync#project_id VaultSecretsSync#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

