// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationaws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultSecretsIntegrationAwsConfig struct {
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
	// Capabilities enabled for the integration. See the Vault Secrets documentation for the list of supported capabilities per provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_aws#capabilities VaultSecretsIntegrationAws#capabilities}
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// The Vault Secrets integration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_aws#name VaultSecretsIntegrationAws#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS IAM key pair used to authenticate against the target AWS account. Cannot be used with `federated_workload_identity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_aws#access_keys VaultSecretsIntegrationAws#access_keys}
	AccessKeys *VaultSecretsIntegrationAwsAccessKeysA `field:"optional" json:"accessKeys" yaml:"accessKeys"`
	// (Recommended) Federated identity configuration to authenticate against the target AWS account. Cannot be used with `access_keys`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_aws#federated_workload_identity VaultSecretsIntegrationAws#federated_workload_identity}
	FederatedWorkloadIdentity *VaultSecretsIntegrationAwsFederatedWorkloadIdentityA `field:"optional" json:"federatedWorkloadIdentity" yaml:"federatedWorkloadIdentity"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_secrets_integration_aws#project_id VaultSecretsIntegrationAws#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

