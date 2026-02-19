// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegrationtwilio

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VaultSecretsIntegrationTwilioConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_twilio#capabilities VaultSecretsIntegrationTwilio#capabilities}
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// The Vault Secrets integration name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_twilio#name VaultSecretsIntegrationTwilio#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// HCP project ID that owns the HCP Vault Secrets integration. Inferred from the provider configuration if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_twilio#project_id VaultSecretsIntegrationTwilio#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Twilio API key parts used to authenticate against the target Twilio account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_integration_twilio#static_credential_details VaultSecretsIntegrationTwilio#static_credential_details}
	StaticCredentialDetails *VaultSecretsIntegrationTwilioStaticCredentialDetails `field:"optional" json:"staticCredentialDetails" yaml:"staticCredentialDetails"`
}

