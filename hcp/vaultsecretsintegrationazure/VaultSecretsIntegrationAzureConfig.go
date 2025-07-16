// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationazure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultSecretsIntegrationAzureConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_azure#capabilities VaultSecretsIntegrationAzure#capabilities}
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// The Vault Secrets integration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_azure#name VaultSecretsIntegrationAzure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Azure client secret used to authenticate against the target Azure application. Cannot be used with `federated_workload_identity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_azure#client_secret VaultSecretsIntegrationAzure#client_secret}
	ClientSecret *VaultSecretsIntegrationAzureClientSecretA `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// (Recommended) Federated identity configuration to authenticate against the target Azure application. Cannot be used with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_azure#federated_workload_identity VaultSecretsIntegrationAzure#federated_workload_identity}
	FederatedWorkloadIdentity *VaultSecretsIntegrationAzureFederatedWorkloadIdentityA `field:"optional" json:"federatedWorkloadIdentity" yaml:"federatedWorkloadIdentity"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_integration_azure#project_id VaultSecretsIntegrationAzure#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

