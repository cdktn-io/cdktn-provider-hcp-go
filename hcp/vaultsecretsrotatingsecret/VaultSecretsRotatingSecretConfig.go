// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsrotatingsecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultSecretsRotatingSecretConfig struct {
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
	// Vault Secrets application name that owns the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#app_name VaultSecretsRotatingSecret#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The Vault Secrets integration name with the capability to manage the secret's lifecycle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#integration_name VaultSecretsRotatingSecret#integration_name}
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The Vault Secrets secret name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#name VaultSecretsRotatingSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the rotation policy that governs the rotation of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#rotation_policy_name VaultSecretsRotatingSecret#rotation_policy_name}
	RotationPolicyName *string `field:"required" json:"rotationPolicyName" yaml:"rotationPolicyName"`
	// The third party platform the dynamic credentials give access to. One of `aws` or `gcp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#secret_provider VaultSecretsRotatingSecret#secret_provider}
	SecretProvider *string `field:"required" json:"secretProvider" yaml:"secretProvider"`
	// AWS configuration to manage the access key rotation for the given IAM user. Required if `secret_provider` is `aws`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#aws_access_keys VaultSecretsRotatingSecret#aws_access_keys}
	AwsAccessKeys *VaultSecretsRotatingSecretAwsAccessKeys `field:"optional" json:"awsAccessKeys" yaml:"awsAccessKeys"`
	// Azure configuration to manage the application password rotation for the given application. Required if `secret_provider` is `Azure`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#azure_application_password VaultSecretsRotatingSecret#azure_application_password}
	AzureApplicationPassword *VaultSecretsRotatingSecretAzureApplicationPassword `field:"optional" json:"azureApplicationPassword" yaml:"azureApplicationPassword"`
	// Confluent configuration to manage the cloud api key rotation for the given service account. Required if `secret_provider` is `confluent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#confluent_service_account VaultSecretsRotatingSecret#confluent_service_account}
	ConfluentServiceAccount *VaultSecretsRotatingSecretConfluentServiceAccount `field:"optional" json:"confluentServiceAccount" yaml:"confluentServiceAccount"`
	// GCP configuration to manage the service account key rotation for the given service account. Required if `secret_provider` is `gcp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#gcp_service_account_key VaultSecretsRotatingSecret#gcp_service_account_key}
	GcpServiceAccountKey *VaultSecretsRotatingSecretGcpServiceAccountKey `field:"optional" json:"gcpServiceAccountKey" yaml:"gcpServiceAccountKey"`
	// MongoDB Atlas configuration to manage the user password rotation on the given database. Required if `secret_provider` is `mongodb_atlas`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#mongodb_atlas_user VaultSecretsRotatingSecret#mongodb_atlas_user}
	MongodbAtlasUser *VaultSecretsRotatingSecretMongodbAtlasUser `field:"optional" json:"mongodbAtlasUser" yaml:"mongodbAtlasUser"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#project_id VaultSecretsRotatingSecret#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Twilio configuration to manage the api key rotation on the given account. Required if `secret_provider` is `twilio`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_secrets_rotating_secret#twilio_api_key VaultSecretsRotatingSecret#twilio_api_key}
	TwilioApiKey *VaultSecretsRotatingSecretTwilioApiKey `field:"optional" json:"twilioApiKey" yaml:"twilioApiKey"`
}

