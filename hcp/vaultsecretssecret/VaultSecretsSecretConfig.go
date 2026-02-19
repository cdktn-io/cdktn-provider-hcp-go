// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretssecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultSecretsSecretConfig struct {
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
	// The name of the application the secret can be found in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_secret#app_name VaultSecretsSecret#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The name of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_secret#secret_name VaultSecretsSecret#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// The value of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_secret#secret_value VaultSecretsSecret#secret_value}
	SecretValue *string `field:"required" json:"secretValue" yaml:"secretValue"`
	// The ID of the HCP project where the HCP Vault Secrets secret is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_secret#project_id VaultSecretsSecret#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

