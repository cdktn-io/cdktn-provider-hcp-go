// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultSecretsIntegrationConfig struct {
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
	// Capabilities enabled for the integration. See the Vault Secrets documentation for the list of supported capabilities per provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#capabilities VaultSecretsIntegration#capabilities}
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// The Vault Secrets integration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#name VaultSecretsIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The provider or 3rd party platform the integration is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#provider_type VaultSecretsIntegration#provider_type}
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// AWS IAM key pair used to authenticate against the target AWS account. Cannot be used with `federated_workload_identity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#aws_access_keys VaultSecretsIntegration#aws_access_keys}
	AwsAccessKeys *VaultSecretsIntegrationAwsAccessKeys `field:"optional" json:"awsAccessKeys" yaml:"awsAccessKeys"`
	// (Recommended) Federated identity configuration to authenticate against the target AWS account. Cannot be used with `access_keys`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#aws_federated_workload_identity VaultSecretsIntegration#aws_federated_workload_identity}
	AwsFederatedWorkloadIdentity *VaultSecretsIntegrationAwsFederatedWorkloadIdentity `field:"optional" json:"awsFederatedWorkloadIdentity" yaml:"awsFederatedWorkloadIdentity"`
	// Azure client secret used to authenticate against the target Azure application. Cannot be used with `federated_workload_identity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#azure_client_secret VaultSecretsIntegration#azure_client_secret}
	AzureClientSecret *VaultSecretsIntegrationAzureClientSecret `field:"optional" json:"azureClientSecret" yaml:"azureClientSecret"`
	// (Recommended) Federated identity configuration to authenticate against the target Azure application. Cannot be used with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#azure_federated_workload_identity VaultSecretsIntegration#azure_federated_workload_identity}
	AzureFederatedWorkloadIdentity *VaultSecretsIntegrationAzureFederatedWorkloadIdentity `field:"optional" json:"azureFederatedWorkloadIdentity" yaml:"azureFederatedWorkloadIdentity"`
	// Confluent API key used to authenticate for cloud apis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#confluent_static_credentials VaultSecretsIntegration#confluent_static_credentials}
	ConfluentStaticCredentials *VaultSecretsIntegrationConfluentStaticCredentials `field:"optional" json:"confluentStaticCredentials" yaml:"confluentStaticCredentials"`
	// (Recommended) Federated identity configuration to authenticate against the target GCP project. Cannot be used with `service_account_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#gcp_federated_workload_identity VaultSecretsIntegration#gcp_federated_workload_identity}
	GcpFederatedWorkloadIdentity *VaultSecretsIntegrationGcpFederatedWorkloadIdentity `field:"optional" json:"gcpFederatedWorkloadIdentity" yaml:"gcpFederatedWorkloadIdentity"`
	// GCP service account key used to authenticate against the target GCP project. Cannot be used with `federated_workload_identity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#gcp_service_account_key VaultSecretsIntegration#gcp_service_account_key}
	GcpServiceAccountKey *VaultSecretsIntegrationGcpServiceAccountKey `field:"optional" json:"gcpServiceAccountKey" yaml:"gcpServiceAccountKey"`
	// GitLab access token used to authenticate against the target GitLab account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#gitlab_access VaultSecretsIntegration#gitlab_access}
	GitlabAccess *VaultSecretsIntegrationGitlabAccess `field:"optional" json:"gitlabAccess" yaml:"gitlabAccess"`
	// MongoDB Atlas API key used to authenticate against the target project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#mongodb_atlas_static_credentials VaultSecretsIntegration#mongodb_atlas_static_credentials}
	MongodbAtlasStaticCredentials *VaultSecretsIntegrationMongodbAtlasStaticCredentials `field:"optional" json:"mongodbAtlasStaticCredentials" yaml:"mongodbAtlasStaticCredentials"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#project_id VaultSecretsIntegration#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Twilio API key parts used to authenticate against the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration#twilio_static_credentials VaultSecretsIntegration#twilio_static_credentials}
	TwilioStaticCredentials *VaultSecretsIntegrationTwilioStaticCredentials `field:"optional" json:"twilioStaticCredentials" yaml:"twilioStaticCredentials"`
}

