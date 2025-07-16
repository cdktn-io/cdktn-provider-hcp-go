// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultSecretsAppConfig struct {
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
	// The Vault Secrets App name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_app#app_name VaultSecretsApp#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The Vault Secrets app description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_app#description VaultSecretsApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the HCP project where the HCP Vault Secrets app is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_app#project_id VaultSecretsApp#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Set of sync names to associate with this app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_app#sync_names VaultSecretsApp#sync_names}
	SyncNames *[]*string `field:"optional" json:"syncNames" yaml:"syncNames"`
}

