// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsdynamicsecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultSecretsDynamicSecretConfig struct {
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
	// Vault Secrets application name that owns the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#app_name VaultSecretsDynamicSecret#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The Vault Secrets integration name with the capability to manage the secret's lifecycle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#integration_name VaultSecretsDynamicSecret#integration_name}
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The Vault Secrets secret name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#name VaultSecretsDynamicSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The third party platform the dynamic credentials give access to. One of `aws` or `gcp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#secret_provider VaultSecretsDynamicSecret#secret_provider}
	SecretProvider *string `field:"required" json:"secretProvider" yaml:"secretProvider"`
	// AWS configuration to generate dynamic credentials by assuming an IAM role. Required if `secret_provider` is `aws`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#aws_assume_role VaultSecretsDynamicSecret#aws_assume_role}
	AwsAssumeRole *VaultSecretsDynamicSecretAwsAssumeRole `field:"optional" json:"awsAssumeRole" yaml:"awsAssumeRole"`
	// TTL the generated credentials will be valid for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#default_ttl VaultSecretsDynamicSecret#default_ttl}
	DefaultTtl *string `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// GCP configuration to generate dynamic credentials by impersonating a service account. Required if `secret_provider` is `gcp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#gcp_impersonate_service_account VaultSecretsDynamicSecret#gcp_impersonate_service_account}
	GcpImpersonateServiceAccount *VaultSecretsDynamicSecretGcpImpersonateServiceAccount `field:"optional" json:"gcpImpersonateServiceAccount" yaml:"gcpImpersonateServiceAccount"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_dynamic_secret#project_id VaultSecretsDynamicSecret#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

